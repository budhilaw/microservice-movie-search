FROM golang:alpine

RUN apk update
WORKDIR /go/src/app
COPY ./service-1 .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o service-1 ./app
RUN mv service-1 /usr/bin/service-1
RUN chmod +x /usr/bin/service-1

CMD ["service-1"]

EXPOSE 5555
