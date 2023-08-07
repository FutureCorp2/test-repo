package main_test

import (
	"testing"

	main "github.com/FutureCorp2/test-repo"

	"github.com/stretchr/testify/assert"
)

func TestCoverFunc(t *testing.T) {
	t.Run("Just Check", func(t *testing.T) {
		assert.Equal(t, main.CoverFunc(), 15)
	})
}
