// Package mjpeg implements mjpeg streaming handlers with a simple API.
package mjpeg

import (
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

// A Handler is an http.Handler that streams mjpeg using an image stream. Encoding
// quality can be controlled using the Options parameters.
type Handler struct {
	Stream  func() image.Image
	Options *jpeg.Options
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	for {
		_, err := io.WriteString(w, "\r\n--frame\r\nContent-Type: image/jpeg\r\n\r\n")
		if err != nil {
			break
		}

		err = jpeg.Encode(w, h.Stream(), h.Options)
		if err != nil {
			break
		}

		_, err = io.WriteString(w, "\r\n")
		if err != nil {
			break
		}
	}
}
