# Build on Alpine GoLang Docker Container
FROM golang:1.15-alpine as build
WORKDIR /home
ADD . /home/
RUN apk add --no-cache git
RUN go build -o application

# Copy Build Artifact to Clean Alpine Container
FROM alpine:3.12
COPY --from=build /home/application /app/bin/application
ENTRYPOINT [ "/app/bin/application", "-config=/app/conf/application.yaml" ]