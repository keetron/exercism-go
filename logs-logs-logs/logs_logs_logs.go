package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var apps = map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, c := range log {
		r := apps[c]
		if r != "" {
			return r
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	r := ""
	for _, c := range log {
		if c == oldRune {
			r = r + string(newRune)
		} else {
			r = r + string(c)
		}
	}

	return r
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	}
	return false
}
