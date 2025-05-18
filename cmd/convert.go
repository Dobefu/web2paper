package cmd

import (
	"github.com/Dobefu/web2paper/internal/converter"
	"github.com/spf13/cobra"
)

func NewConvertCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:   "convert",
		Short: "Convert an HTML document to PDF",
		RunE:  runConvertCmd,
	}

	cmd.Flags().StringP("input", "i", "", "The file to process")
	cmd.Flags().StringP("output", "o", "", "The file to write")

	err = cmd.MarkFlagRequired("input")

	if err != nil {
		return cmd, err
	}

	err = cmd.MarkFlagRequired("output")

	if err != nil {
		return cmd, err
	}

	return cmd, nil
}

func init() {
	cmd, _ := NewConvertCmd()

	rootCmd.AddCommand(cmd)
}

var converterNew = converter.New

func runConvertCmd(cmd *cobra.Command, args []string) (err error) {
	input, err := cmd.Flags().GetString("input")

	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")

	if err != nil {
		return err
	}

	converterNew(input, output)

	return nil
}
