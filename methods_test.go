package main

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetStaff(t *testing.T) {

	assert := assert.New(t)

	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(a, b, "The two words should be the same.")
}
