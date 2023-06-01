FROM golang:1.20

WORKDIR /mongo-go
COPY . .
RUN git config --global --add safe.directory /mongo-go
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /fly/bin/start ./cmd/start
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /fly/bin/start_admin_server ./cmd/admin_server

FROM mongo:6.0

COPY --from=0 /fly/bin/* /usr/local/bin
COPY scripts/* /usr/local/bin
COPY fly-start.sh /usr/local/bin/
COPY mongod.conf /etc/mongod.conf

ENTRYPOINT [ "fly-start.sh" ]