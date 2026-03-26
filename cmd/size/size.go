package size

import (
	"fmt"
	"log"
	"os"
)

func GetSize(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
        if os.IsNotExist(err) {
            log.Fatalf("Файл или ссылка '%s' не существует", path)
        } else {
            log.Fatalf("Произошла ошибка: %v", err)
        }
    }

	fmt.Printf("  Ссылка ведет на: %s\n", fileInfo)
	return fileInfo.Size()
}