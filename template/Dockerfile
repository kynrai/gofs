FROM golang:latest as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/server/main.go

FROM gcr.io/distroless/static-debian12:latest

COPY --from=build /go/bin/app /
EXPOSE 8080
CMD ["/app"]
