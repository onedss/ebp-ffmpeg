package main

import (
	"fmt"
	"github.com/onedss/ebp-ffmpeg/app"
	"github.com/onedss/ebp-ffmpeg/buildtime"
	"github.com/onedss/ebp-ffmpeg/mytool"
	"log"
)

var (
	gitCommitCode string
	buildDateTime string
)

func main() {
	log.SetPrefix("[Ebp-FFmpeg] ")
	if mytool.Debug {
		log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	} else {
		log.SetFlags(log.LstdFlags)
	}
	buildtime.BuildVersion = fmt.Sprintf("%s.%s", buildtime.BuildVersion, gitCommitCode)
	buildtime.BuildTimeStr = fmt.Sprintf("<%s> %s", buildtime.BuildTime.Format(mytool.DateTimeLayout), buildDateTime)
	mytool.Info("BuildVersion:", buildtime.BuildVersion)
	mytool.Info("BuildTime:", buildtime.BuildTimeStr)
	app.StartApp()
}
