package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/phplaber/finder/configs"
)

func save(username string, hostname string, applist []string) {
	db, err := sql.Open(configs.DriverName, configs.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into user_app_collect(username, hostname, app) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, appname := range applist {
		appname = strings.TrimSpace(appname)
		if len(appname) > 0 {
			// 去重插入
			var uname string
			err := db.QueryRow("select username from user_app_collect where username=? and app=?", username, appname).Scan(&uname)
			switch {
			case err == sql.ErrNoRows:
				if _, err := stmt.Exec(username, hostname, appname); err != nil {
					log.Fatal(err)
				}
			case err != nil:
				log.Fatal(err)
			default:
				// nothing to do
			}
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
