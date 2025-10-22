package tanya

import (
	"strings"
	"testing"
)

func TestIsQuestion(t *testing.T) {
	t.Parallel()

	cases := []struct {
		q    string
		want bool
	}{
		// --- explicit / canonical questions
		{"apa itu knowledge graph", true},
		{"bagaimana cara reset password gmail", true},
		{"perbedaan redux vs zustand", true},
		{"kapan sidang mk hari ini", true},
		{"si andi pergi kemana ya", true},
		{"kok servernya error pas deploy", true},
		{"ya nggak sih performanya drop", true},
		{"jelasin cara mukbang", true},
		{"update kematian mahasiswa unud", true},

		// --- abbreviations / slang normalization
		{"gmn cara scrape instagram", true}, // gmn -> gimana
		{"knp server down semalem", true},   // knp -> kenapa
		{"dmn lokasi konser", true},         // dmn -> dimana
		{"brp harga langganan", false},      // brp -> berapa -> price intent => non-question

		// --- particles at the end (colloquial endings)
		{"mau makan kemana siang ini", true},
		{"dia tadi ke kantor dimana", true},
		{"ini kenapa ya", true},
		{"ini apa sih", true},
		{"performanya turun ya kan", true},

		// --- -kah suffix
		{"bisakah presiden diganti", true},
		{"mungkinkah ini berhasil", true},
		{"adakah solusi cepatnya", true},

		// --- how-to variants
		{"cara deploy ke production docker", true},
		{"bagaimana cara memperbaiki error 500", true},
		{"cara  cepat  push ke github  ", true}, // extra spaces

		// --- comparison signals
		{"bagusan mana mirrorless atau dslr", true},
		{"A vs B untuk data pipeline", true},
		{"versus airflow vs dagster", true},

		// --- definition / explain variants
		{"apa arti resilien", true},
		{"apa maksud zero copy", true},
		{"explain RAG pls", true},
		{"penjelasan implementasi RAG", true}, // explanation intent

		// --- update/time intent
		{"terkini erupsi bromo", true},
		{"perkembangan kasus x sekarang", true},

		// --- punctuation / emoji / casing
		{"KENAPA SERVER LEMOT", true},
		{"Kenapa server lemot?", true}, // explicit '?'
		{"kenapa server lemot 🤔", true},
		{"  Bagaimana Cara Reset Password  ", true},

		// --- URL / noise in a query
		{"cara setting oauth di https://example.com/docs", true},
		{"update api rate limit v2 (lihat changelog)", true},

		// --- obvious non-questions
		{"toyota", false},
		{"harga paket premium", false},
		{"kontak cs kumparan", false},
		{"download aplikasi android", false},
		{"grab promo kupon", false},
		{"", false},
		{"   \t  ", false},

		// --- tricky near-misses / should remain non-question
		{"vs code extensions", false}, // 'vs' as product word, not comparison
		{"mana store", false},         // 'mana' as noun chunk; intended info/browse
	}

	for _, tc := range cases {
		tc := tc
		name := tc.q
		if strings.TrimSpace(name) == "" {
			name = "<empty or whitespace>"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := IsQuestion(tc.q)
			if got != tc.want {
				t.Fatalf("IsQuestion(%q) = %v, want %v", tc.q, got, tc.want)
			}
		})
	}
}

func TestClassifyIntent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		q    string
		want Intent
	}{
		{"update kematian mahasiswa unud", IntentUpdate},
		{"jelasin cara mukbang", IntentExplain},
		{"arti overfitting", IntentExplain},
		{"bagaimana cara reset password gmail", IntentHowTo},
		{"apa itu knowledge graph", IntentDefinition},
		{"perbedaan redux vs zustand", IntentComparison},
		{"rekomendasi laptop 10 jutaan untuk desain", IntentRecommendation},
		{"kok servernya error pas deploy", IntentTroubleshoot},
		{"alamat kantor kumparan dimana ya", IntentLocation},
		{"mau makan kemana siang ini", IntentLocation},
		{"kapan jadwal konser hari ini", IntentTime},
		{"berapa harga paket premium", IntentPrice},
		{"kontak cs atau nomor wa resmi", IntentContact},
		{"toyota", IntentOther},
		{"download aplikasi android", IntentOther},

		// mixed signals
		{"update berita gempa vs banjir hari ini", IntentUpdate},
		{"apa sih lebih bagus mana A vs B", IntentQuestion},
		{"gmn cara beli tiket konser", IntentHowTo},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.q, func(t *testing.T) {
			t.Parallel()
			got := ClassifyIntent(tc.q)
			if got != tc.want {
				t.Fatalf("ClassifyIntent(%q) = %s, want one of %v", tc.q, got, tc.want)
			}
		})
	}
}

func BenchmarkIsQuestion(b *testing.B) {
	queries := []string{
		"apa itu knowledge graph",
		"bagaimana cara reset password gmail",
		"update kematian mahasiswa unud",
		"jelasin cara mukbang",
		"perbedaan redux vs zustand",
		"toyota",
		"harga paket premium",
		"kontak cs kumparan",
		"download aplikasi android",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, q := range queries {
			_ = IsQuestion(q)
		}
	}
}
