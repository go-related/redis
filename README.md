# README #

Here I will demonstrate communication between 2 microservices, for simplicity reason we will have both of them into the same repository.
To communicate we will use redis and to store data we will use postgresSQL.
* Service 1
  * this will be a simple service with a Crud and will publish an event when create book call will happen.
* Service 2
  * this will be the interested. once the event will be published this will send email(just log sm) to appropriate addresses.


### Prerequisites ###

* Redis docker 
* configuration for the service (there will be a sample on the repository)

### Notes ###
* build the application
  ``` code 
  go build -o bin/redis
   ```

* run service1
  ``` code 
  ./bin/redis service1 
  ```
