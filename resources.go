package main

import (
  "log"
	"net/http"
	"os"
	"io"
)

func resources(w http.ResponseWriter, req *http.Request) {
    if (logLevel > 1) {
        log.Printf("\n ----- \n" + req.Method + " file " + req.URL.Path[1:] + "\n ----- \n %s\n", req)
    } else {
        log.Printf(req.Method + " file " + req.URL.Path[1:])
    }

    switch req.Method {
    case "GET":
        urlPath := req.URL.Path[1:]
        http.ServeFile(w, req, urlPath)
        return
    case "POST":
        req.ParseMultipartForm(32 << 20)
        file, handler, err := req.FormFile("uploadfile")
        if err != nil {
           log.Println("Cannot upload empty file")
           http.Error(w, "cannot upload empty file", http.StatusBadRequest)
           return
        }
        defer file.Close()
        fileInfo, err := os.Stat(os.Getenv(sharedDirectory) + handler.Filename)
        if err == nil {
           log.Println("File already exists: " + fileInfo.Name())
           http.Error(w, "file already exists", http.StatusConflict)
           return
        }
        f, err := os.OpenFile(os.Getenv(sharedDirectory) + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
           log.Println("Could not open file")
           http.Error(w, "could not open file", http.StatusInternalServerError)
           return
        }
        defer f.Close()
        io.Copy(f, file)
        http.Redirect(w, req, "/", http.StatusSeeOther)
    case "DELETE":
        if(os.Getenv(token) != req.Header.Get("token")){
            w.WriteHeader(http.StatusUnauthorized)
            http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
            return
        }
        var err = os.Remove(req.URL.Path[1:])
        if err != nil {
           w.WriteHeader(http.StatusNotFound)
           http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
        }
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    }
}