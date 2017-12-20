package main

import (
    "encoding/json"
    "fmt"   
    "github.com/boltdb/bolt"
    "time"
    "strconv"
)

type server struct {
    db *bolt.DB
}

var s server
var currId int

func newRepo(dbfile string) bool{
    var err error
    // s = &server{}
    s.db, err = bolt.Open(dbfile, 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        return false
    }
    s.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("visits"))
        if b == nil {
            currId = 0
        } else {
            currId = b.Stats().KeyN - 1
        }
        return nil
    })
    return true
}

func addValue(key string) error{
    var id = []byte(strconv.Itoa(currId))
    currId++
    
    return s.db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte("visits"))
        if err != nil {
            return err
        }
        return b.Put(id, []byte(key));
    })

}
func getById(key string) (ct string, data []byte, err error) {
    s.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("visits"))
        if (b==nil){
            return nil
        }
        r := b.Get([]byte(key))
        if r != nil {
            data = make([]byte, len(r))
            copy(data, r)
        }

        r = b.Get([]byte(fmt.Sprintf("%s-ContentType", key)))
        ct = string(r)
        return nil
    })
    return
}

/**
 * gets all current values
 * @return {[]byte} json encoded byte array, list of objects
 **/
func getValues() []byte {

    // create slice of visits
    visits := []Visit{}
    s.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("visits"))

        // nothing posted yet
        if b==nil {
            return nil
        }

        // loop through table to create array of all datapoints
        b.ForEach(func(ip, location []byte) error {

            // read in byte stream to visits object
            v := Visit{string(location), string(ip)}

            // add it to the slice of foods
            visits = append(visits, v)

      return nil
    })
        return nil
    })

    temp , _ := json.Marshal(visits)
    return temp
}

