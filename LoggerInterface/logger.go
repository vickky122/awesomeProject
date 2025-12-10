package main

import (
	"fmt"
	"os"
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
	consoleServer := NewServer(ConsoleLogger{})
	consoleServer.Start()

	f, _ := os.Create("app.log")
	fileServer := NewServer(FileLogger{File: f})
	fileServer.Start()
}
