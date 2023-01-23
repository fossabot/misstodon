package api

import (
	"github.com/gizmo-ds/misstodon/api/httperror"
	"github.com/gizmo-ds/misstodon/api/middleware"
	"github.com/gizmo-ds/misstodon/api/nodeinfo"
	"github.com/gizmo-ds/misstodon/api/oauth"
	"github.com/gizmo-ds/misstodon/api/v1"
	"github.com/gizmo-ds/misstodon/api/wellknown"
	"github.com/gizmo-ds/misstodon/internal/global"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {
	e.HTTPErrorHandler = httperror.ErrorHandler
	e.Use(middleware.SetContextData)
	if global.Config.Logger.RequestLogger {
		e.Use(middleware.Logger)
	}
	for _, group := range []*echo.Group{
		e.Group(""),
		e.Group("/:proxyServer"),
	} {
		wellknown.Router(group)
		nodeinfo.Router(group)
		v1Api := group.Group("/api/v1", echomiddleware.CORS())
		oauth.Router(group)
		v1.InstanceRouter(v1Api)
		v1.AccountsRouter(v1Api)
		v1.ApplicationRouter(v1Api)
	}
}
