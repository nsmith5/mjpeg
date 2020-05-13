// Package mjpeg implements mjpeg streaming handlers with a simple API.
package mjpeg

import (
	"errors"
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

// ErrorEndOfStream signals the end of the Motion JPEG frames from a MJPEG
// stream
var ErrorEndOfStream = errors.New("End of Motion JPEG Stream")

// A Handler is an http.Handler that streams mjpeg using an image stream. Encoding
// quality can be controlled using the Options parameters.
type Handler struct {
	Next    func() (image.Image, error)
	Options *jpeg.Options
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	boundary := "\r\n--frame\r\nContent-Type: image/jpeg\r\n\r\n"
	for {
		img, err := h.Next()
		if err != nil {
			return
		}

		n, err := io.WriteString(w, boundary)
		if err != nil || n != len(boundary) {
			return
		}

		err = jpeg.Encode(w, img, h.Options)
		if err != nil {
			return
		}

		n, err = io.WriteString(w, "\r\n")
		if err != nil || n != 2 {
			return
		}
	}
}
