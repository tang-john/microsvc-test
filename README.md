# microsvc-test
Deploy a simple microservice written in Golang into a Istio service mesh on a Kubernetes cluster with Prometheus, Grafana, and Jaeger support. 

## Requirements
### Golang binaries https://golang.org/doc/install
### Vagrant
### VirtualBox
### Ubuntu 16.04

## Clone Repository 
* Start a command shell
* go get github.com/tang-john/microsvc-test

## Test Appication 
* cd ~/src/github.com/tang-john/microsvc-test
* go run main.go
* open a browser and navigate to the following URLs
** http://localhost:8080/
** http://localhost:8080/james
** http://localhost:8080/microsvc-test
* The first and second URL will return a JSON response.
* The last URL will return a text response.

## Create Vagrantfile
* mkdir -p /data/vm/vagrant/kubernetes/02-cluster-istio/logs
* cd /data/vm/vagrant/kubernetes/02-cluster-istio
* Copy the Vagrantfile to /data/vm/vagrant/kubernetes/02-cluster-istio/

