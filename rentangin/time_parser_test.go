package rentangin

import (
	"testing"
	"time"
)

func wibLoc() *time.Location {
	return time.FixedZone("WIB", 7*3600)
}

func mustRange(t *testing.T, q string, now time.Time) Range {
	t.Helper()
	r, ok, err := Parse(q, now)
	if err != nil {
		t.Fatalf("Parse(%q) unexpected err: %v", q, err)
	}
	if !ok {
		t.Fatalf("Parse(%q) expected ok=true, got ok=false", q)
	}
	return r
}

func mustNoRange(t *testing.T, q string, now time.Time) {
	t.Helper()
	_, ok, err := Parse(q, now)
	if err != nil {
		t.Fatalf("Parse(%q) unexpected err: %v", q, err)
	}
	if ok {
		t.Fatalf("Parse(%q) expected ok=false, got ok=true", q)
	}
}

func assertRangeEq(t *testing.T, got Range, want Range) {
	t.Helper()
	if !got.Start.Equal(want.Start) || !got.End.Equal(want.End) {
		t.Fatalf("got [%s..%s) want [%s..%s)",
			got.Start.Format(time.RFC3339), got.End.Format(time.RFC3339),
			want.Start.Format(time.RFC3339), want.End.Format(time.RFC3339),
		)
	}
}

func TestParse_NoRange(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	cases := []string{
		"",
		"   ",
		"a",
		"berita politik",
		"not a date or a time",
		"Message me in 2 minutes", // english not supported here
		"10",
		"17",
		"10:am",
		"uu 24/2024",     // should not accidentally become year filter
		"pp 12/2019",     // ditto
		"iphone 12 2020", // no event hint
	}
	for _, q := range cases {
		t.Run(q, func(t *testing.T) {
			mustNoRange(t, q, now)
		})
	}
}

func TestParse_Words(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "gempa hari ini jakarta", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "banjir kemarin bandung", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 3, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "promo besok", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 6, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_UnitModifier(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc()) // Wed

	r := mustRange(t, "minggu ini pilkada", now)
	// Week starts Monday; 2026-02-04 (Wed) => starts 2026-02-02
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 2, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 9, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "bulan lalu saham", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "tahun depan event", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2027, 1, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2028, 1, 1, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_RelativeN_SingleDay(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "3 hari lalu demo", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 2, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "10 hari ke depan konser", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 14, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 15, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_LastN_Range(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "gempa 7 hari terakhir jakarta", now)
	// inclusive today: Feb 4 => start Jan 29; end Feb 5
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 29, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "2 bulan terakhir", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 1, 0, 0, 0, 0, wibLoc()), // monthRange(now) start is Feb 1; - (2-1) months => Jan 1
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_AwalAkhirBulanIni(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "awal bulan ini", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 2, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "akhir bulan ini", now)
	// 2026 bukan kabisat => Feb 28
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 28, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_Tahun_ExplicitKeyword(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "berita tahun 2024 pemilu", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2025, 1, 1, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_ExplicitDates(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "banjir 4 feb 2026 jakarta", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "laporan februari 2026 ekonomi", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "2026 feb inflasi", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "update 2026-02-04", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_RangeFormsAnywhere(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	r := mustRange(t, "data dari 1 feb 2026 sampai 10 feb 2026 foo", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "laporan 1 feb 2026 - 10 feb 2026 foo", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})

	r = mustRange(t, "laporan 1 feb 2026 s.d. 10 feb 2026", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_AmbiguousPastBias(t *testing.T) {
	loc := wibLoc()

	// now = 2026-02-04
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, loc)

	// "februari" => Feb 2026 (<= now)
	r := mustRange(t, "laporan februari", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, loc),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, loc),
	})

	// If now is January, "februari" should resolve to last year.
	nowJan := time.Date(2026, 1, 10, 10, 0, 0, 0, loc)
	r = mustRange(t, "februari", nowJan)
	assertRangeEq(t, r, Range{
		Start: time.Date(2025, 2, 1, 0, 0, 0, 0, loc),
		End:   time.Date(2025, 3, 1, 0, 0, 0, 0, loc),
	})

	// "15 maret" with now April 2026 => 15 Mar 2026
	nowApr := time.Date(2026, 4, 10, 10, 0, 0, 0, loc)
	r = mustRange(t, "15 maret", nowApr)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 3, 15, 0, 0, 0, 0, loc),
		End:   time.Date(2026, 3, 16, 0, 0, 0, 0, loc),
	})

	// "15 maret" with now Feb 2026 => 15 Mar 2025 (past-biased)
	nowFeb := time.Date(2026, 2, 10, 10, 0, 0, 0, loc)
	r = mustRange(t, "15 maret", nowFeb)
	assertRangeEq(t, r, Range{
		Start: time.Date(2025, 3, 15, 0, 0, 0, 0, loc),
		End:   time.Date(2025, 3, 16, 0, 0, 0, 0, loc),
	})
}

func TestParse_BareYear_EventHint_Allows(t *testing.T) {
	loc := wibLoc()
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, loc)

	r := mustRange(t, "pemenang piala dunia 1998", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(1998, 1, 1, 0, 0, 0, 0, loc),
		End:   time.Date(1999, 1, 1, 0, 0, 0, 0, loc),
	})

	r = mustRange(t, "pemilu 2024", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, loc),
		End:   time.Date(2025, 1, 1, 0, 0, 0, 0, loc),
	})

	// Without event hint should not parse bare year.
	mustNoRange(t, "angka 1998", now)
}

func TestParse_BareYear_TitleContext_Blocked(t *testing.T) {
	loc := wibLoc()
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, loc)

	// "film" + "review" => treat 2012 as title, not published_at filter
	mustNoRange(t, "review film 2012", now)
	mustNoRange(t, "film 2012 trailer", now)
}

func TestParse_BareYear_TopicIntent_Blocked(t *testing.T) {
	loc := wibLoc()
	now := time.Date(2025, 12, 1, 10, 0, 0, 0, loc)

	// topic intent + future year => do NOT prefilter by year
	mustNoRange(t, "prediksi gaya rambut 2026", now)
	mustNoRange(t, "tren 2026", now)
}

func TestParse_BareYear_FutureYear_BlockedEvenWithEventHint(t *testing.T) {
	loc := wibLoc()
	now := time.Date(2025, 12, 1, 10, 0, 0, 0, loc)

	// future year relative to nowYear(2025) => do NOT prefilter
	mustNoRange(t, "piala dunia 2026", now)
}

func TestParse_BestMatchWins(t *testing.T) {
	now := time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc())

	// Has both "bulan ini" and "7 hari terakhir" -> prefer 7 hari terakhir (score higher).
	r := mustRange(t, "bulan ini 7 hari terakhir", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 29, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	// Explicit range should beat everything else.
	r = mustRange(t, "7 hari terakhir dari 1 feb 2026 sampai 10 feb 2026", now)
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})
}
