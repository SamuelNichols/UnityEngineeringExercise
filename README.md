# UnityEngineeringExercise
## Part 1 - The Web Service
Write a web service in any language that takes in a JSON payload, does some basic validation against an expected message format and content, puts that payload into a queue of your choice, and stores the relevant info in a relational DB.

Example valid payload:
```
{
	"ts": "1530228282",
	"sender": "testy-test-service",
	"message": {
		"foo": "bar",
		"baz": "bang"
	},
	"sent-from-ip": "1.2.3.4",
	"priority": 2

}
```

Validation rules:
- “ts” must be present and a valid Unix timestamp
- “sender” must be present and a string
- “message” must be present, a JSON object, and have at least one field set
- If present, “sent-from-ip” must be a valid IPv4 address
- All fields not listed in the example above are invalid, and should result in the message being rejected.

## Part 2 - Dockerize It
Write a Dockerfile to produce a container that will run your service, and one to spin up an instance of the queue you used in Part 1

## Part 3 - Kubeify it
Write Kubernetes manifests to create a Service and Deployment object for both your service and the queue

## Part 4 - Productionize It
Now that you have your service, describe/implement your approach to making it ready for production at scale (100k RPS).
