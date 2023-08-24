package acectx

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)


type ACEContex struct {
	Domain         string `yaml:"domain"`
	DomainExternal string `yaml:"domainExternal"`
	StorageClass   struct {
		Db string `yaml:"db"`
	} `yaml:"storageClass"`
	Ingress struct {
		Certificate struct {
			Annotations interface{} `yaml:"annotations"`
			Auto        bool        `yaml:"auto"`
		} `yaml:"certificate"`
		LargeFile struct {
			Annotations struct {
				NginxIngressKubernetesIoProxyBodySize         string `yaml:"nginx.ingress.kubernetes.io/proxy-body-size"`
				NginxIngressKubernetesIoProxyBuffering        string `yaml:"nginx.ingress.kubernetes.io/proxy-buffering"`
				NginxIngressKubernetesIoProxyReadTimeout      string `yaml:"nginx.ingress.kubernetes.io/proxy-read-timeout"`
				NginxIngressKubernetesIoProxyRequestBuffering string `yaml:"nginx.ingress.kubernetes.io/proxy-request-buffering"`
				NginxIngressKubernetesIoProxySendTimeout      string `yaml:"nginx.ingress.kubernetes.io/proxy-send-timeout"`
			} `yaml:"annotations"`
		} `yaml:"largeFile"`
		Vouch struct {
			Annotations struct {
				NginxIngressKubernetesIoAuthResponseHeaders  string `yaml:"nginx.ingress.kubernetes.io/auth-response-headers"`
				NginxIngressKubernetesIoAuthSignin           string `yaml:"nginx.ingress.kubernetes.io/auth-signin"`
				NginxIngressKubernetesIoAuthSnippet          string `yaml:"nginx.ingress.kubernetes.io/auth-snippet"`
				NginxIngressKubernetesIoAuthURL              string `yaml:"nginx.ingress.kubernetes.io/auth-url"`
				NginxIngressKubernetesIoConfigurationSnippet string `yaml:"nginx.ingress.kubernetes.io/configuration-snippet"`
			} `yaml:"annotations"`
		} `yaml:"vouch"`
	} `yaml:"ingress"`
}

func New() ACEContex {

	var ac ACEContex
	return ac
	
}

func (aceCtx *ACEContex) LoadAceContext(contextYamlFile string){

	aceContextFile, err := os.ReadFile(contextYamlFile)
	if err != nil {
			log.Printf("error reading aceContextFile [%s], err: %v ", contextYamlFile, err)
	}

	err = yaml.Unmarshal(aceContextFile, aceCtx)
	if err != nil {
			log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("DEBUG: aceContext: \n%v", aceCtx )

}