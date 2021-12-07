package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// cmdArguments := []string{"-i", "D:/work/test.mp4", "test1.gif"}
	// cmd := exec.Command("ffmpeg", cmdArguments...)

	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Printf("command output: %q", out.String())
	// TODO 加入图片压缩
	PicToGif2()
}

func PicToGif2() {
	var path, output string
	var delay int

	path = "./res"
	output = "test2.gif"
	delay = 30

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	anim := gif.GIF{}
	for _, info := range files {
		//fmt.Println("add one")
		if !strings.HasSuffix(info.Name(), "jpg") {
			continue
		}
		fmt.Println("show file:", info.Name())
		f, err := os.Open(path + "/" + info.Name())
		if err != nil {
			fmt.Printf("Could not open file %s. Error: %s\n", info.Name(), err)
			return
		}
		defer f.Close()
		img, s, err := image.Decode(f)
		if err != nil {
			fmt.Printf("s:%s,decode fail:%v", s, err)
			return
		}

		paletted := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.Point{})

		anim.Image = append(anim.Image, paletted)
		anim.Delay = append(anim.Delay, delay)
	}

	f, _ := os.Create(output)
	defer f.Close()
	gif.EncodeAll(f, &anim)

}
