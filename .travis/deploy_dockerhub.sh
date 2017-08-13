#!/bin/sh
if [ "$TRAVIS_BRANCH" = "master" ]; then
    TAG="latest"
else
    TAG="$TRAVIS_BRANCH"
fi
docker login -u="$DOCKER_USER" -p="$DOCKER_PASS"
docker build -f Dockerfile -t $TRAVIS_REPO_SLUG:$TAG .

docker tag $TRAVIS_REPO_SLUG:$COMMIT $TRAVIS_REPO_SLUG:$TAG
docker tag $TRAVIS_REPO_SLUG:$COMMIT $TRAVIS_REPO_SLUG:travis-$TRAVIS_BUILD_NUMBER

docker push $TRAVIS_REPO_SLUG