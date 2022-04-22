package koan

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"os"
	"strings"
	"testing"
)

var tl *Logger

func TestLogger_Info(t *testing.T) {
	want := INFOCOLOR + "INFO: Test info message" + color.Reset

	tl = createTestLogger()
	got := captureOutput(func() {
		tl.Info("Test info message")
	})
	assertEqual(t, got, want)
}

func TestLogger_Warn(t *testing.T) {
	want := WARNCOLOR + "WARN: Test warning message" + color.Reset

	tl = createTestLogger()
	got := captureOutput(func() {
		tl.Warn("Test warning message")
	})
	assertEqual(t, got, want)
}

func TestLogger_Error(t *testing.T) {
	err := errors.New("test error message")
	msg := fmt.Sprintf("%s (%v)", "ERROR: Test message", err)
	want := ERRCOLOR + msg + color.Reset

	tl = createTestLogger()
	got := captureOutput(func() {
		tl.Error("Test message", err)
	})
	assertEqual(t, got, want)
}

// test helpers
func createTestLogger() *Logger {
	return &Logger{}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)

	output := strings.Split(buf.String(), " ")

	// remove timestamp & newline
	return strings.TrimSuffix(strings.Join(output[2:], " "), "\n")
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		fmt.Printf("want: %q, got: %q", want, got)
		t.Errorf("Fail wanted %s got %s", want, got)
	}
}
