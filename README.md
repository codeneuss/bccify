# BCCify

A command-line tool written in Go that helps you create email recipient lists in BCC format from CSV files.

## Overview

BCCify imports contact data from CSV files and allows you to filter recipients based on custom criteria. It's designed to help you quickly generate BCC recipient lists for bulk email sending.

## Features

- **CSV Import**: Load contacts from CSV files with support for headers
- **Flexible Filtering**: Filter recipients using custom filter functions
- **Recipient Formatting**: Automatically formats contacts as `"Name" <email@example.com>`
- **Multiple Export Formats**:
  - **MailTo Export** (default): Opens your default email client with recipients in BCC field
  - **StdOut Export**: Outputs recipients to stdout for manual copying
- **Multi-Email Support**: Handles contacts with multiple email addresses separated by semicolons

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

By default, BCCify opens your default email client with all recipients pre-filled in the BCC field:

```bash
bccify contacts.csv
```

This creates a mailto link and opens it in your system's default email application with all contacts ready to send.

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
│   ├── mailto.go          # MailTo exporter implementation (opens email client)
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

// Export using MailTo (opens email client)
exporter := exporter.MailToExporter{Recipients: converter.Recipents}
exporter.Export()

// Or export to stdout
stdoutExporter := exporter.StdOutExporter{Recipients: converter.Recipents}
stdoutExporter.Export()
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

## How It Works

1. **Import**: BCCify reads your CSV file and parses the contacts
2. **Convert**: Contacts are converted to the proper email format: `"Name" <email@example.com>`
3. **Export**: By default, creates a mailto URL with all recipients in the BCC field and opens your default email client

### Export Options

BCCify currently supports two export methods:

- **MailToExporter** (default): Generates a mailto link and opens it using the system's `open` command (macOS/Linux)
- **StdOutExporter**: Prints the recipient list to stdout, semicolon-separated

## Roadmap

Future improvements planned:

- Support for additional import formats (Excel, vCard)
- Cross-platform support for mailto export (Windows compatibility)
- Interactive CLI for building custom filters
- Email validation and deduplication
- Custom column mapping configuration
- Command-line flags to choose export format
