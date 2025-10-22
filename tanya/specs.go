package tanya

type (
	Intent    string
	MatchType string
)

const (
	IntentUpdate         Intent = "update"
	IntentExplain        Intent = "explain"
	IntentHowTo          Intent = "how_to"
	IntentDefinition     Intent = "definition"
	IntentComparison     Intent = "comparison"
	IntentRecommendation Intent = "recommendation"
	IntentTroubleshoot   Intent = "troubleshoot"
	IntentLocation       Intent = "location"
	IntentTime           Intent = "time"
	IntentPrice          Intent = "price"
	IntentContact        Intent = "contact"
	IntentQuestion       Intent = "question" // general fallback
	IntentOther          Intent = "other"

	MatchTypeContains        MatchType = "contains"
	MatchTypeStarts          MatchType = "starts"
	MatchTypeEnds            MatchType = "ends"
	MatchTypeEndsTokenSuffix MatchType = "ends_token_suffix"
)

type (
	Rule struct {
		Terms     []string
		Weight    int
		MatchType MatchType
	}

	IntentSpec struct {
		Intent   Intent
		Priority int
		Rules    []Rule
	}
)

func terms(ss ...string) []string { return ss }

var intentTable = []IntentSpec{
	{IntentUpdate, 95, []Rule{
		{terms("update", "perkembangan", "terbaru", "terkini", "progress", "lanjutan", "pembaruan"), 3, "contains"},
		{terms("hari ini", "sekarang", "terkini banget"), 1, "contains"},
	}},
	{IntentExplain, 90, []Rule{
		{terms("jelaskan", "jelasin", "penjelasan", "uraikan", "explain"), 3, "contains"},
		{terms("arti", "artinya", "maksud", "makna", "definisi"), 2, "contains"},
	}},
	{IntentHowTo, 80, []Rule{
		{terms("bagaimana cara ", "gimana cara "), 3, "starts"},
		{terms("cara "), 3, "starts"},
		{terms(" cara ", " langkah ", " step "), 1, "contains"},
	}},
	{IntentDefinition, 75, []Rule{
		{terms("apa itu "), 3, "starts"},
		{terms("apa arti", "apa maksud"), 2, "contains"},
	}},
	{IntentComparison, 70, []Rule{
		{terms(" vs ", " versus "), 2, "contains"},
		{terms("perbedaan ", "beda "), 2, "contains"},
		{terms("bagusan mana", "lebih bagus mana", "pilih mana"), 2, "contains"},
	}},
	{IntentRecommendation, 65, []Rule{
		{terms("rekomendasi", "rekom", "saran"), 2, "contains"},
		{terms("bagusan mana", "pilih mana", "cocok yang mana"), 2, "contains"},
	}},
	{IntentTroubleshoot, 60, []Rule{
		{terms("kenapa", "mengapa"), 2, "contains"},
		{terms("kok "), 2, "starts"},
		{terms("error", "gagal", "bug", "crash", "macet", "hang"), 2, "contains"},
		{terms("solusi ", "fix ", "gimana sih", "kenapa sih"), 1, "contains"},
	}},
	{IntentLocation, 55, []Rule{
		{terms("dimana", "di mana", "kemana", "ke mana", "lokasi", "alamat"), 2, "contains"},
		{terms(" kemana", " dimana", " alamat", " lokasi"), 2, "ends"},
	}},
	{IntentTime, 55, []Rule{
		{terms("kapan", "jadwal", "jam berapa", "pukul berapa"), 2, "contains"},
		{terms("hari ini", "minggu ini", "sekarang", "besok", "nanti sore", "malam ini"), 1, "contains"},
	}},
	{IntentPrice, 50, []Rule{
		{terms("harga", "biaya", "tarif", "fee", "ongkir"), 2, "contains"},
	}},
	{IntentContact, 50, []Rule{
		{terms("kontak", "contact", "telepon", "telp", "nomor", "email", "whatsapp", "wa"), 2, "contains"},
	}},
	// fallback tanya umum
	{IntentQuestion, 10, []Rule{
		{terms("apa", "apakah", "bagaimana", "gimana", "kapan", "siapa", "dimana", "di mana", "kemana", "ke mana", "berapa", "mana"), 2, "contains"},
		{terms(" vs ", " versus "), 1, "contains"},
		{terms("kah"), 1, "ends_token_suffix"},
		{terms("ya ga sih", "ya gak sih", "ya nggak sih", "ya kan", "apa sih", "gimana sih", "kenapa sih"), 2, "contains"},
		{terms(" kok "), 2, "contains"},
		{terms("?"), 3, "contains"},
	}},
}

var abbrevMap = map[string]string{
	" gmn ": " gimana ", " knp ": " kenapa ", " dmn ": " dimana ", " brp ": " berapa ",
}
