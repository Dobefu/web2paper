package cmd

import (
	"fmt"
	"os"

	"github.com/Dobefu/web2paper/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an HTML document to PDF",
	Run:   runConvertCmd,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	log := logger.New(logger.LogLevel(viper.GetInt("log.level")), os.Stdout)

	convertCmd.Flags().StringP("input", "i", "", "The file to process")
	convertCmd.Flags().StringP("output", "o", "", "The file to write")

	err := convertCmd.MarkFlagRequired("input")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = convertCmd.MarkFlagRequired("output")

	if err != nil {
		log.Fatal(err.Error())
	}
}

func runConvertCmd(cmd *cobra.Command, args []string) {
	log := logger.New(logger.LogLevel(viper.GetInt("log.level")), os.Stdout)

	input, err := cmd.Flags().GetString("input")

	if err != nil {
		log.Fatal(err.Error())
	}

	output, err := cmd.Flags().GetString("output")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(input)
	fmt.Println(output)
}
