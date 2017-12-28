# DavidWebsiteBackend

A RESTful Go backend to track website visits. Each website visit contains a) the ip address of the user b) geographical information about that ip address (i.e. lat-lon, city, zipcode) and c) the date the user visited the website. The stack consists of a Go app deployed through heroku, and a [bolt.db](https://github.com/boltdb/bolt) database deployed on the server. Heroku deployment is simple and bolt db makes requests fast and gets rid of the need for additional querying through sql. 

# Routes

### /visits
| Endpoint        | Method         | Description |
| :------------- | :-------------| :--------------------- |
| /visits        | POST         | Adds a new visit        |

The body of the request should have two parameters, `IpAddress` and `Location`, both as strings. There is no specific JSON structure assigned, which adds added flexability for different use cases. For example, in one project I could write location as simply a city name i.e. 'Minneapolis' and in another I can strinify a bunch of JSON.

Example Request:
```sh
curl -H \
 "Content-Type: application/json" \
 -X POST -d '{"ipAddress": "123.456.789.1", "location": "TEST_LOCATION"}' \
 "http://localhost:5000/visits"
 # Response : {"ipAddress": "123.456.789.1", "location": "TEST_LOCATION"}
```

### /visits/ip/{ip}

| Endpoint        | Method         | Description |
| :------------- | :-------------| :--------------------- |
| /visits/ip/:ip | GET          | Gets visit by ip address|

A quick way to get all the visits of by a speicifc IP address. Note that there is an added value Timestamp

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

### /visits

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

Serve locally

```sh
go install
heroku login
heroku local
```

The app should run on http://localhost:5000

# Testing

Each endpoint is tested using a suite of integration tests. They can be executed by running the following commands from the root directory of the project:

```sh
test/launchTestServer.sh {port number}
# from a new window
test/run_tests.sh {port number}
```

You should see several `--- SUCCESS ---` messages outputted to your terminal.

*Note - to rerun tests, you must restart the test server i.e. `test/launchTestServer.sh {port number}`*
