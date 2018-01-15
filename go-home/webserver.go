package main

import (
  "fmt"
  "strconv"
  "strings"
	"html/template"
  "io/ioutil"
	"net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func loadAsset(path string) ([]byte, error) {
  content, err := ioutil.ReadFile(path)
  if err != nil {
    return []byte{}, err
  }
  return content, err
}

func getMimeType(extension string) (string) {
  switch extension {
  case "js":
    return "text/javascript"
  case "css":
    return "text/css"
  case "woff":
  case "woff2":
    return "application/font-woff"
  default:
    return ""
  }
  return ""
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "go-home"}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Login"}
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(w, p)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
  content, err := loadAsset(r.URL.Path[1:])
  chunkedUrl := strings.Split(r.URL.Path, ".")
  mime := getMimeType(chunkedUrl[len(chunkedUrl)-1])
  w.Header().Set("Content-Type", mime)
  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, fmt.Sprintf("Could not find %s", r.URL.Path))
  }
  w.Header().Set("Content-Length", strconv.Itoa(len(content)))
  w.Write(content)
}

func startWebServer(port int) {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/login/", loginHandler)
  http.HandleFunc("/assets/", assetsHandler)
  http.ListenAndServe(":8080", nil)
}
