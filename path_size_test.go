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
	got, err := getSize("./testdata", false, true)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(7171)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}

func TestGetSizeFlagNotAll(t *testing.T) {
	got, err := getSize("./testdata", false, false)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(5442)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}

func TestGetSizeFlagRecursive(t *testing.T) {
	got, err := getSize("./testdata", true, false)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(7960)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}

func TestGetSizeError(t *testing.T) {
    _, err := getSize("./test", true, false)
    
    if err == nil {
        t.Fatal("ожидали ошибку, но получили nil")
    }

    want := "lstat ./test: no such file or directory"
    if err.Error() != want {
        t.Errorf("получили %q, а хотели %q", err.Error(), want)
    }
}

func TestGetSizeEmptyDir(t *testing.T) {
  	got, err := getSize("./testdata/testEmpty", true, false)
	if err != nil {
		t.Fatalf("GetSize(file) error = %v", err)
	}

	want := int64(0)
	if got != want {
		t.Fatalf("GetSize(file) = %d, want %d", got, want)
	}
}

func TestFormatSizeKB(t *testing.T) {
	got := formatSize(2048, true)
	want := "2.0KB"
	if got != want {
		t.Fatalf("FormatSize(size) = %s, want %s", got, want)
	}
}

func TestFormatSizeMB(t *testing.T) {
	got := formatSize(2097152, true)
	want := "2.0MB"
	if got != want {
		t.Fatalf("FormatSize(size) = %s, want %s", got, want)
	}
}

func TestFormatSizeGB(t *testing.T) {
	got := formatSize(2147483648, true)
	want := "2.0GB"
	if got != want {
		t.Fatalf("FormatSize(size) = %s, want %s", got, want)
	}
}