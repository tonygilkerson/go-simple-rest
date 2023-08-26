package env

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Environment struct {
	ContextFile    string
	AceContext aceContex
}

type aceContex struct {
	AceEnvVersion  string `yaml:"aceEnvVersion,omitempty"`
	Domain         string `yaml:"domain,omitempty"`
	DomainExternal string `yaml:"domainExternal,omitempty"`
	StorageClass   struct {
		Db string `yaml:"db,omitempty"`
	} `yaml:"storageClass,omitempty"`
}

func NewEnv(contextFile string) Environment {

	var env Environment
	env.ContextFile = contextFile
	return env

}

func (env *Environment) loadAceContext() error {

	aceContextYamlFile := env.ContextFile
	log.Printf("Using aceContextYamlFile: %v", aceContextYamlFile)

	aceContext, err := os.ReadFile(aceContextYamlFile)
	if err != nil {
		log.Printf("error reading aceContext [%s], err: %v ", aceContext, err)
		return err
	}

	err = yaml.Unmarshal(aceContext, &env.AceContext)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}

	return nil

}

func (env *Environment) EnvContextHandler(w http.ResponseWriter, r *http.Request) {

	env.loadAceContext()
	b, _ := json.Marshal(env.AceContext)
	out := fmt.Sprintf("%v", string(b))
	fmt.Fprint(w, out)

}

func (env *Environment) EnvVersionHandler(w http.ResponseWriter, r *http.Request) {

	env.loadAceContext()
	aceEnvVer := env.AceContext.AceEnvVersion
	out := fmt.Sprintf("%v", aceEnvVer)
	fmt.Fprint(w, out)

}

func (env *Environment) EnvVersionBadgeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store")

	env.loadAceContext()

	t, _ := template.New("EnvVersion").Parse(
		`<svg xmlns="http://www.w3.org/2000/svg" width="188" height="20">
			<linearGradient id="b" x2="0" y2="100%">
					<stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
					<stop offset="1" stop-opacity=".1"/>
			</linearGradient>
			<mask id="a">
					<rect width="188" height="20" rx="3" fill="#fff"/>
			</mask>
			<g mask="url(#a)">
									<path fill="#555" d="M0 0 h130 v20 H0 z"/>
									<path fill="#44cc11" d="M108 0 h188 v20 H108 z"/>
									<path fill="url(#b)" d="M0 0 h188 v20 H0 z"/>
			</g>
			<g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
					<text x="55" y="15" fill="#010101" fill-opacity=".3">ACE Environment</text>
					<text x="55" y="14">ACE Environment</text>
					<text x="140" y="15" fill="#010101" fill-opacity=".3">v{{ .AceEnvVersion }}</text>
					<text x="140" y="14">v{{ .AceEnvVersion }}</text>
			</g>
		</svg>`)

	t.Execute(w, env.AceContext)
}
