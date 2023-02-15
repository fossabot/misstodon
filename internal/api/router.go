package api

import (
	"github.com/gizmo-ds/misstodon/internal/api/httperror"
	middleware2 "github.com/gizmo-ds/misstodon/internal/api/middleware"
	"github.com/gizmo-ds/misstodon/internal/api/nodeinfo"
	"github.com/gizmo-ds/misstodon/internal/api/oauth"
	v12 "github.com/gizmo-ds/misstodon/internal/api/v1"
	"github.com/gizmo-ds/misstodon/internal/api/wellknown"
	"github.com/gizmo-ds/misstodon/internal/global"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.HTTPErrorHandler = httperror.ErrorHandler
	e.Use(
		middleware2.SetContextData,
		middleware2.Recover())
	if global.Config.Logger.RequestLogger {
		e.Use(middleware2.Logger)
	}
	for _, group := range []*echo.Group{
		e.Group(""),
		e.Group("/:proxyServer"),
	} {
		wellknown.Router(group)
		nodeinfo.Router(group)
		oauth.Router(group)
		v1Api := group.Group("/api/v1", middleware2.CORS())
		group.GET("/static/missing.png", v12.MissingImageHandler)
		v12.InstanceRouter(v1Api)
		v12.AccountsRouter(v1Api)
		v12.ApplicationRouter(v1Api)
		v12.StatusesRouter(v1Api)
		v12.StreamingRouter(v1Api)
	}
}