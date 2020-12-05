package result

import (
	"github.com/AmitMichaely/stateful_handler/handler_context"
	"github.com/AmitMichaely/stateful_handler/state"
)

type Result interface {
	Success() bool
	FailedStage() string
	Error() error
	LastState() state.State
	FinalContext() handler_context.HandlerContext
}

type result struct {
	ctx         handler_context.HandlerContext
	success     bool
	failedStage string
	state       state.State
	err         error
}

func New(finalContext handler_context.HandlerContext, success bool, failedStage string, state state.State, err error) *result {
	return &result{
		ctx:         finalContext,
		success:     success,
		failedStage: failedStage,
		state:       state,
		err:         err,
	}
}

func (r *result) Success() bool {
	return r.success
}

func (r *result) FailedStage() string {
	return r.failedStage
}

func (r *result) Error() error {
	return r.err
}

func (r *result) LastState() state.State {
	return r.state
}

func (r *result) FinalContext() handler_context.HandlerContext {
	return r.ctx
}
