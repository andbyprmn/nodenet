package logging

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Logger struct {
	logFilePath string
}

func NewLogger(logFilePath string) *Logger {
	return &Logger{
		logFilePath: logFilePath,
	}
}

// LogEvent mencatat pesan ke dalam file log dengan timestamp.
func (l *Logger) LogEvent(event string) error {
	file, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	timestamp := time.Now().Format(time.RFC3339)
	logLine := fmt.Sprintf("[%s] %s\n", timestamp, event)

	if _, err := file.WriteString(logLine); err != nil {
		return fmt.Errorf("failed to write log: %v", err)
	}

	return nil
}

// GetLastLogs mengambil n baris terakhir dari file log.
func (l *Logger) GetLastLogs(n int) []string {
	file, err := os.Open(l.logFilePath)
	if err != nil {
		return []string{fmt.Sprintf("failed to open log file: %v", err)}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < n {
		return lines
	}

	return lines[len(lines)-n:]
}
