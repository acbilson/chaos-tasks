#!/bin/bash
. .env

ENVIRONMENT=$1

case $ENVIRONMENT in

dev)
  echo "starts local development container..."
  echo "container name is ${IMAGE_NAME}"
  echo "code resides... src: ${CODE_SOURCE_SRC} and dst: ${CODE_SOURCE_DST}"
  echo "content resides... src: ${DEV_CONTENT_SRC} and dst: ${CONTENT_DST}"
  echo "on exposed port ${EXPOSED_PORT} and container port ${CONTAINER_PORT}"
  docker run --rm -it \
    --expose ${EXPOSED_PORT} -p ${EXPOSED_PORT}:${CONTAINER_PORT} \
    -v ${CODE_SOURCE_SRC}:${CODE_SOURCE_DST} \
    -v ${DEV_CONTENT_SRC}:${CONTENT_DST} \
    --name ${IMAGE_NAME} \
    ${USER_NAME}/${DEV_IMAGE_NAME}:${IMAGE_TYPE}
;;

# I run uat/prod with podman, not Docker. This will start the container remotely
# Note: commands run removely via ssh -t must be run one-at-a-time.
uat)
  echo "starts container in UAT..."
  ssh -t ${UAT_HOST} \
    sudo podman run --rm -d \
    --expose ${EXPOSED_PORT} -p ${EXPOSED_PORT}:${CONTAINER_PORT} \
    -e "CONTENT_PATH=${UAT_CONTENT_PATH}" \
    -v ${UAT_CONTENT_SRC}:${CONTENT_DST} \
    --name ${UAT_IMAGE_NAME} \
  ${USER_NAME}/${UAT_IMAGE_NAME}:${IMAGE_TYPE}
;;

# Podman containers can be run with systemd services for all the reliability
# of a production service without the Docker daemon.
prod)
  echo "enabling production service..."
  ssh -t ${PROD_HOST} sudo systemctl daemon-reload
  ssh -t ${PROD_HOST} sudo systemctl enable --now container-${IMAGE_NAME}.service
;;

*)
  echo "please provide one of the following as the first argument: dev, uat, prod."
  exit 1

esac
