package mjpeg

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"net/http"
	"testing"

	"code.nfsmith.ca/nsmith/mjpeg"
)

func TestHandler(t *testing.T) {
	stream := func() image.Image {
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

	mux := http.NewServeMux()
	mux.Handle("/stream", Handler{stream, nil})
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Example() {
	stream := func() image.Image {
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

	mux := http.NewServeMux()
	mux.Handle("/stream", mjpeg.Handler{stream, &jpeg.Options{60}})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
