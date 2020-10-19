package exampleAPI

type DocType int

const (
	DOCUMENT_TYPE_SOURCES DocType = iota
	DOCUMENT_TYPE_OTHERS
)

func (d DocType) String() string {
	return [...]string{
		"Sources",
		"Add others as required...",
	}[d]
}

func (d DocType) Topic() string {
	return [...]string{
		"sourcesExampleTopic",
		"Add others as required...",
	}[d]
}
