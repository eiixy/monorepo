package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/spf13/cast"
)

const (
	HeaderAccountID = "X-Account-ID"
)

func GetAccountID(ctx context.Context) int {
	if tr, ok := transport.FromServerContext(ctx); ok {
		accountId := tr.RequestHeader().Get(HeaderAccountID)
		return cast.ToInt(accountId)
	}
	return 0
}
