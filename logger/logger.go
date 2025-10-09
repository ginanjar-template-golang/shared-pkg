package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

// ==================================================
// CONFIGURATION
// ==================================================

type Config struct {
	LogglyToken string // Token dari Loggly
	LogglyTag   string // Nama tag misalnya: "user-service"
	Enabled     bool   // true = kirim ke Loggly, false = hanya tampil di console
}

var (
	cfg  Config
	once sync.Once
)

// ==================================================
// INITIALIZATION
// ==================================================

func Init(c Config) {
	once.Do(func() {
		cfg = c
		fmt.Println("Logger initialized")
	})
}

// ==================================================
// COLOR CONSTANTS
// ==================================================

var (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[36m"
)

// ==================================================
// FORMATTER (CLEAN CONSOLE OUTPUT)
// ==================================================

func format(entry zapcore.Entry, fields map[string]any) string {
	var sb strings.Builder

	levelColor := colorGreen
	switch entry.Level {
	case zapcore.ErrorLevel:
		levelColor = colorRed
	case zapcore.WarnLevel:
		levelColor = colorYellow
	case zapcore.DebugLevel:
		levelColor = colorBlue
	}

	// HEADER
	sb.WriteString("───────────────────────────────────────────────────────────────\n")
	sb.WriteString(fmt.Sprintf("%s[%s]%s %s\n",
		levelColor,
		strings.ToUpper(entry.Level.String()),
		colorReset,
		entry.Time.Format(time.RFC3339),
	))
	// REQUEST ID
	if rid, ok := fields["request_id"]; ok {
		sb.WriteString(fmt.Sprintf("RequestID: %v\n", rid))
	}

	sb.WriteString(fmt.Sprintf("Message: %s\n", entry.Message))

	// OTHER DATA
	if len(fields) > 0 {
		jsonBytes, _ := json.MarshalIndent(fields, "", "  ")
		sb.WriteString(fmt.Sprintf("Data: %s\n", string(jsonBytes)))
	}

	// ORIGINAL ERROR (if any, no color)
	if errVal, ok := fields["error"]; ok && errVal != nil {
		sb.WriteString(fmt.Sprintf("Error: %v\n", errVal))
		delete(fields, "error")
	}

	sb.WriteString("───────────────────────────────────────────────────────────────\n")
	return sb.String()
}

// ==================================================
// CORE LOG FUNCTION
// ==================================================

func log(level zapcore.Level, msg string, fields map[string]any) {
	entry := zapcore.Entry{
		Level:   level,
		Time:    time.Now(),
		Message: msg,
	}

	fmt.Fprint(os.Stdout, format(entry, fields))

	if cfg.Enabled && cfg.LogglyToken != "" {
		go sendToLoggly(entry, fields)
	}
}

// ==================================================
// PRIVATE: LOGGLY TRANSPORTER
// ==================================================

func sendToLoggly(entry zapcore.Entry, fields map[string]any) {
	url := fmt.Sprintf("https://logs-01.loggly.com/inputs/%s/tag/%s/",
		cfg.LogglyToken,
		cfg.LogglyTag,
	)

	payload := map[string]any{
		"timestamp": entry.Time.Format(time.RFC3339),
		"level":     entry.Level.String(),
		"message":   entry.Message,
		"fields":    fields,
		"hostname":  getHostname(),
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 3 * time.Second}
	client.Do(req)
}

func getHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "unknown-host"
	}
	return host
}

// ==================================================
// PUBLIC LOGGER FUNCTIONS
// ==================================================

// Info — log informasi umum
func Info(msg string, fields map[string]any) {
	log(zapcore.InfoLevel, msg, fields)
}

// Warn — log warning, bisa kirim error opsional
func Warn(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(zapcore.WarnLevel, msg, fields)
}

// Error — log error, bisa kirim error langsung atau di fields
func Error(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(zapcore.ErrorLevel, msg, fields)
}

// Debug — log untuk debugging (bisa tampil error juga)
func Debug(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(zapcore.DebugLevel, msg, fields)
}
