package mjpeg_test

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"net/http"

	"code.nfsmith.ca/nsmith/mjpeg"
)

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
	mux.Handle("/stream", mjpeg.ImageStream(stream))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
