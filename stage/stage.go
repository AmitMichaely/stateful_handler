package stage

import (
	"github.com/AmitMichaely/stateful_handler/handler_context"
	"github.com/AmitMichaely/stateful_handler/state"
	"reflect"
	"runtime"
	"strings"
)

type Stage func(ctx handler_context.HandlerContext, state state.State) error

func (s Stage) Name() string {
	pathToName := strings.Split(runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name(), ".")

	return pathToName[len(pathToName)-1]
}

func (s Stage) Run(ctx handler_context.HandlerContext, state state.State) error {
	return s(ctx, state)
}
