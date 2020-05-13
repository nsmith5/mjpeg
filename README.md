# Simple Motion JPEG Streaming

[![GoDoc](https://godoc.org/github.com/nsmith5/mjpeg?status.svg)](https://godoc.org/github.com/nsmith5/mjpeg)

Super simple mjpeg streaming in Go.

## Getting Started

Get to package with `go get github.com/nsmith5/mjpeg`. An MJPeg stream
can be built using any function that takes no arguments and returns an image.

```go
package main

import (
    "log"
    "net/http"

    "github.com/nsmith5/mjpeg"
)

func main() {
    stream := mjpeg.Handler{
        Stream: func()  image.Image {
            img := image.NewGray(image.Rect(0, 0, 100, 100))
            for i := 0; i < 100; i++ {
                for j := 0; j < 100; j++ {
                    n := rand.Intn(256)
                    gray := color.Gray{uint8(n)}
                    img.SetGray(i, j, gray)
                }
            }
            return img
        },
        Options: &jpeg.Options{Quality: 80},
    }

    mux := http.NewServeMux()
    mux.Handle("/stream", stream)
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```
