package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/nkarpenko/playlog-test/api/app"
	"github.com/nkarpenko/playlog-test/api/conf"
	"github.com/nkarpenko/playlog-test/api/service"
	"github.com/spf13/cobra"
)

// RootCmd execution.
func RootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "playlog",
		Short: "Playlog Test API Service.",
		Long:  `Playlog Test API Service."`,
		Run: func(cmd *cobra.Command, args []string) {
			configFile, err := cmd.Flags().GetString("config")
			if err != nil {
				log.Fatalf("invalid CLI flags, please use the -h flag to see all available options: %+v", err)
				return
			}
			config, err := conf.Load(configFile)
			if err != nil {
				log.Fatalf("failed to load configration file: %+v", err)
				return
			}
			Start(config)
		},
	}

	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "Specify local configuration file.")
	rootCmd.AddCommand(versionCmd)

	return rootCmd
}

// Start the service
func Start(c *conf.Configuration) {
	log.Infof("%s v(%s)", c.Name, app.Version)
	log.Infof("%s", c.Desc)

	s, err := service.New(c)
	if err != nil {
		log.Fatalf("Service initialization failed. Error: %v", err)
		return
	}

	// Start the service
	s.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until quit signal is received
	<-stop

	// Create shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(s.Config().ShutdownTime))
	defer cancel()

	// Shutdown the service
	s.Stop(ctx)
	log.Info("Shutting down the service.")
	os.Exit(0)
}
