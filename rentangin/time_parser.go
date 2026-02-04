package rentangin

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Range is [Start, End) (End exclusive).
type Range struct {
	Start time.Time
	End   time.Time
}

// IsZero returns true if empty
func (r Range) IsZero() bool { return r.Start.IsZero() && r.End.IsZero() }

// Parse scans query and returns the BEST date range it can extract.
// ok=false means "no range found" (not an error).
//
// Timezone is taken from now.Location(). So pass now in the timezone you want.
func Parse(query string, now time.Time) (r Range, isTimeRage bool, err error) {
	s := normalizeID(strings.TrimSpace(query))
	if s == "" {
		return Range{}, false, nil
	}

	hasEventHint := containsEventHint(s)
	topicIntent := containsTopicIntent(s)
	nowYear := now.Year()

	bestScore := -1
	var best Range

	// Scan by token boundary (space).
	for i := 0; i < len(s); i++ {
		if i != 0 && s[i-1] != ' ' {
			continue
		}
		sub := s[i:]
		prev := prevToken(s, i)

		cand, score, found := parseBestAtStart(sub, now, prev, hasEventHint, topicIntent, nowYear)
		if !found {
			continue
		}
		if score > bestScore {
			bestScore = score
			best = cand
		}
	}

	if bestScore < 0 {
		return Range{}, false, nil
	}
	return best, true, nil
}

/* -------------------------
   Core parsing at start
--------------------------*/

func parseBestAtStart(s string, now time.Time, prev string, hasEventHint bool, topicIntent bool, nowYear int) (Range, int, bool) {
	// Highest: explicit "dari ... sampai ..."
	if r, ok := parseFromRangeAtStart(s, now, hasEventHint, topicIntent, nowYear); ok {
		return r, scoreFromRange, true
	}
	// Next: inline "A sampai B" / "A - B" / "A sd B"
	if r, ok := parseInlineRangeAtStart(s, now, hasEventHint, topicIntent, nowYear); ok {
		return r, scoreInlineRange, true
	}
	// Single expression with specificity scoring
	if r, score, ok := parseOneExprFromStart(s, now, prev, hasEventHint, topicIntent, nowYear); ok {
		return r, score, true
	}
	return Range{}, -1, false
}

func parseFromRangeAtStart(s string, now time.Time, hasEventHint bool, topicIntent bool, nowYear int) (Range, bool) {
	m := rxFromRange.FindStringSubmatchIndex(s)
	if m == nil || m[0] != 0 {
		return Range{}, false
	}
	startExpr := strings.TrimSpace(s[m[2]:m[3]])
	endExpr := strings.TrimSpace(s[m[6]:m[7]])

	rs, _, ok := parseOneExprAny(startExpr, now, "", hasEventHint, topicIntent, nowYear)
	if !ok {
		return Range{}, false
	}
	re, _, ok := parseOneExprAny(endExpr, now, "", hasEventHint, topicIntent, nowYear)
	if !ok {
		return Range{}, false
	}

	r := Range{Start: rs.Start, End: re.Start}
	if !r.End.After(r.Start) {
		// Fallback to start only if user input is weird.
		return rs, true
	}
	return r, true
}

func parseInlineRangeAtStart(s string, now time.Time, hasEventHint bool, topicIntent bool, nowYear int) (Range, bool) {
	m := rxInlineRange.FindStringSubmatchIndex(s)
	if m == nil || m[0] != 0 {
		return Range{}, false
	}
	left := strings.TrimSpace(s[m[2]:m[3]])
	right := strings.TrimSpace(s[m[6]:m[7]])

	rl, _, ok := parseOneExprAny(left, now, "", hasEventHint, topicIntent, nowYear)
	if !ok {
		return Range{}, false
	}
	rr, _, ok := parseOneExprAny(right, now, "", hasEventHint, topicIntent, nowYear)
	if !ok {
		return Range{}, false
	}

	r := Range{Start: rl.Start, End: rr.Start}
	if !r.End.After(r.Start) {
		return rl, true
	}
	return r, true
}

