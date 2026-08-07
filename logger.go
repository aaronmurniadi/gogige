package gogige

import (
	"fmt"

	"github.com/rs/zerolog"
)

type zerologAdapter struct{ z zerolog.Logger }

// Zerolog wraps a zerolog.Logger as a Logger.
func Zerolog(z zerolog.Logger) Logger { return zerologAdapter{z: z} }

func (a zerologAdapter) Debug(msg string, kv ...any) { a.event(a.z.Debug(), msg, kv...) }
func (a zerologAdapter) Info(msg string, kv ...any)  { a.event(a.z.Info(), msg, kv...) }
func (a zerologAdapter) Warn(msg string, kv ...any)  { a.event(a.z.Warn(), msg, kv...) }
func (a zerologAdapter) Error(msg string, kv ...any) { a.event(a.z.Error(), msg, kv...) }

func (a zerologAdapter) event(e *zerolog.Event, msg string, kv ...any) {
	for i := 0; i+1 < len(kv); i += 2 {
		key := fmt.Sprint(kv[i])
		e = e.Interface(key, kv[i+1])
	}
	e.Msg(msg)
}

var _ Logger = zerologAdapter{}
