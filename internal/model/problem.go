package model

type Problem struct {
	Index     int
	ID        string
	Title     string
	URL       string
	Statement ProblemStatement
}

type ProblemStatement struct {
	Description   string
	SampleDataset string
	SampleOutput  string
}
