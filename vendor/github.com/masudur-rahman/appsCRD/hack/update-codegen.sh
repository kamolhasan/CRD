#!/bin/sh
vendor/k8s.io/code-generator/generate-groups.sh all \
	github.com/masudur-rahman/appsCRD/pkg/client \
	github.com/masudur-rahman/appsCRD/pkg/apis \
	apps.crd:v1alpha1