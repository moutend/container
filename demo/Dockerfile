FROM golang:1.13-alpine AS build
LABEL stage=temporary
WORKDIR /go/src/
COPY main.go /go/src/
RUN CGO_ENABLED=0 go build -o /go/bin/demo

FROM scratch
COPY --from=build /go/bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
