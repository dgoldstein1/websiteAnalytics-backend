# DavidWebsiteBackend

A REST backend to track website visits

| Route        | Method           | Description  |
| ------------- |:-------------:| -----:|
| / OR /visits      | GET | lists all website visits |
| /visits/${ip}    | GET      |  lists all visits of a specific ip address |
| /visits | POST     |  adds a visit to the database |


# Dev

Get dependencies and build

```sh
go install
```

Serve locally

```sh
heroku login
heroku local
```




