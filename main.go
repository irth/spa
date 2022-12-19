package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func clean(root string, dirty string) string {
	return filepath.Join(root, filepath.FromSlash(path.Clean("/"+dirty)))
}

func tryIndex(dirPath string) (string, error) {
	filePath := filepath.Join(dirPath, "index.html")
	s, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	if s.IsDir() {
		return filePath, os.ErrNotExist
	}
	return filePath, nil
}

func pickFile(filePath string) (string, error) {
	s, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	if s.IsDir() {
		return tryIndex(filePath)
	}
	return filePath, nil
}

func handler(dir string) http.Handler {
	if dir == "" {
		dir = "."
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fallback := false
		defer func() {
			fb := ""
			if fallback {
				fb = " [fallback]"
			}
			log.Printf("%s %s %s%s\n", r.RemoteAddr, r.Method, r.URL, fb)
		}()

		filePath, err := pickFile(clean(dir, r.URL.Path))
		if errors.Is(err, os.ErrNotExist) {
			fallback = true
			http.ServeFile(w, r, filepath.Join(dir, "/index.html"))
			return
		}
		if err != nil {
			w.WriteHeader(500)
			return
		}

		http.ServeFile(w, r, filePath)
	})
}

func main() {
	var addr string
	var dir string

	flag.StringVar(&addr, "addr", ":8000", "address to listen on")
	flag.StringVar(&dir, "dir", ".", "dir to serve")
	flag.Parse()

	p, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("getting abs path to \"%s\" failed: %s", dir, err)
	}
	log.Printf("Serving %s on %s...", p, addr)
	if err := http.ListenAndServe(addr, handler(p)); err != nil {
		log.Fatal(err)
	}
}
