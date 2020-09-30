# Build on Alpine GoLang Docker Container
FROM golang:1.15-alpine as build
WORKDIR /home
ADD . /home/
RUN apk add --no-cache git && go build

# Copy Build Artifact to Clean Alpine Container
FROM alpine:3.12
COPY --from=build /home/hello-go /app/bin/app
ENTRYPOINT [ "/app/bin/app" ]