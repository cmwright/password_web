package main

import (
  // "fmt"
  "os"
  "github.com/hoisie/web"
  "os/exec"
  "encoding/json"
  "log"
)

func main() {
  f, err := os.Create("/dev/null")
  if err != nil {
    println(err.Error())
    return
  }
  logger := log.New(f, "testing123", log.Ldate|log.Ltime)
  web.SetLogger(logger)

  web.Post("/fetch/(.*)", fetch_password)
  web.Post("/set/(.*)", set_password)

  web.Run("0.0.0.0:9999")
}

func make_json_response(domain, response string) []byte {
  response_map := make(map[string]string)
  response_map["domain"] = domain
  response_map["result"] = response

  json_output, _ := json.Marshal(response_map)
  return json_output
}

func fetch_password(ctx *web.Context, domain string) {
  key := get_key(ctx.Params)

  cmd := exec.Command("password_store", "get", domain, key)
  out, _ := cmd.Output()

  json_response := make_json_response(domain, string(out))

  ctx.SetHeader("Content-Type", "application/json; charset=utf-8", true)
  ctx.Write(json_response)
}

func set_password(ctx *web.Context, domain string) {
  key := get_key(ctx.Params)

  cmd := exec.Command("password_store", "set", domain, "20", key)
  out, _ := cmd.Output()

  json_response := make_json_response(domain, string(out))

  ctx.SetHeader("Content-Type", "application/json; charset=utf-8", true)
  ctx.Write(json_response)
}


func get_key(params map[string]string) string {
  return params["key"]
}
