package main

import (
  "io/ioutil"
  "text/template"
  "net/http"
  "github.com/knieriem/markdown"
  "bytes"
  "os"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage loads the markdown for a page from templates/title.markdown, and 
// converts it to HTML, returning a new Page with
// Page.Title = title string
// Page.Body = new html
func loadPage(title string) (*Page, error) {
  // get markdown for page specified
  filename := "templates/" + title + ".markdown"
  markdown_file, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  defer markdown_file.Close()

  // parse markdown into Page.Body
  var html_buf bytes.Buffer
  parser := markdown.NewParser(&markdown.Extensions{Smart: true})
  parser.Markdown(markdown_file, markdown.ToHTML(&html_buf))

  return &Page{Title: title, Body: html_buf.Bytes()}, nil
}

func loadPageIfExists(title string) (*Page) {
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: "nope"}
  }
  return p
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/application.html")
  t.Execute(w, loadPageIfExists("home"))
}

func templatePath(f string) {
  return "templates/" + f + ".html"
}

func blogPostPath(f string) {
  return "blog_posts/" + f + ".text"
}

// blogPageHandler is the default handler for blog posts.  They should be available
// in a file called templates/blog_post.text
func blogPageHandler(w http.ResponseWriter, r *http.Request, f string) {
  t, _ := template.ParseFiles(templatePath("blog_post"))
  t.Execute(w, loadPageIfExists(blogPostPath(f)))
}

func main() {
  blog_pages = make([]string,
    "home",
    "vim_usage",
    "vim_plugins",
    "postgres_analyzer")

  for i, page := range blog_pages {
    // handle page
  }
  http.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8080", nil)
}
