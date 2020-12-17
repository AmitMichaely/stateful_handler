# Stateful Handler

Stateful Handler is an implementation of the Railway pattern.
It allows you to create complex-multistage-handlers that will be able to re-run from the last successful stage.
It also provides an interface (State) that allows you to pass arguments between each stage-function.

## Requirements
- Go 1.13 or higher.

## Usage

```go
func stage1(handlerContext, state) error {
  // stage1 logic
}

func stage2(handlerContext, state) error {
// stage2 logic
}

func stage3(handlerContext, state) error {
// stage3 logic
}

func main() {
    previousContext := readPreviousContext() // read the previous context from where you previously saved it (`nil` is allowed)
    multiStageHandler := New(previousContext, stage1, stage2, stage3)

    result := multiStageHandler.Run()
    if !result.Success() {
      // handle result.Error(), save your result.LastState() and result.FinalContext() etc..
    }
}
```
