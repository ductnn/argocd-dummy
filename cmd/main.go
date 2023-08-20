package main

import (
	"flag"
	"fmt"

	"github.com/ductnn/argocd-dummy/internal/driver/argocd"
)

func main() {
	// Define command-line flags
	address := flag.String("address", "localhost:8080", "ArgoCD server address")
	token := flag.String("token", "my-foo-account-token", "ArgoCD API token")
	projectName := flag.String("project", "foo", "ArgoCD project name")

	// Parse the command-line flags
	flag.Parse()

	connection := argocd.Connection{
		Address: *address,
		Token:   *token,
	}

	client, err := argocd.NewClient(&connection)
	if err != nil {
		panic(err)
	}

	createProject, err := client.CreateProject(*projectName)
	if err != nil {
		panic(err)
	}

	fmt.Println(createProject.UID)

	err = client.AddDestination(createProject.Name, "server", "namespace", "name")
	if err != nil {
		panic(err)
	}

	getProject, err := client.GetProject(*projectName)
	if err != nil {
		panic(err)
	}

	fmt.Println(getProject.Namespace)

	err = client.DeleteProject(getProject.Name)
	if err != nil {
		panic(err)
	}

	clusters, err := client.GetClusters()
	if err != nil {
		panic(err)
	}

	for _, cluster := range clusters {
		fmt.Println(cluster.Name)
	}
}
