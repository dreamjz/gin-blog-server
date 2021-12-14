FROM gcr.io/distroless/base
WORKDIR /app
COPY ./bin/blog-app-server /app
EXPOSE 9090
ENTRYPOINT ["/app/blog-app-server"]