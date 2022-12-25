package mok

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	testCases := []struct {
		name         string
		config       Config
		configReader func(t *testing.T, conf Config) io.Reader
		expErrorMsg  string
	}{
		{
			name: "ok",
			config: Config{
				Servers: []ServerConfig{testServerConfig(7000, ServerTypeHTTP)},
			},
			configReader: func(t *testing.T, conf Config) io.Reader {
				content, err := conf.DumpConfig()
				require.NoError(t, err)
				return bytes.NewReader(content)
			},
		},
		{
			name:   "error: not yaml content",
			config: Config{},
			configReader: func(t *testing.T, conf Config) io.Reader {
				return bytes.NewReader([]byte("qwerty"))
			},
			expErrorMsg: "failed to read config yaml file",
		},
		{
			name:   "error: invalid schema",
			config: Config{},
			configReader: func(t *testing.T, conf Config) io.Reader {
				return bytes.NewReader([]byte("servers: asd"))
			},
			expErrorMsg: "failed to decode config yaml file",
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			expConfig := tc.config

			result, err := ReadConfig(tc.configReader(t, expConfig))
			if len(tc.expErrorMsg) > 0 {
				require.Error(t, err)
				require.Nil(t, result)
				assert.ErrorContains(t, err, tc.expErrorMsg)
			} else {
				require.NotNil(t, result)
				require.NoError(t, err)
				assert.EqualValues(t, expConfig, *result)
			}
		})
	}
}

func testServerConfig(port uint16, _type ServerType) ServerConfig {
	serverType := ServerTypeToString[_type]

	return ServerConfig{
		Name:      fmt.Sprintf("%s-%d", serverType, port),
		IP:        "0.0.0.0",
		Port:      port,
		Type:      serverType,
		IsStopped: false,
		Routes: []RouteConfig{
			{
				URL:        "/some-url",
				Type:       "req-resp",
				IsDisabled: false,
				Expectations: []ExpectationConfig{
					{ID: "1"},
					{ID: "2"},
					{ID: "3"},
				},
			},
		},
	}
}
