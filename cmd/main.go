package main

import (
	"github.com/XT3RM1NATOR/full_auth/config"
	"github.com/XT3RM1NATOR/full_auth/internal/app/db"
	"github.com/XT3RM1NATOR/full_auth/internal/app/server"
	"github.com/XT3RM1NATOR/full_auth/internal/app/storage"
)

func main() {
	cfg := config.Load()
	mongodb := db.ConnectToDB(cfg)
	str := storage.ConnectToStorage(cfg)

	server.RunHTTPServer(cfg, mongodb, str)
}
