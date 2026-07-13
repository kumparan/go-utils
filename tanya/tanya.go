package tanya

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsQuestion returns true if a query is a question
func IsQuestion(q string) bool {
	intent := ClassifyIntent(q)
	switch intent { // nolint:exhaustive
	case IntentPrice, IntentOther:
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
	case MatchTypeTokenSuffix:
		minLen := r.MinTokenLen
		if minLen <= 0 {
			minLen = defaultMinTokenLen
		}
		for _, tok := range tokenize(q) {
			if len(tok) < minLen {
				continue
			}
			for _, suf := range r.Terms {
				if strings.HasSuffix(tok, suf) {
					return true
				}
			}
		}
	}

	return false
}

func normalize(s string) string {
	s = strings.ToLower(strings.Join(strings.Fields(s), " "))
	return expandAbbreviations(" " + s + " ")
}

func expandAbbreviations(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if repl, ok := abbrevMap[w]; ok {
			words[i] = repl
			continue
		}
		// handle punctuation like "knp?" or "dmn," etc.
		base := strings.TrimRight(w, "?.!,")
		suffix := w[len(base):]
		if repl, ok := abbrevMap[base]; ok {
			words[i] = repl + suffix
		}
	}
	return strings.Join(words, " ")
}

// tokenize splits on whitespace and trims leading/trailing non-letters/digits per token.
// keeps tokens simple & fast (no regex).
func tokenize(s string) []string {
	raw := strings.Fields(s)
	out := make([]string, 0, len(raw))
	for _, t := range raw {
		t = trimNonAlphaNum(t)
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}

func trimNonAlphaNum(s string) string {
	start, end := 0, len(s)
	for start < end {
		r := rune(s[start])
		if isAlphaNum(r) {
			break
		}
		_, w := utf8.DecodeRuneInString(s[start:])
		start += w
	}
	for end > start {
		r, w := utf8.DecodeLastRuneInString(s[:end])
		if isAlphaNum(r) {
			break
		}
		end -= w
	}
	return s[start:end]
}

func isAlphaNum(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) }
