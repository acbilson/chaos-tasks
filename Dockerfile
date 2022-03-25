#############
# Base
#############

FROM golang:alpine3.14 as base
WORKDIR /mnt/src


#############
# Development
#############

FROM base as dev

# requires a volume mount
ENTRYPOINT ["go", "run", "."]


#############
# Build
#############

FROM base as build
COPY src .
RUN go version && go install
ENTRYPOINT ["tail", "-f", "/dev/null"]


#####
# UAT
#####

FROM build as uat
ENTRYPOINT ["/go/bin/server"]


############
# Production
############

FROM build as prod
ENTRYPOINT ["/go/bin/server"]
