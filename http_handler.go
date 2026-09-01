package zap

import (
	"io"
	"net/http"

	"go.uber.org/zap/zapcore"
)

func (lvl AtomicLevel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (lvl AtomicLevel) serveHTTP(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func decodePutRequest(contentType string, r *http.Request) (zapcore.Level, error) {
	_ = "STUB: not implemented"
	return *new(zapcore.Level), nil
}

func decodePutURL(r *http.Request) (zapcore.Level, error) {
	_ = "STUB: not implemented"
	return *new(zapcore.Level), nil
}

func decodePutJSON(body io.Reader) (zapcore.Level, error) {
	_ = "STUB: not implemented"
	return *new(zapcore.Level), nil
}
