# Part 1 - The Web Service
## Build and Usage

docker-compose is being used to make the deployment of the sql and rabbitmq service faster and more reliable. To begin testing, run the following command.
```
docker-compose up
```

Wait for the mysql and rabbitmq services to finish spinning up then run the webservice.
```
pushd WebService && go run main.go && popd
```
You can now test the webservice
- Using an app such as Postman, you can POST to the endpoint `localhost:8081/payload`
- You can view the sql database with adminer at `http://localhost:8080/`
    - server: mysql_container
    - username: root
    - password: rootpassword
    - database: mydb
- You can view the queue at `http://localhost:15672/#/queues/%2F/payloads`
    - username: guest
    - password: guest
