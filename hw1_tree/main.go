package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

var depth int

func dirTree(out *os.File, path string, printFiles bool) (error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name() // Use '<' for ascending
	})

	for _, file := range files {
		for i := 0; i < depth; i++ {
			fmt.Print("│\t")
		}
		if file == files[len(files)-1]{
			fmt.Println("└───" + file.Name())
		} else {
			fmt.Println("├───" + file.Name())
		}
		
		if file.IsDir() {
			depth++
			childFolder := filepath.Join(path, file.Name())
			dirTree(out, childFolder, printFiles)
		}
		
	}
	depth--
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
