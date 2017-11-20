package model

// Story ...
type Story struct {
	ID        string
	Title     string
	Summary   string
	Text      string
	Image     string
	Language  string
	Created   string
	CreatedBy string
}

// FilterStories returns all stories.
func FilterStories(filter func(*Story) bool) []*Story {
	var stories []*Story
	storyChannel := DB.All("Story")

	for storyObj := range storyChannel {
		story := storyObj.(*Story)

		if filter(story) {
			stories = append(stories, story)
		}
	}

	return stories
}

// AllStories returns all stories.
func AllStories() []*Story {
	var stories []*Story
	storyChannel := DB.All("Story")

	for storyObj := range storyChannel {
		story := storyObj.(*Story)
		stories = append(stories, story)
	}

	return stories
}
