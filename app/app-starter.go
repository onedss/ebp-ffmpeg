package app

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/onedss/ebp-ffmpeg/mytool"
	"github.com/onedss/ebp-ffmpeg/service"
	"log"
	"os"
)

func StartApp() {
	log.Println("ConfigFile -->", mytool.ConfFile())
	sec := mytool.Conf().Section("service")
	svcConfig := &service.Config{
		Name:        sec.Key("name").MustString("EbpFFmpeg_Service"),
		DisplayName: sec.Key("display_name").MustString("EbpFFmpeg_Service"),
		Description: sec.Key("description").MustString("EbpFFmpeg_Service"),
	}

	httpPort := mytool.Conf().Section("http").Key("port").MustInt(51182)
	oneHttpServer := NewOneHttpServer(httpPort)
	p := &application{}
	p.AddServer(oneHttpServer)

	var s, err = service.New(p, svcConfig)
	if err != nil {
		log.Println(err)
		mytool.PauseExit()
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" || os.Args[1] == "stop" {
			figure.NewFigure("Ebp-Proxy", "", false).Print()
		}
		log.Println(svcConfig.Name, os.Args[1], "...")
		if err = service.Control(s, os.Args[1]); err != nil {
			log.Println(err)
			mytool.PauseExit()
		}
		log.Println(svcConfig.Name, os.Args[1], "ok")
		return
	}
	figure.NewFigure("Ebp-Proxy", "", false).Print()
	if err = s.Run(); err != nil {
		log.Println(err)
		mytool.PauseExit()
	}
}
