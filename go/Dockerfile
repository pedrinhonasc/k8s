FROM golang:alpine as builder

WORKDIR /go/src

COPY ./main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/greeting main.go

FROM scratch

COPY --from=builder /go/bin/greeting /go/bin/greeting

ENTRYPOINT [ "/go/bin/greeting" ]

EXPOSE 8000