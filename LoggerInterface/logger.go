package main

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
	fmt.Println("[CONSOLE]", msg)
}

type FileLogger struct {
	File *os.File
}

func (f FileLogger) Log(msg string) {
	fmt.Fprintln(f.File, "[FILE]", msg)
}

type TimestampLogger struct {
	Wrapped Logger
}

func (t TimestampLogger) Log(msg string) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	t.Wrapped.Log("[" + ts + "] " + msg)
}

type Server struct {
	Logger
}

func NewServer(logger Logger) *Server {
	return &Server{Logger: logger}
}

func (s *Server) Start() {
	s.Log("Server started")
}

func main() {
	consoleLogger := TimestampLogger{Wrapped: ConsoleLogger{}}
	consoleServer := NewServer(consoleLogger)
	consoleServer.Start()

	f, _ := os.Create("app.log")
	fileLogger := TimestampLogger{Wrapped: FileLogger{File: f}}
	fileServer := NewServer(fileLogger)
	fileServer.Start()
}
