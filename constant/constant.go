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

package constant

type ContextKey string

const (
	// UserIDKey represents key for storing authenticated user ID.
	UserIDKey = "user_id"

	// EncryptionKeyKey represents key for website encryption key.
	EncryptionKeyKey ContextKey = "encryption_key"

	// PublicKeyHeaderKey represents key for the Public Key header.
	PublicKeyHeaderKey = "X-Public-Key"
)
