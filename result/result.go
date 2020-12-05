package result

type Result struct {}

func (r *Result) Success() bool {
	return true
}

func (r *Result) FailedStage() string {
	return ""
}
