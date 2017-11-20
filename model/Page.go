package model

// Page represents a single markdown based page.
type Page struct {
	ID string

	// Key is the ISO 639-3 language code, value is the text translation in markdown
	Text map[string]string
}
