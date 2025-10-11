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
	LogglyUrl   string // URL from Loggly https://logs-01.loggly.com/inputs/%s/tag/%s/
	LogglyToken string // Token from Loggly
	LogglyTag   string // Nama tag example: "user-service"
	Environment string // dev, staging, prod
	AllLogLevel bool   // optional for show all level log ignoring the nvirontment
}

var (
	cfg         Config
	logLevel    zapcore.Level
	once        sync.Once
	initialized bool
)

// ==================================================
// INITIALIZATION
// ==================================================

func Init(c Config) {
	once.Do(func() {
		cfg = c
		logLevel = detectLevel(c.Environment)
		initialized = true
		fmt.Println("Logger initialized with env:", c.Environment)
	})
}

func detectLevel(env string) zapcore.Level {
	switch strings.ToLower(env) {
	case "dev", "development":
		return TraceLevel
	case "staging":
		return zapcore.InfoLevel
	case "prod", "production":
		return zapcore.WarnLevel
	default:
		return TraceLevel
	}
}

// ==================================================
// COLOR CONSTANTS
// ==================================================

var (
	colorReset   = "\033[0m"
	colorGreen   = "\033[32m"
	colorRed     = "\033[31m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[36m"
	colorMagenta = "\033[35m"
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
	case zapcore.Level(-2): // TRACE custom level
		levelColor = colorMagenta
	}

	var levelTag string
	switch entry.Level {
	case zapcore.Level(-2): // TRACE custom level
		levelTag = strings.ToUpper("TRACE")
	default:
		levelTag = strings.ToUpper(entry.Level.String())
	}

	sb.WriteString("───────────────────────────────────────────────────────────────\n")
	sb.WriteString(fmt.Sprintf("%s[%s]%s %s\n",
		levelColor,
		levelTag,
		colorReset,
		entry.Time.Format(time.RFC3339),
	))

	if rid, ok := fields["request_id"]; ok {
		sb.WriteString(fmt.Sprintf("RequestID: %v\n", rid))
	}

	sb.WriteString(fmt.Sprintf("Level: %s\n", levelTag))
	sb.WriteString(fmt.Sprintf("Message: %s\n", entry.Message))

	if len(fields) > 0 {
		jsonBytes, _ := json.MarshalIndent(fields, "", "  ")
		sb.WriteString(fmt.Sprintf("Data: %s\n", string(jsonBytes)))
	}

	if errVal, ok := fields["error"]; ok && errVal != nil {
		sb.WriteString(fmt.Sprintf("Error: %v\n", errVal))
	}

	sb.WriteString("───────────────────────────────────────────────────────────────\n")
	return sb.String()
}

// ==================================================
// CORE LOG FUNCTION
// ==================================================

func log(level zapcore.Level, msg string, fields map[string]any) {
	if !initialized {
		Init(Config{Environment: "dev", AllLogLevel: true})
	}

	// Filter by environment level
	if !cfg.AllLogLevel && level < logLevel {
		return
	}

	entry := zapcore.Entry{
		Level:   level,
		Time:    time.Now(),
		Message: msg,
	}

	fmt.Fprint(os.Stdout, format(entry, fields))

	// Kirim ke Loggly hanya jika environment bukan dev
	if cfg.LogglyToken != "" && cfg.Environment != "dev" {
		go sendToLoggly(entry, fields)
	}
}

// ==================================================
// PRIVATE: LOGGLY TRANSPORTER
// ==================================================

func sendToLoggly(entry zapcore.Entry, fields map[string]any) {
	url := fmt.Sprintf(cfg.LogglyUrl,
		cfg.LogglyToken,
		cfg.LogglyTag,
	)

	payload := map[string]any{
		"timestamp": entry.Time.Format(time.RFC3339),
		"level":     entry.Level.String(),
		"message":   entry.Message,
		"fields":    fields,
		"hostname":  getHostname(),
		"env":       cfg.Environment,
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

func Info(msg string, fields map[string]any) {
	log(zapcore.InfoLevel, msg, fields)
}

func Warn(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(zapcore.WarnLevel, msg, fields)
}

func Error(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(zapcore.ErrorLevel, msg, fields)
}

func Debug(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(zapcore.DebugLevel, msg, fields)
}

// ==================================================
// TRACE LEVEL (custom)
// ==================================================

const TraceLevel zapcore.Level = -2

func Trace(msg string, fields map[string]any, err ...error) {
	if fields == nil {
		fields = make(map[string]any)
	}
	if len(err) > 0 && err[0] != nil {
		fields["error"] = err[0].Error()
	}
	log(TraceLevel, msg, fields)
}
