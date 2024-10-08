package http_test

import (
	"github.com/my-otp/otp-rest/internal/http"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHandler(t *testing.T) {
	newHandler := http.NewHandler()
	require.NotNil(t, newHandler)
}
