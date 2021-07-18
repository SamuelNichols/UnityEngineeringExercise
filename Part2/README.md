# Part 2 - Dockerize It
## Build and usage steps
As with part 1, docker-compose is being used to make the service deploying process easier. To test the environment run the following command.
```
./build_and_run.sh
```
This shell script will build the Dockerfile residing in the `WebService` folder then spin up rabbitmq, mysql, and the webservice with docker-compose. 

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

Note: The repeat messages of form below are a byproduct of the healthcheck implementation in the compose to allow for the webservice to wait for rabbitmq to be up. A better wait function would be ideal.
```
rabbitmq             | 2021-07-18 02:49:34.692 [info] <0.1111.0> Closing all channels from connection '127.0.0.1:34369 -> 127.0.0.1:5672' because it has been closed
```