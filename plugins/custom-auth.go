package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/Kong/go-pdk"
    "net/http"
    "time"
)

type config struct {
}

const (
    authSvc = "http://auth-svc"
)

// New ...
func New() interface{} {
    return &config{}
}

// Access ...
func (conf config) Access(kong *pdk.PDK) {

    bearerToken, err := kong.Request.GetHeader("Authorization")
    headers := make(map[string][]string)
    headers["Content-Type"] = append(headers["Content-Type"], "application/json")
    if err != nil {
        kong.Log.Err("No authorization token provided!")
        kong.Response.Exit(http.StatusUnauthorized, "No Authorization Token Provided", headers)
        return
    }

    // Requesting external auth service
    url := fmt.Sprintf("%s/api/v1/me", authSvc)
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        kong.Log.Err("Auth svc requesting error: " + err.Error())
        kong.Response.Exit(http.StatusInternalServerError, err.Error(), headers)
    }
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("Authorization", bearerToken)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    request = request.WithContext(ctx)
    client := http.Client{}

    resp, err := client.Do(request)
    if err != nil {
        kong.Log.Err("Auth svc do request error: " + err.Error())
        kong.Response.Exit(http.StatusInternalServerError, err.Error(), headers)
        return
    }
    defer resp.Body.Close()

    var result struct {
        Me struct {
            Email string `json:"email"`
            Exp   int    `json:"exp"`
            Name  string `json:"name"`
        } `json:"me"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        kong.Log.Err(fmt.Sprintf("Auth svc json decode error: %s", err.Error()))
        kong.Response.Exit(http.StatusInternalServerError, err.Error(), headers)
    }
    if resp.StatusCode != http.StatusOK {
        kong.Log.Err(fmt.Sprintf("Auth svc error code: %d", resp.StatusCode))
        kong.Response.Exit(resp.StatusCode, "Unauthorized", headers)
        return
    }

    // Modifying header for requested service
    kong.ServiceRequest.SetHeader("Email", result.Me.Email)

}
