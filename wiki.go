package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

//func main() {
//  p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.\nTo be noted, the filename 'TestPage' is case sensitive in the url.")}
//  p1.save()
//  p2, _ := loadPage("TestPage")
//  fmt.Println(string(p2.Body))
//}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.ListenAndServe(":8080", nil)
}
