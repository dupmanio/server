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

package resources

// String resource constants.
const (
	AccessDenied          = "access denied"
	InvalidCredentials    = "invalid credentials"
	UnableToCreateToken   = "unable to generate token"
	InvalidToken          = "invalid token"
	FailedHashingPassword = "failed hashing password"
	KeyIsRequired         = "Key '%s' is required"
	ValueIsLessThenMin    = "Value of field '%s' should be at least %s characters"
	ValueIsNotEmail       = "Value of field '%s' is not a valid Email address"
	UsernameIsTaken       = "Username is taken"
	EmailIsTaken          = "Email is taken"
	ValueIsNotURL         = "Value of field '%s' is not a valid URL address"
	UnableToDecodeKey     = "unable to decode key"
)
