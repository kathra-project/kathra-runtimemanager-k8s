package main


import (
	environment "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services/environment"
)

func main() {
	environment.Instance.GetEnvironments()
}
