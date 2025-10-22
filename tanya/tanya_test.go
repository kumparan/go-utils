package tanya

import "testing"

func TestIsQuestion(t *testing.T) {
	t.Parallel()

	cases := []struct {
		q    string
		want bool
	}{
		// questions
		{"apa itu knowledge graph", true},
		{"bagaimana cara reset password gmail", true},
		{"perbedaan redux vs zustand", true},
		{"kapan sidang mk hari ini", true},
		{"si andi pergi kemana ya", true},
		{"kok servernya error pas deploy", true},
		{"ya nggak sih performanya drop", true},
		{"jelasin cara mukbang", true},
		{"update kematian mahasiswa unud", true},

		// abbreviations
		{"gmn cara scrape instagram", true},
		{"knp server down semalem", true},
		{"dmn lokasi konser", true},
		{"brp harga langganan", false},

		// non-question
		{"toyota", false},
		{"harga paket premium", false},
		{"kontak cs kumparan", false},
		{"download aplikasi android", false},
		{"", false},
		{"   \t  ", false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.q, func(t *testing.T) {
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
