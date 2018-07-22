# microsvc-test
Deploy a simple microservice written in Golang into a Istio service mesh on a Kubernetes cluster with Prometheus, Grafana, and Jaeger support. 

## Requirements
* Golang binaries https://golang.org/doc/install
* Vagrant
* VirtualBox
* Ubuntu 16.04
* 4 available cores: 2 cores for kubemaster, 1 core each for kubenode1 and kubenode2
* 6 GB RAM
* IP Addresses listed below
** 172.16.0.40 for kubemaster
** 172.16.0.41 for kubenode1
** 172.16.0.42 for kubenode2
** 172.16.0.130 - 172.16.0.140 for Kubernetes LoadBalancer and Ingress use.

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
### Inject Envoy Proxy into Deployment
* cd /vagrant
* ls -ltr
* You should see the YAML files
* kubectl apply -f <(istioctl kube-inject -f /vagrant/1-microsvc-deployment.yaml)
* kubectl get pods
* You will see two pods that begin with "microsvc-test". Assume one of them is microsvc-test-845b49968-jnq84
* kubeclt describe pod/microsvc-test-845b49968-jnq84
* Replace the command above with the pod that is for your installation. Kubernetes will generate a random string after "microsvc-test". Make sure you see "Started container".


### Create Ingress 
* kubectl apply -f /vagrant/2-microsvc-ingress.yaml
* kubectl get ingress
* You should see an entry for microsvc-test. Notice the IP address. It should be a value like 172.16.0.131. The port should be 80.  
* Visit the URLs below

** http://172.16.0.131/

** http://172.16.0.131/james

** http://172.16.0.131/microsvc-test

