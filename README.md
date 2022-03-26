# Golang Web Service

This project documents patterns and modules for building web services with go.

## API

REST: [Gin](https://github.com/gin-gonic/gin)

## Data persistence

SQL: [PostgreSQL](https://www.postgresql.org/) through [GORM](https://gorm.io/)

## Other

This project mostly follows the [Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Dependency injection is currently achieved without extra modules,
but [Wire](https://github.com/google/wire) might be used in the future.

Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
