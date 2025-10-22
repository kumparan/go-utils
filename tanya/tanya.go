package tanya

import (
	"sort"
	"strings"
	"unicode"
)

// IsQuestion returns true if a query is a question
func IsQuestion(q string) bool {
	intent := ClassifyIntent(q)
	switch intent { // nolint:exhaustive
	case IntentPrice, IntentContact, IntentOther:
		return false
	default:
		return true
	}
}

// ClassifyIntent returns the most likely intent for the given query
func ClassifyIntent(q string) Intent {
	q = normalize(q)
	if q == "" {
		return IntentOther
	}
	type scored struct {
		intent      Intent
		score, prio int
	}
	var candidates []scored

	for _, spec := range intentTable {
		score := 0
		for _, r := range spec.Rules {
			if matchByType(q, r) {
				score += r.Weight
			}
		}
		if score != 0 {
			candidates = append(candidates, scored{spec.Intent, score, spec.Priority})
		}
	}
	if len(candidates) == 0 {
		return IntentOther
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].score == candidates[j].score {
			return candidates[i].prio > candidates[j].prio
		}
		return candidates[i].score > candidates[j].score
	})
	return candidates[0].intent
}

func matchByType(q string, r Rule) bool {
	switch r.MatchType {
	case MatchTypeContains:
		for _, t := range r.Terms {
			if strings.Contains(q, t) {
				return true
			}
		}
	case MatchTypeStarts:
		for _, t := range r.Terms {
			if strings.HasPrefix(q, t) {
				return true
			}
		}
	case MatchTypeEnds:
		for _, t := range r.Terms {
			if strings.HasSuffix(q, t) {
				return true
			}
		}
	}
	return false
}

func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(collapseSpaces(s)))
	s = " " + s + " "
	for k, v := range abbrevMap {
		s = strings.ReplaceAll(s, k, v)
	}
	return strings.TrimSpace(collapseSpaces(s))
}

func collapseSpaces(s string) string {
	var b strings.Builder
	sp := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			if !sp {
				b.WriteByte(' ')
				sp = true
			}
		} else {
			b.WriteRune(r)
			sp = false
		}
	}
	return strings.TrimSpace(b.String())
}
