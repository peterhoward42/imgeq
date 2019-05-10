package imgeq

import (
	"fmt"
	"image"
	"os"

	// These imports are not used explicitly in the code below,
	// but are imported for their initialization side-effect, which allows
	// image.Decode to understand those image formats.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// AreEqual determines if two image files are identical pixel for pixel.
// It supports .png .jpeg and .gif files.
func AreEqual(fileA string, fileB string) (areEqual bool, err error) {
	if fileA == fileB {
		return false, fmt.
			Errorf("You specified the same file name for both files: %s", fileA)
	}
	readerA, err := os.Open(fileA)
	if err != nil {
		return false, fmt.Errorf("os.Open(): %v", err)
	}
	defer readerA.Close()

	readerB, err := os.Open(fileB)
	if err != nil {
		return false, fmt.Errorf("os.Open(): %v", err)
	}
	defer readerB.Close()

	imgA, _, err := image.Decode(readerA)
	if err != nil {
		return false, fmt.Errorf("image.Decode(): %v", err)
	}

	imgB, _, err := image.Decode(readerB)
	if err != nil {
		return false, fmt.Errorf("image.Decode(): %v", err)
	}

	boundsA := imgA.Bounds()
	boundsB := imgB.Bounds()

	if boundsA != boundsB {
		return false, nil
	}

	for x := boundsA.Min.X; x < boundsA.Max.X; x++ {
		for y := boundsA.Min.Y; y < boundsA.Max.Y; y++ {
			colorA := imgA.At(x, y)
			colorB := imgB.At(x, y)
			// Fail fast
			if colorA != colorB {
				return false, nil
			}
		}
	}
	return true, nil
}
