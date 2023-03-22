docker build -f deploy/Dockerfile . -t rounak316/kaniko_test:nightly
docker push rounak316/kaniko_test:nightly

# docker build -f deploy/Dockerfile . -t 192.168.29.6:5000/meddler-xyz/watchdog
# docker push  192.168.29.6:5000/meddler-xyz/watchdog