package logging_for_otus

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestHwAccepted_LogMessage(t *testing.T) {
	hwAccepted := NewHwAccepted(1, 25)
	message := hwAccepted.LogMessage()

	expected := fmt.Sprintf("%s accepted %d %d", time.Now().Format("2006-01-02"), 1, 25)
	if message != expected {
		t.Errorf("messages must match, %s - %s", message, expected)
	}

	expected = fmt.Sprintf("%s accepted %d %d", time.Now().Format("2006-01-02"), 3, 12)
	if message == expected {
		t.Errorf("messages not must match, %s - %s", message, expected)
	}
}

func TestHwSubmitted_LogMessage(t *testing.T) {
	comment := "please take a look at my homework"
	hwAccepted := NewHwSubmitted(1, "", comment)
	message := hwAccepted.LogMessage()

	expected := fmt.Sprintf("%s submitted %d %s", time.Now().Format("2006-01-02"), 1, comment)
	if message != expected {
		t.Errorf("messages must match, %s - %s", message, expected)
	}

	expected = fmt.Sprintf("%s submitted %d %s", time.Now().Format("2006-01-02"), 5, "bad comment")
	if message == expected {
		t.Errorf("messages not must match, %s - %s", message, expected)
	}
}

func TestLogOtusEvent(t *testing.T) {
	var b bytes.Buffer

	hwAccepted := NewHwAccepted(1, 25)
	LogOtusEvent(hwAccepted, &b)

	message := b.String()
	expected := fmt.Sprintf("%s accepted %d %d", time.Now().Format("2006-01-02"), 1, 25)
	if message != expected {
		t.Errorf("messages must match, %s - %s", message, expected)
	}
}
