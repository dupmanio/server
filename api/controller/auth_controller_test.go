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

package controller_test

import (
	"net/http"
	"testing"

	"github.com/dupman/server/resources"
	"github.com/dupman/server/test/helper"
	"github.com/dupman/server/test/seeder"
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthControllerSuite struct {
	suite.Suite
	e *httpexpect.Expect

	seeders seeder.Seeders
}

func TestAuthControllerSuite(t *testing.T) {
	t.Parallel()

	helper.BootstrapSuite(t, new(AuthControllerSuite))
}

func (s *AuthControllerSuite) SetupSuite() {
	setup := helper.SetupTester(s.T())

	s.e = setup.Expect
	s.seeders = setup.Seeders

	s.seeders.Up()
}

func (s *AuthControllerSuite) TearDownSuite() {
	s.seeders.Down()
}

func (s *AuthControllerSuite) Test_Register_emptyRequestBody() {
	res := s.e.POST("/auth/register").
		WithJSON(gin.H{}).
		Expect()

	res.Status(http.StatusBadRequest)
	res.JSON().Object().ValueEqual("code", http.StatusBadRequest)
	res.JSON().Object().ContainsKey("error")
}

func (s *AuthControllerSuite) Test_Register_notUnique() {
	res := s.e.POST("/auth/register").
		WithJSON(gin.H{
			"username":  "user_1",
			"email":     "user1@dup.man",
			"firstName": "user1",
			"lastName":  "user1",
			"password":  "password",
		}).
		Expect()

	res.Status(http.StatusBadRequest)
	res.JSON().Object().ValueEqual("code", http.StatusBadRequest)
	res.JSON().Object().Value("error").Array().ContainsOnly(resources.EmailIsTaken, resources.UsernameIsTaken)
}

func (s *AuthControllerSuite) Test_Register_success() {
	res := s.e.POST("/auth/register").
		WithJSON(gin.H{
			"username":  "new_user",
			"email":     "new_user@dup.man",
			"firstName": "new_user",
			"lastName":  "new_user",
			"password":  "password",
		}).
		Expect()

	res.Status(http.StatusCreated)
	res.JSON().Object().ValueEqual("code", http.StatusCreated)
	res.JSON().Object().ContainsKey("data")
}

func (s *AuthControllerSuite) Test_Login_emptyRequestBody() {
	res := s.e.POST("/auth/login").
		WithJSON(gin.H{}).
		Expect()

	res.Status(http.StatusBadRequest)
	res.JSON().Object().ValueEqual("code", http.StatusBadRequest)
	res.JSON().Object().ContainsKey("error")
}

func (s *AuthControllerSuite) Test_Login_userNotFound() {
	res := s.e.POST("/auth/login").
		WithJSON(gin.H{"username": "unknown_user", "password": "password"}).
		Expect()

	res.Status(http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("code", http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("error", resources.InvalidCredentials)
}

func (s *AuthControllerSuite) Test_Login_passwordIsIncorrect() {
	res := s.e.POST("/auth/login").
		WithJSON(gin.H{"username": "user_1", "password": "wrong_password"}).
		Expect()

	res.Status(http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("code", http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("error", resources.InvalidCredentials)
}

func (s *AuthControllerSuite) Test_Login_successWithUsername() {
	res := s.e.POST("/auth/login").
		WithJSON(gin.H{"username": "user_1", "password": "password"}).
		Expect()

	res.Status(http.StatusOK)
	res.JSON().Object().ValueEqual("code", http.StatusOK)
	res.JSON().Object().Value("data").Object().ContainsKey("token")
}

func (s *AuthControllerSuite) Test_Login_successWithEmail() {
	res := s.e.POST("/auth/login").
		WithJSON(gin.H{"username": "user1@dup.man", "password": "password"}).
		Expect()

	res.Status(http.StatusOK)
	res.JSON().Object().ValueEqual("code", http.StatusOK)
	res.JSON().Object().Value("data").Object().ContainsKey("token")
}
