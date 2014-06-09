package controllers

import (
	"github.com/donnpebe/todoapirevel/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

var Dbm *gorm.DB

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

func (c *GormController) Begin() revel.Result {
	txn := Dbm.Begin()
	checkPANIC(txn.Error)
	c.Txn = txn
	return nil
}

func (c *GormController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	checkPANIC(c.Txn.Commit().Error)
	c.Txn = nil
	return nil
}

func (c *GormController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	checkPANIC(c.Txn.Rollback().Error)
	c.Txn = nil
	return nil
}

func InitDB() {
	var driver, spec string
	var found bool
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("No db.spec found.")
	}
	revel.INFO.Println(spec)

	dbcon, err := gorm.Open(driver, spec)
	checkPANIC(err)

	dbcon.SetLogger(gorm.Logger{revel.INFO})
	dbcon.LogMode(true)

	Dbm = &dbcon
	revel.INFO.Println("Connection made to DB")
}

func SetupTables() {
	revel.INFO.Println("Setting up Prod DB")
	addTables()
}

func SetupDevDB() {
	revel.INFO.Println("Setting up Dev DB")
	dropTables()
	addTables()

	models.FillTables(Dbm)
}

func addTables() {
	revel.INFO.Println("AutoMigrate tables")
	models.AddTables(Dbm)
}

func dropTables() {
	revel.INFO.Println("Dropping tables")
	models.DropTables(Dbm)
}
