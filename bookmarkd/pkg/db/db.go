// Copyright 2018 github.com/ucirello
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db // import "cirello.io/bookmarkd/pkg/db"

import (
	"cirello.io/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
)

// Config defines the environment for the database
type Config struct {
	Filename string
}

// Connect dials to the database.
func Connect(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", config.Filename)
	if err != nil {
		return nil, errors.E(err)
	}
	return db, nil
}
