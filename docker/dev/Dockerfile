FROM cosmtrek/air

#COPY ./onramp /app
ENV GO_PATH=/app
ENV air_wd=/app
env PG_USERNAME="postgres"
env PG_PASSWORD="onramp"
env PG_HOSTNAME="postgres"
env PG_DATABASE="postgres"
env PG_PORT="5432"
env PG_SCHEMA="onramp"
env PG_URI="postgres://postgres:onramp@postgres:5432/postgres?options=--search-path%3Donramp,public"


WORKDIR /app
#RUN go mod download
#RUN go build -o /bin/onramp .
#RUN apk add --no-cache ca-certificates bash
#COPY docker/dev/env /app/.env
#COPY ./docs /app/docs
#COPY ./bin/cmd /bin/cmd
#COPY ./bin/shell /bin/shell

#ENTRYPOINT ["/bin/cmd"]
#CMD [ "/bin/cmd" ]

