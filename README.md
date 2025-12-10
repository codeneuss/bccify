# BCCify

A command-line tool written in Go that helps you create email recipient lists in BCC format from CSV files.

## Overview

BCCify imports contact data from CSV files and allows you to filter recipients based on custom criteria. It's designed to help you quickly generate BCC recipient lists for bulk email sending.

## Features

- **CSV Import**: Load contacts from CSV files with support for headers
- **Flexible Filtering**: Filter recipients using custom filter functions
- **Recipient Formatting**: Automatically formats contacts as `Name <email@example.com>`
- **Column Discovery**: View available columns from your CSV data

## Installation

```bash
go install github.com/codeneuss/bccify@latest
```

Or clone and build from source:

```bash
git clone https://github.com/codeneuss/bccify.git
cd bccify
go build
```

## Usage

Run BCCify by providing a CSV file as an argument:

```bash
bccify contacts.csv
```

### CSV File Format

Your CSV file should include headers in the first row. For example:

```csv
name,emailaddress
John Doe,john@example.com
Jane Smith,jane@example.com
Jack Wilson,jack@example.com
```

### Example

The current implementation includes a demo filter that selects contacts whose names start with "J":

```bash
bccify contacts.csv
```

This will output the email addresses of all contacts matching the filter criteria.

## Project Structure

```
bccify/
├── main.go                 # Main application entry point
├── importer/
│   ├── importer.go        # Importer interface definition
│   └── csv.go             # CSV importer implementation
└── README.md
```

## API

### Importer Interface

The `Importer` interface defines the contract for data importers:

```go
type Importer interface {
    Import() error
    Filter(func(Record) bool) Records
    FilterColumns() []string
}
```

### CSVImporter

Import CSV files with optional header support:

```go
imp := &importer.CSVImporter{
    Filename:   "contacts.csv",
    HasHeaders: true,
}
```

### Filtering Records

Filter contacts using custom filter functions:

```go
filtered := imp.Filter(func(r importer.Record) bool {
    return strings.HasPrefix(r["name"], "J")
})
```

## Development

### Requirements

- Go 1.25.5 or later

### Building

```bash
go build
```

### Testing

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).

## Roadmap

Future improvements planned:

- Support for additional file formats (Excel, vCard)
- Interactive CLI for building filters
- Output to various email client formats
- Email validation and deduplication
