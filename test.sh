docker build -f deploy/Dockerfile . -t rounak316/kaniko_test
docker run rounak316/kaniko_test  --no-push --context=git://github.com/Djkusik/AsynchronousCORSScanner.git --force
# docker build -f deploy/Dockerfile . -t 192.168.29.6:5000/meddler-xyz/watchdog
# docker push  192.168.29.6:5000/meddler-xyz/watchdog