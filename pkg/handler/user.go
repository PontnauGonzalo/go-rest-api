package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/PontnauGonzalo/go-rest-api/internal/user"
	"github.com/PontnauGonzalo/go-rest-api/pkg/transport"
	"github.com/PontnauGonzalo/go-rest-api/response"
	"github.com/gin-gonic/gin"
)

func NewUserHTTPServer(endpoints user.Endpoints) http.Handler {
	router := gin.Default()

	router.POST("/users", transport.GinServer(
		transport.Endpoint(endpoints.Create),
		decodeCreateUser,
		encodeResponse,
		encodeError,
	))
	router.GET("/users", transport.GinServer(
		transport.Endpoint(endpoints.GetAll),
		decodeGetAllUser,
		encodeResponse,
		encodeError,
	))
	router.GET("/users/:id", transport.GinServer(
		transport.Endpoint(endpoints.GetById),
		decodeGetUser,
		encodeResponse,
		encodeError,
	))
	router.PUT("/users/:id", transport.GinServer(
		transport.Endpoint(endpoints.Update),
		decodeUpdateUser,
		encodeResponse,
		encodeError,
	))
	router.DELETE("/users/:id", transport.GinServer(
		transport.Endpoint(endpoints.Delete),
		decodeDeleteUser,
		encodeResponse,
		encodeError,
	))
	return router
}

func decodeGetUser(c *gin.Context) (interface{}, error) {
	if err := tokenVerify(c.Request.Header.Get("Authorization")); err != nil {
		return nil, response.Unauthorized(err.Error())
	}

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return user.GetReq{
		UserID: id,
	}, nil
}

func decodeGetAllUser(c *gin.Context) (interface{}, error) {

	return nil, nil
}
func decodeCreateUser(c *gin.Context) (interface{}, error) {
	token := c.Request.Header.Get("Authorization")

	if err := tokenVerify(token); err != nil {
		return nil, response.Unauthorized(err.Error())
	}

	var data user.CreateRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}
	return data, nil
}
func decodeUpdateUser(c *gin.Context) (interface{}, error) {
	var data user.UpdateRequest

	if err := tokenVerify(c.Request.Header.Get("Authorization")); err != nil {
		return nil, response.Unauthorized(err.Error())
	}

	userId, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}

	data.UserID = userId

	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}

	return data, nil
}

func decodeDeleteUser(c *gin.Context) (interface{}, error) {
	if err := tokenVerify(c.Request.Header.Get("Authorization")); err != nil {
		return nil, response.Unauthorized(err.Error())
	}
	userId, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}
	return user.DeleteReq{
		UserID: userId,
	}, nil
}
func encodeResponse(c *gin.Context, data interface{}) {
	resData := data.(response.Response)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(resData.StatusCode(), data)
}

func encodeError(c *gin.Context, err error) {
	errData := err.(response.Response)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(errData.StatusCode(), errData)
}

func tokenVerify(token string) error {
	if os.Getenv("TOKEN") != token {
		return errors.New("invalid token")
	}
	return nil
}