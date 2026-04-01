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
		    return  0, err 
	    }

    if fileInfo.IsDir() {
        files, err := os.ReadDir(path)
	    if err != nil {
		    return  0, err 
	    }

	    for _, file := range files {
            fileName := file.Name()
            isHidden := strings.HasPrefix(fileName, ".")

            if !all && isHidden {
                continue
            }

		    fullpath := filepath.Join(path, file.Name())

            currentfileInfo, err := os.Lstat(fullpath)

            if err != nil {
                return 0, err
            }

            if currentfileInfo.IsDir() && recursive {
              dirSize, err := getSize(fullpath, all, recursive)
                if err != nil {
                    return 0, err
                }
                size += dirSize
            } else if !currentfileInfo.IsDir() {
                size += currentfileInfo.Size()
            }
	    }

    } else {
        size = fileInfo.Size()
    }

	return size, nil
}

/** Функция форманирования размера файла. */
func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	const step = 1024.0
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	val := float64(size)
	i := 0
	for val >= step && i < len(sizes)-1 {
		val /= step
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%d%s", int64(val), sizes[i])
	}
	return fmt.Sprintf("%.1f%s", val, sizes[i])
}