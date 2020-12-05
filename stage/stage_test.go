package stage

import (
	"github.com/AmitMichaely/stateful_handler/stage/test_mocks/some/path/to/stage_mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStage_Name(t *testing.T) {
	var stageInstance Stage = stage_mock.Rainbow

	assert.Equal(t, "Rainbow", stageInstance.Name())
}
