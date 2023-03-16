docker build --tag workout-api . &&
docker tag workout-api europe-west9-docker.pkg.dev/workout-377920/docker-images/workout-api &&
docker push europe-west9-docker.pkg.dev/workout-377920/docker-images/workout-api &&
kubectl rollout restart deployment workout-server -n default

