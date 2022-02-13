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

package service

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dupman/server/dto"
	"github.com/dupman/server/lib"
	"github.com/dupman/server/model"
	"github.com/dupman/server/resources"
	"github.com/golang-jwt/jwt"
)

var errInvalidToken = errors.New(resources.InvalidToken)

// JWTAuthService data structure.
type JWTAuthService struct {
	config      lib.JWTConfig
	userService UserService
	logger      lib.Logger
}

// NewJWTAuthService creates a new JWTAuthService.
func NewJWTAuthService(config lib.Config, userService UserService, logger lib.Logger) JWTAuthService {
	return JWTAuthService{
		config:      config.JWT,
		userService: userService,
		logger:      logger,
	}
}

// Authorize authorizes the generated token.
func (s JWTAuthService) Authorize(tokenString string) (user model.User, err error) {
	var claims dto.JWTClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		_, publicKey, err := s.getKeys()
		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})

	if err != nil || !token.Valid {
		s.logger.Error(err)

		return user, fmt.Errorf("%s. %w", resources.InvalidToken, err)
	}

	if user, err = s.userService.GetUser(claims.ID); err != nil {
		return user, errInvalidToken
	}

	return user, nil
}

// GenerateToken creates jwt auth token.
func (s JWTAuthService) GenerateToken(user *model.User) (response dto.JWTResponse, err error) {
	now := time.Now()
	expiry := now.Add(s.config.Expiration)

	claims := dto.JWTClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	}

	privateKey, _, err := s.getKeys()
	if err != nil {
		s.logger.Error(err)

		return response, err
	}

	response.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(privateKey)
	if err != nil {
		s.logger.Error(err)

		return response, err
	}

	response.TokenType = "Bearer"
	response.ExpiresIn = int64(s.config.Expiration.Seconds())

	return response, nil
}

func (s JWTAuthService) getKeys() (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, err error) {
	privateKeyBytes, err := ioutil.ReadFile(s.config.PrivateKeyLocation)
	if err != nil {
		return nil, nil, err
	}

	publicKeyBytes, err := ioutil.ReadFile(s.config.PublicKeyLocation)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, publicKey, nil
}
