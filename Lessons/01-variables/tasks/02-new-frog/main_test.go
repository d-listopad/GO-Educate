package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// captureOutput запускает f и возвращает всё, что f напечатала на экран.
func captureOutput(t *testing.T, f func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	saved := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = saved }()

	f()
	w.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("чтение вывода: %v", err)
	}
	return buf.String()
}

func TestNewFrog(t *testing.T) {
	got := captureOutput(t, main)
	want := "Имя: \"\", мух: 0\nИмя: \"Квакуша\", мух: 2\n"

	if got != want {
		t.Errorf("программа напечатала %q, ожидалось %q", got, want)
	}
}
