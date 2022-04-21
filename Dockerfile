#
# Build
#
FROM golang:alpine AS build-env
# Set proxy
ENV GOPROXY https://goproxy.cn,direct
COPY . /app
WORKDIR /app
RUN apk update && apk add git && go mod tidy && go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /blog-app-server

#
# Deploy
#
FROM gcr.io/distroless/static:latest
WORKDIR /app/config
COPY --from=build-env /blog-app-server /app
WORKDIR /app
EXPOSE 9090
ENTRYPOINT ["/app/blog-app-server"]
