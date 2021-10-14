FROM golang:1.17 AS build
WORKDIR /go/src
COPY webserver ./webserver
COPY database ./database
COPY main.go .
COPY go.mod .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /go/src/swagger ./
COPY pages ./pages
COPY static ./static

EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]
