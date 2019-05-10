# A Go package to assess if two image files are identical

## Target use-case

To support the regression testing of code that generates images
programmatically. Specifically to check that the generated images
remain unchanged from a known-correct image that you have stored.

Particularly useful when the only practical way to validate the images is
for a human to look at them.

You care only that the images differ - i.e. a boolean result.

## Image Types Supported

Whatever golang.org/pkg/image supports. Which at time of writing is:

- PNG
- GIF
- JPEG

## Why not Just Compare the File Bytes or a Digest?

Because image files can contain meta data that doesn't affect what they
look like, but make their bytes different. (Like a timestamp).

## Usage

    go get github.com/peterhoward42/imgeq

    import github.com/peterhoward42/imgeq

    areEqual, err := imgeq.AreEqual(filePathA, filePathB)

