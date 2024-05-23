package avoidMutableGlobals

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// sign_test.go

func TestSigner(t *testing.T) {
	s := newSigner()
	t2 := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	s.now = func() time.Time {
		return t2
	}

	assert.Equal(t, "test"+t2.GoString(), s.Sign("test"))
}
