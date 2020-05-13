package mjpeg

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func noise() image.Image {
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			n := rand.Intn(256)
			gray := color.Gray{uint8(n)}
			img.SetGray(i, j, gray)
		}
	}
	return img
}

func TestOneFrame(t *testing.T) {
	stop := false
	stream := func() (image.Image, error) {
		if !stop {
			stop = true
			return noise(), nil
		}
		return nil, ErrorEndOfStream
	}

	handler := Handler{
		Next:    stream,
		Options: nil,
	}

	req := httptest.NewRequest("GET", "http://example.com/stream", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 on single frame stream")
	}
}

func Example() {
	stream := func() (image.Image, error) {
		img := image.NewGray(image.Rect(0, 0, 100, 100))
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				n := rand.Intn(256)
				gray := color.Gray{uint8(n)}
				img.SetGray(i, j, gray)
			}
		}
		return img, nil
	}

	handler := Handler{
		Next:    stream,
		Options: &jpeg.Options{Quality: 60},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
