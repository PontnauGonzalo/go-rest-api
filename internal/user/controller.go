package user

import (
	"context"
	"errors"

	"github.com/LuchoNicolosi/go-web-response/response"
)

type (
	UserController func(ctx context.Context, data interface{}) (interface{}, error)
	Endpoints      struct {
		Create  UserController
		GetAll  UserController
		GetById UserController
		Update  UserController
		Delete  UserController
	}

	GetReq struct {
		UserID uint64
	}
	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	UpdateRequest struct {
		UserID    uint64
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	DeleteReq struct {
		UserID uint64
	}
)

func MakeEndpoints(ctx context.Context, service UserService) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(service),
		GetAll:  makeGetAllEndpoint(service),
		GetById: makeGetByIdEndpoint(service),
		Update:  makeUpdateEndpoint(service),
		Delete:  makeDeleteEndpoint(service),
	}
}

func makeGetAllEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		users, err := service.GetAll(ctx)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}
		return response.OK("success", users), nil
	}
}
func makeGetByIdEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		result := data.(GetReq)
		user, err := service.GetById(ctx, result.UserID)
		if err != nil {
			if errors.As(err, &ErrNotFound{}) {
				return nil, response.NotFound(err.Error())
			}

			return nil, response.InternalServerError(err.Error())
		}
		return response.OK("success", user), nil
	}
}
func makeDeleteEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		result := data.(DeleteReq)
		if err := service.Delete(ctx, result.UserID); err != nil {
			if errors.As(err, &ErrNotFound{}) {
				return nil, response.NotFound(err.Error())
			}
			return nil, response.InternalServerError(err.Error())
		}
		return response.OK("success", nil), nil
	}
}

func makeCreateEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		reqData := data.(CreateRequest)

		if reqData.FirstName == "" {
			return nil, response.BadRequest(ErrFistNameRequeried.Error())
		}
		if reqData.LastName == "" {
			return nil, response.BadRequest(ErrLastNameRequeried.Error())
		}
		if reqData.Email == "" {
			return nil, response.BadRequest(ErrEmailRequeried.Error())
		}
		user, err := service.Create(ctx, reqData.FirstName, reqData.LastName, reqData.Email)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.Created("success", user), nil
	}
}
func makeUpdateEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		reqData := data.(UpdateRequest)

		if err := service.Update(ctx, reqData.UserID, reqData.FirstName, reqData.LastName, reqData.Email); err != nil {

			if errors.As(err, &ErrNotFound{}) {
				return nil, response.NotFound(err.Error())
			}
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", nil), nil
	}
}