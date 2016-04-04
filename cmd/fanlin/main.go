package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/jobtalk/fanlin/lib/conf"
	"github.com/jobtalk/fanlin/lib/handler"
	"github.com/jobtalk/fanlin/lib/logger"
	"github.com/sirupsen/logrus"
)

var confList = []string{
	"/etc/lvimg.json",
	"/etc/lvimg.cnf",
	"/usr/local/etc/lvimg.json",
	"/usr/local/etc/lvimg.cnf",
	"/usr/local/lvimg/lvimg.json",
	"/usr/local/lvimg/lvimg.cnf",
	"./lvimg.json",
	"./lvimg.cnf",
	"./conf.json",
	"~/.lvimg.json",
	"~/.lvimg.cnf",
}

func main() {
	conf := func() *configure.Conf {
		for _, confName := range confList {
			conf := configure.NewConfigure(confName)
			if conf != nil {
				return conf
			}
		}
		return nil
	}()
	if conf == nil {
		log.Fatalln("Can not read configure.")
	}

	notFoundImagePath := flag.String("nfi", conf.NotFoundImagePath(), "not found image path")
	errorLogPath := flag.String("err", conf.ErrorLogPath(), "error log path")
	accessLogPath := flag.String("log", conf.AccessLogPath(), "access log path")
	port := flag.Int("p", conf.Port(), "port")
	port = flag.Int("port", conf.Port(), "port")
	localImagePath := flag.String("li", conf.LocalImagePath(), "local image path")
	maxProcess := flag.Int("cpu", conf.MaxProcess(), "max process.")
	debug := flag.Bool("debug", false, "debug mode.")
	flag.Parse()
	conf.Set("404_img_path", *notFoundImagePath)
	conf.Set("error_log_path", *errorLogPath)
	conf.Set("access_log_path", *accessLogPath)
	conf.Set("port", *port)
	conf.Set("local_image_path", *localImagePath)
	conf.Set("max_process", *maxProcess)

	loggers := map[string]logrus.Logger{
		"err":    *logger.NewLogger(conf.ErrorLogPath()),
		"access": *logger.NewLogger(conf.AccessLogPath()),
	}

	if *debug {
		fmt.Println(conf)
	}

	runtime.GOMAXPROCS(conf.MaxProcess())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.MainHandler(w, r, conf, loggers)
	})
	http.HandleFunc("/healthCheck", handler.HealthCheckHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port()), nil)
}