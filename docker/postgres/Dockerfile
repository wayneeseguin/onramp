FROM postgres:14

RUN apt-get update && apt-get -y install make vim python3 python3-pip curl \
  && pip3 install xlsx2csv pyexcel pyexcel-xlsx pandas openpyxl lxml xmltodict

ENV POSTGRES_PASSWORD="onramp"

COPY bin /app/bin
COPY sql /app/sql
COPY docker/postgres/env /app/.env
COPY docker/postgres/build /app/bin/build

WORKDIR /app

#RUN /app/bin/build

