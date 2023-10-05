package factory

import (
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

func MustGetenv(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		slog.Error("failed to get enviornment", "key", key)
		os.Exit(1)
	}
	return val
}
