package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("Error loading in-cluster config:", err)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating clientset:", err)
		return
	}

	ctx := context.Background()

	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error listing pods:", err)
		return
	}

	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			err := checkpointContainer(pod.Name, container.Name, pod.Status.HostIP)
			if err != nil {
				fmt.Printf("Error creating checkpoint for container %s in pod %s: %v\n", container.Name, pod.Name, err)
			} else {
				fmt.Printf("Checkpoint created for container %s in pod %s\n", container.Name, pod.Name)
			}
		}
	}
}

func checkpointContainer(podName, containerName, hostIP string) error {
	// Load CA certificate
	caCertPath := "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return fmt.Errorf("could not read CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tokenPath := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return fmt.Errorf("could not read token: %v", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Skip TLS cert. verfication
		RootCAs:            caCertPool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{
		Transport: transport,
	}

	url := fmt.Sprintf("https://%s:10250/checkpoint/default/%s/%s", hostIP, podName, containerName)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("could not create HTTP request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %v", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response from Kubelet: %s", string(body))
	}

	fmt.Printf("Response from Kubelet: %s\n", string(body))
	return nil
}
