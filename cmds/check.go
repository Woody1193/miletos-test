package cmds

import (
	"log"
	"os"

	"github.com/Woody1193/miletos-test/batch"
	"github.com/Woody1193/miletos-test/io"
	"github.com/Woody1193/miletos-test/params"
	"github.com/Woody1193/miletos-test/rules"
	"github.com/Woody1193/miletos-test/types"
	"github.com/spf13/cobra"
	"github.com/xefino/goutils/collections"
)

// CheckCmd allows the user to verify invoice data against receivables data
func CheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check receivables data",
		Long:  "Checkes that a receivables file matches an invoices file",
		Run: func(cmd *cobra.Command, args []string) {
			errResults := make([]*types.ErrorResult, 0)

			// First, attempt to read the invoice data; if this fails, exit with an error.
			invoiceData, err := tryReadCsv[*types.InvoiceItem](params.InvoiceFile, &errResults)
			if err != nil {
				log.Fatalf("Failed to read invoice data, error: %v", err)
			}

			// Attempt to read the receivables data; if this fails, exit with an error.
			receivablesData, err := tryReadCsv[*types.ReceivablesItem](params.ReceivablesFile, &errResults)
			if err != nil {
				log.Fatalf("Failed to read receivables data, error: %v", err)
			}

			// Next, attempt to write the error file; if this fails, exit with an error.
			if err := tryWriteCsv(params.ErrorFile, errResults...); err != nil {
				log.Fatalf("Failed to write error file, error: %v", err)
			}

			// Now, check the data and record any errors that occur
			results := batch.NewCheckBatch(invoiceData, receivablesData,
				rules.InvoiceExists, rules.AmountsEqual, rules.PaidOnTime,
				rules.DateNotInFuture, rules.NotPastDue).Check()

			// Finally, write the check results to the output file; if this fails, exit with an error.
			if err := tryWriteCsv(params.OutputFile, results...); err != nil {
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

// Helper function to read a CSV file and return the data as a map
func tryReadCsv[TItem io.Keyer[string]](path string, errs *[]*types.ErrorResult) (*collections.IndexedMap[string, TItem], error) {

	// First, attempt to open the file; if this fails, return the error.
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Ensure that the file is closed when the function exits so we don't
	// leak file descriptors.
	defer file.Close()

	// Next, attempt to read the CSV file; if this fails, return the error.
	data, errResults, err := io.ReadCSV[string, TItem](file, path)
	if err != nil {
		return nil, err
	}

	// Now, append any errors that occurred to the error slice
	*errs = append(*errs, errResults...)

	// Finally, return the data
	return data, nil
}

// Helper function to write a CSV file
func tryWriteCsv[TItem any](path string, data ...TItem) error {

	// First, attempt to create the file; if this fails, return the error.
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// Next, ensure that the file is closed when the function exits so we don't
	// leak file descriptors.
	defer file.Close()

	// Finally, attempt to write the CSV file; if this fails, return the error.
	return io.WriteCsv(file, data...)
}
