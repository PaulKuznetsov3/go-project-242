package code

import (
    "fmt"
    "os"
    "strings"
)

func GetPathSize(path string, human, all, recursive bool) (string, error) {
	size, err := getSize(path, all, recursive )
	if err != nil {
		return "", err
	}
	return formatSize(size, human), nil
}

/** Функция вычисляющая размер файла или верхнего уровня директории. */
func getSize(path string, all bool, recursive bool) (int64, error) {
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
            } 

         
            size += currentfileInfo.Size()
	    }


    } else {
        size = fileInfo.Size()
    }

	return size, nil
}

/** Функция форманирования размера файла. */
func formatSize(size int64, human bool) (string) {
    /** Форматы размеров. */
    sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    /** Индекс по умолчанию для sizes. */
    defaultIndex := 0
    /** Итоговый результат. */
    var resultSize string
    /** Делитель. */
    var divider int64 = 1024
    
    if !human {
        return fmt.Sprintf("%.1f%s", float64(size), sizes[defaultIndex])
    }
  
    for i, s := range sizes {
	    if size < divider || i == len(sizes)-1 {
            resultSize = fmt.Sprintf("%.1f%s", float64(size), s)
            break
        } else {
            size = size / divider
        }
	}

   return resultSize 
}