package cmds

import (
	"log"

	"github.com/Woody1193/miletos-test/batch"
	"github.com/Woody1193/miletos-test/io"
	"github.com/Woody1193/miletos-test/params"
	"github.com/Woody1193/miletos-test/rules"
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

			invoiceData, errResults, err := io.ReadCSV[string, *types.InvoiceItem](params.InvoiceFile)
			if err != nil {
				log.Fatalf("Failed to read invoice data, error: %v", err)
			}

			receivablesData, eh, err := io.ReadCSV[string, *types.ReceivablesItem](params.ReceivablesFile)
			if err != nil {
				log.Fatalf("Failed to read receivables data, error: %v", err)
			}

			errResults = append(errResults, eh...)
			if err := io.WriteCsv(params.ErrorFile, errResults...); err != nil {
				log.Fatalf("Failed to write error file, error: %v", err)
			}

			results := batch.NewCheckBatch(invoiceData, receivablesData,
				rules.InvoiceExists, rules.AmountsEqual, rules.PaidOnTime, rules.DateNotInFuture).Check()

			if err := io.WriteCsv(params.OutputFile, results...); err != nil {
				log.Fatalf("Failed to write output file, error: %v", err)
			}

			log.Printf("Check complete. Results written to %q", params.OutputFile)
		},
	}

	cmd.Flags().StringVar(&params.InvoiceFile, "invoice", "invoice.csv",
		"Name of the CSV file containing invoice data")
	cmd.Flags().StringVar(&params.ReceivablesFile, "receivables", "receivables.csv",
		"Name of the CSV file containing receivables data")
	cmd.Flags().StringVar(&params.OutputFile, "output", "output.csv",
		"Name of the CSV file that should contain the results of the check")
	cmd.Flags().StringVar(&params.ErrorFile, "error", "errors.csv",
		"Name of the CSV file that should contain any errors that occur")
	return cmd
}
