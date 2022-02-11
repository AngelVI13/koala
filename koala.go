package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	subDirToSkip := "skip"

    walkDir := "C:/Users/avili/Documents/"
	fmt.Printf("Walk dir %+v:", walkDir)
    err := filepath.Walk(walkDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return nil  // if file can't be accessed move to the next one
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", walkDir, err)
		return
	}
}
