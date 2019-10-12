package mgo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

var db *mgo.Database
var session *mgo.Session

var url = "mongodb://127.0.0.1:27017/local"
var personCol = "person"
var logCol = "log"
