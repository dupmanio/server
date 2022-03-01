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

package helper

type Pagination struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

const (
	DefaultLimit = 10
	MaxLimit     = 500
)

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = DefaultLimit
	} else if p.Limit > MaxLimit {
		p.Limit = MaxLimit
	}

	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}

	return p.Page
}
