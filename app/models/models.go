package models

import "github.com/jinzhu/gorm"

func AddTables(db *gorm.DB) error {
	err := db.AutoMigrate(Task{}).Error
	if err != nil {
		return err
	}
	return nil
}

func DropTables(db *gorm.DB) error {
	err := db.DropTable(Task{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FillTables(db *gorm.DB) error {
	for _, t := range fakeFillTables {
		err := db.Save(t).Error
		if err != nil {
			return err
		}
	}
	return nil
}

var fakeFillTables = []*Task{
	NewTask("Prepare today's agenda", false),
	NewTask("Define new product", true),
	NewTask("Schedule Doctor's appointment", false),
}
