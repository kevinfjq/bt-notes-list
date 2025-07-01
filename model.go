package main

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state uint
	//store Store
	//textarea.Model
	// ... other fields as needed
}
