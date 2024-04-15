package env

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	reader := strings.NewReader("KEY1=Value1\nKEY2=Value2")
	
	LoadEnv(reader)

	if val := os.Getenv("KEY1"); val != "Value1" {
		t.Errorf("Expected KEY1 to be 'Value1', got '%s'", val)
	}

	if val := os.Getenv("KEY2"); val != "Value2" {
		t.Errorf("Expected KEY2 to be 'Value2', got '%s'", val)
	}
}

type errReader struct {}

func (e errReader) Read([]byte) (int, error) {
	return 0, errors.New("lol")
}

func TestLoadEnvErr(t *testing.T) {
	reader := errReader{}
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected non-nil error")
		}
		if r.(error).Error() != "lol" {
			t.Errorf("got %s but expected lol", r.(error).Error())
		}
	}()
	LoadEnv(reader)
}