package logstash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransportProtocol string

const (
	HTTP  TransportProtocol = "http"
	HTTPS TransportProtocol = "https"
)

type Logstash struct {
	protocol TransportProtocol
	host     string
	port     int
	config   ApplicationConfig
}

type LogMessage struct {
	Message string
	Type    *string
	Level   Level
}

type ApplicationConfig struct {
	AppHost     string
	AppName     string
	Version     string
	Extension   string
	Environment *string
	LogLevel    *Level
}

func (ls *Logstash) Write(pid *string, record LogMessage) error {
	message := ls.messageToStringMap(record)
	if pid != nil {
		message["pid"] = pid
	}
	output, err := json.Marshal(message)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s://%s:%d", ls.protocol, ls.host, ls.port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(output))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (ls *Logstash) pointerToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func (ls *Logstash) stringToPointer(value string) *string {
	return &value
}

func (ls *Logstash) messageToStringMap(record LogMessage) map[string]interface{} {
	message := make(map[string]interface{})
	message = map[string]interface{}{
		"message":     record.Message,
		"type":        ls.pointerToString(record.Type),
		"level":       Level.String(record.Level),
		"host":        ls.config.AppHost,
		"application": ls.config.AppName,
	}
	if ls.config.Environment != nil {
		message["version"] = ls.config.Environment
	}
	return message
}

func applicationConfigToPointer(config *ApplicationConfig, defaultConfig ApplicationConfig) ApplicationConfig {
	if config != nil {
		return *config
	}
	return defaultConfig
}
