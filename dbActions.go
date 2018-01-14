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
    collection = sess.DB("websitevisits").C("visits")
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
    // TODO add to and from 
    visits := []Visit{}
    err := collection.Find(nil).Sort("-visit_date").All(&visits)
    if (err != nil) {
        return nil, err
    }
    // marshal data and return
    return json.Marshal(visits)
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
    return json.Marshal(visits)
}