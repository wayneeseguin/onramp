FROM cosmtrek/air

ENV air_wd=/app/src

ENV PG_URI="postgres://postgres:onramp@postgres:5432/postgres?options=--search-path%3Donramp,public"
ENV PORT=9090

WORKDIR /app

COPY docker/dev/env /app/.env
COPY ./bin/cmd /bin/cmd

ENTRYPOINT ["/bin/cmd"]

