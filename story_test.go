package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStoryIDFromStorySlug(t *testing.T) {
	assert.Equal(t, "1tags0a6FDA", GetStoryIDFromStorySlug("uluran-tangan-kita-bagi-umkm-dan-masyarakat-yang-terdampak-corona-1tags0a6FDA"))
	assert.Equal(t, "1tRa4F1yakl", GetStoryIDFromStorySlug("thr-dari-kita-untuk-para-marbut-masjid-penghafal-al-quran-hingga-guru-ngaji-1tRa4F1yakl"))
	assert.Equal(t, "1sz5EeIPqiX", GetStoryIDFromStorySlug("kisah-pilu-guru-honorer-surabaya-diberhentikan-karena-pembekuan-darah-di-otak-1sz5EeIPqiX"))
}
