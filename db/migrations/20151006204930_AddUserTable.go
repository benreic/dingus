package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20151006204930(txn *sql.Tx) {

	sql := "create table driver (" +
		"driver_id int(11) unsigned not null auto_increment, " +
		"email varchar(255) not null, " +
		"handle varchar(255) not null, " +
		"created_on datetime not null, " +
		"PRIMARY KEY (driver_id)" +
		") ENGINE=InnoDB"

	txn.Exec(sql)

}

// Down is executed when this migration is rolled back
func Down_20151006204930(txn *sql.Tx) {

	txn.Exec("DROP TABLE driver")

}
