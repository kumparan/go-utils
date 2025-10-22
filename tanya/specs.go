package tanya

type (
	// Intent is the intent of a query
	Intent string
	// MatchType is the type of match
	MatchType string
)

// Intent and MatchType constants
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
	MatchTypeEndsTokenSuffix MatchType = "ends_token_suffix" // nolint:gosec
)

type (
	// Rule is a rule for matching a query to intent
	Rule struct {
		Terms     []string
		Weight    int
		MatchType MatchType
	}

	// IntentSpec is a specification for intent
	IntentSpec struct {
		Intent   Intent
		Priority int
		Rules    []Rule
	}
)

func terms(ss ...string) []string { return ss }

var intentTable = []IntentSpec{
	{IntentUpdate, 95, []Rule{
		{terms("update", "perkembangan", "terbaru", "terkini", "progress", "lanjutan", "pembaruan"), 3, MatchTypeContains},
		{terms("hari ini", "sekarang", "terkini banget"), 1, MatchTypeContains},
	}},
	{IntentExplain, 90, []Rule{
		{terms("jelaskan", "jelasin", "penjelasan", "uraikan", "explain"), 3, MatchTypeContains},
		{terms("arti", "artinya", "maksud", "makna", "definisi"), 2, MatchTypeContains},
	}},
	{IntentHowTo, 80, []Rule{
		{terms("bagaimana cara ", "gimana cara "), 3, MatchTypeStarts},
		{terms("cara "), 3, MatchTypeStarts},
		{terms(" cara ", " langkah ", " step "), 1, MatchTypeContains},
	}},
	{IntentDefinition, 75, []Rule{
		{terms("apa itu "), 3, MatchTypeStarts},
		{terms("apa arti", "apa maksud"), 2, MatchTypeContains},
	}},
	{IntentComparison, 70, []Rule{
		{terms(" vs ", " versus "), 2, MatchTypeContains},
		{terms("perbedaan ", "beda "), 2, MatchTypeContains},
		{terms("bagusan mana", "lebih bagus mana", "pilih mana"), 2, MatchTypeContains},
	}},
	{IntentRecommendation, 65, []Rule{
		{terms("rekomendasi", "rekom", "saran"), 2, MatchTypeContains},
		{terms("bagusan mana", "pilih mana", "cocok yang mana"), 2, MatchTypeContains},
	}},
	{IntentTroubleshoot, 60, []Rule{
		{terms("kenapa", "mengapa"), 2, MatchTypeContains},
		{terms("kok "), 2, MatchTypeStarts},
		{terms("error", "gagal", "bug", "crash", "macet", "hang"), 2, MatchTypeContains},
		{terms("solusi ", "fix ", "gimana sih", "kenapa sih"), 1, MatchTypeContains},
	}},
	{IntentLocation, 55, []Rule{
		{terms("dimana", "di mana", "kemana", "ke mana", "lokasi", "alamat"), 2, MatchTypeContains},
		{terms(" kemana", " dimana", " alamat", " lokasi"), 2, MatchTypeEnds},
	}},
	{IntentTime, 55, []Rule{
		{terms("kapan", "jadwal", "jam berapa", "pukul berapa"), 2, MatchTypeContains},
		{terms("hari ini", "minggu ini", "sekarang", "besok", "nanti sore", "malam ini"), 1, MatchTypeContains},
	}},
	{IntentPrice, 50, []Rule{
		{terms("harga", "biaya", "tarif", "fee", "ongkir"), 2, MatchTypeContains},
	}},
	{IntentContact, 50, []Rule{
		{terms("kontak", "contact", "telepon", "telp", "nomor", "email", "whatsapp", "wa"), 2, MatchTypeContains},
	}},
	// fallback tanya umum
	{IntentQuestion, 10, []Rule{
		{terms("apa", "apakah", "bagaimana", "gimana", "kapan", "siapa", "dimana", "di mana", "kemana", "ke mana", "berapa"), 2, MatchTypeContains},
		{terms(" vs ", " versus "), 1, MatchTypeContains},
		{terms(" yang mana "), 2, MatchTypeContains},
		{terms(" mana"), 2, MatchTypeEnds},
		{terms("kah"), 1, MatchTypeEndsTokenSuffix},
		{terms("ya ga sih", "ya gak sih", "ya nggak sih", "ya kan", "apa sih", "gimana sih", "kenapa sih"), 2, MatchTypeContains},
		{terms(" kok "), 2, MatchTypeContains},
		{terms("?"), 3, MatchTypeContains},
	}},
}

var abbrevMap = map[string]string{
	" gmn ": " gimana ", " knp ": " kenapa ", " dmn ": " dimana ", " brp ": " berapa ",
}
