package routes

import (
	"github.com/autokz/go-http-server-helper/httpHelper"
	"github.com/autokz/go-http-server-helper/server_example/handler"
	"github.com/autokz/go-http-server-helper/server_example/middleware"
)

var TestRoutes = httpHelper.Routes{
	UrlPrefix: "/v1",
	Routes: []httpHelper.Route{
		Route1,
		Route2,
	},
	Middlewares: []httpHelper.Middleware{
		middleware.MainMiddleware1,
		middleware.MainMiddleware2,
	},
}

var Route1 = httpHelper.Route{
	UrlPattern: "/test",
	Method:     "POST",
	Action:     handler.Test,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.M1,
		middleware.M2,
	},
}

var Route2 = httpHelper.Route{
	UrlPattern: "/test2",
	Method:     "POST",
	Action:     handler.Test,
	RouteMiddlewares: []httpHelper.Middleware{
		middleware.M1,
	},
}
