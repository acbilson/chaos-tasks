set dotenv-load

# builds a development docker image.
build:
  COMMIT_ID=$(git rev-parse --short HEAD); \
  docker build \
  --target dev \
  -t acbilson/tasks:latest \
  -t acbilson/tasks:${COMMIT_ID} .

# starts a development docker image.
start:
	docker run --rm -it \
	--expose ${EXPOSED_PORT} -p ${EXPOSED_PORT}:5000 \
	-v ${SOURCE_SRC}:/mnt/src \
	-v ${CONTENT_SRC}:/mnt/data \
	--name tasks \
	acbilson/tasks:latest
