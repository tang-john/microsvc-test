REPO=xxx
APP=microsvc-test
USER=xxx
PASS=xxx

docker login -u $USER -p $PASS
docker push $USER/$APP
