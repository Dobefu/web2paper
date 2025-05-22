package cmd

import (
	"github.com/Dobefu/web2paper/internal/converter"
	"github.com/spf13/cobra"
)

func NewConvertCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "convert",
		Short: "Convert an HTML document to PDF",
		RunE:  runConvertCmd,
	}

	cmd.Flags().StringP("input", "i", "", "The file to process")
	cmd.Flags().StringP("output", "o", "", "The file to write")

	_ = cmd.MarkFlagRequired("input")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func init() {
	rootCmd.AddCommand(NewConvertCmd())
}

var converterNew = converter.New

func runConvertCmd(cmd *cobra.Command, args []string) error {
	// Error checking is not needed here.
	// Cobra ensures that this code will not be reached if these are missing.
	input, _ := cmd.Flags().GetString("input")
	output, _ := cmd.Flags().GetString("output")

	pdfConverter, err := converterNew(input, output)

	if err != nil {
		return err
	}

	pdfConverter.AddPage(converter.PdfSize(converter.PdfSizeA0))
	err = pdfConverter.Convert()

	if err != nil {
		return err
	}

	return nil
}
