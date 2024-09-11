package util

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"

	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

/*
// info:[{inclusion 0.41357138752937317 [118 34 13 29]}
{inclusion 0.4607578217983246 [33 5 14 31]}
{inclusion 0.8377113342285156 [116 72 22 41]}
{inclusion 0.8540953397750854 [114 111 21 90]}]
*/

// Draw a square on the image
func DrawSquare(img image.Image, info []DefectData, squareColor color.RGBA) (url string) {
	// Create a new RGBA image to draw the squares on
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	// Define the square properties
	// squareWidth := 2

	// Draw a square for each defect
	for _, defect := range info {
		// Get the position of the defect
		pos := defect.Position

		// Define the rectangle for the square
		// rect := image.Rectangle{Min: image.Point{X: pos[0], Y: pos[1]}, Max: image.Point{X: pos[2], Y: pos[3]}}
		rect := image.Rect(pos[0], pos[1], pos[2], pos[3])

		// Draw the square
		//set the (i,j) of the square to the squareColor
		for i := rect.Min.X; i <= rect.Max.X; i++ {
			rgba.Set(i, rect.Min.Y, squareColor)
			rgba.Set(i, rect.Max.Y, squareColor)
		}
		for i := rect.Min.Y; i <= rect.Max.Y; i++ {
			rgba.Set(rect.Min.X, i, squareColor)
			rgba.Set(rect.Max.X, i, squareColor)
			// rgba.Set(i, rect.Min.X, squareColor)
			// rgba.Set(i, rect.Max.X, squareColor)

		}
		// conf := defect.Conf
		class := defect.Class

		//add the string of conf and class to the image
		// Create a new Point to draw the text near the square
		pt := fixed.P(pos[0], pos[1])

		// pt := fixed.P(10, 10)

		// Set the font and color for the text
		// fontSize := 10.0
		fontFace := basicfont.Face7x13

		textColor := color.RGBA{255, 0, 0, 255}

		// Draw the text onto the image
		drawer := &font.Drawer{
			Dst:  rgba,
			Src:  image.NewUniform(textColor),
			Face: fontFace,
			Dot:  pt,
		}
		drawer.DrawString(class)
		// drawer.DrawString(fmt.Sprintf("%s, %f", class, conf))

	}

	// Save the modified image to a new file
	output, err := os.Create("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	// Encode the image as jpg
	if err := jpeg.Encode(output, rgba, nil); err != nil {
		log.Fatal(err)
	}
	//save the image to the local
	// output, err := os.Create("output.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer output.Close()
	// // Encode the image as jpg
	// if err := jpeg.Encode(output, rgba, nil); err != nil {
	// 	log.Fatal(err)
	// }
	imgResult, err := os.Open("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgResult.Close()

	//upload image to Tencent COS
	outputMode := "result"
	tmp, err := UploadImageToTencentCloudCos(imgResult, outputMode)
	if err != nil {
		fmt.Println("Upload image to Tencent Cloud COS failed!")
		return "Error!"
	}

	fmt.Println("The square image was created successfully!")
	return tmp
}
