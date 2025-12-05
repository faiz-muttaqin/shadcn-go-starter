package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type CSVFormatter struct {
	// DisableHTMLEscape allows disabling html escaping in output
	DisableHTMLEscape bool
	ForceColors       bool
	IncludeHeader     bool
	TimestampFormat   string
	once              bool
}

func (f *CSVFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer

	if f.TimestampFormat == "" {
		f.TimestampFormat = "2006-01-02T15:04:05.000Z07:00"
	}

	// Add header row once
	if f.IncludeHeader && !f.once {
		header := []string{"level", "time", "msg", "caller"}
		for k := range entry.Data {
			header = append(header, k)
		}
		b.WriteString(strings.Join(header, ",") + "\n")
		f.once = true
	}

	// Format log fields
	csvFields := []string{
		entry.Level.String(),
		entry.Time.Format(f.TimestampFormat),
		" " + strings.ReplaceAll(entry.Message, ",", ";") + " ", // sanitize commas
	}

	// Add shortened caller (remove current working dir)
	caller := ""
	if entry.HasCaller() {
		wd, err := os.Getwd()
		if err == nil {
			cleanPath := strings.TrimPrefix(entry.Caller.File, wd+"/")
			caller = fmt.Sprintf("%s:%d", cleanPath, entry.Caller.Line)
		} else {
			caller = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		}
	}
	csvFields = append(csvFields, caller)

	// Add extra fields
	for _, v := range entry.Data {
		csvFields = append(csvFields, fmt.Sprint(v))
	}

	b.WriteString(strings.Join(csvFields, ",") + "\n")
	return b.Bytes(), nil
}
