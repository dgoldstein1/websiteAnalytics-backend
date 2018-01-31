// dbActions.go

package main

/**
 * Created by David Goldstein 12/2017
 * manages interactions with mongo db
 **/

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
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
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		return false
	}
	sess.SetSafe(&mgo.Safe{})
	// database = visits, collection = visits
	collection = sess.DB("websitevisits").C("visits")
	return true
}

/**
 * creates mongodb query from filters
 * @param {string} ip string filter
 * @param {string} city filter
 * @param {string} country code string filter
 * @param {string} country name string filter
 * @param {float64} latitude filter
 * @param {float64} longitude filter
 * @param {int} metro code filter
 * @param {string} region code filter
 * @param {string} time zone (i.e 'America/New_York')
 * @param {string} zipcode filter
 *
 * @return {bson.M{} bytes} {error}
 **/
func createQueryFromFilters(ip string, city string, country_code string, country_name string, latitudeFloat float64, longitudeFloat float64, metroCodeInt int, region_code string, time_zone string, zip_code string) (bson.M, error) {
	query := bson.M{}
	query["$and"] = []bson.M{}
	if ip != NO_INPUT {
		query["$and"] = append(query["$and"].([]bson.M), bson.M{"ip": ip})
	}
	return query, nil
}

/**
 * reads all rows with filter
 * @param {bson.M} query
 * @param {int} to date (0 = now)
 * @param {int} from date (-7 = ) seven days ago
 * @return {[]byte} array of visits
 **/
func readAllRows(query bson.M, to int, from int) ([]byte, error) {
	visits := []Visit{}
	err := collection.Find(query).Sort("-visit_date").All(&visits)
	if err != nil {
		return nil, err
	}
	// marshal data and return
	return json.Marshal(visits)
}

/**
 * adds an entry into the data
 * @param {json} visit to append to "visits" bucket
 * @return {json} visit, {error} error
 **/
func insertRow(visit Visit) (Visit, error) {
	if testMode != "true" { // do not add date for test mode in order to have static data
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
