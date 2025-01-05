package transport

import (
	"context"
	"net/http"
	"strings"
)

type Transport interface {
	Server(
		endpoint Endpoint,
		decode func(ctx context.Context, req *http.Request) (interface{}, error),
		encode func(ctx context.Context, res http.ResponseWriter, data interface{}) error,
		encodeError func(ctx context.Context, err error, res http.ResponseWriter),
	)
}

type Endpoint func(ctx context.Context, data interface{}) (interface{}, error)

type transport struct {
	res http.ResponseWriter
	req *http.Request
	ctx context.Context
}

func New(res http.ResponseWriter, req *http.Request, ctx context.Context) Transport {
	return &transport{
		res,
		req,
		ctx,
	}
}
func (t *transport) Server(
	endpoint Endpoint,
	decode func(ctx context.Context, req *http.Request) (interface{}, error),
	encode func(ctx context.Context, res http.ResponseWriter, data interface{}) error,
	encodeError func(ctx context.Context, err error, res http.ResponseWriter),
) {
	data, err := decode(t.ctx, t.req)
	if err != nil {
		encodeError(t.ctx, err, t.res)
		return
	}
	resultData, err := endpoint(t.ctx, data)
	if err != nil {
		encodeError(t.ctx, err, t.res)
		return
	}

	if err := encode(t.ctx, t.res, resultData); err != nil {
		encodeError(t.ctx, err, t.res)
		return
	}
}

func Clean(url string) ([]string, int) {
	if url[0] != '/' {
		url = "/" + url
	}
	if url[len(url)-1] != '/' {
		url = url + "/"
	}

	parts := strings.Split(url, "/")

	return parts, len(parts)
}
