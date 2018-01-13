// dbActions.go

package main

/**
 * Created by David Goldstein 12/2017
 * manages interactions with mongo db
 **/

import (
    "time"
    "encoding/json"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "fmt"
    "os"
)

var currId int
var collection *mgo.Collection
var testMode string

/**
 * initializes db
 * @param {string} mongo db name to connect to, i.e 'mmongodb://localhost/visits'
 * @return {bool} success
 **/
func connectToDb(uri string) bool {
    testMode = os.Getenv("TEST_MODE")
    sess, err := mgo.Dial(uri)
    if (err != nil) {
            fmt.Printf("Can't connect to mongo, go error %v\n", err)
            return false;
    }
    sess.SetSafe(&mgo.Safe{})
    // database = visits, collection = visits
    collection = sess.DB("visits").C("visits")
    return true;
}

/**
 * reads all rows
 * @param {string} ip filter
 * @param {int} to date (0 = now)
 * @param {int} from date (-7 = ) seven days ago
 * @return {[]byte} array of visits
 **/
func readAllRows(ip string, to int, from int) ([]byte, error) {
    visits := []Visit{}
    err := collection.Find(nil).Sort("-visit_date").All(&visits)
    if (err != nil) {
        return nil, err
    }
    // marshal data and return
    return json.Marshal(visits)
}

/**
 * compares two RFC3339 date strings
 * @param {string} date
 * @param {int} number of 'days ago' i.e (-7 = seven days ago)
 * @return {string} "before" "after" "equal" or "error"
 **/
func compareRFC3339(timestamp string, daysAgo int) string {
    // parse as RFC3339
    aConv, aErr := time.Parse(time.RFC3339, timestamp)
    if (aErr != nil) {
        return "error"
    }
    bConv, _ := time.Parse(time.RFC3339, time.Now().AddDate(0, 0, daysAgo).Format(time.RFC3339))

    // compare times
    if (aConv.Before(bConv)) {
        return "before"
    }
    if (aConv.After(bConv)) {
        return "after"
    }
    return "equal"
}

/**
 * adds an entry into the data
 * @param {json} visit to append to "visits" bucket
 * @return {json} visit, error
 **/
func insertRow(visit Visit) (Visit, error) {
    if (testMode != "true") { // do not add date for test mode in order to have static data
        t := time.Now()
        visit.Visit_Date = t
    }
    err := collection.Insert(visit)
    return visit, err
}

/**
 * retrieves all visits from a specific ip address
 * @param {string} ip address
 * @return {[]byte} array of visits
 **/
func readByIp(ip string) ([]byte, error) {
    // TODO
    visits := []Visit{}
    err := collection.Find(bson.M{ "ip" : ip }).Sort("-visit_date").All(&visits)
    if (err != nil) {
        return nil, err
    }
    return json.Marshal(visits)
}