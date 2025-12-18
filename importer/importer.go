package importer

type Importer interface {
	Import() error
	Filter(func(Record) bool) (Records, error)
	FilterColumns() ([]string, error)
}

type Record map[string]string

type Records []Record
