package arena

import "github.com/fatih/color"

// Cleaner binds a callback function and a label to help us to identify the tasks that are executed during the cleaning stack
type Cleaner struct {
	name     string
	callback func(bool)
}

var cleanerStack []Cleaner

// RegisterCleaner creates a stack of call back functions to be executed when the main process ends
func RegisterCleaner(name string, callback func(bool)) {
	if cleanerStack == nil {
		cleanerStack = make([]Cleaner, 0)
	}
	cleaner := Cleaner{name, callback}
	cleanerStack = append([]Cleaner{cleaner}, cleanerStack...)
}

// Cleanup executes all cleaner functions in the stack passing the arg that tells if the process was interrupted or not
func Cleanup(interrupted bool) {
	LogNotice("Cleaning up")
	for _, cleaner := range cleanerStack {
		LogInfo("Closing: %s", cleaner.name)
		cleaner.callback(interrupted)
	}
	LogNotice("-- All cleaners done -- ")
	color.Unset()
}
