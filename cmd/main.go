package main

import (
	"fmt"

	"github.com/pverma911/go-gin-tonic/internal/config"
	"github.com/pverma911/go-gin-tonic/internal/router"
)

func main() {
    db:=config.InitializeDb()
    config.CreateOrUpdateTables(db)
    r := router.SetupRouter(db)
    r.Run("127.0.0.1:8080")
    fmt.Println("Server started")
}
