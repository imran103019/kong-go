package main

import (
    "github.com/Kong/go-pdk"
    "net/http"
)

type config struct {
    APIKey string
}

// New ...
func New() interface{} {
    return &config{}
}

// Access ...
func (conf config) Access(kong *pdk.PDK) {

    apiKey, err := kong.Request.GetHeader("Api-Key")
    headers := make(map[string][]string)
    headers["Content-Type"] = append(headers["Content-Type"], "application/json")
    if err != nil {
        kong.Log.Err("No Api-Key provided!")
        kong.Response.Exit(http.StatusUnauthorized, "No Api-Key provided!", headers)
        return
    }

    if conf.APIKey != apiKey {
        kong.Log.Err("Invalid Api-Key provided!")
        kong.Response.Exit(http.StatusUnauthorized, "Unauthorized", headers)
        return
    }

}