func prevToken(s string, start int) string {
	// start is token boundary index in s
	if start <= 0 {
		return ""
	}
	j := start - 1
	for j >= 0 && s[j] == ' ' {
		j--
	}
	if j < 0 {
		return ""
	}
	k := j
	for k >= 0 && s[k] != ' ' {
		k--
	}
	return s[k+1 : j+1]
}

/* -------------------------
   Normalization
--------------------------*/

func normalizeID(s string) string {
	ls := strings.ToLower(s)

	// connectors
	ls = strings.ReplaceAll(ls, "s.d.", "sd")
	ls = strings.ReplaceAll(ls, "s.d", "sd")
	ls = strings.ReplaceAll(ls, "s / d", "sd")
	ls = strings.ReplaceAll(ls, "s/d", "sd")

	// multi-token phrases
	ls = strings.ReplaceAll(ls, "hari ini", "hariini")
	ls = strings.ReplaceAll(ls, "yang lalu", "yanglalu")
	ls = strings.ReplaceAll(ls, "ke depan", "kedepan")

	// relative anchors
	ls = strings.ReplaceAll(ls, "dari sekarang", "darisekarang")
	ls = strings.ReplaceAll(ls, "dari hariini", "darihariini")

	ls = strings.Join(strings.Fields(ls), " ")
	return ls
}

/* -------------------------
   One expression parsing + scoring
--------------------------*/

const (
	scoreFromRange   = 100
	scoreInlineRange = 90

	scoreLastNRange = 85 // "7 hari terakhir"

	scoreDay       = 80 // day-specific (ymd, dmy, dm past-biased, edge month day)
	scoreMonth     = 70 // month-specific (my/ym/month-only)
	scoreYear      = 60 // "tahun 2024"
	scoreRelNumber = 50 // "3 hari lalu" (single day)
	scoreRelative  = 40 // hariini/kemarin/besok, unit modifiers
)

func parseOneExprAny(expr string, now time.Time, prev string, hasEventHint bool, topicIntent bool, nowYear int) (Range, int, bool) {
	e := normalizeID(strings.TrimSpace(expr))
	return parseOneExprFromStart(e, now, prev, hasEventHint, topicIntent, nowYear)
}

