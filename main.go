// webcamqrscanner is an example of how to use a webcam to read a QR code.
// It used gocv for capturing image from the webcam and gozxing for finding the
// QR code and decoding it value.

package main // import "lazyhacker.dev/webcamqrscanner"

import (
	"flag"
	"fmt"
	"log"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"gocv.io/x/gocv"
)

var (
	deviceID = flag.Int("device-id", 0, "device ID of webcam")
)

func main() {

	flag.Parse()

	// Initialize the webcam.
	webcam, err := gocv.VideoCaptureDevice(*deviceID)
	defer webcam.Close()
	if err != nil {
		log.Fatalf("Unable to get video capture device. %v", err)
	}

	// Create a window on the desktop to show what the camera is capturing.
	window := gocv.NewWindow("Webcam QR Scanner")
	mat := gocv.NewMat()

	qrReader := qrcode.NewQRCodeReader()

	// Keep scanning the webam images until a QR code is found.
	for {
		webcam.Read(&mat)
		window.IMShow(mat)
		img, err := mat.ToImage()
		if err != nil {
			log.Fatalf("Unable to convert webcam data to image. %v", err)
		}

		// Scan for QR code from the webcam image.
		src := gozxing.NewLuminanceSourceFromImage(img)
		if err != nil {
			log.Fatalf("Unable to convert Luminance source. %v", err)
		}
		bmp, err := gozxing.NewBinaryBitmap(gozxing.NewGlobalHistgramBinarizer(src))
		//bmp, err := gozxing.NewBinaryBitmapFromImage(mat.ToImage())

		if result, err := qrReader.Decode(bmp, nil); err == nil {
			// QR code found!  Show the image used that found the QR code.
			window.IMShow(mat)
			fmt.Println(result)
			break
		}
		window.WaitKey(1)

	}

	// Wait until a key is pressed to stop the app.
	for {
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
