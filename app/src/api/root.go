package api

import (
  "fmt"
  "os"
  "time"
  "net/http"
  "html/template"
  "path/filepath"

  "context"
  "onramp/pg"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
  var (
    err error
    timestamp string
  )
  stmt := "SELECT CURRENT_TIMESTAMP::varchar"
  err = pg.DB.QueryRow(context.Background(), stmt).Scan(&timestamp)
  if err != nil {
    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    return
  }

  //accept := request.Header.Get("Accept-Language")
  //if r.FormValue("lang") != "" {
  //  lang := r.FormValue("lang")
  //} else {
  //  lang := langauge.English.String() //"en"
  //}
  //localizer := i18n.NewLocalizer(bundle, language.English.String(), accept)
  // TODO: ^^^ use init()to set the language for the template, "en" vvv

  p := filepath.Join(os.Getenv("APP_ROOT"),"public","en","html","root.html")
  tpl, err := template.ParseFiles(p)
  if err != nil {
    fmt.Fprintf(os.Stderr, "template.ParseFiles err: %v", err)
    return
  }

  tpl.Execute(w, struct { 
    Message string
    Timestamp string
    PGTimestamp string 
    Lang string
  } {
    Message: "Hello, Business World!!!",
    Timestamp: time.Now().Format(time.RFC3339),
    PGTimestamp: timestamp,
    Lang: "en", // For now
  })
}

