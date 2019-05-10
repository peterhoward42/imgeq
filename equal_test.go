package imgeq

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/image/colornames"
)

var backgroundColor color.Color = colornames.Red
var spotColor color.Color = colornames.Yellow
var differentSpotColor color.Color = colornames.Green

func TestErrorHandlingOfIOErrors(t *testing.T) {
	assert := assert.New(t)
	invalidFilePath := ""
	_, err := AreEqual(invalidFilePath, "foo")
	assert.EqualError(err, "os.Open(): open : no such file or directory")
}

func TestErrorHandlingForInvalidFileContents(t *testing.T) {
	assert := assert.New(t)
	corruptFileName := "iscorrupt.png"
	err := ioutil.WriteFile(corruptFileName, []byte("notdotpngbytes"), 0644)
	assert.NoError(err)
	defer os.Remove(corruptFileName)

	otherFileName := "other.png"
	err = makeRectWithSpotPng(
		200, 200, otherFileName, backgroundColor, spotColor)
	assert.NoError(err)
	defer os.Remove(otherFileName)

	_, err = AreEqual(corruptFileName, otherFileName)
	assert.EqualError(err, "image.Decode(): image: unknown format")
}

func TestErrorHandlingWhenBothFilesAreSameName(t *testing.T) {
	assert := assert.New(t)
	_, err := AreEqual("fileA", "fileA")
	assert.EqualError(err,
		"You specified the same file name for both files: fileA")
}

func TestWhenAreEqualReturnsTrue(t *testing.T) {
	assert := assert.New(t)

	err := makeRectWithSpotPng(
		200, 200, "A.png", backgroundColor, spotColor)
	assert.NoError(err)
	defer os.Remove("A.png")

	err = makeRectWithSpotPng(
		200, 200, "B.png", backgroundColor, spotColor)
	assert.NoError(err)
	defer os.Remove("B.png")

	areEqual, err := AreEqual("A.png", "B.png")
	assert.NoError(err)
	assert.True(areEqual)
}

func TestWhenDifferReturnsFalse(t *testing.T) {
	assert := assert.New(t)

	err := makeRectWithSpotPng(
		200, 200, "A.png", backgroundColor, spotColor)
	assert.NoError(err)
	defer os.Remove("A.png")

	err = makeRectWithSpotPng(
		200, 200, "B.png", backgroundColor, differentSpotColor)
	assert.NoError(err)
	defer os.Remove("B.png")

	areEqual, err := AreEqual("A.png", "B.png")
	assert.NoError(err)
	assert.False(areEqual)
}

// makeRectWithSpotPng saves a .png file for an image filled with
// one solid colour, except for a small region that is a different
// colour.
func makeRectWithSpotPng(width int, height int, fileName string,
	backgroundColor color.Color, spotColor color.Color) error {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{backgroundColor},
		image.ZP, draw.Src)
	srcImg := &image.Uniform{spotColor}
	sr := srcImg.Bounds()
	dp := image.Point{10, 10}
	r := image.Rectangle{dp, dp.Add(sr.Size())}
	draw.Draw(m, r, srcImg, sr.Min, draw.Src)
	outF, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("os.Create(): %v", err)
	}
	defer outF.Close()
	png.Encode(outF, m)
	return nil
}
