package stage

import (
	"reflect"
	"runtime"
	"strings"
)

type Stage func()

func (s Stage) Name() string {
	pathToName := strings.Split(runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name(), ".")

	return pathToName[len(pathToName)-1]
}
