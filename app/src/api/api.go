package api

import (
  "fmt"
  "os"
  "net/http"

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

