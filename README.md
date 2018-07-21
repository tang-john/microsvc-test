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

## Create Files
* mkdir -p /data/vm/vagrant/kubernetes/01-cluster-prometheus/logs
* Copy the Vagrantfile, 1-microsvc-deployment.yaml, and 2-microsvc-ingress.yaml to /data/vm/vagrant/kubernetes/01-cluster-prometheus/



## Start Kubernetes Cluster
* cd /data/vm/vagrant/kubernetes/01-cluster-prometheus/
* vagrant up
* wait for VMs to start.  The Kubernetes cluster will start once the VMs are all running. 

## SSH Into Kubemster
* cd /data/vm/vagrant/kubernetes/01-cluster-prometheus/
* vagrant ssh kubemaster


