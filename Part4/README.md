# Part 4 - Productionize It
## Description/Implementation
In part 3, we moved the webservice, queue, and sql store to Kubernets using deployment and service clusters. Kubernetes was made to scale and what we did in part 3 has pretty much gotten us to the point where we can scale to 100k+ RPS. From the beginning, we utilized a message queue to allow a single stable point that we could push messages to. Now all we would have to do to scale our webservice would be to increase the number of replicas in the `webservice-deployment.yaml` and Kubernetes would handle the routing and loadbalancing. We could simply implement this by adding replicas to the `webservice-deployment.yaml` spec and statically increasing the number to match a max load (currently, replicas is set to 3). A much better solution would be to use the autoscale feature supplied by `kubectl`. Below would be an example of this command to scale the webservice based on load.
```
kubectl autoscale deployment webservice --max 1000 --min 10 --cpu-percent 50
```


To test this, simply use the instructions from part 3 then utilize the above command after everything has been set up.
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
