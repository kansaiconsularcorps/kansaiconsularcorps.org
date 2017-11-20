package model

// Page represents a single markdown based page.
type Page struct {
	ID string

	// Key is the ISO 639-1 language code, value is the text translation in markdown
	Title map[string]string
	Text  map[string]string
}
