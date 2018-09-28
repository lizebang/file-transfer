package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/lizebang/file-transfer/handle"
	"github.com/lizebang/file-transfer/ip"
	"github.com/lizebang/file-transfer/qr"
)

var (
	dirpath  string
	filename string
	host     string
	port     string
)

func init() {
	host = ip.IP()
	flag.Usage = help
	flag.StringVar(&filename, "f", "", "File Path")
	flag.StringVar(&port, "p", "8080", "Port")
}

func help() {
	fmt.Println("File Transfer")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if len(filename) == 0 || len(dirpath) == 0 {
		log.Fatalln("invalid path")
	}

	abspath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalln(err)
	}
	base := filepath.Base(abspath)
	ext := filepath.Ext(abspath)
	qr.QR("http://" + host + ":" + port + "/" + base)

	http.Handle("/"+base, &handle.File{Ext: ext, Path: abspath})
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
