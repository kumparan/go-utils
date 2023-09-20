package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StandardizeSpaces(t *testing.T) {
	s := `test
	1
	2
	3`

	assert.Equal(t, "test 1 2 3", StandardizeSpaces(s))
}

func Test_StringToBool(t *testing.T) {
	assert.Equal(t, true, StringToBool("true"))
	assert.Equal(t, false, StringToBool("false"))
	assert.Equal(t, false, StringToBool("0"))
	assert.Equal(t, true, StringToBool("1"))
	assert.Equal(t, false, StringToBool("bebek"))
}

func Test_StringToInt64(t *testing.T) {
	assert.Equal(t, int64(10), StringToInt64("10"))
	assert.Equal(t, int64(20), StringToInt64("20"))
	assert.Equal(t, int64(0), StringToInt64("20abc"))
}

func Test_StringToInt64WithDefault(t *testing.T) {
	assert.Equal(t, int64(10), StringToInt64WithDefault("10", 0))
	assert.Equal(t, int64(20), StringToInt64WithDefault("20", 0))
	assert.Equal(t, int64(999), StringToInt64WithDefault("20abc", 999))
}

func Test_StringPointerToString(t *testing.T) {
	var s *string
	assert.Equal(t, "", StringPointerToString(s))
	ss := "bengbeng"
	s = &ss
	assert.Equal(t, "bengbeng", StringPointerToString(s))
	*s = ""
	assert.Equal(t, "", StringPointerToString(s))
}

func Test_StringPointerToFloat64(t *testing.T) {
	var s *string
	assert.Equal(t, float64(0), StringPointerToFloat64(s))
	ss := "12.22"
	s = &ss
	assert.Equal(t, float64(12.22), StringPointerToFloat64(s))
	*s = ""
	assert.Equal(t, float64(0), StringPointerToFloat64(s))
}

func Test_StringPointerToInt64(t *testing.T) {
	var s *string
	assert.Equal(t, int64(0), StringPointerToInt64(s))
	ss := "12"
	s = &ss
	assert.Equal(t, int64(12), StringPointerToInt64(s))
	*s = ""
	assert.Equal(t, int64(0), StringPointerToInt64(s))
}

func Test_ArrayStringPointerToArrayInt64(t *testing.T) {
	var ps1, ps2 *string
	s1 := "123"
	s2 := "321"
	ps1 = &s1
	ps2 = &s2

	s := &[]*string{ps1, ps2}
	as := ArrayStringPointerToArrayInt64(s)

	assert.Contains(t, as, int64(123))
	assert.Contains(t, as, int64(321))
}

func Test_TruncateString(t *testing.T) {
	str1 := "A-tisket a-tasket A green and yellow basket"
	str2 := "Peter Piper picked a peck of pickled peppers"

	assert.Equal(t, "A-tisket...", TruncateString(str1, 11))
	assert.Equal(t, "Peter Piper...", TruncateString(str2, 14))
	assert.Equal(t, "A-tisket a-tasket A green and yellow basket", TruncateString(str1, len(str1)))
	assert.Equal(t, "A-tisket a-tasket A green and yellow basket", TruncateString(str1, len(str1)+2))
	assert.Equal(t, "A", TruncateString("A-", 1))
	assert.Equal(t, "Ab", TruncateString("Absolutely Longer", 2))
	assert.Equal(t, "A...", TruncateString("Absolutely Longer", 4))
}

func Test_EscapeQuote(t *testing.T) {
	strWithoutQuote := "this is brazil"
	strWithQuote := "\"bek bek bek\""
	strWithEscapedQuote := "\\\"bek\\\" bek\\\" bek\\\""

	testCases := map[string]string{
		strWithoutQuote:     "this is brazil",
		strWithQuote:        "\\\"bek bek bek\\\"",
		strWithEscapedQuote: "\\\"bek\\\" bek\\\" bek\\\"",
	}

	for in, out := range testCases {
		assert.Equal(t, out, EscapeQuote(in))
	}
}

func Test_TrimSpacePointerString(t *testing.T) {
	str := "this is america"
	strWithSpace := "     this is america  "

	testCases := map[*string]*string{
		&str:          &str,
		&strWithSpace: &str,
	}

	for in, out := range testCases {
		TrimSpacePointerString(in)
		assert.Equal(t, out, in)
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	assert.NotEqual(t, "", uuid)
	assert.Equal(t, 36, len(uuid))
}
