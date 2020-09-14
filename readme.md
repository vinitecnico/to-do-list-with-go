## Docker

Build image docker:

```
docker-compose up --build
```

## Go

```
http://localhost:8080/
```

## Postgre

```
http://localhost:16543/browser/
user: admin@admin.com
password: "admin"


CREATE DATABASE dev
CREATE TABLE products
(
    id  SERIAL,
    name character varying(50) COLLATE pg_catalog."default",
    description character varying(50) COLLATE pg_catalog."default",
    price numeric,
    quantity integer
)
Insert into products (name, description, price, quantity) values ('Samsung', '16gb i7', 2000, 2)
```