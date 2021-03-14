package cli

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/deployer"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/services/depositer"
	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	cfg := config.NewConfig(kv.MustFromEnv())
	log = cfg.Log()

	app := kingpin.New("odin-deposit-ether-svc", "")

	runCmd := app.Command("run", "run command")
	deposit := runCmd.Command("deposit", "run deposit service")
	deploy := runCmd.Command("deploy", "run deploy service")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case deposit.FullCommand():
		svc := depositer.New(cfg)
		err := svc.Run(context.Background())
		if err != nil {
			log.WithError(err).Error("failed to run depositer")
			return false
		}
	case deploy.FullCommand():
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
