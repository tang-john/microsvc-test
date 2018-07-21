Employee microservice for CRUD operations

Modify 1-build.sh and 2-push.sh to use your Docker repository.

Execute 1-build.sh to create the Docker image.
Execute 2-push.sh to push the Docker image to Docker hub.

Run the docker container: docker run -d -p 8080:8080 <image name>

Visit (the first three URLs are valid) :

http://localhost:8080
http://localhost:8080/james
http://localhost:8080/employee
http://localhost:8080/enis
