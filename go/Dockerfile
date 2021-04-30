FROM golang:alpine as builder

WORKDIR /go/src

COPY ./main.go .

RUN GOOS=linux go build -o /go/bin/greeting main.go

FROM scratch

COPY --from=builder /go/bin/greeting /go/bin/greeting

EXPOSE 8001

ENTRYPOINT [ "/go/bin/greeting" ]