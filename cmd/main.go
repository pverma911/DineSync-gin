package main

import (
    "myapp/internal/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}
