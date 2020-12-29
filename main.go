package main

import "./imgdow"

func main() {
	imageURL := imgdow.GetStringLine("Enter a URL: ")

	err := imgdow.DownloadImage(imageURL)
	if err != nil {
		panic(err)
	}
}
