# BCCify

A command-line tool written in Go that helps you create email recipient lists in BCC format from CSV files.

## Overview

BCCify imports contact data from CSV files and allows you to filter recipients based on custom criteria. It's designed to help you quickly generate BCC recipient lists for bulk email sending.

## Features

- **CSV Import**: Load contacts from CSV files with support for headers
- **Flexible Filtering**: Filter recipients using custom filter functions
- **Recipient Formatting**: Automatically formats contacts as `Name <email@example.com>`
- **Multiple Export Formats**: Output recipients to various formats (currently stdout)

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

### Output

The tool outputs all recipients in BCC format, separated by semicolons:

```bash
bccify contacts.csv
```

Output example:

```
John Doe <john@example.com>;Jane Smith <jane@example.com>;Jack Wilson <jack@example.com>
```

## Project Structure

```
bccify/
├── cmd/
│   └── main.go            # Application entry point
├── bccify.go              # Core BCCify logic and converter
├── importer/
│   ├── importer.go        # Importer interface definition
│   └── csv.go             # CSV importer implementation
├── models/
│   ├── recipient.go       # Recipient model
│   └── recipients.go      # Recipients collection and methods
├── exporter/
│   ├── exporter.go        # Exporter interface definition
│   └── stdout.go          # Stdout exporter implementation
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

### Exporter Interface

The `Exporter` interface defines the contract for data exporters:

```go
type Exporter interface {
    Export() error
}
```

### Models

**Recipient**: Represents a single email recipient with name and email address.

**Recipients**: A collection of recipients formatted as strings (`Name <email@example.com>`).

### Example Usage

```go
// Import from CSV
importer := &importer.CSVImporter{
    Filename:   "contacts.csv",
    HasHeaders: true,
}
err := importer.Import()

// Get all records (or filter them)
records := importer.Filter(nil)

// Convert to recipients
converter := RecipientConverter{Records: records}
converter.Convert()

// Export to stdout
exporter := exporter.StdOutExporter{Recipents: converter.Recipents}
exporter.Export()
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

- Support for additional import formats (Excel, vCard)
- Support for additional export formats (file output, email clients)
- Interactive CLI for building custom filters
- Email validation and deduplication
- Custom column mapping configuration
