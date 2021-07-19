# Part 3 - Kubeify it
## Build and Usage
This is going to be a local kubernetes deployment so we will be using minikube. Start by starting the minikube environment.
```
minikube start
```

Next we are going to want to add environment variable that will allow minikube to see locally built docker images
```
eval $(minikube -p minikube docker-env) 
```

Now run the build_and_run script. This is going to build the webservice Dockerfile locally and spin up the Kubernetes manifests in `/Kubernetes`.
```
./build_and_run.sh
```

It is going to take a moment for all of the services to fully start up (the webservice will restart roughly 4 times as it waits for the requisite services to spin up). Once all services have successfully started, in a separate termainal, you will want to expose adminer (localhost:8080) and the webservice (localhost:8081/payloads).
```
// in a separate terminal run
minikube tunnel
```
Since we are running this kubernetes cluster locally with minikube, this step is required to expose loadbalanced enpoints.

Now you can simply test the webservice using an app such as Postman
