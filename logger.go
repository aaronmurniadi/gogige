package gogige

import (
	"context"
	"fmt"
	"log/slog"
)

type slogAdapter struct{ z *slog.Logger }

// Slog wraps a log/slog.Logger as a Logger. Pass slog.Default() or any handler:
//
//	gogige.Open(ctx, ip, gogige.WithLogger(gogige.Slog(slog.Default())))
func Slog(z *slog.Logger) Logger { return slogAdapter{z: z} }

func (a slogAdapter) Debug(msg string, kv ...any) { a.log(slog.LevelDebug, msg, kv...) }
func (a slogAdapter) Info(msg string, kv ...any)  { a.log(slog.LevelInfo, msg, kv...) }
func (a slogAdapter) Warn(msg string, kv ...any)  { a.log(slog.LevelWarn, msg, kv...) }
func (a slogAdapter) Error(msg string, kv ...any) { a.log(slog.LevelError, msg, kv...) }

func (a slogAdapter) log(level slog.Level, msg string, kv ...any) {
	if a.z == nil {
		return
	}
	attrs := make([]slog.Attr, 0, len(kv)/2)
	for i := 0; i+1 < len(kv); i += 2 {
		key, ok := kv[i].(string)
		if !ok {
			key = fmt.Sprint(kv[i])
		}
		attrs = append(attrs, slog.Any(key, kv[i+1]))
	}
	a.z.LogAttrs(context.Background(), level, msg, attrs...)
}

var _ Logger = slogAdapter{}
