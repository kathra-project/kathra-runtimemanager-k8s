package services

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jeremywohl/flatten"
	"gopkg.in/yaml.v2"
)

// HelmCatalogRepository struct
type HelmCatalogRepository struct {
	Name     string
	Url      string
	Username string
	Password string
}

// HelmList struct
type HelmList struct {
	Releases []HelmRelease `yaml:"Releases,omitempty"`
}

// HelmRelease struct
type HelmRelease struct {
	Name       string `yaml:"name,omitempty"`
	Revision   string `yaml:"revision,omitempty"`
	Updated    string `yaml:"updated,omitempty"`
	Status     string `yaml:"status,omitempty"`
	Chart      string `yaml:"chart,omitempty"`
	AppVersion string `yaml:"app_version,omitempty"`
	Namespace  string `yaml:"namespace,omitempty"`
}

// HelmReleaseValues struct
type HelmReleaseValues struct {
	Key   string `yaml:"Name,omitempty"`
	Value string `yaml:"Revision,omitempty"`
}

// HelmClientService struct
type HelmClientService struct {
	repositories []HelmCatalogRepository
	binary       string
}

// NewHelmClientService create new instance
func NewHelmClientService() HelmClientService {
	svc := HelmClientService{repositories: []HelmCatalogRepository{}, binary: "helm3"}
	//svc.HelmInitKathraRepository()
	//svc.HelmUpdate()
	svc.GetReleasesFromNamespace("my-team")
	return svc
}

// HelmClientServiceInstance instance
var HelmClientServiceInstance = NewHelmClientService()

// HelmInitKathraRepository init local repository
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

// HelmFindLocalRepository return local repository name
func (service HelmClientService) HelmFindLocalRepository(catalogRepository HelmCatalogRepository) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", service.binary+" repo list  | awk '{if ($2 == \""+catalogRepository.Url+"\") {print $1;}}'")
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

// HelmUpdate update local repo
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

// HelmEntry helm entry
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

// HelmAddRepository add new repository
func (service HelmClientService) HelmAddRepository(catalogRepository HelmCatalogRepository) error {
	cmdAddRepo := exec.Command("/bin/sh", "-c", "helm repo add "+catalogRepository.Name+" --username="+catalogRepository.Username+" --password="+catalogRepository.Password+" "+catalogRepository.Url+"")
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

// GetReleasesFromNamespace get all release from namespace
func (service HelmClientService) GetReleasesFromNamespace(namespace string) ([]*HelmRelease, error) {
	output, err := exec.Command(service.binary, "list", "--namespace", namespace, "--output", "yaml").Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var helmReleases []HelmRelease
	errUnmarshal := yaml.Unmarshal(output, &helmReleases)
	if errUnmarshal != nil {
		log.Printf("error: %v", errUnmarshal)
		return nil, errUnmarshal
	}
	var releases []*HelmRelease
	for i := range helmReleases {
		releases = append(releases, &helmReleases[i])
	}
	return releases, nil
}

// GetReleaseFromNamespace get all release from namespace and release name
func (service HelmClientService) GetReleaseFromNamespace(namespace string, release string) (*HelmRelease, error) {
	output, err := exec.Command(service.binary, "list", "--namespace", namespace, "--output", "yaml").Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var helmReleases []HelmRelease
	errUnmarshal := yaml.Unmarshal(output, &helmReleases)
	if errUnmarshal != nil {
		log.Printf("error: %v", errUnmarshal)
		return nil, errUnmarshal
	}
	for i := range helmReleases {
		if helmReleases[i].Name == release {
			return &helmReleases[i], nil
		}
	}
	return nil, nil
}

// InstallRelease install release
func (service HelmClientService) InstallRelease(namespace string, release string, chart string, version string, values map[string]string) error {
	var cmd = service.binary + " install " + release + " " + chart + " --namespace " + namespace
	if version != "" {
		cmd = cmd + " --version=" + version
	}
	for k, v := range values {
		cmd = cmd + " --set " + k + "=" + v
	}
	log.Println(cmd)
	_, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteReleaseFromNamespace delete release
func (service HelmClientService) DeleteReleaseFromNamespace(namespace string, release string) error {
	// Helm 3
	_, err := exec.Command(service.binary, "delete", "--namespace", namespace, release, "--purge").Output()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetReleaseValueFromNamespace get release values
func (service HelmClientService) GetReleaseValueFromNamespace(namespace string, release string) (map[string]string, error) {

	valuesAsString := make(map[string]string)
	output, err := exec.Command(service.binary, "get", "values", "--namespace", namespace, release, "--output", "yaml").Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if strings.TrimSpace(string(output)) == "null" {
		return valuesAsString, nil
	}

	flat, errFlatten := flatten.FlattenString(string(output), "", flatten.DotStyle)
	if errFlatten != nil {
		log.Printf("error: %v", errFlatten)
		return nil, errFlatten
	}

	values := make(map[string]flexString)
	errUnmarshal := json.Unmarshal([]byte(flat), &values)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	for k, v := range values {
		valuesAsString[k] = v.Value
	}
	return valuesAsString, nil
}

type flexString struct {
	Value string
}

func (fi *flexString) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		if b[0] == 't' {
			*fi = flexString{Value: "true"}
		} else if b[0] == 'f' {
			*fi = flexString{Value: "false"}
		} else {
			var i int
			if err := json.Unmarshal(b, &i); err != nil {
				return err
			}
			*fi = flexString{Value: strconv.Itoa(i)}
		}
	} else {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*fi = flexString{Value: s}
	}
	return nil
}
