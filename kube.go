/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var clientset *kubernetes.Clientset

func connect_k8s() {
	var kubeconfig *string

	if os.Getenv("ENV") == "K8S" {

	}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
func testListdeployments() {
	deps, _ := clientset.AppsV1().Deployments("wordpresses").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=wordpress",
	}) //label selector: app.kubernetes.io/name=wordpress
	for _, dep := range deps.Items {
		fmt.Printf("dep: %s\n", dep.Name)

		/*zero := int32(0)
		dep.Spec.Replicas = &zero
		clientset.AppsV1().Deployments("wordpresses").Update(context.TODO(), &dep,metav1.UpdateOptions{})*/
	}
}
