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
* vi Vagrantfile
* paste the following 
Vagrant.configure("2") do |config|

  config.vm.define "kubemaster" do |v|
    v.vm.box = "johntang/kubemaster-prometheus"
    v.vm.box_version = "1.0.0"
    v.vm.hostname = "kubemaster"
    v.vm.network "private_network", ip: "172.16.0.40"
    v.vm.synced_folder ".", "/vagrant", :mount_options => ["dmode=777", "fmode=777"]
    config.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "2048"]
      vb.customize ["modifyvm", :id, "--cpus", "2"]
    end
  end

  config.vm.define "kubenode1" do |v|
    v.vm.box = "johntang/kubenode1-prometheus"
    v.vm.box_version = "1.0.0"
    v.vm.hostname = "kubenode1"
    v.vm.network "private_network", ip: "172.16.0.41"
    v.vm.synced_folder ".", "/vagrant", :mount_options => ["dmode=777", "fmode=777"]
    config.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "2048"]
      vb.customize ["modifyvm", :id, "--cpus", "1"]
    end
  end

  config.vm.define "kubenode2" do |v|
    v.vm.box = "johntang/kubenode2-prometheus"
    v.vm.box_version = "1.0.0"
    v.vm.hostname = "kubenode2"
    v.vm.network "private_network", ip: "172.16.0.42"
    v.vm.synced_folder ".", "/vagrant", :mount_options => ["dmode=777", "fmode=777"]
    config.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "2048"]
      vb.customize ["modifyvm", :id, "--cpus", "1"]
    end
  end


end
```

