package request

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	nethttp "net/http"
)

func Decoder() http.DecodeRequestFunc {
	return func(r *nethttp.Request, v interface{}) error {
		if _, ok := v.(*emptypb.Empty); ok {
			return nil
		}
		codec, ok := http.CodecForRequest(r, "Content-Type")
		if !ok {
			return fmt.Errorf("codec: %s", r.Header.Get("Content-Type"))
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if err = codec.Unmarshal(data, v); err != nil {
			return err
		}
		return nil
	}
}
