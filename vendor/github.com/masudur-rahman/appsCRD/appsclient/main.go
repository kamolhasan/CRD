package appsclient

import (
	"fmt"
	"github.com/masudur-rahman/appsCRD/pkg/apis/apps.crd/v1alpha1"
	cdclientset "github.com/masudur-rahman/appsCRD/pkg/client/clientset/versioned"
	cor1v1 "k8s.io/api/core/v1"
	crdapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"time"

	//"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func initiate() *rest.Config {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal("Error building config file")
	}
	return config
	//clientSet := crdclientset.NewForConfigOrDie(config)
	//return clientSet
}

func CreateCRD() {
	log.Println("Creating Custom Resource Definition")
	clientSet := crdclientset.NewForConfigOrDie(initiate())

	apicrd := &crdapi.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CustomResourceDefinition",
			APIVersion: "apiextensions.k8s.io/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "customdeployments.apps.crd",
		},
		Spec: crdapi.CustomResourceDefinitionSpec{
			Group: "apps.crd",
			Versions: []crdapi.CustomResourceDefinitionVersion{
				{
					Name:    "v1alpha1",
					Served:  true,
					Storage: true,
				},
			},
			Scope: crdapi.NamespaceScoped,
			Names: crdapi.CustomResourceDefinitionNames{
				Plural:   "customdeployments",
				Singular: "customdeployment",
				Kind:     "CustomDeployment",
				ShortNames: []string{
					"customdeploy",
					"cd",
					"cdeploy",
				},
			},
		},
	}

	time.Sleep(1 * time.Second)

	_, err := clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Create(apicrd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("CustomResourceDefinition Created...!")

}

func DeleteCRD() {
	log.Println("Deleting CustomResourceDefinition..!")
	clientSet := crdclientset.NewForConfigOrDie(initiate())

	time.Sleep(1 * time.Second)

	if err := clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(
		"customdeployments.apps.crd",
		&metav1.DeleteOptions{},
	); err != nil {
		log.Println(err)
	}

	log.Println("CustomResourceDefinition Deleted..!")
}

func CreateCustomDeployment() {
	log.Println("Creating CustomDeployment")

	var int322 = int32(5)

	clientset := cdclientset.NewForConfigOrDie(initiate())

	customDeploy := &v1alpha1.CustomDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CustomDeployment",
			APIVersion: "v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-new-custom-deploy",
			Labels: map[string]string{
				"app": "appscrd",
			},
		},
		Spec: v1alpha1.CustomDeploymentSpec{
			Replicas: &int322,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "appscrd",
				},
			},
			Template: v1alpha1.CustomPodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "custom-deploy-pod-template",
					Labels: map[string]string{
						"app": "appscrd",
					},
				},
				Spec: cor1v1.PodSpec{
					Containers: []cor1v1.Container{
						{
							Name:            "appscodeserver",
							Image:           "masudjuly02/appscodeserver",
							ImagePullPolicy: "IfNotPresent",
						},
					},
					RestartPolicy: "Alwyas",
				},
			},
		},
	}
	_, err := clientset.AppsV1alpha1().CustomDeployments("default").Create(customDeploy)

	time.Sleep(1 * time.Second)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("CustomDeployment created successfully..!")
}

func DeleteCustomDeployment() {
	log.Println("Deleting CustomDeployment..!")
	time.Sleep(1 * time.Second)

	clientset := cdclientset.NewForConfigOrDie(initiate())

	err := clientset.AppsV1alpha1().CustomDeployments("default").Delete("my-new-custom-deploy", &metav1.DeleteOptions{})

	if err != nil {
		log.Println(err)
	}

	log.Println("CustomDeployment deleted successfully..!")
}

func GetCustomDeployment() {
	log.Println("Showing all CustomDeployments...")
	time.Sleep(1*time.Second)

	clientset := cdclientset.NewForConfigOrDie(initiate())

	deployList, err := clientset.AppsV1alpha1().CustomDeployments("default").List(metav1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CustomDeployment names:")
	for _, deploy := range deployList.Items {
		fmt.Println("\t", deploy.Name)
	}
}