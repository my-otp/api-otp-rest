package config_test

import (
	"github.com/my-otp/otp-rest/internal/config"
	"github.com/restartfu/gophig"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const mockConfigContent = "http:\n    address: 127.0.0.1\n    port: 8080\ndatabase:\n    database: my-otp\n    url: mongodb://localhost:27017"

var yaml = gophig.YAMLMarshaler{}

func TestLoadConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		root := "./../../assets/tmp"
		filePath := root + "/config.yaml"

		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		require.NoError(t, err)

		err = os.WriteFile(filePath, []byte(mockConfigContent), os.ModePerm)
		require.NoError(t, err)

		conf, err := gophig.LoadConf[config.Config](filePath, yaml)
		require.NoError(t, err)

		err = os.RemoveAll(filepath.Dir(filePath))
		require.NoError(t, err)

		buf, err := yaml.Marshal(conf)
		require.NoError(t, err)
		require.Equal(t, mockConfigContent, strings.TrimSpace(string(buf)))
	})
}
