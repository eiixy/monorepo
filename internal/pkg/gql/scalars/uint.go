package scalars

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/spf13/cast"
	"io"
	"strconv"
)

func MarshalUint(t uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(int64(t), 10))
	})
}

func UnmarshalUint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case string, int, int64, uint, uint64:
		return cast.ToUintE(v)
	case json.Number:
		u64, err := v.Int64()
		return uint(u64), err
	}
	return 0, fmt.Errorf("%T is not an uint", v)
}
