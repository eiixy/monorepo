package middleware

import "github.com/go-kratos/gateway/middleware"

func init() {
	middleware.Register("monorepo-auth", AuthMiddleware)
	middleware.Register("monorepo-permission", PermissionMiddleware)
}
