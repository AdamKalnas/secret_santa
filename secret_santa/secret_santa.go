package secret_santa

import (
	"math/rand"
	"time"
)

func GetMatches(participants []string) map[string]string {
	matches := make(map[string]string)
	if len(participants) > 1 {
		return repeatUntilValid(participants)
	}
	return matches
}

func repeatUntilValid(originals []string) (matches map[string]string) {
	for {
		matches = obtainMatches(originals)
		if !containsReflexiveMatches(matches) {
			return matches
		}
	}
}

func containsReflexiveMatches(matches map[string]string) bool {
	for key, value := range matches {
		if key == value {
			return true
		}
	}
	return false
}

func obtainMatches(originals []string) map[string]string {
	matches := make(map[string]string)
	shuffled := shuffle(originals)
	for i, original := range originals {
		matches[original] = shuffled[i]
	}
	return matches
}

func shuffle(original []string) []string {
	shuffled := make([]string, len(original))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(original))
	for i, v := range perm {
		shuffled[v] = original[i]
	}
	return shuffled
}
