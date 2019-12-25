package main

import (
  "github.com/tttmaximttt/gobot/botApp"
  "net/http"

  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/tttmaximttt/gobot/config"
)

var conf config.Config

func setLogLevel(logLevel string) {
  switch logLevel {
    case "debug":
      log.SetLevel(log.DebugLevel)
    case "warn":
      log.SetLevel(log.WarnLevel)
    case "panic":
      log.SetLevel(log.PanicLevel)
    case "fatal":
      log.SetLevel(log.FatalLevel)
    case "error":
      log.SetLevel(log.ErrorLevel)
    default:
      log.SetLevel(log.InfoLevel)
  }
}

func main() {
  conf, err := config.LoadConfig("APP")

  if err != nil {
    log.Error("Config is not define", err)
  }

  if conf.Logger.LogLevel == "" {
    log.Error("LogLevel not define")
  }

  setLogLevel(conf.Logger.LogLevel)

  defer func () {
    if err := recover(); err != nil {
      fmt.Println("Panic hapend cause:", err)
    }
  }()

  if err != nil {
    log.Error(err)
  }

  bot, err := botApp.New(*conf)

  if err != nil {panic(err)}

  // bot.Debug = true
  log.WithFields(log.Fields{
    "account": bot.Self.UserName,
  }).Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

  updates := bot.ListenForWebhook("/")
  updatesList := bot.ListenForWebhook("/list")

  go http.ListenAndServe(fmt.Sprintf(":%d", 8383), nil)
  log.Debug(fmt.Sprintf("start listen :%d", conf.Port))


  botApp.Run(*bot, updates)
  botApp.Run(*bot, updatesList)
}
