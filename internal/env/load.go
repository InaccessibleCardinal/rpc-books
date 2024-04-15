package env

import (
	"io"
	"os"
	"strings"
)

func LoadEnv(reader io.Reader) {
	bts, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bts), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "=")

		os.Setenv(parts[0], strings.TrimSpace(parts[1]))
	}
}