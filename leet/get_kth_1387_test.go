package leet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getKth(t *testing.T) {
	lo := 12
	hi := 15
	k := 2

	num := getKth(lo, hi, k)
	assert.Equal(t, num, 13)
}
