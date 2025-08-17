package docs

import (
	"os"
	"testing"
)

func TestFetchMany(t *testing.T) {
	var testCases = []struct {
		urls     []string
		expected string
	}{
		{urls: []string{}, expected: "No documentation URLs provided"},
		{urls: []string{"https://raw.githubusercontent.com/AstraBert/workflows-go/main/.gitignore"}, expected: "## Documentation for https://raw.githubusercontent.com/AstraBert/workflows-go/main/.gitignore\n\n# env variables\n.env\n\n\n---\n\n"},
		{urls: []string{"https://raw.githubusercontent.com/AstraBert/workflows-go/main/.gitignore", "https://raw.githubusercontent.com/AstraBert/workflows-go/main/.gitignore"}, expected: "## Documentation for https://raw.githubusercontent.com/AstraBert/workflows-go/main/.gitignore\n\n# env variables\n.env\n\n\n---\n\n## Documentation for https://raw.githubusercontent.com/AstraBert/workflows-go/main/.gitignore\n\n# env variables\n.env\n\n\n---\n\n"},
	}

	for _, tc := range testCases {
		result := FetchMany(tc.urls)
		if result != tc.expected {
			t.Errorf("Testing FetchMany: expecting %s, got %s", tc.expected, result)
		}
	}
}

func TestWriteFileContent(t *testing.T) {
	var testCases = []struct {
		filePath string
		content  string
		expected error
	}{
		{"hello.txt", "Hello world, it is nice to meet you!", nil},
		{"document.md", "## This is an official document", nil},
	}

	for _, tc := range testCases {
		err := WriteFileContent(tc.filePath, tc.content)
		if err != tc.expected {
			t.Error("Testing WriteFileContent: Expecting no error, got: " + err.Error())
		}
		readContent, fail := os.ReadFile(tc.filePath)
		if fail != nil {
			t.Errorf("Expecting %s to be readable, but it is not", tc.filePath)
		}
		if string(readContent) != tc.content {
			t.Errorf("Expecting %s content to be %s, got %s", tc.filePath, tc.content, string(readContent))
		}
		os.Remove(tc.filePath)
	}
}
