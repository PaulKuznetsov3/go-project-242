package code

import (
	"math"
    "fmt"
    "os"
    "strings"
)

/** Функция вычисляющая размер файла или верхнего уровня директории. */
func GetPathSize(path string, all bool, recursive bool) (int64, error) {
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

		    fullpath :=  fmt.Sprint(path,"/",file.Name())

            if fileInfo.IsDir() && recursive {
             dirSize, err := GetPathSize(fullpath, all, recursive)
                if err != nil {
                    return 0, err
                }
                size += dirSize
            } 

            fileInfo, err := os.Lstat(fullpath)

             if err != nil {
                    return 0, err
                }
            size += fileInfo.Size()
	    }


    } else {
        size = fileInfo.Size()
    }

	return size, nil
}

/** Функция форманирования размера файла. */
func FormatSize(size int64) (string) {
    /** Форматы размеров. */
    sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    /** Итоговый результат. */
    var resultSize string
    /** Делитель. */
    var divider int64 = 1024

    for _, s := range sizes {
	    if size < divider {
            resultSize = fmt.Sprintf("%.1f%s", math.Floor(float64(size)), s)
            return resultSize
        } else {
           size = size / divider
        }

	}
    
    return resultSize
}