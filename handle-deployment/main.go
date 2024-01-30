package main

import (
	"log"
	"path/filepath"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	corev1 "k8s.io/api/core/v1"
)

func main() {
	homePath := homedir.HomeDir()
	if homePath == "" {
		log.Fatal("failed to get the home dir")
	}

	kubeconfig := filepath.Join(homePath, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	dpClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

	log.Println("create Deployment")
	if err := createDeployment(dpClient); err != nil {
		log.Fatal(err)
	}

	<-time.Tick(1 * time.Minute)

	log.Println("update Deployment")
	if err := updateDeployment(dpClient); err != nil {
		log.Fatal(err)
	}

	<-time.Tick(1 * time.Minute)

	log.Println("delete Deployment")
	if err := deleteDeployment(dpClient); err != nil {
		log.Fatal(err)
	}

	<-time.Tick(1 * time.Minute)

	log.Println("end")
}
