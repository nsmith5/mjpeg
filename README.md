# Simple Motion JPEG Streaming

[![GoDoc](https://godoc.org/github.com/gorilla/websocket?status.svg)](https://godoc.org/github.com/gorilla/websocket)

Super simple mjpeg streaming in Go.

## Getting Started

Get to package with `go get code.nfsmith.ca/nsmith/mjpeg`. An MJPeg stream
can be built using any function that takes no arguments and returns an image.

```go
package main

func stream() image.Image {
    img := image.NewGray(image.Rect(0, 0, 100, 100))
    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            n := rand.Intn(256)
            gray := color.Gray{uint8(n)}
            img.SetGray(i, j, gray)
        }
    }
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/stream", mjpeg.Handler{Stream: stream, Option: nil})
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```
