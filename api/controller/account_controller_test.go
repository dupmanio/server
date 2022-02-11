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
	"fmt"
	"net/http"
	"testing"

	"github.com/dupman/server/resources"
	"github.com/dupman/server/test/helper"
	"github.com/dupman/server/test/seeder"
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AccountControllerSuite struct {
	suite.Suite
	e *httpexpect.Expect

	seeders seeder.Seeders
}

func TestAccountControllerSuite(t *testing.T) {
	t.Parallel()

	helper.BootstrapSuite(t, new(AccountControllerSuite))
}

func (s *AccountControllerSuite) SetupSuite() {
	setup := helper.SetupTester(s.T())

	s.e = setup.Expect
	s.seeders = setup.Seeders

	s.seeders.Up()
}

func (s *AccountControllerSuite) TearDownSuite() {
	s.seeders.Down()
}

func (s *AccountControllerSuite) Test_Login_withoutAuthHeader() {
	res := s.e.GET("/account").
		Expect()

	res.Status(http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("code", http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("error", resources.AccessDenied)
}

func (s *AccountControllerSuite) Test_Login_invalidToken() {
	res := s.e.GET("/account").
		WithHeader("Authorization", "Bearer invalid_token").
		Expect()

	res.Status(http.StatusUnauthorized)
	res.JSON().Object().ValueEqual("code", http.StatusUnauthorized)
	res.JSON().Object().Value("error").String().Contains(resources.InvalidToken)
}

func (s *AccountControllerSuite) Test_Login_success() {
	res := s.e.POST("/auth/login").
		WithJSON(gin.H{"username": "user_1", "password": "password"}).
		Expect()

	authToken := res.JSON().Object().Value("data").Object().Value("token").String().Raw()

	res = s.e.GET("/account").
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", authToken)).
		Expect()

	res.Status(http.StatusOK)
	res.JSON().Object().ValueEqual("code", http.StatusOK)
	res.JSON().Object().Value("data").Object().ValueEqual("username", "user_1")
}