// parseOneExprFromStart parses if expression begins at start of s.
// Ambiguous expressions default to PAST (latest date <= now).
func parseOneExprFromStart(s string, now time.Time, prev string, hasEventHint bool, topicIntent bool, nowYear int) (Range, int, bool) {
	// 0) "awal bulan ini" / "akhir bulan ini" (day-range)
	if m := rxEdgeMonthThis.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		edge := s[m[2]:m[3]] // "awal"|"akhir"
		if r, ok := edgeMonthThis(edge, now); ok {
			return r, scoreDay, true
		}
	}

	// 1) "7 hari terakhir" / "N hari terakhir" (range)
	if m := rxLastNUnit.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		n, _ := strconv.Atoi(s[m[2]:m[3]])
		unit := s[m[4]:m[5]]
		if r, ok := lastN(unit, n, now); ok {
			return r, scoreLastNRange, true
		}
	}

	// 2) One-word relative
	if m := rxOneWord.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		w := s[m[2]:m[3]]
		if r, ok := oneWord(w, now); ok {
			return r, scoreRelative, true
		}
	}

	// 3) Unit modifier: minggu/bulan/tahun ini/lalu/depan
	if m := rxUnitModifier.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		unit := s[m[2]:m[3]]
		mod := s[m[4]:m[5]]
		if r, ok := unitModifier(unit, mod, now); ok {
			return r, scoreRelative, true
		}
	}

	// 4) Relative numeric (single day/week/month/year anchor)
	if m := rxRelNUnit.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		n, _ := strconv.Atoi(s[m[2]:m[3]])
		unit := s[m[4]:m[5]]
		rel := s[m[6]:m[7]]
		if r, ok := relN(n, unit, rel, now); ok {
			return r, scoreRelNumber, true
		}
	}

	// 5) "tahun 2024"
	if m := rxTahunYYYY.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		y, _ := strconv.Atoi(s[m[2]:m[3]])
		if okYear(y) {
			return yearRange(y, now.Location()), scoreYear, true
		}
		return Range{}, -1, false
	}

	// 5.5) Bare year: "1998" (only if event-hint and not topic/title context and not future-year)
	if m := rxBareYear.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		y, _ := strconv.Atoi(s[m[2]:m[3]])
		next := nextTokenAfterPrefix(s[m[1]:])
		if okYear(y) && allowBareYear(prev, next, hasEventHint, topicIntent, y, nowYear) {
			return yearRange(y, now.Location()), scoreYear - 1, true
		}
	}

	// 6) Numeric YMD: 2026-02-04 / 2026/02/04
	if m := rxYMDNumeric.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		yr, _ := strconv.Atoi(s[m[2]:m[3]])
		moN, _ := strconv.Atoi(s[m[4]:m[5]])
		day, _ := strconv.Atoi(s[m[6]:m[7]])
		if okYear(yr) && 1 <= moN && moN <= 12 && okDay(day) {
			start := time.Date(yr, time.Month(moN), day, 0, 0, 0, 0, now.Location())
			return dayRange(start), scoreDay, true
		}
		return Range{}, -1, false
	}

	// 7) D M Y (month name): "4 feb 2026"
	if m := rxDMY.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		day, _ := strconv.Atoi(s[m[2]:m[3]])
		mon := s[m[4]:m[5]]
		yr, _ := strconv.Atoi(s[m[6]:m[7]])
		mo, ok := monthID(mon)
		if ok && okDay(day) && okYear(yr) {
			start := time.Date(yr, mo, day, 0, 0, 0, 0, now.Location())
			return dayRange(start), scoreDay, true
		}
		return Range{}, -1, false
	}

	// 8) M Y: "februari 2026"
	if m := rxMY.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		mon := s[m[2]:m[3]]
		yr, _ := strconv.Atoi(s[m[4]:m[5]])
		mo, ok := monthID(mon)
		if ok && okYear(yr) {
			start := time.Date(yr, mo, 1, 0, 0, 0, 0, now.Location())
			return monthRange(start), scoreMonth, true
		}
		return Range{}, -1, false
	}

	// 9) Y M: "2026 feb"
	if m := rxYM.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		yr, _ := strconv.Atoi(s[m[2]:m[3]])
		mon := s[m[4]:m[5]]
		mo, ok := monthID(mon)
		if ok && okYear(yr) {
			start := time.Date(yr, mo, 1, 0, 0, 0, 0, now.Location())
			return monthRange(start), scoreMonth, true
		}
		return Range{}, -1, false
	}

	// 10) D M (no year): "15 maret" => past-biased (latest <= now)
	if m := rxDM.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		day, _ := strconv.Atoi(s[m[2]:m[3]])
		mon := s[m[4]:m[5]]
		if mo, ok := monthID(mon); ok && okDay(day) {
			if r, ok := dayMonthPast(day, mo, now); ok {
				return r, scoreDay, true
			}
		}
	}

	// 11) Month only: "februari" => past-biased (latest <= now)
	if m := rxMonthOnly.FindStringSubmatchIndex(s); m != nil && m[0] == 0 {
		mon := s[m[2]:m[3]]
		if mo, ok := monthID(mon); ok {
			return monthOnlyPast(mo, now), scoreMonth, true
		}
	}

	return Range{}, -1, false
}

func nextTokenAfterPrefix(rest string) string {
	rest = strings.TrimSpace(rest)
	if rest == "" {
		return ""
	}
	if i := strings.IndexByte(rest, ' '); i >= 0 {
		return rest[:i]
	}
	return rest
}

