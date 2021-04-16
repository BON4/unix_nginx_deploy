package main

import (
	"os"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func createRGB(w, h int, wr io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	if err := png.Encode(wr, img); err != nil {
		return err
	}

	return nil
}

func createGrey(w, h int, wr io.Writer) error {
	imgByte := make([]byte, w*h)
	if _, err := rand.Read(imgByte); err != nil {
		return err
	}

	img := image.NewGray(image.Rect(0, 0, w, h))
	img.Pix = imgByte

	if err := png.Encode(wr, img); err != nil {
		return err
	}

	return nil
}

func main() {
	port, ok := os.LookupEnv("PROXY_IN_PORT")
	if !ok {
		port = "8080"
	}

	http.HandleFunc("/grey", func(w http.ResponseWriter, r *http.Request) {
		err := createGrey(1000, 1000, w)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	http.HandleFunc("/rgb", func(w http.ResponseWriter, r *http.Request) {
		err := createRGB(1000, 1000, w)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	log.Printf("Listening on port :%s", port)
	log.Println(http.ListenAndServe(":"+port, nil))
}