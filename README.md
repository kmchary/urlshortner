# urlshortner
Simple URL Shortener Service that stores the short url code in redis. Currently it is not storing the actual url so
redirecting service and other stuff is not implemented. I am using hexagonal (ports and adapters) architecture to 
structure the code. 

This will support redirect service as well.




##How to run
To run test this service we need to first run redis in docker container using the below command.

* clone the repository
* run redis in docker container (use default port)
  * sudo docker run --name my-redis -d redis
* go to the urlshortner/cmd/service directory
* run the below command 
  * go run main.go
* post  a request to localhost:8080 with the JSON body as shown below
  * {"Url":"www.google.com", "UserId":"kmchary"}
* take the short url code generated, and hit localhost:8080/short_url_code, you will be redirected
to the actual site


##How to run the app in Docker container
* clone the repository
* run redis in docker container (use default port)
  * sudo docker run --name my-redis -d redis
* go to the urlshortner directory
* build the docker image using the following command
  * docker build --tag urlshortenerapp .
* run the docker container using the following command
  * docker run --network=host -p 8080:8080 --name my_url_app urlshortenerapp:latest
* post  a request to localhost:8080 with the JSON body as shown below
  * {"Url":"www.google.com/something/something/something/text.htm", "UserId":"kmchary"}

docker run --network=host -p 8080:8080 --name my_url_app urlshortenerapp:latest