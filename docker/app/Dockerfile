FROM golang:alpine3.16 AS builder
COPY ./onramp /app
ENV GO_PATH=/app
WORKDIR /app
RUN go mod download
RUN go build -o /bin/onramp .

FROM golang:alpine3.16
RUN apk add --no-cache ca-certificates bash
RUN mkdir /app
WORKDIR /app
COPY docker/app/env /app/.env
COPY ./docs /app/docs
COPY ./bin/cmd /bin/cmd
COPY ./bin/shell /bin/shell
COPY --from=builder /bin/onramp /app/bin/onramp

EXPOSE 12345

ENTRYPOINT ["/bin/cmd"]

CMD [ "/bin/cmd" ]

