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


#####
# UAT
#####

FROM alpine:3.15 as uat
COPY --from=build /mnt/src/templates /etc/chaos-tasks/templates
COPY --from=build /go/bin/server /usr/local/bin/server
ENTRYPOINT ["/usr/local/bin/server"]


############
# Production
############

FROM alpine:3.15 as prod
COPY --from=build /mnt/src/templates /etc/chaos-tasks/templates
COPY --from=build /go/bin/server /usr/local/bin/server
ENTRYPOINT ["/usr/local/bin/server"]
