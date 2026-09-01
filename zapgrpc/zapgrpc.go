package zapgrpc

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	grpcLvlInfo int = iota
	grpcLvlWarn
	grpcLvlError
	grpcLvlFatal
)

var _grpcToZapLevel = map[int]zapcore.Level{
	grpcLvlInfo:  zapcore.InfoLevel,
	grpcLvlWarn:  zapcore.WarnLevel,
	grpcLvlError: zapcore.ErrorLevel,
	grpcLvlFatal: zapcore.FatalLevel,
}

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) { _ = "STUB: not implemented"; return }

func WithDebug() Option { _ = "STUB: not implemented"; return *new(Option) }

func withWarn() Option { _ = "STUB: not implemented"; return *new(Option) }

func NewLogger(l *zap.Logger, options ...Option) *Logger { _ = "STUB: not implemented"; return nil }

type printer struct {
	enab   zapcore.LevelEnabler
	level  zapcore.Level
	print  func(...interface{})
	printf func(string, ...interface{})
}

func (v *printer) Print(args ...interface{}) { _ = "STUB: not implemented"; return }

func (v *printer) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (v *printer) Println(args ...interface{}) { _ = "STUB: not implemented"; return }

type Logger struct {
	delegate     *zap.SugaredLogger
	levelEnabler zapcore.LevelEnabler
	print        *printer
	fatal        *printer
}

func (l *Logger) Print(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Println(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Info(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Infoln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Warning(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Warningln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Warningf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Errorln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Fatalln(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Logger) V(level int) bool { _ = "STUB: not implemented"; return false }

func sprintln(args []interface{}) string { _ = "STUB: not implemented"; return "" }
