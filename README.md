# Logstash client for Go
go-tologstash is an easy-to-use, with next to no configuration Go module, that
allows you to send various log messages to your Logstash server.

## Installation
```bash
go get github.com/Faldon/go-tologstash
```

## Quickstart
```go
import (
	logstash "github.com/Faldon/go-tologstash"
)

func ExampleLog() {
	var Log logstash.LogHandler
	if err := Log.Init(logstash.HTTP, "logstash.example.com", 8080, nil); err != nil {
		panic(err)
	}
	// Write an info log message
	Log.Info("this is a log message")
	// Write an error message
	Log.Error("help me rhonda")
}
```
