FROM golang:1.22-alpine AS build
RUN apk --no-cache add gcc g++ make git sqlite-dev

WORKDIR /go/src/backend

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV ARCH=amd64
ENV CGO_CFLAGS="-I/usr/include"

COPY . .
RUN go mod download
RUN go build -o ./bin/backend ./cmd/app/main.go

FROM alpine AS runtime
WORKDIR /

COPY --from=build /go/src/backend/bin/backend /go/bin/

EXPOSE 8000
ENTRYPOINT /go/bin/backend