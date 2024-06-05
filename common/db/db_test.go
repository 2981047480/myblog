package db_test

import (
	"context"
	"testing"
	"vblog/common/db"
)

func TestGetConn(t *testing.T) {
	var c *context.Context
	database := db.CreateNewDBimpl()
	database.GetConn(c)
}
