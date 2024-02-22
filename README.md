# Finvoice to CSV Converter

This command line tool converts Finvoice XML files to CSV format to be imported
into [Plain Text Accounting](https://plaintextaccounting.org/).

## Usage

**Installing**

```sh
go install github.com/joneskoo/finvoice-to-csv@latest
```

To run the application, provide the directory containing the Finvoice XML files as a command-line argument:

```sh
finvoice-to-csv /path/to/directory
```

The CSV output will be written to stdout. To save the output to a file, you can redirect it:

```sh
finvoice-to-csv /path/to/directory > output.csv
```
