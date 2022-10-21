package user

import (
    "errors"

    "github.com/asdine/storm/v3"
    "gopkg.in/mgo.v2/bson"
)

// Holds data for a single user
type User struct {
    ID bson.ObjectId `json:"id"`
    Name string `json:"name"`
    Role string `json:"role"`
}

const (
    dbPath = "user.db"
)

var (
    ErrRecordInvalid = errors.New("record is invalid")
)

// Returns all user records from the database
func All() ([]User, error) {
    db, err := storm.Open(dbPath)
    if err != nil {
        return nil, err
    }

    defer db.Close()
    users := []User{}
    err = db.All(&users)

    if err != nil {
        return nil, err
    }
    return users, nil
}

// Returns a single user record from the database
func One(id bson.ObjectId) (*User, error) {
    db, err := storm.Open(dbPath)
    if err != nil {
        return nil, err
    }

    defer db.Close()
    user := new(User)
    err = db.One("ID", id, user)

    if err != nil {
        return nil, err
    }
    return user, nil
}

// Deletes a given record from the database
func Delete(id bson.ObjectId) error {
    db, err := storm.Open(dbPath)
    if err != nil {
        return err
    }

    defer db.Close()
    user := new(User)
    err = db.One("ID", id, user)

    if err != nil {
        return err
    }
    return db.DeleteStruct(user)
}

// Function is a member of the user struct, updates or creates a given record
func (user *User) Save() error {
    if err := user.validate(); err != nil {
        return err
    }

    db, err := storm.Open(dbPath)
    if err != nil {
        return err
    }

    defer db.Close()
    return db.Save(user)
}

// Makes sure that the record is valid
func (user *User) validate() error {
    if user.Name == "" {
        return ErrRecordInvalid
    }
    return nil
}
