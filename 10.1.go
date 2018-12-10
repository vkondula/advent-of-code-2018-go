package main

import (
	"./lib"
	"strconv"
	"regexp"
	"image"
	"image/gif"
	"image/color"
	"os"
)

const OFFSET = 2000
const SIDE = 500

type Point struct {
	X int
	Y int
	VelocityX int
	VelocityY int
}

func (point *Point) Move(steps int) *Point{
	point.X += point.VelocityX * steps
	point.Y += point.VelocityY * steps
	return point
}

func Put (x int, y int, frame *image.Paletted) bool {
	offset := (y - frame.Rect.Min.Y) * frame.Stride + (x - frame.Rect.Min.X)
	if offset >= SIDE * SIDE || offset < 0 { return false }
	frame.Pix[offset] = 1
	return true
}


func main() {
	lines := lib.GetItems("inputs/10.txt")
	points := make([]Point, len(lines))
	re, _ := regexp.Compile(`position=< ?(-?\d+), *(-?\d+)> velocity=< ?(-?\d+), *(-?\d+)>`)
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		velocityX, _ := strconv.Atoi(matches[3])
		velocityY, _ := strconv.Atoi(matches[4])
		points[i] = Point{x, y, velocityX, velocityY}
	}
	rectangle := image.Rect(0, 0, SIDE, SIDE)
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
	}
	var frames []*image.Paletted
	steps := 100
	i := 0
	for i < 11000 {
		frame := image.NewPaletted(rectangle, palette)
		changed := false
		for i, point := range points {
			changed = Put(point.X, point.Y, frame) || changed
			points[i] = *point.Move(steps)
		}
		i = i + steps
		if changed {
			steps = 1
			frames = append(frames, frame)
		} else {
			steps = 100
		}
	}
	f, _ := os.OpenFile("10.1.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	delay := make([]int, len(frames))
	for i := range delay { delay[i] = 0 }
	gif.EncodeAll(f, &gif.GIF{
		Image: frames,
		Delay: delay,
	})
}