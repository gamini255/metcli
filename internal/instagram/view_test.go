package instagram

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestParseUsername(t *testing.T) {
	cases := map[string]string{
		"@sportg33k":                         "sportg33k",
		" sportg33k ":                        "sportg33k",
		"https://www.instagram.com/foo/":     "foo",
		"https://instagram.com/bar/":         "bar",
		"https://www.instagram.com/baz/reel": "baz",
	}
	for input, want := range cases {
		if got := ParseUsername(input); got != want {
			t.Fatalf("ParseUsername(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestBuildItemsCopiesFields(t *testing.T) {
	profile := Profile{
		ProfilePicURL:   "pic",
		ProfilePicURLHD: "hd",
		Media: []MediaItem{
			{
				URL:       "m1",
				IsVideo:   false,
				Shortcode: "s1",
				TakenAt:   1,
				Username:  "user",
				Caption:   "cap",
			},
			{
				URL:     "m2",
				IsVideo: true,
			},
		},
	}
	items := BuildItems(profile, true, false)
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].Kind != "avatar" || items[0].URL != "hd" {
		t.Fatalf("expected avatar hd, got kind=%q url=%q", items[0].Kind, items[0].URL)
	}
	if items[1].URL != "m1" || items[1].Username != "user" || items[1].Caption != "cap" {
		t.Fatalf("expected media with fields, got url=%q user=%q cap=%q", items[1].URL, items[1].Username, items[1].Caption)
	}
}

func TestEnsurePNG(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 255, A: 255})
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatalf("encode png: %v", err)
	}
	data, err := EnsurePNG(buf.Bytes())
	if err != nil {
		t.Fatalf("EnsurePNG: %v", err)
	}
	if len(data) < 8 || !bytes.HasPrefix(data, []byte("\x89PNG\r\n\x1a\n")) {
		t.Fatalf("expected png header")
	}
}
