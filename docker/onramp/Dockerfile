FROM golang:1.19 as builder

RUN mkdir -p /app/src /app/bin

COPY ./app/src /app/src

WORKDIR /app/src

RUN go mod download
RUN go build -o /bin/onramp . # ./cmd/onramp

FROM golang:1.19
COPY --from=builder /bin/onramp /bin/onramp

ENV PORT=9090

WORKDIR /app

RUN echo '#!/usr/bin/env bash' > /app/.env

COPY ./bin/cmd /app/bin/cmd

ENTRYPOINT ["/app/bin/cmd"]

