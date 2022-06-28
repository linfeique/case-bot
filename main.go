package main

import (
	"io"
	"os"

	youtube "github.com/kkdai/youtube/v2"
)

func main() {
	client := youtube.Client{Debug: true}

	video, err := client.GetVideo("https://www.youtube.com/watch?v=BaW_jenozKc")
	if err != nil {
		panic(err)
	}

	// Typically youtube only provides separate streams for video and audio.
	// If you want audio and video combined, take a look a the downloader package.
	format := video.Formats.FindByQuality("medium")
	reader, _, err := client.GetStream(video, format)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		panic(err)
	}
}
