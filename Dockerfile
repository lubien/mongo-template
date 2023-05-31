FROM golang:1.20 AS gostage

WORKDIR /go/src/github.com/fly-apps/fly-postgres
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /fly/bin/start ./cmd/start
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /fly/bin/start_admin_server ./cmd/admin_server

FROM mongo:6.0

COPY --from=gostage /fly/bin/* /usr/local/bin
COPY fly-start.sh /usr/local/bin/
COPY mongod.conf /etc/mongod.conf
COPY checks.py /checks.py
COPY checks.service /etc/systemd/system/checks.service

ENTRYPOINT [ "fly-start.sh" ]