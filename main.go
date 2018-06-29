package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// An ImageStream is a function that, when called repeatedly, returns successive
// images
type ImageStream func() image.Image

func (s ImageStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	for {
		time.Sleep(time.Second)

		_, err := io.WriteString(w, "\r\n--frame\r\nContent-Type: image/jpeg\r\n\r\n")
		if err != nil {
			break
		}

		err = jpeg.Encode(w, s(), &jpeg.Options{Quality: 10})
		if err != nil {
			break
		}

		_, err = io.WriteString(w, "\r\n")
		if err != nil {
			break
		}
	}
}

func stream() image.Image {
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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", ImageStream(stream))
	log.Fatal(http.ListenAndServe(":8080", mux))
	return
}
