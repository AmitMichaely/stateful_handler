package handler_context

import "errors"

type HandlerContext interface {
	ShouldSkip(runnable) bool
	MarkAsDone(runnable) error
	Stop(string)
	Stopped() bool
	StopReason() string
}

type runnable interface {
	Name() string
}

type handlerContext struct {
	doneStages map[string]bool
	stopped bool
	stopReason string
}

func CreateFrom(previousCtx HandlerContext) *handlerContext {
	if previousCtx == nil {
		return &handlerContext{}
	}

	if castedCtx, ok := previousCtx.(*handlerContext); ok {
		return castedCtx
	}

	return &handlerContext{}
}

func (ctx *handlerContext) ShouldSkip(stage runnable) bool {
	_, stageDone := ctx.doneStages[stage.Name()]
	return stageDone
}

func (ctx *handlerContext) MarkAsDone(stage runnable) error {
	if _, doneStage := ctx.doneStages[stage.Name()]; doneStage {
		return errors.New("stage was already marked as done")
	}

	ctx.doneStages[stage.Name()] = true
	return nil
}

func (ctx *handlerContext) Stop(reason string) {
	ctx.stopped = true
	ctx.stopReason = reason
}

func (ctx *handlerContext) Stopped() bool {
	return ctx.stopped
}

func (ctx *handlerContext) StopReason() string {
	return ctx.stopReason
}
