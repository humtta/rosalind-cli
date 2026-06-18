package model

type Problem struct {
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
