package cmds

import (
	"log"

	"github.com/Woody1193/miletos-test/io"
	"github.com/Woody1193/miletos-test/params"
	"github.com/Woody1193/miletos-test/types"
	"github.com/spf13/cobra"
)

// CheckCmd allows the user to verify invoice data against receivables data
func CheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check receivables data",
		Long:  "Checkes that a receivables file matches an invoices file",
		Run: func(cmd *cobra.Command, args []string) {

			eh := new(io.ErrorHandler)

			invoiceData, err := io.ReadCSV[types.InvoiceItem](params.InvoiceFile, eh)
			if err != nil {
				log.Fatalf("Failed to read invoice data, error: %v", err)
			}

			receivablesData, err := io.ReadCSV[types.ReceivablesItem](params.ReceivablesFile, eh)
			if err != nil {
				log.Fatalf("Failed to read receivables data, error: %v", err)
			}
		},
	}

	cmd.Flags().StringVar(&params.InvoiceFile, "invoice", "invoice",
		"Name of the CSV file containing invoice data")
	cmd.Flags().StringVar(&params.ReceivablesFile, "receivables", "receivables",
		"Name of the CSV file containing receivables data")
	cmd.Flags().StringVar(&params.OutputFile, "output", "output",
		"Name of the CSV file that should contain the results of the check")
	cmd.Flags().StringVar(&params.ErrorFile, "error", "error",
		"Name of the CSV file that should contain any errors that occur")
	return cmd
}
