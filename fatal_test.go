package ansi

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestAnsi_Fatal(t *testing.T) {
	t.Run("test ansi.Fatal() with exit code 0", func(t *testing.T) {
		cmd := "go"
		args := []string{"run", "example/fatal.go"}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program:\nerr: %v\nout: %v", err, actual)
		}
		expectedAnsiResetCodes := []byte{27, 91, 48, 109, 99, 111, 100, 101, 58, 32, 48, 27, 91, 48, 109}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%s' %v\n"+
				"expected: '%v'\n", string(actual), actual, expectedAnsiResetCodes)
		}
	})
	t.Run("test ansi.Fatal() with exit code 1", func(t *testing.T) {
		code := 1
		cmd := "go"
		args := []string{"run", "example/fatal.go", "-code", fmt.Sprintf("%d", code)}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil && err.Error() != fmt.Sprintf("exit status %d", code) {
			t.Fatalf("Failed to run test program:\nerr: %v\nout: %v", err, actual)
		}
		expectedAnsiResetCodes := []byte{27, 91, 48, 109, 99, 111, 100, 101, 58, 32, 49, 27, 91, 48, 109}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%s' %v\n"+
				"expected: '%v'\n", string(actual), actual, expectedAnsiResetCodes)
		}
	})
	t.Run("test ansi.Fatal() with exit code 2", func(t *testing.T) {
		code := 2
		cmd := "go"
		args := []string{"run", "example/fatal.go", "-code", fmt.Sprintf("%d", code)}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil && err.Error() != fmt.Sprintf("exit status %d", 1) {
			t.Fatalf("Failed to run test program:\nerr: %v\nout: %v", err, actual)
		}
		expectedAnsiResetCodes := []byte{27, 91, 48, 109, 99, 111, 100, 101, 58, 32, 50, 27, 91, 48, 109}
		if !bytes.Equal(actual, expectedAnsiResetCodes) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%s' %v\n"+
				"expected: '%v'\n", string(actual), actual, expectedAnsiResetCodes)
		}
	})
}
