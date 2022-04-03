/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package seeder

import (
	"github.com/dupman/server/lib"
	"github.com/dupman/server/model"
	"github.com/dupman/server/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	userService service.UserService
	database    lib.Database
	logger      lib.Logger
}

func NewUserSeeder(
	userService service.UserService,
	database lib.Database,
	logger lib.Logger,
) UserSeeder {
	return UserSeeder{
		userService: userService,
		database:    database,
		logger:      logger,
	}
}

func (s UserSeeder) Up() {
	var users []model.User

	passwordB, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	password := string(passwordB)

	users = append(users, model.User{
		Username:  "user_1",
		FirstName: "user1",
		LastName:  "user1",
		Email:     "user1@dup.man",
		Password:  password,
	})

	users = append(users, model.User{
		Username:  "user_2",
		FirstName: "user2",
		LastName:  "user2",
		Email:     "user2@dup.man",
		Password:  password,
	})

	for _, user := range users {
		go func(u model.User) {
			if err := s.userService.Create(&u); err != nil {
				s.logger.Error(err)
			}
		}(user)
	}
}

func (s UserSeeder) Down() {
	s.database.
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().
		Delete(&model.User{})
}
