
# See Makefile
ARG IMAGE_GOLANG=golang:1.18

FROM $IMAGE_GOLANG AS builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

RUN go build -o app

FROM scratch

COPY --from=builder /app/app /usr/local/bin/app

ENTRYPOINT ["app"]
