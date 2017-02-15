# Go API Blueprint

This project demonstrates how to structure and build an API using the Go language without a framework.

## Prerequisites

- Go >= v1.5.x (**BUT Go v.1.7.x more recommended**)
- MySQL
- MongoDB

If you are on Go 1.5, you need to set GOVENDOREXPERIMENT to 1.
If you are on Go 1.4 or earlier, the code will not work because it uses the vendor folder.

## Structure

The majority of the code is in the **vendor/app** folder. There were a lot of users trying to use the code on their own, but had to change all the imports path for it to work properly.
The only downside is godoc does not work with the vendor folder method. Luckily, all the code can be moved out of the vendor folder and then a quick find and replace will get it working again if you want.

By dafault we are using MongoDB with authentication, you need to update DSN in **vendor/app/commons/mongodb/mongodb.go** if your MongoDB server does not have authentication.

## Rules for Consistency

Rules for mapping HTTP methods to CRUD:

```
POST   - Create (add record into database)
GET    - Read (get record from the database)
PUT    - Update (edit record in the database)
DELETE - Delete (remove record from the database)
```

Rules for status codes:

```
* Create something - 201 (Created)
* Read something - 200 (OK)
* Update something - 200 (OK)
* Delete something - 200 (OK)
* Create but missing info - 400 (Bad Request)
* Any other error - 500 (Internal Server Error)
```

Rules for messages:

```
* 201 - item created
* 200 - item found; no items to find; items deleted; no items to delete; etc
* 400 - [field] is missing; [field] needs to be type: [type]
* 500 - an error occurred, please try again later (should also log error because it's a programming or server issue)
```

## Goals for this project

Integrate security similar to Parse: http://blog.parse.com/learn/secure-your-app-one-class-at-a-time/

Code generation for the following:
* Controllers with routes
* Models
* Endpoint tests
* Swagger spec

## Thanks for the awesome project

- [josephspurrier/gowebapi](https://github.com/josephspurrier/gowebapi)
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [gorilla/context](https://github.com/gorilla/context)
- [spf13/viper](https://github.com/spf13/viper)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [jmoiron/sqlx](https://github.com/jmoiron/sqlx)
- [justinas/alice](https://github.com/justinas/alice)

