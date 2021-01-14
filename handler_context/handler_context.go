package handler_context

type HandlerContext interface {
	UpdateCurrentStage(runnable)
	Stop(string)
	Stopped() bool
	StopReason() string
	EntryStage() string
}

type runnable interface {
	Name() string
}

type handlerContext struct {
	entryStage string
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

func (ctx *handlerContext) UpdateCurrentStage(stage runnable) {
	ctx.entryStage = stage.Name()
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

func (ctx *handlerContext) EntryStage() string {
	return ctx.entryStage
}
