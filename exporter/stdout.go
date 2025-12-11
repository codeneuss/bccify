package exporter

import (
	"fmt"
	"strings"

	"github.com/codeneuss/bccify/models"
)

type StdOutExporter struct {
	Recipients models.Recipents
}

func (e StdOutExporter) Export() error {
	fmt.Println(strings.Join(e.Recipients, ";"))

	return nil
}
