// dbActions.go

package main

/**
 * Created by David Goldstein 12/2017
 * manages interactions with bolt db
 **/

import (
    "github.com/boltdb/bolt"
    "time"
    "encoding/json"
)

// database to use
type server struct {
    db *bolt.DB
}

var s server
var currId int

/**
 * initializes db
 * @param {string} name of local db file in root directory
 * of this project
 * @return {bool} success
 **/
func newRepo(dbfile string) bool{
    var err error
    // s = &server{}
    s.db, err = bolt.Open(dbfile, 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        return false
    }
    err = s.db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("visits"))
        if err != nil {
            return err
        }
        return nil
    })
    return true
}

/**
 * reads all rows
 * @param {string} ip filter
 * @param {int} to date (0 = now)
 * @param {int} from date (-7 = ) seven days ago
 * @return {[]byte} array of visits
 **/
func readAllRows(ip string, to int, from int) ([]byte, error) {
    visitEntries := []VisitEntry{}
    // view transaction
    s.db.View(func(tx *bolt.Tx) error {
        // get bucket with for visits
        b := tx.Bucket([]byte("visits"))
        // loop through table to create array
        b.ForEach(func(timestamp, data []byte) error {
            // conv []byte => json for visit
            v := Visit{}
            json.Unmarshal(data, &v)
            ve := VisitEntry{string(data), string(timestamp)}
            // check equal to passed ip
            if (ip != NO_INPUT && v.IpAddress != ip) {
                return nil
            }
            // if timestamp is before the 'from' date, return
            if (from != NO_INPUT_INT && compareRFC3339(string(timestamp[:]), from) == "before") {
                return nil
            }
            // if timestamp is after the 'to'date, return
            if (to != NO_INPUT_INT && compareRFC3339(string(timestamp[:]), to) == "after") {
                return nil
            }
            // if pases all conditions, add to valid entries
            visitEntries = append(visitEntries, ve)
            // marshal data into 
            return nil
        })
        return nil
    })
    // cast to json
    return json.Marshal(visitEntries)
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
    // start an update transaction
    err := s.db.Update(func(tx *bolt.Tx) error {
        // retrieve the visits bucket
        b := tx.Bucket([]byte("visits"))
        // generate id
        buf, err := json.Marshal(visit)
        if (err != nil) {
            return err
        }
        // Persist bytes to bucket
        return b.Put([]byte(time.Now().Format(time.RFC3339)), buf)
    })
    return visit, err
}

/**
 * retrieves all visits from a specific ip address
 * @param {string} ip address
 * @return {[]byte} array of visits
 **/
func readByIp(ip string) ([]byte, error) {
    visitEntries := []VisitEntry{}
    s.db.View(func(tx *bolt.Tx) error {
        // get bucket with for visits
        b := tx.Bucket([]byte("visits"))
        // loop through table to create array
        b.ForEach(func(timestamp, data []byte) error {
            // conv []byte => json for visit
            v := Visit{}
            json.Unmarshal(data, &v)
            // if IPs match, add to response
            if (v.IpAddress == ip) {
                // read in byte stream to visits object
                ve := VisitEntry{string(data), string(timestamp)}
                // add it to the slice of foods
                visitEntries = append(visitEntries, ve)
            }
            return nil
        })
        return nil
    })
    return json.Marshal(visitEntries)
}