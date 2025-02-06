package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(msg string, fields ...zapcore.Field) {
	zapLogger().Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	zapLogger().Debug(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	zapLogger().Warn(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	zapLogger().Fatal(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	zapLogger().Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	zapLogger().Panic(msg, fields...)
}

func DPanic(msg string, fields ...zapcore.Field) {
	zapLogger().DPanic(msg, fields...)
}

// wrap zapcore.Field
func Any(key string, value interface{}) zapcore.Field {
	return zap.Any(key, value)
}

func Binary(key string, val []byte) zapcore.Field {
	return zap.Binary(key, val)
}

func Bool(key string, val bool) zapcore.Field {
	return zap.Bool(key, val)
}

func Boolp(key string, val *bool) zapcore.Field {
	return zap.Boolp(key, val)
}

func ByteString(key string, val []byte) zapcore.Field {
	return zap.ByteString(key, val)
}

func Complex128(key string, val complex128) zapcore.Field {
	return zap.Complex128(key, val)
}
func Complex128p(key string, val *complex128) zapcore.Field {
	return zap.Complex128p(key, val)
}
func Complex64(key string, val complex64) zapcore.Field {
	return zap.Complex64(key, val)
}
func Complex64p(key string, val *complex64) zapcore.Field {
	return zap.Complex64p(key, val)
}
func Duration(key string, val time.Duration) zapcore.Field {
	return zap.Duration(key, val)
}
func Durationp(key string, val *time.Duration) zapcore.Field {
	return zap.Durationp(key, val)
}
func Float32(key string, val float32) zapcore.Field {
	return zap.Float32(key, val)
}
func Float32p(key string, val *float32) zapcore.Field {
	return zap.Float32p(key, val)
}
func Float64(key string, val float64) zapcore.Field {
	return zap.Float64(key, val)
}
func Float64p(key string, val *float64) zapcore.Field {
	return zap.Float64p(key, val)
}
func Int(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}
func Int16(key string, val int16) zapcore.Field {
	return zap.Int16(key, val)
}
func Int16p(key string, val *int16) zapcore.Field {
	return zap.Int16p(key, val)
}
func Int32(key string, val int32) zapcore.Field {
	return zap.Int32(key, val)
}
func Int32p(key string, val *int32) zapcore.Field {
	return zap.Int32p(key, val)
}
func Int64(key string, val int64) zapcore.Field {
	return zap.Int64(key, val)
}
func Int64p(key string, val *int64) zapcore.Field {
	return zap.Int64p(key, val)
}
func Int8(key string, val int8) zapcore.Field {
	return zap.Int8(key, val)
}
func Int8p(key string, val *int8) zapcore.Field {
	return zap.Int8p(key, val)
}
func Intp(key string, val *int) zapcore.Field {
	return zap.Intp(key, val)
}
func Namespace(key string) zapcore.Field {
	return zap.Namespace(key)
}
func Object(key string, val zapcore.ObjectMarshaler) zapcore.Field {
	return zap.Object(key, val)
}
func Reflect(key string, val interface{}) zapcore.Field {
	return zap.Reflect(key, val)
}
func Skip() zapcore.Field {
	return zap.Skip()
}
func Stack(key string) zapcore.Field {
	return zap.Stack(key)
}
func String(key string, val string) zapcore.Field {
	return zap.String(key, val)
}
func Stringer(key string, val fmt.Stringer) zapcore.Field {
	return zap.Stringer(key, val)
}
func Stringp(key string, val *string) zapcore.Field {
	return zap.Stringp(key, val)
}
func Strings(key string, ss []string) zapcore.Field {
	return zap.Strings(key, ss)
}
func Stringv(key string, v interface{}) zapcore.Field {
	return zap.String(key, fmt.Sprintf("%v", v))
}
func Stringf(key string, format string, a ...interface{}) zapcore.Field {
	return zap.String(key, fmt.Sprintf(format, a...))
}
func Time(key string, val time.Time) zapcore.Field {
	return zap.Time(key, val)
}
func Timep(key string, val *time.Time) zapcore.Field {
	return zap.Timep(key, val)
}
func Uint(key string, val uint) zapcore.Field {
	return zap.Uint(key, val)
}
func Uint16(key string, val uint16) zapcore.Field {
	return zap.Uint16(key, val)
}
func Uint16p(key string, val *uint16) zapcore.Field {
	return zap.Uint16p(key, val)
}
func Uint32(key string, val uint32) zapcore.Field {
	return zap.Uint32(key, val)
}
func Uint32p(key string, val *uint32) zapcore.Field {
	return zap.Uint32p(key, val)
}
func Uint64(key string, val uint64) zapcore.Field {
	return zap.Uint64(key, val)
}
func Uint64p(key string, val *uint64) zapcore.Field {
	return zap.Uint64p(key, val)
}
func Uint8(key string, val uint8) zapcore.Field {
	return zap.Uint8(key, val)
}
func Uint8p(key string, val *uint8) zapcore.Field {
	return zap.Uint8p(key, val)
}
func Uintp(key string, val *uint) zapcore.Field {
	return zap.Uintp(key, val)
}
func Uintptr(key string, val uintptr) zapcore.Field {
	return zap.Uintptr(key, val)
}
func Uintptrp(key string, val *uintptr) zapcore.Field {
	return zap.Uintptrp(key, val)
}
func Errorf(format string, a ...interface{}) zapcore.Field {
	return zap.String("error", fmt.Sprintf(format, a...))
}
func Errorv(err error) zapcore.Field {
	return Stringv("error", err)
}
