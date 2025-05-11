# BookShelf

## Needed `env` variables (in file `./config/.env` for docker)

```dotenv
DB_PATH="/path/to/db/sqlite3.db"
MEDIA_PATH="/path/to/media/for/excel"
```

## Deployment

### _1. Go to deployment directory_

```shell
cd ./deployment
```

### _2. Build and run docker compose (in background)_

```shell
docker compose up --build -d
```

### _3. Up all DB migrations_

```shell
docker exec -it book_shelf_app /bin/sh -c "/app/migrator up"
```

## TODO

- [x] Create massive interface for export to Excel (just prototype yet)
- [x] Set up docker and docker compose
