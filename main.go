package main

import (
	"github.com/appscode/kutil/tools/clientcmd"
	"github.com/kamolhasan/CRD/pkg/apis/crd.com/v1"
	customclient "github.com/kamolhasan/CRD/pkg/client/clientset/versioned"
	apiexv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

func CreateCustomResource() {
	log.Println("Creating Custom Resource...!")

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := crdclientset.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	customRD := apiexv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "custompods.crd.com",
		},
		Spec: apiexv1beta1.CustomResourceDefinitionSpec{
			Group: "crd.com",
			Versions: []apiexv1beta1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
				},
			},
			Names: apiexv1beta1.CustomResourceDefinitionNames{
				Kind:       "CustomPod",
				Plural:     "custompods",
				Singular:   "custompod",
				ShortNames: []string{"cp"},
			},
			Scope: "Namespaced",
		},
	}

	_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&customRD)

	if err != nil {
		panic(err)

	}
	log.Println("Custom Resource Created!")
}

func CreateCustomPod() {

	log.Println("Creating Custom Pod...!")

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := customclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	cPod := v1.CustomPod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mycustompod",
			Labels: map[string]string{
				"app": "custom",
			},
			Namespace: "default",
		},
		Spec: v1.CustomPodSpec{
			Containers: []v1.CustomContainer{
				{
					Name:            "busy-box",
					Image:           "busybox",
					ImagePullPolicy: "always",
				},
			},
			RestartPolicy: "always",
		},
	}

	_, err = clientset.CrdV1().CustomPods("default").Create(&cPod)
	if err != nil {
		panic(err)
	}
	log.Println("Custom Pod Created!")

}

func DeleteCRD() {

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := customclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	err = clientset.CrdV1().CustomPods("default").Delete("mycustompod", &metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}

	clientset2, err := crdclientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	err = clientset2.ApiextensionsV1beta1().CustomResourceDefinitions().Delete("custompods.crd.com", &metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}

	log.Println("CleanUp completed!")
}

func main() {

	CreateCustomResource()

	log.Println("Press ctrl+C to create Pod")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	CreateCustomPod()

	log.Println("Press ctrl+C to cleanup everything")
	signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	DeleteCRD()

}
