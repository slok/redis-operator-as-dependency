package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	rfv1alpha2cli "github.com/spotahome/redis-operator/client/k8s/clientset/versioned/typed/redisfailover/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	log.Printf("listing all redis failover from on the default cluster (kubeconfig configuration)...")

	// Get local configuration (only for testing).
	kubehome := filepath.Join(homedir.HomeDir(), ".kube", "config")
	cfg, err := clientcmd.BuildConfigFromFlags("", kubehome)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load configuration: %s", err)
		os.Exit(1)
	}

	// Create a redis failover client.
	cli, err := rfv1alpha2cli.NewForConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed creating client: %s", err)
		os.Exit(1)
	}

	// List all failovers.
	rfs, err := cli.RedisFailovers("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed getting redis failovers: %s", err)
		os.Exit(1)
	}

	for _, rf := range rfs.Items {
		log.Printf("[%s] %s\n", rf.Namespace, rf.Name)
	}
}
