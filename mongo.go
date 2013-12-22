/*
Created by zhhigh
date:2013-11-10
*/

/*
Worker side api for gearman

usage:
w = worker.New(worker.Unlimited)
w.AddFunction("foobar", foobar)
w.AddServer("127.0.0.1:4730")
w.Work() // Enter the worker's main loop

The definition of the callback function 'foobar' should suit for the type 'JobFunction'.
It looks like this:

func foobar(job *Job) (data []byte, err os.Error) {
    //sth. here
    //plaplapla...
    return
}
*/

package o2o

import (
	"labix.org/v2/mgo"
	"fmt"
	//"reflect"
)

type MongoDBConn struct {
	session *mgo.Session
	dbName  string
	tableName string
}

var ResultCount int

func NewMongoDBConn() *MongoDBConn {
	return &MongoDBConn{}
}

func (m *MongoDBConn) Connect(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
		//return nil
	}
	m.session = session
	session.SetMode(mgo.Monotonic, true)
	return m.session
}

func (m *MongoDBConn) Stop() {
	m.session.Close()
	fmt.Println("The Mongodb is Closed!!!")
}


func (m *MongoDBConn) SetDBName(dbName string){
     m.dbName = dbName
}

func (m *MongoDBConn) SetTableName(tableName string){
     m.tableName = tableName
}

/*only find One
*/
/*
func (m *MongoDBConn) FindOne(query interface{},result interface{})(interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Find(query).One(&result)
	if err != nil {
		panic(err)
		return nil
	}
    //fmt.Println(result.)
	return result
} */


func (m *MongoDBConn) FindOne(query interface{},result interface{})(interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Find(query).One(result)
	if err != nil {
		//panic(err)
		return nil
	}
	//fmt.Println(result.)
	return result
}




/*find all
*/
func (m *MongoDBConn) FindAll(result interface{})(interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Find(nil).All(result)


	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n",err)
		return nil
	}
    return result
}

/*find one or more from condition
*/
func (m *MongoDBConn) Find(query interface{},result interface{})(interface{}){

	Collection := m.session.DB(m.dbName).C(m.tableName)
	Collection.Find(query).All(result)

	////
	ResultCount = 0
	ResultCount,_ = Collection.Count()
	//
	var p *[]interface {}//判断result是否有返回
	p = result.(*[]interface{})
	if (*p == nil){
		return nil
	}
    return result
}

/*insert data*/
func (m *MongoDBConn) Insert(data interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Insert(data)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n",err)
	}
}

func (m *MongoDBConn) Update(selector interface{},change interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Update(selector,change)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n",err)
	}
}

func (m *MongoDBConn) Delete(data interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.Remove(data)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n",err)
	}
}


func (m *MongoDBConn) DeleteID(id interface{}){
	Collection := m.session.DB(m.dbName).C(m.tableName)
	err := Collection.RemoveId(id)
	if err != nil {
		panic(err)
		fmt.Printf("Log Info======:\n",err)
	}
}

