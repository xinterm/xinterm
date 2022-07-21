package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// SimpleFormatter is a simple formatter for logrus
type simpleFormatter struct {
	TimeFormat string
	Utc        bool
}

// Format the log string
func (f simpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	var logTime time.Time
	if f.Utc {
		logTime = entry.Time.UTC()
	} else {
		logTime = entry.Time
	}
	timeText := logTime.Format(f.TimeFormat)

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	caller := ""
	if entry.HasCaller() {
		caller = fmt.Sprintf("%s:%d %s()",
			entry.Caller.File, entry.Caller.Line, entry.Caller.Function)
	}

	fmt.Fprintf(b, "[%s](%s)%s %s", timeText, levelText, caller, entry.Message)

	b.WriteByte('\n')

	return b.Bytes(), nil
}
