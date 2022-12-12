package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC]|^\[DBG]|^\[INF]|^\[WRN]|^\[ERR]|^\[FTL]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	c := 0
	for _, l := range lines {
		if re.MatchString(l) {
			c = c + 1
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+\w+`)
	for i, l := range lines {
		if re.MatchString(l) {
			n := regexp.MustCompile(`\s+`).Split(re.FindString(l), -1)[1]
			lines[i] = "[USR] " + n + " " + l
		}
	}
	return lines
}
