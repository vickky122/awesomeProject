package main

import (
	"testing"
	"time"
)

func TestChannelFlow(t *testing.T) {
	done := make(chan bool)

	go func() {
		main()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("deadlock detected — main did not exit within timeout")
	}
}
