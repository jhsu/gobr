package gobr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBranches(t *testing.T) {
	assert.Equal(t, []string{"master"}, Branches())
}
