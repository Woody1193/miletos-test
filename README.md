[![Unit Tests Status](https://github.com/Woody1193/miletos-test/actions/workflows/test.yml/badge.svg)](https://github.com/Woody1193/miletos-test/actions)

# miletos-test
This repository contains code that should serve as a solution to the case study presented to Miletos. This should not be used to inform potential candidates for future interviews, but may serve as an example of good coding practices for Go. The repository is intended to contain code for a fictional CLI that should contain a variety of invoicing tools.

## Languages
[![English](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/Woody1193/miletos-test/blob/main/README.md)
[![日本語](https://img.shields.io/badge/lang-jp-green.svg)](https://github.com/Woody1193/miletos-test/blob/main/README.jp.md)

## check command
The check command verifies that the transactions between an invoice file and a receivables file match according to several criteria. The following checks are currently supported:

- That, for a given receivable, an associated invoice exists
- That the invoiced amount matches the received amount
- That the receipt date is before the invoice due date
- That the receipt date is less than one month in the future
- That, if an invoice exists which has not been paid, that the invoice is not yet due

### Command Syntax
The syntax for the check command is as follows:

``` bash
invtools check --invoice <invoice_file> --receivables <receivables_file> --output <output_file> --error <error_file>
```

### Command Input
The check command requires two input files and two output files. If these are not provided explicitly then `invoice.csv`, `receivables.csv`, `output.csv` and `errors.csv` will be produced for the invoice file, the receivalbles file, the output file and the errors file, respectively.

The `invoice_file` parameter specifies the file path to the invoice file. This file should be a CSV file that contains the following columns:

- **ID (string):** A unique identifier for the invoice.
- **Amount (integer):** The amount of the invoice in JPY.
- **Due Date (date):** The date by which the invoice must be paid, in `YYYY-MM-DDThh:mm:ssZ` format.

The `receivables_file` parameter specifies the file path to the receivables file. This file should be a CSV file that contains the following columns:

- **ID (string):** A unique identifier for the receivable.
- **Amount (integer):** The amount received in JPY.
- **Date (date):** The date on which the amount was received, in `YYYY-MM-DDThh:mm:ssZ` format.

### Command Output
The check command will not produce any output if there were no errors. The `output_file` will contain all instances where any rule check failed.

The `output_file` will include the following information:

- **ID (string):** The ID of the invoice for which an error occurred
- **Invoices File Line (integer):** The line in the invoices file corresponding to this error
- **Receivables File Line (integer):** The line in the receivables file corresponding to this error
- **Description (string):** A description of the issue(s) that occurred. There may be more than one.

The `error_file` will also be produced by this command and shall contain information on data that could not be handled and why. If data is written to this file then it will not be checked by the command, as doing so would be impossible. Data will be written to this file under the following conditions:

- If the invoice ID is missing
- If the amount is not an integer value
- If the due date or date cannot be parsed according to the format specified
- If more than one item can be found for the same invoice ID. Note that, in this case, the first occurrence of the invoice ID will be accepted and all subsequent occurrences will be rejected.

The `error_file` will include the following information:

- **File (string):** The name of the file that contained the bad data
- **Line (integer):** The line containing the error
- **Description (string):** A description of the error

### Error Conditions
The check command may generate errors under the following conditions:

- If either the `invoice_file` or the `receivables_file` parameter is not specified or the file does not exist, an error message will be generated.
- If the CSV files are not in the correct format, an error message will be generated.
- If either the `output_file` or `error_file` could not be written, then an error message will be generated.