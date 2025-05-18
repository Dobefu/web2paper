package cmd

import (
	"os"

	"github.com/Dobefu/web2paper/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd, _ = NewRootCmd()

	logLevel int
)

func NewRootCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:   "web2paper",
		Short: "Quickly and easily convert an HTML document to PDF",
	}

	viper.Set("log.level", logLevel)
	cmd.PersistentFlags().IntVarP(&logLevel, "loglevel", "l", logger.LogLevelInfo, "The log level to use. 0=TRACE, 1=DEBUG, 2=INFO, 3=WARN, 4=ERROR, 5=FATAL")

	return cmd, nil
}

func Execute() {
	log := logger.New(logger.LogLevel(viper.GetInt("log.level")), os.Stdout)

	err := rootCmd.Execute()

	if err != nil {
		_, _ = log.Error(err.Error())
	}
}
