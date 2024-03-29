package config

import (
	"github.com/eiixy/monorepo/internal/pkg/request"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"time"
)

type Server struct {
	Network  string
	Addr     string
	Timeout  int // 毫秒
	Metadata map[string]string
}

func (s Server) GrpcOptions(logger log.Logger, opts ...grpc.ServerOption) []grpc.ServerOption {
	opts = append(opts, grpc.Middleware(
		recovery.Recovery(),
		logging.Server(logger),
	))
	if s.Network != "" {
		opts = append(opts, grpc.Network(s.Network))
	}
	if s.Addr != "" {
		opts = append(opts, grpc.Address(s.Addr))
	}
	if s.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(s.Timeout)*time.Millisecond))
	}
	return opts
}

func (s Server) HttpOptions(logger log.Logger, opts ...http.ServerOption) []http.ServerOption {
	opts = append(opts, http.Middleware(
		recovery.Recovery(),
		logging.Server(logger),
		request.Validator,
	))

	if s.Network != "" {
		opts = append(opts, http.Network(s.Network))
	}
	if s.Addr != "" {
		opts = append(opts, http.Address(s.Addr))
	}
	if s.Timeout != 0 {
		opts = append(opts, http.Timeout(time.Duration(s.Timeout)*time.Millisecond))
	}
	return opts
}
