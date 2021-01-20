package main

import (
	"log"

	"github.com/la880703/goav/avcodec"
	"github.com/la880703/goav/avdevice"
	"github.com/la880703/goav/avfilter"
	"github.com/la880703/goav/avformat"
	"github.com/la880703/goav/avutil"
	"github.com/la880703/goav/swresample"
	"github.com/la880703/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
