package dao

import (
	"log"
	"fmt"
	"time"

	. "github.com/ogrima/go-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
) 

type MoviesDAO struct {
	Server   string
	Database string
	Username string
	Password string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
	hosts      = "mongodb-rh:27017"
    database   = "movies_db"
    username   = "mongo"
    password   = "mongo"
)

func (m *MoviesDAO) Connect() {

	info := &mgo.DialInfo{
        Addrs:    []string{hosts},
        Timeout:  60 * time.Second,
        Database: database,
        Username: username,
        Password: password,
    }

	//session, err := mgo.Dial(m.Server)
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ultimo Passo.")
	db = session.DB(m.Database)
}

func (m *MoviesDAO) GetAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesDAO) GetByID(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesDAO) Create(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *MoviesDAO) Update(id string, movie Movie) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &movie)
	return err
}
