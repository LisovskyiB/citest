#!/usr/bin/env bash
docker login -u $DOCKER_USER -p $DOCKER_PASS
export TAG=`if [ "$TRAVIS_BRANCH" == "master" ]; then echo "latest"; else echo $TRAVIS_BRANCH ; fi`
docker build -f Dockerfile -t $TRAVIS_REPO_SLUG:$TAG .
docker tag $TRAVIS_REPO_SLUG:$TAG $TRAVIS_REPO_SLUG:$TAG
#docker tag $TRAVIS_REPO_SLUG:$TAG $TRAVIS_REPO_SLUG:travis-$TRAVIS_BUILD_NUMBER
docker push $TRAVIS_REPO_SLUG