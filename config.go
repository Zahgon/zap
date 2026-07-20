package zap

import (
	"go.uber.org/zap/zapcore"
)

type SamplingConfig struct {
	Initial    int                                           `json:"initial" yaml:"initial"`
	Thereafter int                                           `json:"thereafter" yaml:"thereafter"`
	Hook       func(zapcore.Entry, zapcore.SamplingDecision) `json:"-" yaml:"-"`
}

type Config struct {
	Level AtomicLevel `json:"level" yaml:"level"`

	Development bool `json:"development" yaml:"development"`

	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`

	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`

	Sampling *SamplingConfig `json:"sampling" yaml:"sampling"`

	Encoding string `json:"encoding" yaml:"encoding"`

	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`

	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`

	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`

	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`
}

func NewProductionEncoderConfig() zapcore.EncoderConfig {
	_ = "STUB: not implemented"
	return *new(zapcore.EncoderConfig)
}

func NewProductionConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

func NewDevelopmentEncoderConfig() zapcore.EncoderConfig {
	_ = "STUB: not implemented"
	return *new(zapcore.EncoderConfig)
}

func NewDevelopmentConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

func (cfg Config) Build(opts ...Option) (*Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfg Config) buildOptions(errSink zapcore.WriteSyncer) []Option {
	_ = "STUB: not implemented"
	return nil
}

func (cfg Config) openSinks() (zapcore.WriteSyncer, zapcore.WriteSyncer, error) {
	_ = "STUB: not implemented"
	return *new(zapcore.WriteSyncer), *new(zapcore.WriteSyncer), nil
}

func (cfg Config) buildEncoder() (zapcore.Encoder, error) {
	_ = "STUB: not implemented"
	return *new(zapcore.Encoder), nil
}