func allowBareYear(prev, next string, hasEventHint bool, topicIntent bool, year int, nowYear int) bool {
	// Topic-intent: year is likely a theme, not a published_at filter.
	if topicIntent {
		return false
	}
	// Future year: often a target topic/prediction, not document time.
	if year > nowYear {
		return false
	}
	// Must have event/time hint (to avoid false positives).
	if !hasEventHint {
		return false
	}
	// Title/entity contexts (film, review, trailer, etc).
	if isEntityBlocker(prev) || isEntityBlocker(next) {
		return false
	}
	return true
}

func isEntityBlocker(tok string) bool {
	switch tok {
	case "film", "movie", "series", "serial", "drama", "anime",
		"album", "lagu", "song", "buku", "novel", "game",
		"review", "ulasan", "sinopsis", "trailer", "subtitle":
		return true
	default:
		return false
	}
}

func containsEventHint(q string) bool {
	// q already normalized & lowercase
	eventPhrases := []string{
		"piala dunia",
		"world cup",
	}
	for _, ph := range eventPhrases {
		if strings.Contains(q, ph) {
			return true
		}
	}
	eventTokens := []string{
		"pemilu", "pilpres", "pilkada",
		"oscar", "grammy",
		"olimpiade", "olympic", "olympics",
		"liga", "turnamen", "juara", "final", "pemenang",
		"gempa", "banjir", "krisis", "inflasi", "resesi",
		"piala", "musim",
	}
	for _, t := range eventTokens {
		if strings.Contains(q, t) {
			return true
		}
	}
	return false
}

func containsTopicIntent(q string) bool {
	// q already normalized & lowercase
	tokens := []string{
		"prediksi", "ramalan", "forecast", "proyeksi", "outlook",
		"tren", "trend", "gaya", "model", "inspirasi",
		"rekomendasi", "tips", "tip", "cara", "panduan", "tutorial",
		"review", "ulasan", "sinopsis", "trailer",
	}
	for _, t := range tokens {
		if strings.Contains(q, t) {
			return true
		}
	}
	return false
}

/* -------------------------
   Semantics
--------------------------*/

// "awal bulan ini" -> day-range for first day of current month
// "akhir bulan ini" -> day-range for last day of current month
func edgeMonthThis(edge string, now time.Time) (Range, bool) {
	loc := now.Location()
	y, m, _ := now.Date()
	first := time.Date(y, m, 1, 0, 0, 0, 0, loc)

	switch edge {
	case "awal":
		return dayRange(first), true
	case "akhir":
		nextMonth := first.AddDate(0, 1, 0)
		lastDay := nextMonth.AddDate(0, 0, -1)
		return dayRange(lastDay), true
	default:
		return Range{}, false
	}
}

// "N hari terakhir" => range including today, ending tomorrow 00:00
// Start = today 00:00 - (n-1) days
// End   = tomorrow 00:00
func lastN(unit string, n int, now time.Time) (Range, bool) {
	if n <= 0 {
		return Range{}, false
	}
	end := truncateDay(now).AddDate(0, 0, 1)

	switch unit {
	case "hari":
		start := truncateDay(now).AddDate(0, 0, -(n - 1))
		return Range{Start: start, End: end}, true
	case "minggu", "pekan":
		start := weekRange(now).Start.AddDate(0, 0, -7*(n-1))
		return Range{Start: start, End: end}, true
	case "bulan":
		start := monthRange(now).Start.AddDate(0, -(n - 1), 0)
		return Range{Start: start, End: end}, true
	case "tahun":
		start := yearRange(now.Year(), now.Location()).Start.AddDate(-(n - 1), 0, 0)
		return Range{Start: start, End: end}, true
	default:
		return Range{}, false
	}
}

func oneWord(w string, now time.Time) (Range, bool) {
	switch w {
	case "sekarang":
		return Range{Start: now, End: now.Add(time.Second)}, true
	case "hariini":
		return dayRange(now), true
	case "kemarin":
		return dayRange(now.AddDate(0, 0, -1)), true
	case "besok":
		return dayRange(now.AddDate(0, 0, 1)), true
	default:
		return Range{}, false
	}
}

