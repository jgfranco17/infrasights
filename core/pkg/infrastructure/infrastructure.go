// kube/kube.go

package kube

import (
	"os"
)

func GetKubeNamespace() string {
	namespace := os.Getenv("NAMESPACE")
	if namespace != "" {
		return namespace
	}
	return "None"
}
