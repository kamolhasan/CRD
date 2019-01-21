package main

import (
	"github.com/appscode/kutil/tools/clientcmd"
	apiexv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"
)

func CreateClientSet() *crdclientset.Clientset {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err)
	}

	clientset, err := crdclientset.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	return clientset

}

func main() {

	clientset := CreateClientSet()

	customRD := apiexv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "customservices.crd.com",
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
				Kind:       "CustomService",
				Plural:     "customservices",
				Singular:   "customservice",
				ShortNames: []string{"cs"},
			},
			Scope:"Namespaced",
		},
	}

	_,err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&customRD)

	if err!= nil{
		panic(err)

	}

}
