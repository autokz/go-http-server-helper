package routes

import (
	"github.com/autokz/go-http-server-helper/httpHelper"
	"github.com/autokz/go-http-server-helper/server_example/handler"
	"github.com/autokz/go-http-server-helper/server_example/middleware"
)

var TestRoutes2 = httpHelper.Routes{
	UrlPrefix: "/v2",
	Routes: []httpHelper.Route{
		Route3,
	},
	Middlewares: []httpHelper.Middleware{
		middleware.MainMiddleware2,
	},
}

var Route3 = httpHelper.Route{
	UrlPattern: "/test",
	Method:     "POST",
	Action:     handler.Test,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.M2,
	},
}

var Route4 = httpHelper.Route{
	UrlPattern: "/test4",
	Method:     "POST",
	Action:     handler.Test,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.M2,
	},
}
