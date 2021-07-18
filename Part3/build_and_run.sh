# Building Webservice Dockerfile
pushd WebService
docker build -t webservice-container .
popd
# Starting docker compose
docker-compose up
