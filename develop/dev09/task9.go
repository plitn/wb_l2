package dev09

import (
	"flag"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"path/filepath"
	"strings"
)

func Wget() {
	out := flag.String("O", ".", "Output path")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Println("wrong usage")
		os.Exit(1)
	}
	url, err := url2.Parse(flag.Arg(0))
	if err != nil {
		log.Println("parse error")
		os.Exit(1)
	}

	if _, err := os.Stat(*out); os.IsNotExist(err) {
		if err := os.MkdirAll(*out, 0755); err != nil {
			log.Println("mk dir error")
			os.Exit(1)
		}
	}

	responce, err := http.Get(url.String())
	if err != nil {
		log.Println("download error")
		os.Exit(1)
	}

	defer responce.Body.Close()
	var outputFile string

	if strings.HasSuffix(url.Path, "/") || url.Path == "" {
		outputFile = filepath.Join(*out, url.Host, url.Path, "index.html")
	} else {
		outputFile = filepath.Join(*out, url.Host, url.Path)
	}
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		log.Println("creating dir error")
		os.Exit(1)
	}
	f, err := os.Create(outputFile)
	if err != nil {
		log.Println("file create error")
		os.Exit(1)
	}
	defer f.Close()

	_, err = io.Copy(f, responce.Body)
	if err != nil {
		log.Println("writing error")
		os.Exit(1)
	}
}
