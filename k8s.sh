# Script that Builds and pushes image to docker so that it can be used in K8S

docker build src/. -t portal-kafka-consumer
docker tag portal-kafka-consumer ubiwhere/portal-kafka-consumer
docker push ubiwhere/portal-kafka-consumer