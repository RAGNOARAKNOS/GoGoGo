# App Notes

## Dependencies

```shell
docker pull postgres
docker pull dpage/pgadmin4
```

```shell
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=devuser -e POSTGRES_PASSWORD=password -e POSTGRES_DB=devdb -d postgres
docker run --name pgadmin -p 82:80 -e 'PGADMIN_DEFAULT_EMAIL=test@dev.com' -e 'PGADMIN_DEFAULT_PASSWORD=pass123' -d dpage/pgadmin4
```

REMEMBER: host.docker.internal is the DNS name of the host from within containers (172.17.0.1)

REMEMBER: To access PGADMIN4 use link [PGAdmin dashboard](http://localhost:82/browser/)

## Config Values

A .ENV file located in the root is read to configure the below values, EnvVars can ALSO be used to override

| Variable | Description | Default |
| --- | --- | --- |
| GAMELIB_DB_HOST | Hostname of Postgres instance | 127.0.0.1 |
| GAMELIB_DB_USER | Postgres username | devuser |
| GAMELIB_DB_PASSWORD | Postgres password | password|
| GAMELIB_DB_NAME | Postgres database instance name | devdb|
| GAMELIB_DB_PORT | Postgres port | 5432 |
| GAMELIB_REST_PORT | Application REST endpoint port | 9999 |
| GAMELIB_RUNTIME | sets deployment location, used for debug flags | dev |
