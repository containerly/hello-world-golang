FROM golang:1.19 AS build
WORKDIR /go/src
COPY . .

ENV CGO_ENABLED=0

RUN go build .

FROM scratch AS runtime
COPY --from=build /go/src/golang-service .
EXPOSE 8080/tcp
ENTRYPOINT ["./golang-service"]
