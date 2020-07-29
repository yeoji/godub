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

func combineMono() {
	lFilePath := path.Join(dataDirectory(), "testing-stereo-left-channel.wav")
	lSegment, _ := godub.NewLoader().Load(lFilePath)

	rFilePath := path.Join(dataDirectory(), "testing-stereo-right-channel.wav")
	rSegment, _ := godub.NewLoader().Load(rFilePath)

	combined, _ := godub.NewAudioSegmentFromMono(lSegment, rSegment)
	fmt.Println(combined)

	// Export to file
	toFilePath := path.Join(tmpDataDirectory(), "combined-testing-stereo.wav")
	godub.NewExporter(toFilePath).
		WithDstFormat("wav").
		Export(combined)
}