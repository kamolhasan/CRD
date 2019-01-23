# AppsCRD

### Code Generation for CustomResources and using them for CustomResourceDefinition


#### Prerequisites before using appsCRD

* `go get -u github.com/masudur-rahman/appsCRD` 
* `cd /home/$USER/go/src/github.com/masudur-rahman/appsCRD`
* `go install`

#### Available Commands

* `appsCRD` - Welcome from AppsCRD
* `appsCRD --help` - to get all available commands and usages
* `appsCRD create` - to create a CustomResourceDefinition named "customdeployments.apps.crd"
* `appsCRD delete` - to delete the created CustomResourceDefinition
* `appsCRD deploy` - to create a CustomDeployment named "my-new-custom-deploy"
* `appsCRD getDeploy` - to get all CustomDeployments
* `appsCRD deleteDeploy` - to delete the created CustomDeployment

##### Also we can get all the CustomDeployments using the following commands:

* `kubectl get customdeployments`
* `kubectl get customdeployement`
* `kubectl get customdeploy`
* `kubectl get cdeploy`
* `kubectl get cd`
