package main

import (
	"fmt"
	"github.com/yeoji/godub"
	"path"
)

func splitToMono() {
	filePath := path.Join(dataDirectory(), "code-geass.mp3")
	segment, _ := godub.NewLoader().Load(filePath)

	monoSegments, _ := segment.SplitToMono()
	fmt.Println(monoSegments)
}