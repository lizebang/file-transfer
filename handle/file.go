package handle

import (
	"io"
	"mime"
	"net/http"
	"os"
)

// File contains the basic information to transfer the file.
type File struct {
	Ext  string
	Path string
}

func (f *File) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(f.Path)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", mime.TypeByExtension(f.Ext))
	w.WriteHeader(http.StatusOK)
}
