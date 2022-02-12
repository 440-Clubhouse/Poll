package models

import (
    "github.com/440-clubhouse/poll/config"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database.
func Init() (err error) {
    DB, err = gorm.Open(sqlite.Open(config.Conf.Database), &gorm.Config{})
    if err != nil {
        return
    }

    err = DB.AutoMigrate()
    if err != nil {
        return
    }

    return
}
