package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	inFile := flag.String("i", "", "the image file that will be processed, using png, jpg or jpeg extension.")
	outFile := flag.String("o", "", "the name of the image file to output to, using png, jpg or jpeg extension.")
	wFlag := flag.Int("w", 8, "width of each cell")
	hFlag := flag.Int("h", 8, "height of each cell")
	tFlag := flag.Bool("t", false, "should the image retain transparency? (default=false)")
	palFlag := flag.Int("p", 0, "palette to use; 0)commodore64, 1)pico8, 2)gameboy (default 0)")
	tiFlag := flag.Int("ti", 0, "index of color in palette that will replace fully transparent pixels if -t=false (default 0)")
	sFlag := flag.Bool("s", true, "if -t=true, smooth edges of image?")

	flag.Parse()

	fmt.Println("input image file:", *inFile)
	inSplit := strings.Split(*inFile, ".")
	inExt := inSplit[len(inSplit)-1]
	if inExt != "png" && inExt != "jpeg" && inExt != "jpg" {
		log.Fatal("Please use png, jpg or jpeg file extension for input file.")
	}

	fmt.Println("output image file:", *outFile)
	if *outFile == "" {
		log.Fatal("Please provide a filename to output to.")
	}

	outSplit := strings.Split(*outFile, ".")
	outExt := outSplit[len(outSplit)-1]
	if outExt != "png" && outExt != "jpeg" && outExt != "jpg" {
		log.Fatal("Please use png, jpg or jpeg file extension for output file.")
	}

	cellW, cellH := *wFlag, *hFlag
	if cellW < 0 {
		fmt.Println("-w cannot be less than 0, defaulting to 8")
		cellW = 8
	} else {
		fmt.Println("cell width:", *wFlag)
	}

	if cellH < 0 {
		fmt.Println("-h cannot be less than 0, defaulting to 8")
		cellH = 8
	} else {
		fmt.Println("cell height:", *hFlag)
	}

	palette := palettes[0]
	if *palFlag < 0 || *palFlag >= len(palettes) {
		fmt.Println("invalid palette selected, defaulting to", palNames[0])
	} else {
		fmt.Println("palette:", *palFlag, ", name:", palNames[*palFlag])
		palette = palettes[*palFlag]
	}

	fmt.Println("keep transparency:", *tFlag)
	useTransparency := *tFlag

	tpIndex := *tiFlag
	if useTransparency {
		palette = append(palette, color.RGBA{0, 0, 0, 0})
		tpIndex = len(palette) - 1
	} else {
		if tpIndex < 0 || tpIndex >= len(palette) {
			fmt.Println("invalid index given to -ti, defaulting to index 0:", palette[0])
			tpIndex = 0
		} else {
			fmt.Println("transparency replacement color:", *tiFlag, ",", palette[*tiFlag])
		}
	}

	fmt.Println("smoothing:", *sFlag)
	smoothing := *sFlag

	//

	imgFile, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()

	out := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += cellH {
		for x := bounds.Min.X; x < bounds.Max.X; x += cellW {
			// Choose 2 colors to use in cell
			var color1, color2 color.Color
			tally := map[int]int{}
			// For each pixel in cell, get nearest color in palette
			for cy := 0; cy < cellH && y+cy < bounds.Max.Y; cy++ {
				for cx := 0; cx < cellW && x+cx < bounds.Max.X; cx++ {
					// If transparent pixel, use preselected 'transparent' color.
					if _, _, _, a := img.At(x+cx, y+cy).RGBA(); a < 65535 {
						tally[tpIndex]++
						if smoothing {
							continue
						}
						break
					}
					// else, figure out which color to use
					nearestDistance := int(^uint(0) >> 1)
					indexOfNearest := 0
					for i, palColor := range palette {
						dist := getDistBetweenColors(img.At(x+cx, y+cy), palColor)
						if dist < nearestDistance {
							indexOfNearest = i
							nearestDistance = dist
						}
					}
					tally[indexOfNearest]++
				}
			}
			// Get 2 colors with most tallies.
			colIndexes := make([]int, 0, len(tally))
			for k := range tally {
				colIndexes = append(colIndexes, k)
			}
			sort.Slice(colIndexes, func(i int, j int) bool {
				return tally[colIndexes[i]] > tally[colIndexes[j]]
			})

			color1 = palette[colIndexes[0]]

			var closestColor color.Color

			// Apply colors
			for cy := 0; cy < cellH && y+cy < bounds.Max.Y; cy++ {
				for cx := 0; cx < cellW && x+cx < bounds.Max.X; cx++ {
					closestColor = color1
					if len(colIndexes) > 1 {
						color2 = palette[colIndexes[1]]
						nearestDistance := getDistBetweenColors(img.At(x+cx, y+cy), color1)

						dist := getDistBetweenColors(img.At(x+cx, y+cy), color2)
						if dist < nearestDistance {
							closestColor = color2
						}
					}
					cr, cg, cb, ca := closestColor.RGBA()
					out.SetRGBA(x+cx, y+cy, color.RGBA{uint8(cr), uint8(cg), uint8(cb), uint8(ca)})

				}
			}
		}
	}

	outImg, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	if outExt == "png" {
		png.Encode(outImg, out)
	} else if outExt == "jpg" || outExt == "jpeg" {
		jpeg.Encode(outImg, out, nil)
	}
	outImg.Close()
}

func getDistBetweenColors(pix, pal color.Color) int {
	r, g, b, a := pix.RGBA()
	pr, pg, pb, pa := pal.RGBA()
	rdist := r - pr
	gdist := g - pg
	bdist := b - pb
	adist := a - pa
	// Alter weights for color correction
	return int((rdist * rdist / 9 * 2) +
		(gdist * gdist / 9 * 6) +
		(bdist * bdist / 9) +
		(adist * adist))
}
