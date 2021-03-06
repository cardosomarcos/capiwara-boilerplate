package db

import (
	"errors"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Session *mgo.Session
var Database = "capi"

type MongoCollection struct {
	*mgo.Collection
}

func (conn *MongoCollection) FindOne(id string) (interface{}, error) {
	var data map[string]interface{}

	if err := conn.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data); err != nil {
		return nil, err
	}

	data["id"] = data["_id"].(bson.ObjectId).Hex()
	return data, nil
}

func MongoRepoBuilder(repoDefinition RepositoryDef) Repository {
	mongoCollection, err := prepareDB(Session, Database, repoDefinition.GetName())
	if err != nil {
		log.Printf("error in mongo builder: %v", err)
		return nil
	}
	return &MongoCollection{
		mongoCollection,
	}
}

func prepareDB(session *mgo.Session, db, dbcollection string) (*mgo.Collection, error) {
	if session == nil {
		return nil, errors.New("session is nil")
	}
	return session.DB(db).C(dbcollection), nil
}

func NewSession(username, password string) error {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Username: username,
		Password: password,
		Database: Database,
		Timeout:  10 * time.Second,
	})

	if err != nil {
		log.Printf("db error: %v", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	Session = session
	return nil
}
