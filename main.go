package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"

	// import _ "net/http/pprof" is profiler, see https://golang.org/pkg/net/http/pprof/
	_ "net/http/pprof"

	logs "github.com/sirupsen/logrus"
)

// global variables used in this module
var _tdir, _base string

// Configuration stores configuration parameters
type Configuration struct {
	Port   int    `json:"port"`   // DAS port number
	Base   string `json:"base"`   // base path of the server
	Static string `json:"static"` // location of DAS templates
}

// Config variable represents configuration object
var Config Configuration

// String returns string representation of Config
func (c *Configuration) String() string {
	return fmt.Sprintf("<Config port=%d static=%s>", c.Port, c.Static)
}

// parseConfig parse given config file
func parseConfig(configFile string) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		logs.WithFields(logs.Fields{"configFile": configFile}).Fatal("Unable to read", err)
		return err
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		logs.WithFields(logs.Fields{"configFile": configFile}).Fatal("Unable to parse", err)
		return err
	}
	return nil
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var templates CMSTemplates
	tmplData := make(map[string]interface{})
	tmplData["Base"] = _base
	tmplData["Time"] = time.Now()
	/* Example of markdown support in templates
	mdfile, err := ioutil.ReadFile(_tdir + "/dashboards.md")
	if err == nil {
		tmplData["Dashboards"] = template.HTML(blackfriday.MarkdownCommon([]byte(mdfile)))
	}
	*/
	tmplData["Dashboards"] = template.HTML(templates.Dashboards(_tdir, tmplData))
	tmplData["Sources"] = template.HTML(templates.Sources(_tdir, tmplData))
	tmplData["Training"] = template.HTML(templates.Training(_tdir, tmplData))
	tmplData["Shifters"] = template.HTML(templates.Shifters(_tdir, tmplData))
	tmplData["Issues"] = template.HTML(templates.Issues(_tdir, tmplData))
	tmplData["Meetings"] = template.HTML(templates.Meetings(_tdir, tmplData))
	page := templates.Main(_tdir, tmplData)
	w.Write([]byte(page))
}

func server(configFile string) {
	err := parseConfig(configFile)
	if err != nil {
		logs.WithFields(logs.Fields{"Time": time.Now(), "Config": configFile}).Error("Unable to parse")
	}
	tdir := Config.Static
	base := Config.Base
	port := Config.Port
	_base = base
	_tdir = fmt.Sprintf("%s/templates", tdir)
	styles := fmt.Sprintf("%s/css", tdir)
	images := fmt.Sprintf("%s/images", tdir)
	scripts := fmt.Sprintf("%s/js", tdir)
	http.Handle(base+"/css/", http.StripPrefix(base+"/css/", http.FileServer(http.Dir(styles))))
	http.Handle(base+"/js/", http.StripPrefix(base+"/js/", http.FileServer(http.Dir(scripts))))
	http.Handle(base+"/images/", http.StripPrefix(base+"/images/", http.FileServer(http.Dir(images))))
	http.HandleFunc("/", defaultHandler)
	server := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
	}
	server.ListenAndServe()
}

func info() string {
	goVersion := runtime.Version()
	tstamp := time.Now()
	return fmt.Sprintf("git={{VERSION}} go=%s date=%s", goVersion, tstamp)
}

func main() {
	var config string
	flag.StringVar(&config, "config", "config.json", "server config file")
	flag.Parse()
	server(config)
}
