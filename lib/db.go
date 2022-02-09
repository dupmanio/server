/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, February 2022
 */

package lib

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database data structure.
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new Database.
func NewDatabase(config Config, logger Logger) Database {
	dbConfig := config.Database
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return Database{
		DB: db,
	}
}
