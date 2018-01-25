# websiteAnalytics-backend

A RESTful Go backend to track website visits. Each website visit contains a) the ip address of the user b) geographical information about that ip address (i.e. lat-lon, city, zipcode) and c) the date the user visited the website. The stack consists of a Go app deployed through heroku, and a [mongo db](www.mongodb.com) database. Heroku deployment is simple and mongo makes requests fast and gets rid of the need for additional querying through sql. 

# Routes

### /visits
| Endpoint        | Method         | Description |
| :------------- | :-------------| :--------------------- |
| /visits        | POST         | Adds a new visit        |

The body of the request should have two parameters, `IpAddress` and `Location`, both as strings. There is no specific JSON structure assigned, which adds added flexibility for different use cases. For example, in one project I could write location as simply a city name i.e. 'Minneapolis' and in another I can strinify a bunch of JSON.

Example Request:
```sh
curl -H \
 "Content-Type: application/json" \
 -X POST -d '{"ipAddress": "123.456.789.1", "location": "TEST_LOCATION"}' \
 "http://localhost:5000/visits"
 # Response : {"ipAddress": "123.456.789.1", "location": "TEST_LOCATION"}
```

### /visits/ip/{ip} (UNDER CONSTRUCTION)

| Endpoint        | Method         | Description |
| :------------- | :-------------| :--------------------- |
| /visits/ip/:ip | GET          | Gets visit by ip address|

A quick way to get all the visits of by a specific IP address. Note that there is an added value Timestamp

Example request:
```sh
curl "http://localhost:5000/visits/ip/123.456.789.1"
# Response :
# [
#  {
#    "Data": "{\"ipAddress\":\"123.456.789.1\",\"location\":\"TEST_LOCATION}",
#    "Timestamp": "2017-12-27T14:08:28-06:00"
#  }
#]
```

### /visits (UNDER CONSTRUCTION)

| Endpoint        | Method         | Description           | Query strings|
| :------------- | :-------------| :---------------------  | :----- |
| /visits      | GET             | Lists all visits        | ip, to, from |

Gets all visits, filterable by ip address and dates as query strings. Note that `to` and `from` are not inputed as dates, but rather negative integers which translate to the number of 'days ago', i.e. `-7` means 'seven days ago'.

Example Request: (all visits from `123.456.789.1` within the last year and yesterday)
```sh
curl "http://localhost:5000/visits?ip=123.456.789.1&to=-1&from=-365"
# Response
# [
#  {
#    "Data": "{\"ipAddress\":\"123.456.789.1\",\"location\":\"TEST_LOCATION}",
#    "Timestamp": "2017-12-27T14:08:28-06:00"
#  },
#  {
#    "Data": "{\"ipAddress\":\"123.456.789.1\",\"location\":\"TEST_LOCATION}",
#    "Timestamp": "2017-10-27T14:08:28-06:00"
#  }
#]
```

# Development

### Setup

1. Download the project using go

```sh
# download src
cd $GOPATH/src
go get github.com/dgoldstein1/websiteAnalytics-backend
# cd into directory
cd $GOPATH/src/github.com/dgoldstein1/websiteAnalytics-backend
# install dependencies
go get
```

2. Launch Using Docker
 
```
docker-compose up -d
```

The app should now be running on http://localhost:5000. Running `curl http://localhost:5000/visits` should give you `[]` as there are no current visits in the mongo db.

### Testing

```
test/run_tests.sh 5000 
# 5000 is the server port, tells the tests where to make their requests
```

You should see the containers reload, and the result of the tests:

```
--- downing containers... --- 
--- done ---
--- loading test data... ---
--- done ---
--- starting dev server ---
Creating websiteanalyticsbackend_mongodb_1 ... done
Creating websiteanalyticsbackend_mongodb_1 ... 
Creating websiteanalyticsbackend_server_1  ... done
--- done ---
--- SUCCESS --- retrieve all visits
--- SUCCESS --- add a new visit
... other tests
--- resetting environment.. ---
Stopping websiteanalyticsbackend_server_1  ... done
Stopping websiteanalyticsbackend_mongodb_1 ... done
Removing websiteanalyticsbackend_server_1  ... done
Removing websiteanalyticsbackend_mongodb_1 ... done
Removing network websiteanalyticsbackend_default
--- done ---
```

