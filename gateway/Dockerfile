FROM golang:alpine

RUN apk update
WORKDIR /go/src/app
COPY ./gateway .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o gateway-service ./app
RUN mv gateway-service /usr/bin/gateway-service
RUN chmod +x /usr/bin/gateway-service

CMD ["gateway-service"]

EXPOSE 5555
