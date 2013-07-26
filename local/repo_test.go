package gobr

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestBranches(t *testing.T) {
  assert.Equal(t, []string{"master"}, Branches())
}
