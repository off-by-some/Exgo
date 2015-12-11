package config

import "github.com/jacobstr/confer"
import "os"
import "fmt"

var File *confer.Config

func init() {
  File = confer.NewConfig()
  appenv := os.Getenv("EXGO_ENV");
  configPath := fmt.Sprintf("%s", os.Getenv("GOPATH")) + "/src/Exgo/config/"

  // Default to development.yml if no environment is set
  paths := []string{configPath + "development.yml"}

  // $EXGO_ENV should be one of "production", "development" or "test"
  if (appenv != "") {
    paths = append(paths, configPath + fmt.Sprintf("%s.yml", appenv))
  }

  // Override other config with local config
  paths = append(paths, configPath + "local.yml")

  if err := File.ReadPaths(paths...); err != nil {
    panic(err)
  }
}
