package logging_for_otus

import (
	"fmt"
	"io"
	"time"
)

type OtusEvent interface {
	LogMessage() string
}

func NewHwAccepted(id int, grade int) *HwAccepted {
	return &HwAccepted{Id: id, Grade: grade}
}

type HwAccepted struct {
	Id    int
	Grade int
}

func (a HwAccepted) LogMessage() string {
	return fmt.Sprintf(
		"%s accepted %d %d",
		time.Now().Format("2006-01-02"),
		a.Id,
		a.Grade,
	)
}

func NewHwSubmitted(id int, code string, comment string) *HwSubmitted {
	return &HwSubmitted{Id: id, Code: code, Comment: comment}
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (s HwSubmitted) LogMessage() string {
	return fmt.Sprintf(
		"%s submitted %d %s",
		time.Now().Format("2006-01-02"),
		s.Id,
		s.Comment,
	)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	w.Write([]byte(e.LogMessage()))
}
