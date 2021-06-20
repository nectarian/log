package log

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/komkom/toml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// --------------------------------------------------------------------------------------------------
const (
	defaultConfigFilePath = "logger.toml"
)

func init() {
	var err error
	conf := Config{}
	var content []byte
	if content, err = ioutil.ReadFile(defaultConfigFilePath); err != nil {
		conf.Default()
		logger = InitLogger(conf)
		return
	}

	decoder := json.NewDecoder(toml.New(bytes.NewBuffer(content)))
	if err = decoder.Decode(&conf); err != nil {
		conf.Default()
		logger = InitLogger(conf)
		return
	}
	conf.Fix()
	logger = InitLogger(conf)
}

// --------------------------------------------------------------------------------------------------

// logger main logger
var logger *zap.Logger

// InitLogger initialize main logger
func InitLogger(conf Config) *zap.Logger {
	// config output
	var ws = []zapcore.WriteSyncer{}
	if conf.Output.File {
		hook := lumberjack.Logger{
			Filename:   conf.Output.FilePath,
			MaxSize:    conf.Output.MaxSize,
			MaxBackups: conf.Output.MaxBackups,
			MaxAge:     conf.Output.MaxAge,
			Compress:   conf.Output.Compress,
		}
		ws = append(ws, zapcore.AddSync(&hook))
	}
	if conf.Output.Console {
		ws = append(ws, zapcore.AddSync(os.Stdout))
	}

	// config Encoder
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// set main logger leave
	var level = zap.NewAtomicLevel()
	if f := level.UnmarshalText([]byte(conf.Level)); f != nil {
		level.SetLevel(zap.InfoLevel)
	}

	// set format
	var format zapcore.Encoder
	if conf.Context.Format == "json" {
		format = zapcore.NewJSONEncoder(config)
	} else {
		format = zapcore.NewConsoleEncoder(config)
	}

	// create uber zap core
	core := zapcore.NewCore(format, zapcore.NewMultiWriteSyncer(ws...), level)

	// set additional fields if exists
	var fields []zapcore.Field
	for k, v := range conf.Context.Fileds {
		fields = append(fields, zap.String(k, v))
	}
	// build main logger
	opts := make([]zap.Option, 0, 3)
	opts = append(opts, zap.Fields(fields...))
	if conf.Context.ShowCaller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	return zap.New(core, opts...)
}

// Debug call main logger Debug function
func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

// Info call main logger Info function
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// Warn call main logger Warn function
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

// Error call main logger Error function
func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

// DPanic call main logger DPanic function
func DPanic(msg string, fields ...zapcore.Field) {
	logger.DPanic(msg, fields...)
}

// Panic call main logger Panic function
func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

// Fatal call main logger Fatal function
func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}

// Sync logger sync
func Sync() error {
	return logger.Sync()
}

// --------------------------------------------------------------------------------------------------
