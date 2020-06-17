package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JoinURL(t *testing.T) {
	baseURL := "https://kumparan.com/trending"

	t.Run("path elements empty", func(t *testing.T) {
		resURL, err := JoinURL(baseURL)
		assert.Equal(t, baseURL, resURL)
		assert.NoError(t, err)
	})

	t.Run("base empty", func(t *testing.T) {
		url := "trending"
		resURL, err := JoinURL("", url)
		assert.Equal(t, url, resURL)
		assert.NoError(t, err)
	})

	t.Run("multiple path elements", func(t *testing.T) {
		resURL, err := JoinURL(baseURL, "category", "news")
		assert.Equal(t, "https://kumparan.com/trending/category/news", resURL)
		assert.NoError(t, err)
	})

	t.Run("multiple path elements in single string", func(t *testing.T) {
		resURL, err := JoinURL(baseURL, "category/news")
		assert.Equal(t, "https://kumparan.com/trending/category/news", resURL)
		assert.NoError(t, err)
	})

	t.Run("if path elements have \"..\"", func(t *testing.T) {
		baseURL := "https://kumparan.com/trending/feed"
		url := "../indonesia/news"
		resURL, err := JoinURL(baseURL, url)
		assert.Equal(t, "https://kumparan.com/trending/indonesia/news", resURL)
		assert.NoError(t, err)

		t.Run("if path elements have \"..\" on the last element", func(t *testing.T) {
			baseURL := "https://kumparan.com/trending/feed"
			url := "sepakbola/.."
			resURL, err := JoinURL(baseURL, url)

			assert.Equal(t, "https://kumparan.com/trending/feed", resURL)
			assert.NoError(t, err)

		})
	})

	t.Run("if path elements have \".\"", func(t *testing.T) {
		baseURL := "https://kumparan.com/trending/feed"
		url := "./indonesia/news"
		resURL, err := JoinURL(baseURL, url)
		assert.Equal(t, "https://kumparan.com/trending/feed/indonesia/news", resURL)
		assert.NoError(t, err)
	})

	t.Run("if url have port", func(t *testing.T) {
		baseURL := "https://192.168.1:3000/trending/feed"
		url := "/story"
		expectedURL := "https://192.168.1:3000/trending/feed/story"
		resURL, err := JoinURL(baseURL, url)
		assert.Equal(t, expectedURL, resURL)
		assert.NoError(t, err)
	})
}
