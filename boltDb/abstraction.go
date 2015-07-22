package boltDb

import (
    "github.com/boltdb/bolt"
    "os/user"
    "fmt"
    "log"
)


var Name string

func Setup(){
    usr, _ := user.Current()
    Name = string(usr.HomeDir)+"/bolt.db"
}

func Update(key []byte, value []byte, bucket []byte){

    db, err := bolt.Open(Name, 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(bucket)
        if err != nil {
            return err
        }

        err = bucket.Put(key, value)
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}

func Find(key []byte, bucketString []byte) []byte {

    db, err := bolt.Open(Name, 0644, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists(bucketString)
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }

    var d []byte

    err = db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket(bucketString)
        if bucket == nil {
            return fmt.Errorf("Bucket %q not found!", bucket)
        }

        val := bucket.Get(key)
        d = val
        //result = string(val)
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }
    return d
}
