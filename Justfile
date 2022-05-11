# builds a production-ready podman image
build:
	podman build --target=prod -t acbilson/tasks:alpine .

# restarts the systemd service
restart:
	systemctl --user restart container-tasks.service

# builds a production-ready podman image
redeploy: build restart