func unitModifier(unit, mod string, now time.Time) (Range, bool) {
	switch unit {
	case "minggu", "pekan":
		switch mod {
		case "ini":
			return weekRange(now), true
		case "lalu":
			return weekRange(now.AddDate(0, 0, -7)), true
		case "depan":
			return weekRange(now.AddDate(0, 0, 7)), true
		}
	case "bulan":
		switch mod {
		case "ini":
			return monthRange(now), true
		case "lalu":
			return monthRange(now.AddDate(0, -1, 0)), true
		case "depan":
			return monthRange(now.AddDate(0, 1, 0)), true
		}
	case "tahun":
		switch mod {
		case "ini":
			return yearRange(now.Year(), now.Location()), true
		case "lalu":
			return yearRange(now.Year()-1, now.Location()), true
		case "depan":
			return yearRange(now.Year()+1, now.Location()), true
		}
	}
	return Range{}, false
}

func relN(n int, unit, rel string, now time.Time) (Range, bool) {
	if n <= 0 {
		return Range{}, false
	}
	switch unit {
	case "hari":
		switch rel {
		case "lalu", "yanglalu":
			return dayRange(now.AddDate(0, 0, -n)), true
		case "lagi", "kedepan", "darisekarang", "darihariini":
			return dayRange(now.AddDate(0, 0, n)), true
		}
	case "minggu", "pekan":
		switch rel {
		case "lalu", "yanglalu":
			return weekRange(now.AddDate(0, 0, -7*n)), true
		case "lagi", "kedepan", "darisekarang", "darihariini":
			return weekRange(now.AddDate(0, 0, 7*n)), true
		}
	case "bulan":
		switch rel {
		case "lalu", "yanglalu":
			return monthRange(now.AddDate(0, -n, 0)), true
		case "lagi", "kedepan", "darisekarang", "darihariini":
			return monthRange(now.AddDate(0, n, 0)), true
		}
	case "tahun":
		switch rel {
		case "lalu", "yanglalu":
			return yearRange(now.Year()-n, now.Location()), true
		case "lagi", "kedepan", "darisekarang", "darihariini":
			return yearRange(now.Year()+n, now.Location()), true
		}
	}
	return Range{}, false
}

/* -------------------------
   Past-biased ambiguity helpers
--------------------------*/

// Month only like "februari" => latest month start <= now (past-biased)
func monthOnlyPast(month time.Month, now time.Time) Range {
	loc := now.Location()
	y := now.Year()
	candidate := time.Date(y, month, 1, 0, 0, 0, 0, loc)
	if candidate.After(truncateDay(now)) {
		candidate = candidate.AddDate(-1, 0, 0)
	}
	return monthRange(candidate)
}

// Day+month without year like "15 maret" => latest date <= now (past-biased)
func dayMonthPast(day int, month time.Month, now time.Time) (Range, bool) {
	if !okDay(day) {
		return Range{}, false
	}

	loc := now.Location()
	y := now.Year()
	candidate := time.Date(y, month, day, 0, 0, 0, 0, loc)

	// If invalid date (e.g. 31 februari), time.Date rolls over; detect by checking Month/Day.
	if candidate.Month() != month || candidate.Day() != day {
		return Range{}, false
	}

	if candidate.After(truncateDay(now)) {
		candidate = candidate.AddDate(-1, 0, 0)
		if candidate.Month() != month || candidate.Day() != day {
			return Range{}, false
		}
	}

	return dayRange(candidate), true
}

/* -------------------------
   Range constructors
--------------------------*/

func truncateDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func dayRange(ref time.Time) Range {
	s := truncateDay(ref)
	return Range{Start: s, End: s.AddDate(0, 0, 1)}
}

func monthRange(ref time.Time) Range {
	s := time.Date(ref.Year(), ref.Month(), 1, 0, 0, 0, 0, ref.Location())
	return Range{Start: s, End: s.AddDate(0, 1, 0)}
}

