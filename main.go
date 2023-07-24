package main

import (
	"os"

	"github.com/nanthaphol2/bigdogshop/config"
	"github.com/nanthaphol2/bigdogshop/modules/servers"
	"github.com/nanthaphol2/bigdogshop/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()

}
