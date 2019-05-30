package mongodb

import (
	"go/build"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMongoDBConnection(t *testing.T) {
	envDir := build.Default.GOPATH + "/src/gitlab.com/go-rest-mux"
	err := godotenv.Load(envDir + ".env")
	if err != nil {
		assert.Error(t, err)
	}

	if testing.Short() {
		t.Skip("Skipping Integration Test on Short Mode")
	}

	t.Run("TestNewMongoDBSession", func(t *testing.T) {
		sess := NewMongoDBSession()
		err := sess.Ping()

		assert.NoError(t, err)
	})
}
