package console

import (
	"fmt"
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"

	"github.com/fajarachmadyusup13/gathering-app/internal/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gathering-app",
	Short: "Gathering App CLI",
	Long:  "CLI Tool for Gathering App",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolP("integration", "i", false, "use integration config file")
	config.GetConf()
	setupLogger()
}

func setupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &log.JSONFormatter{},
		Line:           true,
		File:           true,
	}

	if config.Env() == "development" {
		formatter = runtime.Formatter{
			ChildFormatter: &log.TextFormatter{
				ForceColors:   true,
				FullTimestamp: true,
			},
			Line: true,
			File: true,
		}
	}

	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)

	logLevel, err := log.ParseLevel(config.LogLevel())
	if err != nil {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)
}
