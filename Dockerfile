##
## Build
##
#FROM golang:alpine AS build-env
## set golang proxy
#ENV GOPROXY https://goproxy.cn,direct
#WORKDIR /app
#COPY . /app
#RUN go mod download \
#    && CGO_ENABLED=0 GOOS=linux go build -o /blog-app-server
##
## Deploy
##
#FROM gcr.io/distroless/static
#COPY --from=build-env /blog-app-server /app
#ENV APP_NEW=docker
#EXPOSE 9090
#ENTRYPOINT ["/app/blog-app-server"]

# Local compile
FROM scratch
WORKDIR /app/config
COPY ./bin/blog-app-server /app
WORKDIR /app
ENV APP_ENV=docker
EXPOSE 9090
ENTRYPOINT ["/app/blog-app-server"]
