package importer

type Importer interface {
	Import() error
	Filter(func(Record) bool) Records
	FilterColumns() []string
}

type Record map[string]string

type Records []Record
