package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockScanlnYes(a ...interface{}) (int, error) {
	b, _ := a[0].(*string)
	*b = "yes"
	return 0, nil
}

func TestCheckError(t *testing.T) {
	CheckError(nil)

	e := errors.New("A fake error")
	assert.Panics(t, func() { CheckError(e) })
}

func TestUserPromt(t *testing.T) {
	dstFile := "ad-hoc"
	UserPromt(mockScanlnYes, dstFile)
}
