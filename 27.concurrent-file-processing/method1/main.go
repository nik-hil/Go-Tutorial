package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type pair struct {
	hash string
	path string
}
type filelist []string
type results map[string]filelist

func hashFile(path string) pair {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
	for path := range paths {
		pairs <- hashFile(path)
	}
	done <- true
}

func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)
	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}
	result <- hashes
}

func searchTree(dir string, paths chan<- string) error {
	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}
		if fi.Mode().IsRegular() && fi.Size() > 0 {
			paths <- p
		}
		return nil
	}
	return filepath.Walk(dir, visit)
}
func search(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0)
	paths := make(chan string)
	pairs := make(chan pair)
	done := make(chan bool)
	result := make(chan results)

	for i := 0; i < workers; i++ {
		go processFiles(paths, pairs, done)
	}

	go collectHashes(pairs, result)

	if err := searchTree(dir, paths); err != nil {
		return nil
	}
	close(paths)

	for i := 0; i < workers; i++ {
		<-done
	}
	close(pairs)
	return <-result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing dir name. Run as $ go run . test")
	}
	if hashes := search(os.Args[1]); hashes != nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				fmt.Println(hash[len(hash)-7:], len(files))
				for _, file := range files {
					fmt.Println(" ", file)
				}
			}
		}
	}

}
