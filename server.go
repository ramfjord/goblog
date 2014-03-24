package tinyserver

import (
  "html/template"
  "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  p, err := loadPage("home")
  title := "Thomblog"
  if err != nil {
    // handle errors better please
    // p = &Page{Title: "nope", Body: "nope"}
  }
  t, _ := template.ParseFiles("templates/application.html")
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8080", nil)
}

