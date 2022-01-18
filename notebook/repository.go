package notebook

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Migrate() error {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = db.AutoMigrate(&note{}, &tag{}, &noteTag{}); err != nil {
		return err
	}

	return nil
}
