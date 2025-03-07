package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/microsoft/retina/test/multicloud/test/utils"
)

func TestRetinaEKSIntegration(t *testing.T) {
	t.Parallel()

	opts := &terraform.Options{
		TerraformDir: utils.ExamplesPath + "integration/retina-eks",

		Vars: map[string]interface{}{
			"prefix":               "test",
			"region":               "eu-west-1", // Dublin
			"retina_chart_version": utils.RetinaVersion,
			"retina_values": map[string]interface{}{
				// Example using a public image
				"image": map[string]interface{}{
					"tag":        "65b6244-linux-amd64",
					"repository": "ghcr.io/microsoft/retina/retina-agent",
				},
				"operator": map[string]interface{}{
					"tag": utils.RetinaVersion,
				},
				"logLevel": "info",
			},
		},
	}

	// clean up at the end of the test
	defer terraform.Destroy(t, opts)
	terraform.InitAndApply(t, opts)

	// get outputs
	caCert := utils.FetchSensitiveOutput(t, opts, "cluster_ca_certificate")
	host := utils.FetchSensitiveOutput(t, opts, "host")
	token := utils.FetchSensitiveOutput(t, opts, "access_token")

	// decode the base64 encoded cert
	caCertString := utils.DecodeBase64(t, caCert)

	// build the REST config
	restConfig := utils.CreateRESTConfigWithBearer(caCertString, token, host)

	// create a Kubernetes clientset
	clientSet, err := utils.BuildClientSet(restConfig)
	if err != nil {
		t.Fatalf("Failed to create Kubernetes clientset: %v", err)
	}

	// test the cluster is accessible
	utils.TestClusterAccess(t, clientSet)

	retinaPodSelector := utils.PodSelector{
		Namespace:     "kube-system",
		LabelSelector: "k8s-app=retina",
		ContainerName: "retina",
	}

	timeOut := time.Duration(90) * time.Second
	// check the retina pods are running
	result, err := utils.ArePodsRunning(clientSet, retinaPodSelector, timeOut)
	if !result {
		t.Fatalf("Retina pods did not start in time: %v\n", err)
	}

	// check the retina pods logs for errors
	utils.CheckPodLogs(t, clientSet, retinaPodSelector)

	// TODO: add more tests here
}
