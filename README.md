# OnRamp


## Operations

# Configuration

Currently configuraton for the application is expected to come from env vars:

- `PG_URI` for the postgres database URI

# Dev

```sh
make build start
```
Can then follow the logs

```sh
make logs
```

## Development

### Env Vars
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

