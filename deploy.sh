docker build -f deploy/Dockerfile . -t rounak316/kaniko_test
docker push rounak316/kaniko_test

# docker build -f deploy/Dockerfile . -t 192.168.29.6:5000/meddler-xyz/watchdog
# docker push  192.168.29.6:5000/meddler-xyz/watchdog