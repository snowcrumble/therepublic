package therepublic_test

import (
	"testing"

	"github.com/snowcrumble/therepublic"
	"github.com/stretchr/testify/assert"
)

// 测试医生是否用心治病
func TestBestDoctor(t *testing.T) {
	bestDoctor := therepublic.NewBestDoctor()
	assert.Equal(t, bestDoctor.Heart, bestDoctor.CureAbility())
}
