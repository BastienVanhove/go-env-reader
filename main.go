package main

import (
    "fmt"
    "log"
    "github.com/joho/godotenv"
)

func main() {
    var envs map[string]string
    envs, err := godotenv.Read(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    fmt.Println(envs)
   }