### Deployment

This project is continuously deployed with every push or merge to `master`.

1. Mongo

To deploy a new instance, create a new mongo db on [https://mlab.com/](mongo lab) named 'websitevisits'. A free version is fine as this project does not consume a lot of memory. This will be the deployed app's datastore.

2.  Heroku

```sh
heroku login
# create new app on heroku
heroku create
Creating app... done, ⬢ thawing-inlet-61413
https://thawing-inlet-61413.herokuapp.com/ | https://git.heroku.com/thawing-inlet-61413.git
```

This creates a new heroku app. Before we can deploy code to it, we need to update the configuration settings.

3. Configure Environment Variables

Open up the app configuration settings on the  [heroku dashboard](https://dashboard.heroku.com/apps) and set the following config variables :

| Name        | Value         |
| :------------- | :-------------|
| DATABASE_URL      | mongodb://${dbuser}:${dbpassword}@ds255787.mlab.com:55787/${dbname}          |
| LOGGER      | true              |
| PORT      | 5000              |

For `mongodb://<dbuser>:<dbpassword>@ds255787.mlab.com:55787/websitevisits` git this from `mongo lab` and replace `<dbuser>`, `<dbpassword>`, and `<dbname>` with the user you wish to access the db fromm and the db name. *Note -- these are the credentials from the `add database user` button on mlab*

4. Deploy the app

```sh
# push the code to your remote
git push heroku master
Counting objects: 3, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 326 bytes | 0 bytes/s, done.
Total 3 (delta 2), reused 0 (delta 0)
remote: Compressing source files... done.
remote: Building source:
remote: 
remote: -----> Go app detected
remote: -----> Checking Godeps/Godeps.json file.
remote: -----> Using go1.9.2
remote:  !!    Installing package '.' (default)
remote:  !!    
remote: -----> Running: go install -v -tags heroku . 
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/gin-contrib/sse
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/gin-gonic/gin/json
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/golang/protobuf/proto
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/ugorji/go/codec
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/gopkg.in/go-playground/validator.v8
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/gopkg.in/yaml.v2
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/mattn/go-isatty
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/gopkg.in/mgo.v2/internal/json
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/gopkg.in/mgo.v2/bson
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/gopkg.in/mgo.v2/internal/scram
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/gopkg.in/mgo.v2
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/gin-gonic/gin/binding
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/gin-gonic/gin/render
remote: github.com/dgoldstein1/websiteAnalytics-backend/vendor/github.com/gin-gonic/gin
remote: github.com/dgoldstein1/websiteAnalytics-backend
remote: -----> Discovering process types
remote:        Procfile declares types -> web
remote: 
remote: -----> Compressing...
remote:        Done: 5.7M
remote: -----> Launching...
remote:        Released v9
remote:        https://quiet-brushlands-26130.herokuapp.com/ deployed to Heroku
remote: 
remote: Verifying deploy... done.
To https://git.heroku.com/{your app name}.git
   ad80626..2537239  master -> master

# follow deployment from logs
heroku logs
...
# build success
2018-01-14T14:59:53.000000+00:00 app[api]: Build started by user {your username}
2018-01-14T15:00:11.066497+00:00 app[api]: Deploy ad806265 by user {your username}
2018-01-14T15:00:11.066497+00:00 app[api]: Release v8 created by user {your username}
2018-01-14T14:59:53.000000+00:00 app[api]: Build succeeded
```

Then run `heroku open` to open up the page in your browser. You should see an empty bracket or this readme page.

## Authors

* **David Goldstein** - [DavidCharlesGoldstein.com](davidcharlesgoldstein.com) - [Decipher Technology Studios](deciphernow.com)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details