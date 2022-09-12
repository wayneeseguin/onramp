package api

import (
  "sort"
  "fmt"
  "os"
  "time"
  "net/http"
  "html/template"
  "path/filepath"

  "context"
  "onramp/pg"

  // i18n/l10n:
	//"github.com/BurntSushi/toml"
	//"github.com/nicksnyder/go-i18n/v2/i18n"
	//"golang.org/x/text/language"
  // TODO: Check out https://github.com/vorlif/spreak
)

var (
  //bundle = *i18n.Bundle
)

func init() {
  //bundle := i18n.NewBundle(language.English)
	//bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//bundle.MustLoadMessageFile("active.zh.toml")
}

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
  } {
    Message: "Hello, Business World!!!",
    Timestamp: time.Now().Format(time.RFC3339),
    PGTimestamp: timestamp,
  })
}

func apiV1HeadersHandler(w http.ResponseWriter, req *http.Request) {
  timestamp := time.Now().Format(time.RFC3339)

  headerKeys := make([]string, 0, len(req.Header))
  for name := range req.Header {
    headerKeys = append(headerKeys, name)
  }
  sort.Strings(headerKeys)
  headers := make(map[string]string)
  for _, k := range headerKeys {
    for _,v := range req.Header[k] {
      headers[k] = v
    }
  }

  p := filepath.Join(os.Getenv("APP_ROOT"),"public","en","html","v1_headers.html")
  tpl, err := template.ParseFiles(p)
  if err != nil {
    fmt.Fprintf(os.Stderr, "template.ParseFiles err: %v", err)
    return
  }

  tpl.Execute(w, struct { 
    Headers map[string]string
    Timestamp string
  } {
    Headers: headers,
    Timestamp: timestamp,
  })
}

func Api() (err error) {
  http.HandleFunc("/", rootHandler)
  http.HandleFunc("/api/v1/headers", apiV1HeadersHandler)

  err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
  if err != nil {
    fmt.Fprintf(os.Stderr, "ListenAndServe: %v", err)
    os.Exit(1)
  }

  return
}

