package logstash

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogstash_messageToStringMap(t *testing.T) {
	ls := Logstash{
		protocol: HTTP,
		host:     "localhost",
		port:     80,
		config: ApplicationConfig{
			AppHost: "localhost",
			AppName: "UnitTest",
		},
	}

	// given
	tests := make([]LogMessage, 0)
	tests = append(tests,
		LogMessage{
			Message: "this is a test message",
			Type:    nil,
			Level:   DebugLevel,
		},
		LogMessage{
			Message: "this is a second test message",
			Type:    ls.stringToPointer("general"),
			Level:   ErrorLevel,
		},
	)

	// when
	result := make([]map[string]interface{}, 0)
	for _, record := range tests {
		result = append(result, ls.messageToStringMap(record))
	}

	// then
	expected := make([]map[string]interface{}, 0)
	expected = append(expected,
		map[string]interface{}{
			"message":     "this is a test message",
			"type":        "",
			"level":       "debug",
			"host":        "localhost",
			"application": "UnitTest",
		},
		map[string]interface{}{
			"message":     "this is a second test message",
			"type":        "general",
			"level":       "error",
			"host":        "localhost",
			"application": "UnitTest",
		})

	require.Equal(t, expected, result)
}
