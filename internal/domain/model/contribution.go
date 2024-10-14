package model

type Contribution struct {
	Year int

	IssueCount  int
	PRCount     int
	ReviewCount int

	// Removed because this field might be a huge number than others
	// CommitCount int
}
