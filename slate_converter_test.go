package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToPlainText(t *testing.T) {
	inputJSON := `{"document":{"nodes":[{"object":"block","type":"heading-large","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"10 HP Android Paling Ngebut Versi AnTuTu Februari 2025, Ini Juaranya","marks":[]}]}]},{"object":"block","type":"figure","data":{},"nodes":[{"object":"block","type":"image","data":{"image":{"id":"1741342849091738823","title":"Untitled Image","description":"","publicID":"01jnr1ydtm32jndtmgaytxzp4y","externalURL":"https://blue.kumparan.com/image/upload/v1634025439/01jnr1ydtm32jndtmgaytxzp4y.jpg","awsS3Key":"2025/Mar/image/01jnr1ydtm32jndtmgaytxzp4y/","height":433,"width":768,"locationName":null,"locationLat":0,"locationLon":0,"mediaType":"IMAGE","mediaSourceID":"7","photographer":"","eventDate":"2025-03-07T10:20:49.091685Z","internalTags":[],"__typename":"Media"}},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"","marks":[]}]}]},{"object":"block","type":"caption","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"OnePlus Ace 5 Pro. Foto: OnePlus","marks":[]}]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Platform benchmark AnTuTu meluncurkan laporan baru soal daftar handphone (","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com/topic/hp"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"HP","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":") Android dengan performa terkencang di dunia. Untuk periode Februari 2025, smartphone dengan dapur pacu Snapdragon 8 Elite menjadi juaranya.","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Pengukuran AnTuTu berdasarkan beberapa aspek komponen di ","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com/topic/smartphone"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"smartphone","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":", seperti CPU, GPU, RAM, memori penyimpanan, hingga UX. Skor yang ditampilkan merupakan hasil sejumlah pengujian benchmark perangkat via aplikasi AnTuTu, minimal 1.000 kali dalam sebulan","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Prosesor Snapdragon 8 Elite dari Qualcomm dan Dimensity 9400 buatan MediaTek bersaing ketat dalam daftar 10 HP ","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com/topic/android"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Android","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":" dengan performa tercepat selama Februari 2025. Berikut daftar lengkapnya:","marks":[]}]}]},{"object":"block","type":"numbered-list","data":{},"nodes":[{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"OnePlus Ace 5 Pro","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"vivo X200 Pro","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Red Magic 10 Pro+","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"iQoo 13","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"iQoo Neo 10 Pro","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"OnePlus 13","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Realme GT 7 Pro","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Oppo Find X8 Pro","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Redmi K80 Pro","marks":[]}]}]}]},{"object":"block","type":"list-item","data":{},"nodes":[{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Oppo Find X8","marks":[]}]}]}]}]},{"object":"block","type":"figure","data":{},"nodes":[{"object":"block","type":"image","data":{"image":{"id":"1741342007934594403","title":"Untitled Image","description":"","publicID":"01jnr14renhw6ssnzjpd824eb7","externalURL":"https://blue.kumparan.com/image/upload/v1634025439/01jnr14renhw6ssnzjpd824eb7.jpg","awsS3Key":"2025/Mar/image/01jnr14renhw6ssnzjpd824eb7/","height":556,"width":738,"locationName":null,"locationLat":0,"locationLon":0,"mediaType":"IMAGE","mediaSourceID":"7","photographer":"","eventDate":"2025-03-07T10:06:47.934489Z","internalTags":[],"__typename":"Media"}},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"","marks":[]}]}]},{"object":"block","type":"caption","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Daftar 10 HP Android flagship paling ngebut versi AnTuTu periode Februari 2025. Foto: AnTuTu","marks":[]}]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"Peringkat pertama ditempati OnePlus Ace 5 Pro berbasis Snapdragon 8 Elite, dengan skor AnTuTu mencapai 2.890.600. Sementara itu, runner up-nya adalah vivo X200 Pro yang menggunakan cip Dimensity 9400, dengan skor AnTuTu 2.884.682.","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"RedMagic 10 Pro+ berada di posisi ketiga dengan skor AnTuTu 2.879.356, diikuti oleh iQoo 13 di peringkat keempat dengan skor AnTuTu 2.853.651. Kemudian, peringkat top 5 terakhir ditempati oleh iQoo Neo 10 Pro dengan skor AnTuTu 2.836.633.","marks":[]}]}]}]}}`
	expected := `Platform benchmark AnTuTu meluncurkan laporan baru soal daftar handphone (HP) Android dengan performa terkencang di dunia. Untuk periode Februari 2025, smartphone dengan dapur pacu Snapdragon 8 Elite menjadi juaranya.
Pengukuran AnTuTu berdasarkan beberapa aspek komponen di smartphone, seperti CPU, GPU, RAM, memori penyimpanan, hingga UX. Skor yang ditampilkan merupakan hasil sejumlah pengujian benchmark perangkat via aplikasi AnTuTu, minimal 1.000 kali dalam sebulan.
Prosesor Snapdragon 8 Elite dari Qualcomm dan Dimensity 9400 buatan MediaTek bersaing ketat dalam daftar 10 HP Android dengan performa tercepat selama Februari 2025. Berikut daftar lengkapnya:
OnePlus Ace 5 Pro, vivo X200 Pro, Red Magic 10 Pro+, iQoo 13, iQoo Neo 10 Pro, OnePlus 13, Realme GT 7 Pro, Oppo Find X8 Pro, Redmi K80 Pro, Oppo Find X8. Peringkat pertama ditempati OnePlus Ace 5 Pro berbasis Snapdragon 8 Elite, dengan skor AnTuTu mencapai 2.890.600. Sementara itu, runner up-nya adalah vivo X200 Pro yang menggunakan cip Dimensity 9400, dengan skor AnTuTu 2.884.682.
RedMagic 10 Pro+ berada di posisi ketiga dengan skor AnTuTu 2.879.356, diikuti oleh iQoo 13 di peringkat keempat dengan skor AnTuTu 2.853.651. Kemudian, peringkat top 5 terakhir ditempati oleh iQoo Neo 10 Pro dengan skor AnTuTu 2.836.633.`

	input, err := ParseSlateDocument(inputJSON)
	assert.Nil(t, err)
	assert.NotNil(t, input)

	result, err := input.ToPlainText()
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func Test_ToPlainText_Link(t *testing.T) {
	t.Run("text link not contain space", func(t *testing.T) {
		inputJSON := `{"document":{"nodes":[{"object":"block","type":"heading-large","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"ini adalah judul","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"tanda kemunculan ","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"suzuki fronx","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":" di indonesia.","marks":[]}]}]}]}}`
		expected := `tanda kemunculan suzuki fronx di indonesia.`

		input, err := ParseSlateDocument(inputJSON)
		assert.NoError(t, err)

		result, err := input.ToPlainText()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("text link include space on prefix", func(t *testing.T) {
		inputJSON := `{"document":{"nodes":[{"object":"block","type":"heading-large","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"ini adalah judul","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"tanda kemunculan","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":" suzuki fronx","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":" di indonesia.","marks":[]}]}]}]}}`
		expected := `tanda kemunculan suzuki fronx di indonesia.`

		input, err := ParseSlateDocument(inputJSON)
		assert.NoError(t, err)

		result, err := input.ToPlainText()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("text link include space on suffix", func(t *testing.T) {
		inputJSON := `{"document":{"nodes":[{"object":"block","type":"heading-large","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"ini adalah judul","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"tanda kemunculan ","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"suzuki fronx ","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":"di indonesia.","marks":[]}]}]}]}}`
		expected := `tanda kemunculan suzuki fronx di indonesia.`

		input, err := ParseSlateDocument(inputJSON)
		assert.NoError(t, err)

		result, err := input.ToPlainText()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("text link include space on prefix and suffix", func(t *testing.T) {
		inputJSON := `{"document":{"nodes":[{"object":"block","type":"heading-large","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"ini adalah judul","marks":[]}]}]},{"object":"block","type":"paragraph","data":{},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":"tanda kemunculan","marks":[]}]},{"object":"inline","type":"link","data":{"href":"https://kumparan.com"},"nodes":[{"object":"text","leaves":[{"object":"leaf","text":" suzuki fronx ","marks":[]}]}]},{"object":"text","leaves":[{"object":"leaf","text":"di indonesia.","marks":[]}]}]}]}}`
		expected := `tanda kemunculan suzuki fronx di indonesia.`

		input, err := ParseSlateDocument(inputJSON)
		assert.NoError(t, err)

		result, err := input.ToPlainText()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
