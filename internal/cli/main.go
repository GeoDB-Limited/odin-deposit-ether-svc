package cli

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/deployer"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/depositer"
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

	app := kingpin.New("odin-deposit-ether-svc", "")

	runCmd := app.Command("run", "run command")
	depositerService := runCmd.Command("depositer", "run a service to deposit into Odin")
	deployerService := runCmd.Command("deployer", "run a service to deploy a bridge contract")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case depositerService.FullCommand():
		svc := depositer.New(cfg)
		err := svc.Run(context.Background())
		if err != nil {
			log.WithError(err).Error("failed to run depositer")
			return false
		}
	case deployerService.FullCommand():
		svc := deployer.New(cfg)
		err := svc.Run(context.Background())
		if err != nil {
			log.WithError(err).Error("failed to run deployer")
			return false
		}
	default:
		log.WithField("command", cmd).Error("Unknown command")
		return false
	}

	return true
}
