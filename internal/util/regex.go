package util

import "regexp"

func FindAllMatches(body []byte, re *regexp.Regexp) []string {
	return re.FindAllString(string(body), -1)
}
