package main

import "github.com/pverma911/go-gin-tonic/internal/router"

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}
