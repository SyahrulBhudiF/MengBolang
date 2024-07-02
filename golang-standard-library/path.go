package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("/a/b/c.go"))
	fmt.Println(path.Base("/a/b/c.go"))
	fmt.Println(path.Ext("/a/b/c.go"))
	fmt.Println(path.Join("/a/b", "c.go"))
	fmt.Println()

	fmt.Println(filepath.Dir("/a/b/c.go"))
	fmt.Println(filepath.Base("/a/b/c.go"))
	fmt.Println(filepath.Ext("/a/b/c.go"))
	fmt.Println(filepath.Join("/a/b", "c.go"))
}
