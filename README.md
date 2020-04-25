# webcamqrscanner

webcamqrscanner is an example for capture and decoding a QR code from a webcam.

It relies on [GoCV](http://gocv.io) for getting the image from the webcam and
[GoZxing](http://github.com/makiuchi-d/gozxing) for identifying and decoding the
QR code.

## Prerequisite

webcamqrscanner requires [GoCV](http://gocv.io) which provides a binding to
OpenCV 4.  GoCV requires CGO to compile so there are additional steps to build
and compile OpenCV 4.

## Install

Once [GoCV](http://gocv.io) is installed then install webcamqrscanner with:

```
go get lazyhacker.dev/webcamqrscanner
```
