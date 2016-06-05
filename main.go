package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/html"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var addr = flag.String("b", ":8080", "Bind address, eg. 0.0.0.0:8080")
var confFile = flag.String("f", "conf.yaml", "Configuration file, eg. conf.yaml")

// TODO? https support
// proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

type substitution struct {
	Pattern     *regexp.Regexp
	ReplaceWith string
}

type config struct {
	Substitutions []substitution `yaml:"substitutions"`
}

func (sub *substitution) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var interSubstitution struct {
		Pattern     string `yaml:"pattern"`
		ReplaceWith string `yaml:"replace_with"`
	}

	if err := unmarshal(&interSubstitution); err != nil {
		return err
	}

	sub.Pattern = regexp.MustCompile(interSubstitution.Pattern)
	sub.ReplaceWith = interSubstitution.ReplaceWith
	return nil
}

func readConfig() config {
	file, err := os.Open(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	conf := config{}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func logConfig(conf config) {
	log.Println("Binding to ", *addr)
	log.Println("Applying substitutions:")
	for _, sub := range conf.Substitutions {
		log.Println(sub.Pattern, "->", sub.ReplaceWith)
	}
}

func main() {
	flag.Parse()
	conf := readConfig()
	logConfig(conf)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnResponse().Do(goproxy_html.HandleString(
		func(s string, ctx *goproxy.ProxyCtx) string {
			for _, sub := range conf.Substitutions {
				s = sub.Pattern.ReplaceAllString(s, sub.ReplaceWith)
			}

			return s
		}))

	log.Fatal(http.ListenAndServe(*addr, proxy))
}
