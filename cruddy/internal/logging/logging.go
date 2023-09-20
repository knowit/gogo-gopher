package logging

import (
  "net/http"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

type LoggingMiddleWare struct {
  logger *zap.Logger
}

func InitZap() (*zap.Logger, error) {
  level := zap.NewAtomicLevelAt(zapcore.InfoLevel)

  zapEncoderConfig := zapcore.EncoderConfig{
    TimeKey: "ts",
    LevelKey: "level",
    NameKey: "logger",
    CallerKey: "caller",
    MessageKey: "msg",
    StacktraceKey: "stacktrace",
    LineEnding: zapcore.DefaultLineEnding,
    EncodeLevel: zapcore.LowercaseLevelEncoder,
    EncodeTime: zapcore.ISO8601TimeEncoder,
    EncodeDuration: zapcore.SecondsDurationEncoder,
    EncodeCaller: zapcore.ShortCallerEncoder,
  }

  zapConfig := zap.Config{
    Level: level,
    Development: false,
    Sampling: &zap.SamplingConfig{
      Initial: 100,
      Thereafter: 100,
    },
    Encoding: "json",
    EncoderConfig: zapEncoderConfig,
    OutputPaths: []string{"stderr"},
    ErrorOutputPaths: []string{"stderr"},
  }

  return zapConfig.Build()
}

func NewLoggingMiddleware(logger *zap.Logger) *LoggingMiddleWare {
  return &LoggingMiddleWare{
    logger: logger,
  }
}

func (lm *LoggingMiddleWare) Handler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
    lm.logger.Debug(
      "request started",
      zap.String("proto", request.Proto),
      zap.String("uri", request.RequestURI),
      zap.String("method", request.Method),
      zap.String("remote", request.RemoteAddr),
      zap.String("user-agent", request.UserAgent()),
    )
    next.ServeHTTP(writer, request)
  })
}
