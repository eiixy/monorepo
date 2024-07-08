package response

import (
	"encoding/json"
	"fmt"
	"github.com/eiixy/monorepo/apis/third_party/response"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

type ErrorEncoderOption func(resp *response.Response, err error) *response.Response

func ErrorEncoder(opts ...ErrorEncoderOption) http.EncodeErrorFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
		e := response.FromError(err)
		for _, opt := range opts {
			e = opt(e, err)
		}
		resp := e.GetResponse(nil)
		body, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(int(e.GetHttpCode()))
		_, _ = w.Write(body)
	}
}

func Encoder() http.EncodeResponseFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request, v interface{}) error {
		codec, _ := http.CodecForRequest(r, "Accept")
		data, err := codec.Marshal(v)
		if err != nil {
			return err
		}
		data = []byte(fmt.Sprintf(`{"code":%d,"data":%s,"msg":"%s","reason":"%s"}`,
			response.SuccessCode,
			data,
			response.SuccessMessage,
			response.SuccessReason,
		))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(nethttp.StatusOK)
		_, err = w.Write(data)
		return err
	}
}
