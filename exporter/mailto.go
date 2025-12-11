package exporter

import (
	"net/url"
	"os/exec"
	"strings"

	"github.com/codeneuss/bccify/models"
)

type MailToExporter struct {
	Recipients models.Recipents
}

func (e *MailToExporter) Export() error {
	params := url.Values{}

	params.Add("bcc", strings.Join(e.Recipients, ","))
	mailtoURL := "mailto:?" + strings.Replace(params.Encode(), "+", "%20", -1)

	cmd := exec.Command("open", mailtoURL)
	return cmd.Run()
}
