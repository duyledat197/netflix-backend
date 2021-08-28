FROM golang:1.16-alpine as build-stage

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 && go build -o app-exe ./cmd/srv/...

# production: 
FROM golang:1.16-alpine as production-stage

COPY --from=build-stage /app/app-exe /bin
COPY --from=build-stage /app/google_cloud_key.json /bin

RUN chmod +x /bin/app-exe


EXPOSE 10010

ENTRYPOINT [ "/bin/app-exe", "start" ]
