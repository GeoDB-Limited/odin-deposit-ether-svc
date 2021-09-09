package cli

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/deploy"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/deposit"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/fix"
	"github.com/alecthomas/kingpin"
	"os"
)

func Run(args []string) bool {
	cfg := config.NewConfig(os.Getenv("CONFIG"))
	ctx := context.Background()
	log := cfg.Logger()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("app panicked\n", rvr)
		}
	}()

	app := kingpin.New("odin-deposit-ether-svc", "")

	runCmd := app.Command("run", "run command")
	depositService := runCmd.Command("deposit", "run a service to deposit into Odin")
	deployService := runCmd.Command("deploy", "run a service to deploy a bridge contract")
	fixService := runCmd.Command("fix", "run a service to fix unprocessed transactions")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case depositService.FullCommand():
		deposit.New(ctx, cfg).Run()
	case deployService.FullCommand():
		if err := deploy.New(ctx, cfg).Run(); err != nil {
			log.WithError(err).Error("failed to run deploy service")
			return false
		}
	case fixService.FullCommand():
		fix.New(ctx, cfg).Run()
	default:
		log.WithField("command", cmd).Error("Unknown command")
		return false
	}

	return true
}
