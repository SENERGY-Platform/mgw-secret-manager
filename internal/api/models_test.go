package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryError(t *testing.T) {
	err := MissingQueryError{Parameter: "test"}
	assert.Equal(t, err.Error(), "Query parameter test is missing")
}
