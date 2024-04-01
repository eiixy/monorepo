package helpers

import (
	"bytes"
	"github.com/klauspost/compress/zstd"
	"io"
)

func ZstdEncode(in []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer, err := zstd.NewWriter(&buf)
	if err != nil {
		_ = writer.Close()
		return nil, err
	}
	_, err = writer.Write(in)
	if err != nil {
		_ = writer.Close()
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ZstdDecode(in []byte) ([]byte, error) {
	reader, err := zstd.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(reader)
}
