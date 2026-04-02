package code

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
)

func GetPathSize(path string, human, all, recursive bool) (string, error) {
	size, err := getSize(path, all, recursive )
	if err != nil {
		return "", err
	}
	return formatSize(size, human), nil
}

/** Функция вычисляющая размер файла или верхнего уровня директории. */
func getSize(path string, all, recursive bool) (int64, error) {
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
            dirSize, err := getSize(fullpath, all, recursive)
            if err != nil {
                return 0, err
            }
            size += dirSize
        }
    }

    return size, nil
}


/** Функция форманирования размера файла. */
func formatSize(size int64, human bool) string {
    /** Размерность.*/
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    /** Шаг деления. */
	const step = 1024.0
    /** Размер в формате float64. */
    currentSize := float64(size)
    
	if !human {
		return fmt.Sprintf("%dB", size)
	}
 
	i := 0
	for currentSize >= step && i < len(sizes)-1 {
		currentSize /= step
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%d%s", int64(currentSize), sizes[i])
	}
	return fmt.Sprintf("%.1f%s", currentSize, sizes[i])
}