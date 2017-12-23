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
 * @return {[]byte} array of visits
 **/
func readAllRows() []byte {
    visitEntries := []VisitEntry{}
    // view transaction
    s.db.View(func(tx *bolt.Tx) error {
        // get bucket with for visits
        b := tx.Bucket([]byte("visits"))
        // loop through table to create array
        b.ForEach(func(timestamp, data []byte) error {
            // read in byte stream to visits object
            ve := VisitEntry{string(data), string(timestamp)}
            // add it to the slice of foods
            visitEntries = append(visitEntries, ve)
            return nil
        })
        return nil
    })

    // cast to json
    temp , _ := json.Marshal(visitEntries)
    return temp
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