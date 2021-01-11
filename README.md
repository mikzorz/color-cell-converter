# Color Cell Converter #
## What is this?
I watched this video:
https://www.youtube.com/watch?v=Tfh0ytz8S0k
_('How "oldschool" graphics worked Part 1 - Commodore and Nintendo' by "The 8-bit Guy")_

I thought it was interesting and I wanted to make something, for practice with Go.

## How to use
`go install` or `go build`

`color-cell-converter -i inputFile -o outputFile ...`

### Command Line Flags
Flag | Type | Description
---- | ------ | --------------
-i | string | the image file that will be processed, using png, jpg or jpeg extension.
-o | string | the name of the image file to output to, using png, jpg or jpeg extension.
-w | int | width of each cell (default 8)
-h | int | height of each cell (default 8)
-p | int | palette to use; 0)commodore64, 1)pico8, 2)gameboy (default 0)
-t | bool | should the image retain transparency? (default=false)
-ti | int | index of color in palette that will replace fully transparent pixels if -t=false (default 0)
-s | bool | if -t=true, smooth edges of image? (default true)

### Examples
![Watermill photo by Paul Teysen, with 8x8 cells and C64 palette](https://raw.githubusercontent.com/mikzorz/color-cell-converter/main/examples/example_watermill.png)
- (original) https://unsplash.com/photos/hthCw4I19-A

![Tokyo street photo by Sergio Rola, with large cells and Pico-8 palette](https://raw.githubusercontent.com/mikzorz/color-cell-converter/main/examples/example_tokyo.png)
- (original) https://unsplash.com/photos/bVM7IO7pt7s

### _Problem(s)_
- Green tends to be sparse in the converted images.
- Pictures of big, green fields seem to end up either brown or dark blue.
- The color weighting may be off to some degree.
- I will either try to make it more accurate or just let the user adjust it themselves.
- It could also be the palettes' limited color selection.
