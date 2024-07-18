# docker build --build-arg GOOS=linux --build-arg GOARCH=amd64  -f deploy/Dockerfile . -t rounak316/kaniko_test:nightly
# docker buildx build --platform linux/amd64  -f deploy/Dockerfile  -t rounak316/kaniko_test:nightly
# For linux
docker build --build-arg TARGETOS=linux --build-arg TARGETARCH=amd64  -f deploy/Dockerfile . -t rounak316/kaniko_test:nightly

docker push rounak316/kaniko_test:nightly

# docker build -f deploy/Dockerfile . -t 192.168.29.6:5000/meddler-xyz/watchdog
# docker push  192.168.29.6:5000/meddler-xyz/watchdog