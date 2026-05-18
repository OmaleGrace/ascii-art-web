package main

import (
	"os"
	"strings"
	"testing"
)

// helper — loads the banner once for all tests
var testBannerLines = func() []string {
	lines, err := readBanner("banner/standard.txt")
	if err != nil {
		panic("could not load banner/standard.txt: " + err.Error())
	}
	return lines
}()

// ─────────────────────────────────────────────
// readBanner tests
// ─────────────────────────────────────────────

func TestReadBanner_ReturnsCorrectLineCount(t *testing.T) {
	lines, err := readBanner("banner/standard.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 855 {
		t.Errorf("expected 855 lines, got %d", len(lines))
	}
}

func TestReadBanner_FileNotFound(t *testing.T) {
	_, err := readBanner("nonexistent.txt")
	if err == nil {
		t.Error("expected an error for missing file, got nil")
	}
}

// ─────────────────────────────────────────────
// getLines tests
// ─────────────────────────────────────────────

func TestGetLines_ReturnsEightLines(t *testing.T) {
	art := getLines(testBannerLines, 'A')
	if len(art) != 8 {
		t.Errorf("expected 8 lines for 'A', got %d", len(art))
	}
}

func TestGetLines_SpaceCharacter(t *testing.T) {
	art := getLines(testBannerLines, ' ')
	if len(art) != 8 {
		t.Errorf("expected 8 lines for space, got %d", len(art))
	}
}

func TestGetLines_FirstCharSpace(t *testing.T) {
	// space is ASCII 32 — the very first character in the banner
	art := getLines(testBannerLines, ' ')
	for _, line := range art {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			t.Errorf("expected all lines of space to be blank, got %q", line)
		}
	}
}

func TestGetLines_LastCharTilde(t *testing.T) {
	// tilde ~ is ASCII 126 — the last character in the banner
	art := getLines(testBannerLines, '~')
	if len(art) != 8 {
		t.Errorf("expected 8 lines for '~', got %d", len(art))
	}
}

func TestGetLines_CorrectStartIndex(t *testing.T) {
	// 'A' is ASCII 65 → start = (65-32)*9 = 297
	// art starts at index 298 (skip the blank separator)
	art := getLines(testBannerLines, 'A')
	expected := testBannerLines[298]
	if art[0] != expected {
		t.Errorf("expected first line of 'A' to be %q, got %q", expected, art[0])
	}
}

// ─────────────────────────────────────────────
// printArt tests  (capture stdout)
// ─────────────────────────────────────────────

// captureOutput redirects stdout during fn(), returns what was printed
func captureOutput(fn func()) string {
	// redirect stdout to a pipe
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	buf := new(strings.Builder)
	tmp := make([]byte, 1024)
	for {
		n, err := r.Read(tmp)
		buf.Write(tmp[:n])
		if err != nil {
			break
		}
	}
	return buf.String()
}

func TestPrintArt_SingleWord(t *testing.T) {
	lines := []string{"Hi"}
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	if out == "" {
		t.Error("expected output for 'Hi', got empty string")
	}
	// output should have exactly 8 newlines (8 rows)
	rows := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(rows) != 8 {
		t.Errorf("expected 8 rows for 'Hi', got %d", len(rows))
	}
}

func TestPrintArt_EmptyStringLineProducesBlankLine(t *testing.T) {
	// a single empty string in lines → one blank line
	lines := []string{""}
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	if out != "\n" {
		t.Errorf("expected single newline for empty line, got %q", out)
	}
}

func TestPrintArt_NewlineBetweenWords(t *testing.T) {
	// "Hello\nThere" splits into ["Hello", "There"] → 16 rows total
	lines := strings.Split("Hello\\nThere", "\\n")
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	rows := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(rows) != 16 {
		t.Errorf("expected 16 rows for Hello\\nThere, got %d", len(rows))
	}
}

func TestPrintArt_DoubleNewline(t *testing.T) {
	// "Hello\n\nThere" splits into ["Hello", "", "There"] → 8 + 1 + 8 = 17 rows
	lines := strings.Split("Hello\\n\\nThere", "\\n")
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	rows := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(rows) != 17 {
		t.Errorf("expected 17 rows for Hello\\n\\nThere, got %d", len(rows))
	}
}

func TestPrintArt_NumbersWork(t *testing.T) {
	lines := []string{"123"}
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	if out == "" {
		t.Error("expected output for '123', got empty string")
	}
}

func TestPrintArt_SpecialCharacters(t *testing.T) {
	lines := []string{"!@#"}
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	if out == "" {
		t.Error("expected output for '!@#', got empty string")
	}
}

func TestPrintArt_SpaceOnly(t *testing.T) {
	lines := []string{" "}
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	rows := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(rows) != 8 {
		t.Errorf("expected 8 rows for space, got %d", len(rows))
	}
}

func TestPrintArt_MixedCase(t *testing.T) {
	lines := []string{"HeLlO"}
	out := captureOutput(func() {
		printArt(lines, testBannerLines)
	})
	if out == "" {
		t.Error("expected output for 'HeLlO', got empty string")
	}
}