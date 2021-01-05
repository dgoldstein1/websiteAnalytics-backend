// dbActions.go

package main

/**
 * Created by David Goldstein 12/2017
 * manages interactions with mongo db
 **/

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
)

var currId int
var collection *mongo.Collection
var testMode string

/**
 * initializes db
 * @param {string} mongo db name to connect to, i.e 'mmongodb://localhost/visits'
 * @return {bool} success
 **/
func connectToDb(uri string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Cannot connect to mongo: %v\n", err)
		return false
	}
	fmt.Println("CONNECTED")

	testMode = os.Getenv("TEST_MODE")
	// database = visits, collection = visits
	collection = client.Database(os.Getenv("DB_NAME")).Collection("visits")
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
	if (visitFilters.Href != NO_INPUT) {
		query["$" + query_type] = append(query["$" + query_type].([]bson.M), bson.M{"href": visitFilters.Href})
		valueAdded = true
	}
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
func readAllRows(visitFilters Visit, to int, from int, query_type string) ([]Visit, error) {
	query, _ := createQueryFromFilters(visitFilters, query_type)
	visits := []Visit{}

	// find all documents in which the "name" field is "Bob"
	// specify the Sort option to sort the returned documents by age in ascending order
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{{"visit_date", 1}})
	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
	    return visits, err
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
	    fmt.Printf("Could not read all rows, %v \n", err)
	    return visits, err
	}
	for _, result := range results {
		fmt.Println(result)
	    // visits = append(result, visits)
	}
	// marshal data and return
	return visits, nil
}

/**
 * adds an entry into the data
 * @param {json} visit to append to "visits" bucket
 * @return {json} visit, {error} error
 **/
func insertRow(visit Visit) (Visit, error) {
	return Visit{}, nil
	// if testMode != "true" { // do not add date for test mode in order to have static data
	// 	t := time.Now()
	// 	visit.Visit_Date = t
	// }
	// err := collection.Insert(visit)
	// return visit, err
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
