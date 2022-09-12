# OnRamp


## Operations

# Configuration

Currently configuraton for the application is expected to come from env vars:

- `PG_URI` for the postgres database URI

# Onramp Image

The docker image `onramp:latest` is build by

```sh
make onramp
```

which will locally build an `onramp:latest` container image that can then 
be tagged and pushed to a registry.

# Dev

```sh
make build start
```
Can then follow the logs

```sh
make logs
```

## Development

### Dev Container Env Vars
```
GOLANG_VERSION => 1.19
GOPATH => /go
HOME => /root
HOSTNAME => dev
PATH => /go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
PG_URI => postgres://postgres:onramp@postgres:5432/postgres?options
PWD => /app/src
air_wd => /app/src
```

