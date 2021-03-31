package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetMariaDB(username, password, addr, name string) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"UTC")
	client, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, errors.New("db is nil")
	}

	_ = client.Callback().Update().Remove("gorm:update_time_stamp")
	_ = client.Callback().Create().Remove("gorm:update_time_stamp")

	return client, nil
}
