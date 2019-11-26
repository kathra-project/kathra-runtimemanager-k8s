package services

import (
	"bytes"
	"strconv"
	"log"
	"os"
	"os/exec"
	"gopkg.in/yaml.v2"
	"github.com/jeremywohl/flatten"
	"encoding/json"
)

type HelmCatalogRepository struct {
	Name string 
	Url string
	Username string
	Password string
}

type HelmList struct {
	Next string `yaml:"Next,omitempty"`
	Releases []HelmRelease `yaml:"Releases,omitempty"`
}

type HelmRelease struct {
	Name string `yaml:"Name,omitempty"`
	Revision string `yaml:"Revision,omitempty"`
	Updated string `yaml:"Updated,omitempty"`
	Status string `yaml:"Status,omitempty"`
	Chart string `yaml:"Chart,omitempty"`
	AppVersion string `yaml:"AppVersion,omitempty"`
	Namespace string `yaml:"Namespace,omitempty"`
}

type HelmReleaseValues struct {
	Key string `yaml:"Name,omitempty"`
	Value string `yaml:"Revision,omitempty"`
}

type HelmClientService struct {  
	repositories []HelmCatalogRepository
	binary string
}

func NewHelmClientService() HelmClientService {  
	svc := HelmClientService {repositories: []HelmCatalogRepository{}, binary: "helm"}
	//svc.HelmInitKathraRepository()
	//svc.HelmUpdate()
	svc.GetReleasesFromNamespace("my-team")
	return svc
}

var HelmClientServiceInstance = NewHelmClientService()


func (service HelmClientService) HelmInitKathraRepository() {
	var kathraRepo = service.getKathraCatalogRepository()
	if kathraRepo.Name == "" {
		log.Println("Warn: No kathra repository configured")
		return
	}
	var repoLocalName, err = service.HelmFindLocalRepository(kathraRepo)
	if err != nil {
		log.Panic("Err: Unable to configure local chart repository, check env variable KATHRA_REPO_NAME, KATHRA_REPO_URL, KATHRA_REPO_CREDENTIAL_ID, KATHRA_REPO_SECRET ")
	} else {
		log.Println("Info: Kathra repository configured on Helm, local-name : " + repoLocalName)
	}
}

func (service HelmClientService) HelmFindLocalRepository(catalogRepository HelmCatalogRepository) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", service.binary+" repo list  | awk '{if ($2 == \""+catalogRepository.Url+"\") {print $1;}}'")
	var out bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		println("Err: " + stdErr.String())
		log.Println(err)
		return "", err
	}
	var repoName = out.String()
	if repoName == "" {
		println("Unable to find repository with url : " + catalogRepository.Url + ", add new repository ")
		var errAddRepo = service.HelmAddRepository(catalogRepository)
		if errAddRepo != nil {
			log.Println(errAddRepo)
			return "", errAddRepo
		}
		repoName = catalogRepository.Name
	}
	return repoName, nil
}

func (service HelmClientService) getKathraCatalogRepository() HelmCatalogRepository {
	var repoName = os.Getenv("KATHRA_REPO_NAME")
	if repoName == "" {
		repoName = "kathra-local"
	}
	return HelmCatalogRepository{
		Name:     os.Getenv("KATHRA_REPO_NAME"),
		Url:      os.Getenv("KATHRA_REPO_URL"),
		Username: os.Getenv("KATHRA_REPO_CREDENTIAL_ID"),
		Password: os.Getenv("KATHRA_REPO_SECRET")}
}

func (service HelmClientService) HelmUpdate() error {
	cmd := exec.Command(service.binary, "repo", "update")
	var stdErr, stdOut bytes.Buffer
	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut
	err := cmd.Run()
	log.Println(stdOut.String())
	if err != nil {
		log.Println(stdErr.String())
		log.Println(err)
		return err
	}
	return nil
}

type HelmEntry struct {
	Name          string
	LocalName     string
	VersionChart  string
	VersionApp    string
	Description   string
	RepositoryURL string
}

