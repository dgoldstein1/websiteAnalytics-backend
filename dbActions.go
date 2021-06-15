// dbActions.go

package main

/**
 * Created by David Goldstein 12/2017
 * manages interactions with mongo db
 **/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"context"
	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Cannot connect to mongo: %v\n", err)
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("Ping: %v\n", err)
		panic(err)
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
	query["$"+query_type] = []bson.M{}
	// we need to keep track of if a value was added, if not, then we should return nil for the query
	valueAdded := false

	// go through each param in visit filters and add to query if not empty
	if visitFilters.Href != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"href": visitFilters.Href})
		valueAdded = true
	}
	if visitFilters.Ip != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"ip": visitFilters.Ip})
		valueAdded = true
	}
	if visitFilters.City != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"city": visitFilters.City})
		valueAdded = true
	}
	if visitFilters.Country_Code != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"country_code": visitFilters.Country_Code})
		valueAdded = true
	}
	if visitFilters.Country_Name != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"country_name": visitFilters.Country_Name})
		valueAdded = true
	}
	if visitFilters.Latitude != NO_INPUT_FLOAT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"latitude": visitFilters.Latitude})
		valueAdded = true
	}
	if visitFilters.Longitude != NO_INPUT_FLOAT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"longitude": visitFilters.Longitude})
		valueAdded = true
	}
	if visitFilters.Metro_Code != NO_INPUT_INT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"metro_code": visitFilters.Metro_Code})
		valueAdded = true
	}
	if visitFilters.Region_Code != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"region_code": visitFilters.Region_Code})
		valueAdded = true
	}
	if visitFilters.Time_Zone != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"time_zone": visitFilters.Time_Zone})
		valueAdded = true
	}
	if visitFilters.Zip_Code != NO_INPUT {
		query["$"+query_type] = append(query["$"+query_type].([]bson.M), bson.M{"zip_code": visitFilters.Zip_Code})
		valueAdded = true
	}

	// only return query if we've added a value
	if valueAdded == true {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Find().SetSort(bson.D{{"visit_date", -1}})
	cur, err := collection.Find(ctx, query, opts)
	if err != nil {
		fmt.Printf("Collection.Find(): %v\n", err)
		return visits, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		v := Visit{}
		err := cur.Decode(&v)
		if err != nil {
			fmt.Printf("cur.Decode(): %v\n", err)
			return visits, err
		}
		// append to visits slice
		visits = append(visits, v)
	}
	if err := cur.Err(); err != nil {
		fmt.Printf("cur.Err() %v\n", err)
		return visits, err
	}

	return visits, nil
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, visit)
	return visit, err
}

/**
 * gets total number of documents in db
 **/
func docCount() int64 {
	opts := options.EstimatedDocumentCount().SetMaxTime(2 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	count, err := collection.EstimatedDocumentCount(ctx, opts)
	if err != nil {
		fmt.Printf("Could not get estimated number of documents: %v", err)
		count = -1
	}
	return count
}

// goes through all entries filling in where latitude and longitude is 0
// if cannot update, latitude and longitude is set to -1
func updateAllEmptyEntries() error {
	// get all visits @0,0
	visitFilters := Visit{
		Href:         NO_INPUT,
		Ip:           NO_INPUT,
		City:         NO_INPUT,
		Country_Code: NO_INPUT,
		Country_Name: NO_INPUT,
		Latitude:     0,
		Longitude:    0,
		Metro_Code:   NO_INPUT_INT,
		Region_Code:  NO_INPUT,
		Time_Zone:    NO_INPUT,
		Zip_Code:     NO_INPUT,
	}
	query, err := createQueryFromFilters(visitFilters, "and")
	if err != nil {
		return fmt.Errorf("Could not createQueryFromFilters: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, query, nil)
	if err != nil {
		return fmt.Errorf("Collection.Find(): %v\n", err)
	}
	defer cur.Close(context.Background())
	// go through each record and update
	for cur.Next(context.Background()) {
		v := Visit{}
		err := cur.Decode(&v)
		if err != nil {
			return fmt.Errorf("could not decode visit: %v, %v", v, err)
		}
		fmt.Printf("found visit to update: %v\n", spew.Sdump(v))
		newVisit, err := fetchGeoIP(v)
		if err != nil {
			fmt.Printf("could not get info for new visit %v", err)
			continue
		}
		fmt.Printf("found geoIP info for visit: %v", newVisit)
		// success, merge and update in database
		if err = updateVisit(v.Ip, newVisit); err != nil {
			fmt.Printf("could not update visit: %v", err)
		}
	}
	return nil
}

// updates visit in store based on IP
func updateVisit(ip string, newVisit Visit) error {
	opts := options.Update().SetUpsert(true)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateMany(
		ctx,
		bson.M{"ip": ip},
		bson.D{
			{"$set", bson.D{{"latitude", newVisit.Latitude}}},
			{"$set", bson.D{{"longitude", newVisit.Longitude}}},
			{"$set", bson.D{{"country_code", newVisit.Country_Code}}},
			{"$set", bson.D{{"country_name", newVisit.Country_Name}}},
			{"$set", bson.D{{"city", newVisit.City}}},
			{"$set", bson.D{{"metro_code", newVisit.Metro_Code}}},
			{"$set", bson.D{{"region_code", newVisit.Region_Code}}},
			{"$set", bson.D{{"time_zone", newVisit.Time_Zone}}},
			{"$set", bson.D{{"zip_code", newVisit.Zip_Code}}},
		},
		opts,
	)
	fmt.Printf("updated docs: %v", spew.Sdump(result))
	return err
}

// fetchInfoFromIP fetches geoIP and merges into object
func fetchGeoIP(v Visit) (Visit, error) {
	// http://api.ipstack.com/check\?access_key\=7eca814a6de384aab338e110c57fef37
	params := url.Values{}
	params.Add("access_key", os.Getenv("IP_STACK_ACCESS_KEY"))
	url := "http://api.ipstack.com/" + url.QueryEscape(v.Ip) + "?" + params.Encode()
	fmt.Println("fetching IP: %s", url)
	r, err := http.Get(url)
	if err != nil {
		return Visit{}, fmt.Errorf("could not fetch visit from %s: %v", url, v)
	}
	if r.StatusCode != http.StatusOK {
		return Visit{}, fmt.Errorf("bad response code: %d\n", r.StatusCode)
	}
	newVisit := Visit{}
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &newVisit)
	if err != nil {
		return Visit{}, fmt.Errorf("could not decode json to visit: %v", err)
	}
	if newVisit.Ip == "" {
		return Visit{}, fmt.Errorf("could not convert bytes to Visit: %s\n", string(bodyBytes))
	}
	if newVisit.Latitude == 0 {
		return Visit{}, fmt.Errorf("could not fetch new lat/lon: %s\n", string(bodyBytes))
	}
	// copy over non-ipstack properties
	newVisit.Href = v.Href
	newVisit.Visit_Date = v.Visit_Date
	return newVisit, nil
}
