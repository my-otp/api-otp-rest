package mongo_test

import (
	"github.com/my-otp/otp-rest/internal/config"
	"github.com/my-otp/otp-rest/internal/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestNewAdapter(t *testing.T) {
	mockConfig := config.DatabaseConfig{
		Database: "test_new_adapter",
		URL:      "mongodb://localhost:27017",
	}
	adapter := mongo.NewAdapter(mockConfig)
	assert.NotNil(t, adapter)
}

func TestAdapter_LoadSecret(t *testing.T) {
	mockConfig := config.DatabaseConfig{
		Database: "test_load_secret",
		URL:      "mongodb://127.0.0.1:27017",
	}

	adapter := mongo.NewAdapter(mockConfig)
	assert.NotNil(t, adapter)

	_, err := adapter.LoadSecret("test_load_secret")
	require.ErrorIs(t, err, mongo2.ErrNoDocuments)
}

func TestAdapter_SaveSecret(t *testing.T) {
	mockConfig := config.DatabaseConfig{
		Database: "test_save_secret",
		URL:      "mongodb://127.0.0.1:27017",
	}

	adapter := mongo.NewAdapter(mockConfig)
	assert.NotNil(t, adapter)

	err := adapter.SaveSecret("test_save_secret", "test_load_secret")
	require.NoError(t, err)
}
