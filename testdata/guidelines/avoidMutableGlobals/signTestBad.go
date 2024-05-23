package avoidMutableGlobals

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	oldTimeNow := _timeNow
	t2 := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	_timeNow = func() time.Time {
		return t2
	}
	defer func() { _timeNow = oldTimeNow }()

	assert.Equal(t, "test"+t2.GoString(), sign("test"))
}
