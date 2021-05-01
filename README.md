# Utilizando K8S

This challenge is split in three parts. All solutions are based on `kubernetes`.

1.NGINX Web Server

Follow the below steps to accomplish the first part of the challenge:
 - Use the Nginx Alpine base image;
 - Make three replicas available;
 - When someone accesses the LoadBalance extern IP of the created service, it should be displayed in the browser: `Code.education Rocks`. 

2.MYSQL Configuration

To accomplish this part of the challenge, the following steps are needed:
 - Configure a MySQL database server using the `kubernetes` concept of `deployment`;
 - Use `secret` along with the environment variable;
 - Use `persistent volume claim` to store the data that come from MySQL.

3.Go Challenge

The steps to accomplish this challenge are:
 - Create a Golang app that starts a web server on port 8000. When the server is accessed a message(in bold) saying `Code.education Rocks!` should be displayed. The responsible for formatting the message is a function named `greeting` that receives a string as parameter and returns it between the HTML bold tags(`<b></b>`);
 - Create a unit test for the greeting function;
 - Start a CI process on `Google Cloud Build` platform to make sure that each created Pull Request triggers the unit test;
 - Build an optmized image of the Golang app and push it to DokcerHub;
 - Using `Kubernetes`, start a LoadBalancer service that allows to access the Go application running in the cluster.

**Note**: The Go project image can be pulled from [here](https://hub.docker.com/r/jpedronascimentofilho/go-k8s).

## Goal
The goals of this challenge is to learn `kubernetes`' concepets and how it works in practice.

## Usage

In order to be able to run this challenge you must have `minikube` and `kubectl` installed previously.

1.NGINX Web Server

Start the cluster:
```bash
$ minikube start
```

Create the objects and run the service:
```bash
$ cd nginx/
$ kubectl apply -f configmap.yaml
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
```

See the objects' status:
```bash
$ kubectl get configmap
$ kubectl get deployment
$ kubectl get service
```

Access the service:
```bash
$ minikube service nginx-k8s-service
```

To delete all the objects and stop the cluster:
```bash
$ kubectl delete configmap nginx-html
$ kubectl delete deployment nginx-k8s
$ kubectl delete service nginx-k8s-service
$ minikube stop
```

2.MYSQL Configuration

Start the cluster:
```bash
$ minikube start
```

Create the objects and run the service:
```bash
$ cd mysql/
$ kubectl apply -f persistent-volume.yaml
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
```

See the objects' status:
```bash
$ kubectl get persistentvolumeclaim
$ kubectl get deployment
$ kubectl get service
```

To delete all the objects and stop the cluster:
```bash
$ kubectl delete persistentvolumeclaim mysql-pv-claim
$ kubectl delete deployment mysql-server
$ kubectl delete service mysql-service
$ minikube stop
```

3.Go Challenge

Start the cluster:
```bash
$ minikube start
```

Create the objects and run the service:
```bash
$ cd go/k8s/
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
```

See the objects' status:
```bash
$ kubectl get deployment
$ kubectl get service
```

Access the service:
```bash
$ minikube service go-service
```

To delete all the objects and stop the cluster:
```bash
$ kubectl delete deployment go-server
$ kubectl delete service go-service
$ minikube stop
```

