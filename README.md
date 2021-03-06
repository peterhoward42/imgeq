# A Go package to assess if two image files are identical, pixel for pixel

## Target use-case

Regression testing that images generated by your code remain identical
to a golden-reference imaged you stored before.

## Image Types Supported

- PNG
- GIF
- JPEG

## Implementation Notes

You might think you could just look at whether the files' content bytes were
the same. But you can't beause image files can contain meta data (e.g.
timestamp) that doesn't affect how they look, but do make their bytes
different.

The implementation is a very thin wrapper around Go's *image* package.

The filename suffices are immaterial; the *image* package sniffs out what
type of file it is from what it finds inside.

## Usage

    go get github.com/peterhoward42/imgeq

    import github.com/peterhoward42/imgeq

    areEqual, err := imgeq.AreEqual(filePathA, filePathB)