var entriesCached = []HelmEntry{}
var entriesAllVersionsCached = []HelmEntry{}

func (service HelmClientService) HelmAddRepository(catalogRepository HelmCatalogRepository) error {
	cmdAddRepo := exec.Command("/bin/bash", "-c", "helm repo add "+catalogRepository.Name+" --username="+catalogRepository.Username+" --password="+catalogRepository.Password+" "+catalogRepository.Url+"")
	errAddRepo := cmdAddRepo.Run()
	var stdErr bytes.Buffer
	cmdAddRepo.Stderr = &stdErr
	if errAddRepo != nil {
		println("Err: " + stdErr.String())
		log.Println(errAddRepo)
		return errAddRepo
	}
	return nil
}

func (service HelmClientService) GetReleasesFromNamespace(namespace string)([]*HelmRelease, error) {
	output, err := exec.Command(service.binary, "list","--namespace",namespace,"--output","yaml").Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var helmList HelmList
	errUnmarshal := yaml.Unmarshal(output, &helmList)
	if errUnmarshal != nil {
		log.Printf("error: %v", errUnmarshal)
		return nil, errUnmarshal
	}
	var releases []*HelmRelease;
	for i := range helmList.Releases {
		releases = append(releases, &helmList.Releases[i])
	}
	return releases, nil
}

func (service HelmClientService) GetReleaseFromNamespace(namespace string, release string)(*HelmRelease, error) {
	output, err := exec.Command(service.binary, "list","--namespace",namespace, release,"--output","yaml").Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var helmList HelmList
	errUnmarshal := yaml.Unmarshal(output, &helmList)
	if errUnmarshal != nil {
		log.Printf("error: %v", errUnmarshal)
		return nil, errUnmarshal
	}
	for i := range helmList.Releases {
		if (helmList.Releases[i].Name == release) {
			return &helmList.Releases[i], nil
		}
	}
	return nil, nil
}

func (service HelmClientService) InstallRelease(namespace string, release string, chart string, version string, values map[string]string) (error) {
	var cmd = service.binary+" install --name "+release+" "+chart+" --namespace "+namespace+" --version="+version
	
	for k, v := range values { 
		cmd = cmd + " --set "+k+"="+v
	}
	_, err := exec.Command("git", "bash", "-c", "/bin/bash -c '" + cmd + "'").Output()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (service HelmClientService) DeleteReleaseFromNamespace(namespace string, release string) (error) {
	// Helm 3
	//_, err := exec.Command(service.binary, "delete","--namespace",namespace, release,"--purge").Output()
	_, err := exec.Command(service.binary, "delete", release,"--purge").Output()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (service HelmClientService) GetReleaseValueFromNamespace(namespace string, release string)(map[string]string, error) {
	// Helm 3
	//output, err := exec.Command(service.binary, "get","values","--namespace",namespace,"--output","yaml").Output()
	output, err := exec.Command(service.binary, "get","values",release,"--output","json").Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	flat, errFlatten := flatten.FlattenString(string(output), "", flatten.DotStyle)
	if errFlatten != nil {
		log.Printf("error: %v", errFlatten)
		return nil, errFlatten
	}

	values := make(map[string]FlexString)
	errUnmarshal := json.Unmarshal([]byte(flat), &values)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	valuesAsString := make(map[string]string)
	for k, v := range values { 
		valuesAsString[k] = v.Value
	}
	return valuesAsString, nil
}

type FlexString struct {
	Value string
}
func (fi *FlexString) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		if (b[0] == 't' ) {
			*fi = FlexString{Value:  "true"}
		} else if (b[0] == 'f') {
			*fi = FlexString{Value:  "false"}
		} else {
			var i int
			if err := json.Unmarshal(b, &i); err != nil {
				return err
			}
			*fi = FlexString{Value:  strconv.Itoa(i)}
		}
	} else {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*fi = FlexString{Value: s}
	}
	return nil
}