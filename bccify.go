package bccify

import (
	"fmt"
	"os"
	"strings"

	ex "github.com/codeneuss/bccify/exporter"
	im "github.com/codeneuss/bccify/importer"
	mo "github.com/codeneuss/bccify/models"
)

func BCCify() {
	fmt.Println("Welcome to BCCify!")
	args := os.Args

	if len(args) < 1 {
		fmt.Println("You need to provide a file as an argument.")
		os.Exit(1)
	}

	fmt.Println(args[1])

	var importer im.Importer
	switch {
	case strings.HasSuffix(args[1], ".csv"):
		importer = &im.CSVImporter{
			Filename:   args[1],
			HasHeaders: true,
		}
	}

	if importer == nil {
		fmt.Println("No Importer found")
		os.Exit(1)
	}

	if err := importer.Import(); err != nil {
		fmt.Println("Error on import: ", err.Error())
		os.Exit(1)
	}

	converter := RecipientConverter{
		Records: importer.Filter(nil),
	}

	if err := converter.Convert(); err != nil {
		fmt.Println("Error converting Recipients", err.Error())
		os.Exit(2)
	}

	exporter := ex.StdOutExporter{Recipents: converter.Recipents}

	if err := exporter.Export(); err != nil {
		fmt.Println("Error exporting recipients", err.Error())
		os.Exit(2)
	}

}

type RecipientConverter struct {
	Records   im.Records
	Recipents mo.Recipents
}

func (t *RecipientConverter) Convert() error {
	for _, rec := range t.Records {
		t.Recipents.Add(mo.NewRecipient(rec["emailaddress"], rec["name"]))
	}

	return nil
}
