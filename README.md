# urlshortner
Simple URL Shortener Service that stores the short url code in redis. Currently it is not storing the actual url so
redirecting service and other stuff is not implemented. I am using hexagonal (ports and adapters) architecture to 
structure the code. It is not production ready code. Actually the logic for generating short url is not so perfect.
We can do it in multiple ways one of the method is mentioned below:

We can get 7 character length random url in base62 
We could just make a random choice for each character and check if this tiny url exists in DB or not. 
If it doesn’t exist return the tiny url else continue rolling/retrying.As more and more 7 characters 
short links are generated in Database, we would require 4 rolls before finding non-existing one 
short link which will slow down tiny url generation process.


##How to run
To run test this service we need to first run redis in docker container using the below command.

* clone the repository
* run redis in docker container (use default port)
  * sudo docker run --name my-redis -d redis
* go to the cmd/service directory
* run go run main.go
* post  a request to localhost:8081 with the JSON body as shown below
  * {"Url":"www.google.com/something/something/something/text.htm", "UserId":"kmchary"}



