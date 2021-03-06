package database

import (
	"gopkg.in/mgo.v2"
)

var (
	mgoSession   *mgo.Session
	databaseName = "myDB"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		err = ensureUsersIndex(mgoSession)
		if err != nil {
			panic(err)
		}
		err = ensureAnnotationsIndex(mgoSession)
		if err != nil {
			panic(err)
		}
		err = ensureAnnotationUserIndex(mgoSession)
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}

func usingCollection(collection string) *mgo.Collection {
	session := getSession()
	return session.DB(databaseName).C(collection)
}
