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
 * @param {Visit} the parameters of Visit to filter against, without 'time'
 * @param {string} the type of query i.e. 'and'
 *
 * @return {bson.M{} bytes} {error}
 **/
func createQueryFromFilters(visitFilters Visit, query_type string) (bson.M, error) {
	// initialize bson object
	query := bson.M{}
	query["$" + query_type] = []bson.M{}
	// we need to keep track of if a value was added, if not, then we should return nil for the query
	valueAdded := false

	// go through each param in visit filters and add to query if not empty
	if (visitFilters.Ip != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"ip": visitFilters.Ip})
		valueAdded = true
	}
	if (visitFilters.City != NO_INPUT) { 
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"city": visitFilters.City})
		valueAdded = true
	}
	if (visitFilters.Country_Code != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"country_code": visitFilters.Country_Code})
		valueAdded = true
	}
	if (visitFilters.Country_Name != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"country_name": visitFilters.Country_Name})
		valueAdded = true
	}
	if (visitFilters.Latitude != NO_INPUT_FLOAT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"latitude": visitFilters.Latitude})
		valueAdded = true
	}
	if (visitFilters.Longitude != NO_INPUT_FLOAT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"longitude": visitFilters.Longitude})
		valueAdded = true
	}
	if (visitFilters.Metro_Code != NO_INPUT_INT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"metro_code": visitFilters.Metro_Code})
		valueAdded = true
	}
	if (visitFilters.Region_Code != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"region_code": visitFilters.Region_Code})
		valueAdded = true
	}
	if (visitFilters.Time_Zone != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"time_zone": visitFilters.Time_Zone})
		valueAdded = true
	}
	if (visitFilters.Zip_Code != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"zip_code": visitFilters.Zip_Code})
		valueAdded = true
	}

	// only return query if we've added a value
	if (valueAdded == true) {
		return query, nil
	}
	return nil, nil
}

/**
 * reads all rows with filter
 * @param {bson.M} query
 * @param {int} to date (0 = now)
 * @param {int} from date (-7 = ) seven days ago
 * @param {string} the type of query i.e. 'and' or 'or'
 * @return {[]byte} array of visits
 **/
func readAllRows(visitFilters Visit, to int, from int, query_type string) ([]byte, error) {
	query, _ := createQueryFromFilters(visitFilters, query_type)
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
