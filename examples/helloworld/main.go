package main

import (
	"io"
	"log"
	"os"

	brainfuck "github.com/zzhaolei/brainfuck"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("need input bf file")
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("close file err: %s", err)
		}
	}()

	code, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err = brainfuck.Brainfuck(code); err != nil {
		log.Fatal(err)
	}
}
