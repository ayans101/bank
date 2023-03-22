package api

import (
	"os"
	"testing"
	"time"

	db "github.com/ayans101/bank/db/sqlc"
	"github.com/ayans101/bank/util"
	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	if err != nil {
		return nil
	}

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
