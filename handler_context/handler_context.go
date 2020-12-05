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

func CreateFrom(previousCtx HandlerContext) *handlerContext {
	if previousCtx == nil {
		return &handlerContext{}
	}

	if castedCtx, ok := previousCtx.(*handlerContext); ok {
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
