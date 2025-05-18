package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: which <binary name>")
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		findBinary(os.Args[1])
	} else {
		fmt.Println("Unsupported OS")
		os.Exit(1)
	}

}

func findBinary(bin string) {
	// fmt.Println(bin)

	pathVariables := os.Getenv("PATH")
	paths := strings.Split(pathVariables, ";")
	wg.Add(len(paths))

	for _, path := range paths {
		// fmt.Println(path)
		go func() {
			defer wg.Done()
			filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
				if err != nil || info == nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				base := info.Name()
				if base == bin || base == bin+".exe" || base == bin+".bat" || base == bin+".cmd" {
					fmt.Println(path)
				}
				return nil
			})
		}()
	}

	wg.Wait()
}
