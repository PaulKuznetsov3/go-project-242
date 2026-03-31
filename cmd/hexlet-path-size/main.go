package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "code"
    "path/filepath"
    "github.com/urfave/cli/v3"
)




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
                &cli.BoolFlag{
                Name:  "all",
                Aliases: []string{"a"},
                Usage: "include hidden files and directories",
            },
                &cli.BoolFlag{
                Name:  "recursive",
                Aliases: []string{"r"},
                Usage: "recursive size of directories",
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
        
        size, err := code.GetSize(path, cmd.Bool("all"), cmd.Bool("recursive"))

        if err != nil {
            return err
        }

        if cmd.Bool("human") {
           resultSize =  code.FormatSize(size)
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