// Package mjpeg implements mjpeg streaming handlers with a simple API.
package mjpeg

import (
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

// An ImageStream is a function that, when called repeatedly, returns successive
// images
type ImageStream func() image.Image

func (s ImageStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	for {
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
