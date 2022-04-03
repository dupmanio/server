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

package dto

// HTTPResponse represents HTTP response data.
// swagger:model HTTPResponse
type HTTPResponse struct {
	// HTTP status code
	//
	// required: true
	Code int `json:"code"`

	// HTTP response body
	//
	// required: true
	Data interface{} `json:"data"`

	// Pagination data
	Pagination interface{} `json:"pagination,omitempty"`
}

// HTTPError represents HTTP error data.
// swagger:model HTTPError
type HTTPError struct {
	// HTTP status code
	//
	// required: true
	Code int `json:"code"`

	// HTTP error body
	//
	// required: true
	Error interface{} `json:"error"`
}
