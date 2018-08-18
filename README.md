# Simple Motion JPEG Streaming

[![GoDoc](https://godoc.org/code.nfsmith.ca/nsmith/mjpeg?status.svg)](https://godoc.org/code.nfsmith.ca/nsmith/mjpeg)

Super simple mjpeg streaming in Go.

## Getting Started

Get to package with `go get code.nfsmith.ca/nsmith/mjpeg`. An MJPeg stream
can be built using any function that takes no arguments and returns an image.

```go
package main

import (
    "log"
    "net/http"

    "code.nfsmith.ca/nsmith/mjpeg"
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
