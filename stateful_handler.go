package stateful_handler

import (
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
		handlerContext: handler_context.CreateFrom(ctx),
		stages:         stages,
		state: state.New(),
	}
}

func (sh *statefulHandler) Run() result.Result {
	for _, currentStage := range sh.stagesToRun() {
		sh.handlerContext.UpdateCurrentStage(currentStage)

		if err := currentStage.Run(sh.handlerContext, sh.state); err != nil {
			return result.New(sh.handlerContext, false, currentStage.Name(), sh.state, err)
		}

		if sh.handlerContext.Stopped() {
			return result.New(sh.handlerContext, true, "", sh.state, nil)
		}
	}

	return result.New(sh.handlerContext, true, "", sh.state, nil)
}

func (sh *statefulHandler) stagesToRun() []Stage {
	if len(sh.handlerContext.EntryStage()) == 0 {
		return sh.stages
	}

	for i, stage := range sh.stages {
		if stage.Name() == sh.handlerContext.EntryStage() {
			return sh.stages[i:]
		}
	}

	return sh.stages
}
