docker build -f deploy/Dockerfile . -t rounak316/kaniko_test
# docker run -v /tmp/store rounak316/kaniko_test  --no-push --context=git://github.com/Djkusik/AsynchronousCORSScanner.git --force --tar-path=/tmp/store/image.tar
docker run -v /tmp/store/:/workspace/  rounak316/kaniko_test --context=git://github.com/meddler-io/test-docker-image.git --force   --no-push --tarPath=/workspace/image.tar
# docker build -f deploy/Dockerfile . -t 192.168.29.6:5000/meddler-xyz/watchdog
# docker push  192.168.29.6:5000/meddler-xyz/watchdog