func yearRange(y int, loc *time.Location) Range {
	s := time.Date(y, 1, 1, 0, 0, 0, 0, loc)
	return Range{Start: s, End: s.AddDate(1, 0, 0)}
}

func weekRange(ref time.Time) Range {
	// Week starts Monday 00:00
	d := truncateDay(ref)
	wd := int(d.Weekday())
	if wd == 0 {
		wd = 7 // Sunday -> 7
	}
	start := d.AddDate(0, 0, -(wd - 1))
	return Range{Start: start, End: start.AddDate(0, 0, 7)}
}

/* -------------------------
   Validation & months
--------------------------*/

func okYear(y int) bool { return 1000 <= y && y <= 9999 }
func okDay(d int) bool  { return 1 <= d && d <= 31 }

func monthID(s string) (time.Month, bool) {
	switch s {
	case "jan", "januari":
		return time.January, true
	case "feb", "februari":
		return time.February, true
	case "mar", "maret":
		return time.March, true
	case "apr", "april":
		return time.April, true
	case "mei":
		return time.May, true
	case "jun", "juni":
		return time.June, true
	case "jul", "juli":
		return time.July, true
	case "agu", "agt", "agustus":
		return time.August, true
	case "sep", "september":
		return time.September, true
	case "okt", "oktober":
		return time.October, true
	case "nov", "november":
		return time.November, true
	case "des", "desember":
		return time.December, true
	default:
		return 0, false
	}
}

/* -------------------------
   Regex
--------------------------*/

var (
	rxFromRange   = regexp.MustCompile(`^dari\s+(.+?)\s+(sampai|hingga|sd|-)\s+(.+)$`)
	rxInlineRange = regexp.MustCompile(`^(.+?)\s+(sampai|hingga|sd|-)\s+(.+)$`)

	rxEdgeMonthThis = regexp.MustCompile(`^(awal|akhir)\s+bulan\s+ini(?:\s|$)`)

	rxLastNUnit = regexp.MustCompile(`^(\d+)\s+(hari|minggu|pekan|bulan|tahun)\s+(terakhir|belakangan)(?:\s|$)`)

	rxOneWord      = regexp.MustCompile(`^(sekarang|hariini|kemarin|besok)(?:\s|$)`)
	rxUnitModifier = regexp.MustCompile(`^(minggu|pekan|bulan|tahun)\s+(ini|lalu|depan)(?:\s|$)`)

	rxRelNUnit = regexp.MustCompile(`^(\d+)\s+(hari|minggu|pekan|bulan|tahun)\s+(lalu|lagi|yanglalu|kedepan|darisekarang|darihariini)(?:\s|$)`)

	rxTahunYYYY = regexp.MustCompile(`^tahun\s+(\d{4})(?:\s|$)`)
	rxBareYear  = regexp.MustCompile(`^((?:19|20)\d{2})(?:\s|$)`)

	rxYMDNumeric = regexp.MustCompile(`^(\d{4})[-/](\d{1,2})[-/](\d{1,2})(?:\s|$)`)

	rxDMY = regexp.MustCompile(`^(\d{1,2})\s+([a-z]+)\s+(\d{4})(?:\s|$)`)
	rxMY  = regexp.MustCompile(`^([a-z]+)\s+(\d{4})(?:\s|$)`)
	rxYM  = regexp.MustCompile(`^(\d{4})\s+([a-z]+)(?:\s|$)`)

	// Ambiguous:
	rxDM = regexp.MustCompile(`^(\d{1,2})\s+([a-z]+)(?:\s|$)`)

	// Month only (limit to known tokens to avoid matching random words)
	rxMonthOnly = regexp.MustCompile(`^(jan|januari|feb|februari|mar|maret|apr|april|mei|jun|juni|jul|juli|agu|agt|agustus|sep|september|okt|oktober|nov|november|des|desember)(?:\s|$)`)
)
