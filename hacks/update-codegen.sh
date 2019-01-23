#!/bin/sh
vendor/k8s.io/code-generator/generate-groups.sh all \
	github.com/kamolhasan/CRD/pkg/client \
	github.com/kamolhasan/CRD/pkg/apis \
	crd.com:v1