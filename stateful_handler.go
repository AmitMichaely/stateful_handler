package stateful_handler

import (
	"errors"
	"fmt"
	"github.com/AmitMichaely/stateful_handler/handler_context"
	"github.com/AmitMichaely/stateful_handler/result"
	"github.com/AmitMichaely/stateful_handler/state"
)

type StatefulHandler interface {
	Run() result.Result
}

type Stage interface {
	Name() string
	Run(ctx handler_context.HandlerContext, state state.State) error
}

type statefulHandler struct {
	handlerContext handler_context.HandlerContext
	stages []Stage
	state state.State
}

func New(ctx handler_context.HandlerContext, stages ...Stage) *statefulHandler {
	return &statefulHandler{
		handlerContext: ctx,
		stages:         stages,
		state: state.New(),
	}
}

func (sh *statefulHandler) Run() result.Result {
	for _, currentStage := range sh.stages {
		if sh.handlerContext.Skip(currentStage){
			continue
		}

		if err := currentStage.Run(sh.handlerContext, sh.state); err != nil {
			return result.New(sh.handlerContext, false, currentStage.Name(), sh.state, err)
		}

		if sh.handlerContext.Stopped() {
			return result.New(sh.handlerContext, true, "", sh.state, nil)
		}

		if err := sh.handlerContext.MarkAsDone(currentStage); err != nil {
			return result.New(sh.handlerContext, false, currentStage.Name(), sh.state, errors.New(fmt.Sprintf("failed marking %s as done", currentStage.Name())))
		}
	}

	return result.New(sh.handlerContext, true, "", sh.state, nil)
}
