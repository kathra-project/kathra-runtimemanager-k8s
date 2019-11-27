/*
 * Kathra Runtime Environment Manager
 *
 * KATHRA Runtime Environment Management API
 *
 * AppService functions
 */
package app_service

import (
	"context"
	"log"
	"os/exec"

	k8sClt "github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	api "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/models"
	svc "gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-go/kathra-runtimemanager-k8s/services"
)

type ContainerService struct {
	client *k8sClt.Client
}

func NewContainerService() ContainerService {
	return ContainerService{client: svc.KubernetesClientServiceInstance.Client}
}

var Instance = NewContainerService()

func (self ContainerService) List(namespaceProviderID string, applicationProviderID string) ([]*api.RuntimeContainer, error) {
	var podList corev1.PodList
	var list = []*api.RuntimeContainer{}
	labels := new(k8sClt.LabelSelector)
	labels.Eq("release", applicationProviderID)
	if err := self.client.List(context.Background(), namespaceProviderID, &podList, labels.Selector()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, pod := range podList.Items {
		for _, container := range mapContainerServiceFromPod(pod) {
			list = append(list, container)
		}
	}
	return list, nil
}

func (self ContainerService) Delete(namespaceProviderID string, applicationProviderID string, containerId string) error {
	var podList corev1.PodList
	labels := new(k8sClt.LabelSelector)
	labels.Eq("release", applicationProviderID)
	if err := self.client.List(context.Background(), namespaceProviderID, &podList, labels.Selector()); err != nil {
		log.Fatal(err)
		return err
	}
	for _, pod := range podList.Items {
		for _, container := range mapContainerServiceFromPod(pod) {
			if container.ProviderID == containerId {
				_, err := exec.Command("kubectl", "-n", namespaceProviderID, "exec", "-it", pod.GetMetadata().GetName(), "-c", container.Name, "--", "/bin/sh", "-c", "kill 1").Output()
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			}
		}
	}
	return nil
}

// Convert RuntimeContainer From Pod
func mapContainerServiceFromPod(pod *corev1.Pod) []*api.RuntimeContainer {
	var containers = []*api.RuntimeContainer{}
	for _, podContainer := range pod.Spec.Containers {
		var containerStatus = ""
		var imageHash = ""
		for _, status := range pod.Status.ContainerStatuses {
			if *status.Name == *podContainer.Name {
				if *status.Ready {
					containerStatus = "READY"
				}
				imageHash = *status.ImageID
			}
		}
		containers = append(containers, &api.RuntimeContainer{
			ProviderID:  pod.Metadata.GetName() + "/" + podContainer.GetName(),
			Image:       podContainer.GetImage(),
			MemoryLimit: podContainer.Resources.Limits["memory"].GetString_(),
			CPULimit:    podContainer.Resources.Limits["cpu"].GetString_(),
			Status:      containerStatus,
			ImageHash:   imageHash,
		})
	}
	return containers
}
