REPO=xxx
APP=microsvc-test
USER=xxx
PASS=xxx

docker build -t $REPO/$APP .

#Remove all docker images without a tag
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
