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

package sqltype

import (
	"context"

	"github.com/dupman/encryptor"
	"github.com/dupman/server/constant"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WebsiteToken string

func (t *WebsiteToken) Decrypt(privateKey string) (decrypted string, err error) {
	rsaEncryptor := encryptor.NewRSAEncryptor()

	err = rsaEncryptor.SetPrivateKey(privateKey)
	if err != nil {
		return decrypted, err
	}

	return rsaEncryptor.Decrypt(string(*t))
}

func (t *WebsiteToken) Encrypt(publicKey string) (encrypted string, err error) {
	rsaEncryptor := encryptor.NewRSAEncryptor()
	if err = rsaEncryptor.SetPublicKey(publicKey); err != nil {
		return encrypted, err
	}

	if encrypted, err = rsaEncryptor.Encrypt(string(*t)); err == nil {
		return encrypted, nil
	}

	return encrypted, err
}

func (t WebsiteToken) GormValue(ctx context.Context, tx *gorm.DB) (expr clause.Expr) {
	if encryptionKey, ok := ctx.Value(constant.EncryptionKeyKey).(string); ok {
		if encrypted, err := t.Encrypt(encryptionKey); err == nil {
			return clause.Expr{SQL: "?", Vars: []interface{}{encrypted}}
		}
	}

	return
}
