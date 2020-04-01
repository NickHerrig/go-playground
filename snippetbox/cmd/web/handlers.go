package main

import (
    "errors"
    "fmt"
    //"html/template"
    "net/http"
    "strconv"

    "nickherrig.com/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    s, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    for _, snippet := range s {
        fmt.Fprintf(w, "%v/n", snippet)
    }

//    files := []string{
//        "./ui/html/home.page.tmpl",
//        "./ui/html/base.layout.tmpl",
//        "./ui/html/footer.partial.tmpl",
//    }
//
//    ts, err := template.ParseFiles(files...)
//    if err != nil {
//        app.serverError(w, err)
//        return
//    }
//
//    err = ts.Execute(w, nil)
//    if err != nil {
//        app.serverError(w, err)
//    }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    s, err := app.snippets.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }

    fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    title :=  "dummy title data"
    content :=  "dummy content data"
    expires := "7"

    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
