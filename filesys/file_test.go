package filesys

import (
	"os"
	"path"
	"testing"
)

func TestFilePath(t *testing.T) {
	h := os.Getenv("HOME")

	tt := []struct {
		name     string
		filename string
		expected string
	}{
		{
			name:     "empty filename",
			filename: "",
			expected: path.Join(h, ".bh"),
		},
		{
			name:     "actual filename",
			filename: "info.json",
			expected: path.Join(h, ".bh", "info.json"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := FilePath(tc.filename)
			if o != tc.expected {
				t.Errorf("expected path %s got %s", tc.expected, o)
			}
		})
	}
}
