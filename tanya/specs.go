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

	MatchTypeContains    MatchType = "contains"
	MatchTypeStarts      MatchType = "starts"
	MatchTypeEnds        MatchType = "ends"
	MatchTypeTokenSuffix MatchType = "token_suffix" // nolint:gosec
)

type (
	// Rule is a rule for matching a query to intent
	Rule struct {
		Terms       []string
		Weight      int
		MatchType   MatchType
		MinTokenLen int // optional: for token_suffix; <=0 => default 4
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
		{terms("update", "perkembangan", "terbaru", "terkini", "progress", "lanjutan", "pembaruan"), 3, MatchTypeContains, 0},
		{terms("hari ini", "sekarang", "terkini banget"), 1, MatchTypeContains, 0},
	}},
	{IntentExplain, 90, []Rule{
		{terms("jelaskan", "jelasin", "penjelasan", "uraikan", "explain"), 3, MatchTypeContains, 0},
		{terms("arti", "artinya", "maksud", "makna", "definisi"), 2, MatchTypeContains, 0},
	}},
	{IntentHowTo, 80, []Rule{
		{terms("bagaimana cara ", "gimana cara "), 3, MatchTypeStarts, 0},
		{terms("cara "), 3, MatchTypeStarts, 0},
		{terms(" cara ", " langkah ", " step "), 1, MatchTypeContains, 0},
		{terms("resep "), 3, MatchTypeStarts, 0},
		{terms(" panduan ", "panduan "), 2, MatchTypeContains, 0},
		{terms(" tutorial ", "tutorial "), 2, MatchTypeContains, 0},
	}},
	{IntentDefinition, 75, []Rule{
		{terms("apa itu "), 3, MatchTypeStarts, 0},
		{terms("apa arti", "apa maksud"), 2, MatchTypeContains, 0},
	}},
	{IntentComparison, 70, []Rule{
		{terms(" vs ", " versus "), 2, MatchTypeContains, 0},
		{terms("perbedaan ", "beda "), 2, MatchTypeContains, 0},
		{terms("bagusan mana", "lebih bagus mana", "pilih mana"), 2, MatchTypeContains, 0},
	}},
	{IntentRecommendation, 65, []Rule{
		{terms("rekomendasi", "rekom", "saran"), 2, MatchTypeContains, 0},
		{terms("bagusan mana", "pilih mana", "cocok yang mana"), 2, MatchTypeContains, 0},
		{terms("menu "), 2, MatchTypeStarts, 0},
		{terms(" ide ", "ide "), 1, MatchTypeContains, 0},
	}},
	{IntentTroubleshoot, 60, []Rule{
		{terms("kenapa", "mengapa"), 2, MatchTypeContains, 0},
		{terms("kok "), 2, MatchTypeStarts, 0},
		{terms("error", "gagal", "bug", "crash", "macet", "hang"), 2, MatchTypeContains, 0},
		{terms("solusi ", "fix ", "gimana sih", "kenapa sih"), 1, MatchTypeContains, 0},
	}},
	{IntentLocation, 55, []Rule{
		{terms("dimana", "di mana", "kemana", "ke mana", "lokasi", "alamat"), 2, MatchTypeContains, 0},
		{terms(" kemana", " dimana", "di mana", " alamat", " lokasi"), 2, MatchTypeEnds, 0},
	}},
	{IntentTime, 55, []Rule{
		{terms("kapan", "jadwal", "jam berapa", "pukul berapa"), 2, MatchTypeContains, 0},
		{terms("hari ini", "minggu ini", "sekarang", "besok", "nanti sore", "malam ini"), 1, MatchTypeContains, 0},
	}},
	{IntentPrice, 50, []Rule{
		{terms("harga", "biaya", "tarif", "fee", "ongkir"), 2, MatchTypeContains, 0},
	}},
	{IntentContact, 50, []Rule{
		{terms("kontak", "contact", "telepon", "telp", "nomor", "email", "whatsapp", "wa"), 2, MatchTypeContains, 0},
	}},
	// fallback tanya umum
	{IntentQuestion, 10, []Rule{
		{terms("apa", "apakah", "bagaimana", "gimana", "kapan", "siapa", "dimana", "di mana", "kemana", "ke mana", "berapa"), 2, MatchTypeContains, 0},
		{terms(" vs ", " versus "), 1, MatchTypeContains, 0},
		{terms(" yang mana "), 2, MatchTypeContains, 0},
		{terms(" mana saja "), 2, MatchTypeContains, 0},
		{terms("yang mana "), 2, MatchTypeStarts, 0},
		{terms(" mana"), 2, MatchTypeEnds, 0},
		{terms("kah"), 1, MatchTypeTokenSuffix, 5},
		{terms("ya ga sih", "ya gak sih", "ya nggak sih", "ya kan", "apa sih", "gimana sih", "kenapa sih"), 2, MatchTypeContains, 0},
		{terms(" kok "), 2, MatchTypeContains, 0},
		{terms("?"), 3, MatchTypeContains, 0},
	}},
}

var abbrevMap = map[string]string{
	"gmn":   "gimana",
	"gmna":  "gimana",
	"bgmn":  "bagaimana",
	"knp":   "kenapa",
	"knpa":  "kenapa",
	"dmn":   "di mana",
	"dmna":  "di mana",
	"dimn":  "di mana",
	"kmn":   "ke mana",
	"kmna":  "ke mana",
	"brp":   "berapa",
	"brpa":  "berapa",
	"kpn":   "kapan",
	"kpan":  "kapan",
	"sapa":  "siapa",
	"sp":    "siapa",
	"syp":   "siapa",
	"sypa":  "siapa",
	"apkh":  "apakah",
	"apakh": "apakah",
}
