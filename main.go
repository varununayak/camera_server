package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
        fmt.Println(img)
	}
}