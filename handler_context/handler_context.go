package handler_context

import "errors"

type HandlerContext interface {
	Skip(runnable) bool
	MarkAsDone(runnable) error
}

type runnable interface {
	Name() string
}

type handlerContext struct {
	doneStages map[string]bool
}

func New(ctx HandlerContext) *handlerContext {
	if ctx == nil {
		return &handlerContext{}
	}

	if castedCtx, ok := ctx.(*handlerContext); ok {
		return castedCtx
	}

	return &handlerContext{}
}

func (ctx *handlerContext) Skip(stage runnable) bool {
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
