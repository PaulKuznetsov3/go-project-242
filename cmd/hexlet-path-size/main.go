package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "path/filepath"
    "github.com/urfave/cli/v3"
)

/** Функция обработки ошибки. */
func checkError(err error, path string) {
    	if err != nil {
    	if os.IsNotExist(err) {
        	log.Fatalf("Файл или ссылка '%s' не существует", path)
    	} else {
        	log.Fatalf("Произошла ошибка: %v", err)
    	}
	}
}

/** Функция вычисляющая размер файла или верхнего уровня директории. */
func GetSize(path string) (int64, error) {
    /** Итоговый размер файла. */
    var size int64
 
	fileInfo, err := os.Lstat(path)
    checkError(err, path)

    if fileInfo.IsDir() {
        files, err := os.ReadDir(path)
	    if err != nil {
		    log.Fatal(err)
	    }

	    for _, file := range files {
		    fullpath :=  fmt.Sprint(path,"/",file.Name())
            fileInfo, err := os.Lstat(fullpath)

            checkError(err, fullpath)
            size += fileInfo.Size()
	    }


    } else {
        size = fileInfo.Size()
    }

	return size, nil
}

/** Точка входа в приложение. */
func main() {
    cmd := &cli.Command{
        Name:  "hexlet-path-size",
        Usage: "print size of a file or directory",
        Action: func(ctx context.Context, cmd *cli.Command) error {
        path := cmd.Args().First()
     
        if path == "" {
            fmt.Println("Пожалуйста, укажите путь к файлу или директории")
            return nil
        }
        
        size, err := GetSize(path)

        if err != nil {
            log.Printf("Ошибка: %v\n", err)
            return err
        }

        /** Последний елемент пути к файлу или директории. */
        last := filepath.Base(path)
        
        fmt.Print(size,"B\\t", last)
            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}