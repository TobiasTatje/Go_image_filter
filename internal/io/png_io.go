package io

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func ReadPng(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func WritePng(name string, img image.Image) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode: %v", err)
	}
	return err
}
