package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "math"
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

/** Функция форманирования размера файла. */
func FormatSize(size int64) (string) {
    /** Формат размеров. */
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


/** Точка входа в приложение. */
func main() {
    cmd := &cli.Command{
        Name:  "hexlet-path-size",
        Usage: "print size of a file or directory",
        Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:  "human",
                Aliases: []string{"H"},
                Usage: "human-readable sizes",
            },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {
        path := cmd.Args().First()
     
        /** Последний елемент пути к файлу или директории. */
        last := filepath.Base(path)

        /** Итоговый результат после форматирования. */
        var resultSize string

        if path == "" {
            fmt.Println("Пожалуйста, укажите путь к файлу или директории")
            return nil
        }
        
        size, err := GetSize(path)

        if err != nil {
            log.Printf("Ошибка: %v\n", err)
            return err
        }

        if cmd.Bool("human") {
           resultSize =  FormatSize(size)
            fmt.Print(resultSize,"\\t", last)
            return nil
        }

        fmt.Print(size,"B\\t", last)
        return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}