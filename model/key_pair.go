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

package model

import (
	"github.com/dupman/server/helper"
	"gorm.io/gorm"
)

type KeyPair struct {
	Base
	PrivateKey string
	PublicKey  string
}

func (e *KeyPair) BeforeCreate(tx *gorm.DB) (err error) {
	if err = e.Base.BeforeCreate(tx); err != nil {
		return err
	}

	encryptor := helper.NewRSAEncryptor()
	if err = encryptor.GenerateKeyPair(); err == nil {
		e.PrivateKey = encryptor.PrivateKey()
		e.PublicKey, err = encryptor.PublicKey()
	}

	return err
}
