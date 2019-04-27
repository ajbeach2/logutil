package logutil

import "testing"

func TestGetLog(t *testing.T) {
	levels := []string{
		"debug",
		"info",
		"warn",
		"error",
	}
	for _, level := range levels {
		level := level
		t.Run(level, func(t *testing.T) {
			logger := GetLogger(level)
			logger.Error("info")
		})
	}
}
