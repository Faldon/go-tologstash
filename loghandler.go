package logstash

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type LogHandler struct {
	Logger   Logstash
	LogLevel Level
	init     bool
	pid      string
}

func (lsh *LogHandler) Init(protocol TransportProtocol, host string, port int, config *ApplicationConfig) error {
	logstash := Logstash{}
	if host == "" {
		return errors.New("no Logstash host provided")
	}
	logstash.host = host
	if port == 0 {
		return errors.New("invalid destination port provided")
	}
	logstash.port = port
	logstash.protocol = protocol

	logLevel := ErrorLevel
	hostname, _ := os.Hostname()
	logstash.config = applicationConfigToPointer(config, ApplicationConfig{
		AppHost:     hostname,
		AppName:     os.Args[0],
		Version:     "",
		Extension:   "",
		Environment: nil,
		LogLevel:    &logLevel,
	})

	lsh.Logger = logstash
	lsh.LogLevel = ErrorLevel
	if config.LogLevel != nil {
		lsh.LogLevel = *config.LogLevel
	}
	lsh.init = true
	lsh.pid = strconv.Itoa(os.Getppid())
	return nil
}

func (lsh *LogHandler) Info(msg string) {
	if lsh.LogLevel < InfoLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   InfoLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (lsh *LogHandler) Error(msg string) {
	if lsh.LogLevel < ErrorLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   ErrorLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (lsh *LogHandler) Debug(msg string) {
	if lsh.LogLevel < DebugLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   DebugLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (lsh *LogHandler) Fatal(msg string) {
	if lsh.LogLevel < FatalLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   FatalLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (lsh *LogHandler) Panic(msg string) {
	if lsh.LogLevel < PanicLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   PanicLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (lsh *LogHandler) Warn(msg string) {
	if lsh.LogLevel < WarnLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   WarnLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (lsh *LogHandler) Trace(msg string) {
	if lsh.LogLevel < TraceLevel {
		return
	}
	record := LogMessage{
		Message: msg,
		Level:   TraceLevel,
	}
	err := lsh.Logger.Write(&lsh.pid, record)
	if err != nil {
		fmt.Println(err.Error())
	}
}
