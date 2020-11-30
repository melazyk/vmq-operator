package vernemq

import (
	"fmt"

	vernemqv1alpha1 "github.com/vernemq/vmq-operator/pkg/apis/vernemq/v1alpha1"
)

func volumeName(name string) string {
	return fmt.Sprintf("%s-db", prefixedName(name))
}

func deploymentName(name string) string {
	return fmt.Sprintf("%s-deployment", prefixedName(name))
}

func prefixedName(name string) string {
	return fmt.Sprintf("%s-%s", vernemqName, name)
}

func getHostname(instance *vernemqv1alpha1.VerneMQ) string {
	clusterName := instance.Spec.ClusterName

	if clusterName == "" {
		clusterName = "cluster.local"
	}
	return fmt.Sprintf("%s.%s.svc.%s", serviceName(instance.Name), instance.Namespace, clusterName)
}
