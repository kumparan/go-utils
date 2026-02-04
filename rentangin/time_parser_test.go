package rentangin

import (
	"testing"
	"time"
)

func wibLoc() *time.Location {
	return time.FixedZone("WIB", 7*3600)
}

func nowWIB() time.Time {
	return time.Date(2026, 2, 4, 10, 0, 0, 0, wibLoc()) // Wed
}

func assertRangeEq(t *testing.T, got, want Range) {
	t.Helper()
	if !got.Start.Equal(want.Start) || !got.End.Equal(want.End) {
		t.Fatalf("got [%s..%s) want [%s..%s)",
			got.Start.Format(time.RFC3339), got.End.Format(time.RFC3339),
			want.Start.Format(time.RFC3339), want.End.Format(time.RFC3339),
		)
	}
}

func TestParse_NoRange(t *testing.T) {
	now := nowWIB()

	cases := []string{
		"",
		"   ",
		"a",
		"berita politik",
		"not a date or a time",
		"Message me in 2 minutes", // english not supported
		"10",
		"17",
		"10:am",
		"uu 24/2024", // ambiguous id-style, we don't parse
	}
	for _, s := range cases {
		t.Run(s, func(t *testing.T) {
			_, ok, err := Parse(s, now)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if ok {
				t.Fatalf("expected ok=false")
			}
		})
	}
}

func TestParse_Words(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("gempa hari ini jakarta", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("banjir kemarin bandung", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 3, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("promo besok", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 6, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_UnitModifier(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("minggu ini pilkada", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	// Week starts Monday; 2026-02-04 is Wed => week starts 2026-02-02
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 2, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 9, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("bulan lalu saham", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("tahun depan event", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2027, 1, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2028, 1, 1, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_RelativeN_SingleDay(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("3 hari lalu demo", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 2, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("10 hari ke depan konser", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 14, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 15, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_TahunYYYY(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("berita tahun 2024 pemilu", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2025, 1, 1, 0, 0, 0, 0, wibLoc()),
	})

	// bare year should NOT parse (safe mode)
	_, ok, err = Parse("2024 pemilu", now)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if ok {
		t.Fatalf("expected ok=false for bare year")
	}
}

func TestParse_ExplicitDates(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("banjir 4 feb 2026 jakarta", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("laporan februari 2026 ekonomi", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("2026 feb inflasi", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("update 2026-02-04", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 4, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_RangeFormsAnywhere(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("data dari 1 feb 2026 sampai 10 feb 2026 foo", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("laporan 1 feb 2026 - 10 feb 2026 foo", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("laporan 1 feb 2026 s.d. 10 feb 2026", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_Last7HariTerakhir(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("gempa 7 hari terakhir jakarta", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}

	// inclusive today: Feb 4 -> start Jan 29, end Feb 5
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 29, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_AwalAkhirBulanIni(t *testing.T) {
	now := nowWIB()

	r, ok, err := Parse("awal bulan ini", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 2, 0, 0, 0, 0, wibLoc()),
	})

	r, ok, err = Parse("akhir bulan ini", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	// 2026 bukan kabisat => Feb 28
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 28, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, wibLoc()),
	})
}

func TestParse_AmbiguousPastBias(t *testing.T) {
	loc := wibLoc()

	// now = 2026-02-04
	now := nowWIB()

	// "februari" => Feb 2026 (<= now)
	r, ok, err := Parse("laporan februari", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, loc),
		End:   time.Date(2026, 3, 1, 0, 0, 0, 0, loc),
	})

	// If now is January, "februari" should resolve to last year.
	nowJan := time.Date(2026, 1, 10, 10, 0, 0, 0, loc)
	r, ok, err = Parse("februari", nowJan)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2025, 2, 1, 0, 0, 0, 0, loc),
		End:   time.Date(2025, 3, 1, 0, 0, 0, 0, loc),
	})

	// "15 maret" with now April 2026 => 15 Mar 2026
	nowApr := time.Date(2026, 4, 10, 10, 0, 0, 0, loc)
	r, ok, err = Parse("15 maret", nowApr)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 3, 15, 0, 0, 0, 0, loc),
		End:   time.Date(2026, 3, 16, 0, 0, 0, 0, loc),
	})

	// "15 maret" with now Feb 2026 => 15 Mar 2025 (past-biased)
	nowFeb := time.Date(2026, 2, 10, 10, 0, 0, 0, loc)
	r, ok, err = Parse("15 maret", nowFeb)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2025, 3, 15, 0, 0, 0, 0, loc),
		End:   time.Date(2025, 3, 16, 0, 0, 0, 0, loc),
	})
}

func TestParse_BestMatchWins(t *testing.T) {
	now := nowWIB()

	// Has both "bulan ini" and "7 hari terakhir" -> prefer 7 hari terakhir (score higher).
	r, ok, err := Parse("bulan ini 7 hari terakhir", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 1, 29, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 5, 0, 0, 0, 0, wibLoc()),
	})

	// Explicit range should beat everything else.
	r, ok, err = Parse("7 hari terakhir dari 1 feb 2026 sampai 10 feb 2026", now)
	if err != nil || !ok {
		t.Fatalf("err=%v ok=%v", err, ok)
	}
	assertRangeEq(t, r, Range{
		Start: time.Date(2026, 2, 1, 0, 0, 0, 0, wibLoc()),
		End:   time.Date(2026, 2, 10, 0, 0, 0, 0, wibLoc()),
	})
}
