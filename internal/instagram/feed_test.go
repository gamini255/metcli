package instagram

import "testing"

func TestItemCaption(t *testing.T) {
	item := feedItem{
		Caption: &feedCaption{Text: " hello "},
	}
	if got := itemCaption(item); got != "hello" {
		t.Fatalf("expected caption from struct, got %q", got)
	}

	item = feedItem{CaptionText: " world "}
	if got := itemCaption(item); got != "world" {
		t.Fatalf("expected caption from text, got %q", got)
	}
}

func TestFeedItemToMediaPhoto(t *testing.T) {
	item := feedItem{
		MediaType: 1,
		ImageVersions: imageVersions{Candidates: []imageCandidate{
			{URL: "small", Width: 10, Height: 10},
			{URL: "large", Width: 40, Height: 20},
		}},
		Code:    "abc",
		TakenAt: 123,
		User:    feedUser{Username: "tester"},
		Caption: &feedCaption{Text: "cap"},
	}
	items := feedItemToMedia(item)
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	got := items[0]
	if got.URL != "large" {
		t.Fatalf("expected best candidate, got %q", got.URL)
	}
	if got.IsVideo {
		t.Fatalf("expected photo")
	}
	if got.Shortcode != "abc" {
		t.Fatalf("expected shortcode abc, got %q", got.Shortcode)
	}
	if got.Username != "tester" || got.Caption != "cap" {
		t.Fatalf("expected username/caption, got %q/%q", got.Username, got.Caption)
	}
	if got.TakenAt != 123 {
		t.Fatalf("expected taken_at 123, got %d", got.TakenAt)
	}
}

func TestFeedItemToMediaVideo(t *testing.T) {
	item := feedItem{
		MediaType:    2,
		ThumbnailURL: "thumb",
		Shortcode:    "xyz",
		User:         feedUser{Username: "video_user"},
		CaptionText:  "hello",
	}
	items := feedItemToMedia(item)
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	got := items[0]
	if got.URL != "thumb" {
		t.Fatalf("expected thumbnail url, got %q", got.URL)
	}
	if !got.IsVideo {
		t.Fatalf("expected video")
	}
	if got.Shortcode != "xyz" {
		t.Fatalf("expected shortcode xyz, got %q", got.Shortcode)
	}
	if got.Username != "video_user" || got.Caption != "hello" {
		t.Fatalf("expected username/caption, got %q/%q", got.Username, got.Caption)
	}
}

func TestFeedItemToMediaCarousel(t *testing.T) {
	item := feedItem{
		MediaType: 8,
		CarouselMedia: []carouselMedia{
			{
				MediaType:     1,
				ImageVersions: imageVersions{Candidates: []imageCandidate{{URL: "c1", Width: 5, Height: 5}}},
			},
			{
				MediaType:    2,
				ThumbnailURL: "c2",
			},
		},
		Code:    "car",
		User:    feedUser{Username: "carousel_user"},
		Caption: &feedCaption{Text: "cap"},
	}
	items := feedItemToMedia(item)
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
	if items[0].URL != "c1" || items[1].URL != "c2" {
		t.Fatalf("unexpected urls: %q %q", items[0].URL, items[1].URL)
	}
	for _, got := range items {
		if got.Shortcode != "car" {
			t.Fatalf("expected shortcode car, got %q", got.Shortcode)
		}
		if got.Username != "carousel_user" || got.Caption != "cap" {
			t.Fatalf("expected username/caption, got %q/%q", got.Username, got.Caption)
		}
	}
}

func TestPickBestCandidate(t *testing.T) {
	candidates := []imageCandidate{
		{URL: "a", Width: 10, Height: 10},
		{URL: "b", Width: 40, Height: 5},
		{URL: "c", Width: 30, Height: 10},
	}
	if got := pickBestCandidate(candidates); got != "c" {
		t.Fatalf("expected best candidate c, got %q", got)
	}
}
