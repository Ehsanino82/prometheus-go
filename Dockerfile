FROM golang:alpine AS build

RUN apk --update add ca-certificates git

WORKDIR /srv/build

COPY go.mod .
COPY go.sum .
COPY cmd/main.go .
COPY metrics ./metrics
COPY pkg ./pkg

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

EXPOSE 9000
EXPOSE 8080

FROM golang:alpine AS runtime

WORKDIR /srv/app

COPY --from=build /srv/build/main /bin/

CMD ["/bin/main"]
