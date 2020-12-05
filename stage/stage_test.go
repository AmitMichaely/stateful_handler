package stage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStage_Name(t *testing.T) {
	var stageInstance Stage = stamStage

	assert.Equal(t, "stamStage", stageInstance.Name())
}

func stamStage() {}