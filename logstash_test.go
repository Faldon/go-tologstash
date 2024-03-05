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

func TestLogStash_pointerToString(t *testing.T) {
	ls := Logstash{}
	s1 := ls.pointerToString(nil)
	require.Equal(t, "", s1)
	s1 = "pointertest"
	s2 := ls.pointerToString(&s1)
	require.Equal(t, s1, s2)
}

func TestLogStash_stringToPointer(t *testing.T) {
	ls := Logstash{}
	s1 := "stringtest"
	s2 := ls.stringToPointer(s1)
	require.Equal(t, s1, *s2)
}

func TestLogHandler_Init_Errors(t *testing.T) {
	// given
	handler := LogHandler{}

	// when
	err := handler.Init(HTTP, "", 0, nil)

	// then
	errmsg := "no Logstash host provided"
	require.Equal(t, errmsg, err.Error())

	// when
	err = handler.Init(HTTP, "localhost", 0, nil)

	// then
	errmsg = "invalid destination port provided"
	require.Equal(t, errmsg, err.Error())

	// when
	err = handler.Init(HTTP, "localhost", 443, nil)

	// then
	require.Equal(t, err, nil)
	require.Equal(t, true, handler.init)
	require.Equal(t, ErrorLevel, handler.LogLevel)
}
