package code

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all )
	if err != nil {
		return "", fmt.Errorf("error processing %s: %v", path, err) 
	}
	return formatSize(size, human), nil
}

/** Функция вычисляющая размер файла или верхнего уровня директории. */
func getSize(path string, recursive, all bool) (int64, error) {
    /** Итоговый размер файла. */
    var size int64

    fileInfo, err := os.Lstat(path)
    if err != nil {
        return 0, err
    }

    if !fileInfo.IsDir() {
        return fileInfo.Size(), nil
    }

    files, err := os.ReadDir(path)
    if err != nil {
        return 0, err
    }

    for _, file := range files {
        fileName := file.Name()
        isHidden := strings.HasPrefix(fileName, ".")

        if !all && isHidden {
            continue
        }

        fullpath := filepath.Join(path, file.Name())

        if !file.IsDir() {
            fileInfo, err := os.Lstat(fullpath)
            if err != nil {
                return 0, err
            }
            size += fileInfo.Size()
            continue
        }

        if recursive {
            dirSize, err := getSize(fullpath, recursive, all)
            if err != nil {
                return 0, err
            }
            size += dirSize
        }
    }

    return size, nil
}


/** Функция форматирования размера файла. */
func formatSize(size int64, human bool) string {

    currentSize := float64(size)

    const step = 1024.0

    const (
        KB = step
        MB = step * KB
        GB = step * MB
        TB = step * GB
        PB = step * TB
        EB = step * PB
    )

    if !human {
        return fmt.Sprintf("%dB", size)
    }

   
    switch {
    case currentSize < KB:
        return fmt.Sprintf("%d%s", int64(currentSize), "B")
    case currentSize < MB:
        return fmt.Sprintf("%.1f%s", currentSize/KB, "KB")
    case currentSize < GB:
        return fmt.Sprintf("%.1f%s", currentSize/MB, "MB")
    case currentSize < TB:
        return fmt.Sprintf("%.1f%s", currentSize/GB, "GB")
    case currentSize < PB:
        return fmt.Sprintf("%.1f%s", currentSize/TB, "TB")
    case currentSize < EB:
        return fmt.Sprintf("%.1f%s", currentSize/PB, "PB")
    default:
        return fmt.Sprintf("%.1f%s", currentSize/EB, "EB")
    }
}
