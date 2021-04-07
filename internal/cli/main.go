package cli

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/deploy"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/deposit"
	"github.com/alecthomas/kingpin"
	"os"
)

func Run(args []string) bool {
	cfg := config.NewConfig(os.Getenv("CONFIG"))
	log := cfg.Logger()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("app panicked\n", rvr)
		}
	}()

	app := kingpin.New("bridge", "")

	runCmd := app.Command("run", "run command")
	depositService := runCmd.Command("deposit", "run a service to deposit into Odin")
	deployService := runCmd.Command("deploy", "run a service to deploy a bridge contract")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case depositService.FullCommand():
		svc := deposit.New(cfg)
		svc.Run(context.Background())
	case deployService.FullCommand():
		svc := deploy.New(cfg)
		err := svc.Run(context.Background())
		if err != nil {
			log.WithError(err).Error("failed to run deploy")
			return false
		}
	default:
		log.WithField("command", cmd).Error("Unknown command")
		return false
	}

	return true
}
