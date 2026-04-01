package code

import (
	"testing"
)

func TestGetSize(t *testing.T) {
	got, err := getSize("./testdata/testFile.txt", false, false)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(5442)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}

func TestGetSizeFlagAll(t *testing.T) {
	got, err := getSize("./testdata", true, false)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(7171)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}
func TestGetSizeFlagRecursive(t *testing.T) {
	got, err := getSize("./testdata", false, true)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(7960)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}

func TestGetSizeError(t *testing.T) {
    _, err := getSize("./test", false, true)
    
    if err == nil {
        t.Fatal("ожидали ошибку, но получили nil")
    }

    want := "lstat ./test: no such file or directory"
    if err.Error() != want {
        t.Errorf("получили %q, а хотели %q", err.Error(), want)
    }
}

func TestFormatSize(t *testing.T) {
	got := formatSize(2048, true)
	want := "2.0KB"
	if got != want {
		t.Fatalf("FormatSize(size) = %s, want %s", got, want)
	}
}
