package gql

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrDeprecated 方法已弃用
var ErrDeprecated = &gqlerror.Error{
	Message: "deprecated",
	Extensions: map[string]interface{}{
		"code": "Deprecated",
	},
}

// ErrNoPermissions 无权限
var ErrNoPermissions = &gqlerror.Error{
	Message: "no permissions",
	Extensions: map[string]interface{}{
		"code": "NoPermission",
	},
}

// ErrAccessDenied 拒绝访问，无角色或身份
var ErrAccessDenied = &gqlerror.Error{
	Message: "access denied",
	Extensions: map[string]interface{}{
		"code": "AccessDenied",
	},
}

// ErrUnauthorized 未授权
var ErrUnauthorized = &gqlerror.Error{
	Message: "Unauthorized",
	Extensions: map[string]interface{}{
		"code": "Unauthorized",
	},
}
