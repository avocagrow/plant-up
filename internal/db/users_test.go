package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var userIds = []string{}

func TestCreateUpdateDeleteUsers(t *testing.T) {
	ctx := context.Background()

	testInst, err := Configure(
		WithNewConnection("postgresql://postgres:postgres@localhost:5432/postgres"),
	)
	require.NoError(t, err)

	// create a user
}
