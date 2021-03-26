package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		want string
	}{
		{
			"basic test",
			"hello golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Value(); got != tt.want {
				assert.Equal(tt.want, got)
			}
		})
	}
}
