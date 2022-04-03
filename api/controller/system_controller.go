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

package controller

import (
	"fmt"
	"net/http"

	"github.com/dupman/server/constant"
	"github.com/dupman/server/dto"
	"github.com/dupman/server/resources"
	"github.com/dupman/server/service"
	sqltype "github.com/dupman/server/sql/type"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// SystemController data type.
type SystemController struct {
	httpService    service.HTTPService
	websiteService service.WebsiteService
	userService    service.UserService
}

// NewSystemController creates a new SystemController.
func NewSystemController(
	httpService service.HTTPService,
	websiteService service.WebsiteService,
	userService service.UserService,
) SystemController {
	return SystemController{
		httpService:    httpService,
		websiteService: websiteService,
		userService:    userService,
	}
}

// Websites returns the list of websites with tokens.
// @todo: Add Swagger docs.
func (c SystemController) Websites(ctx *gin.Context) {
	publicKey := ctx.GetHeader(constant.PublicKeyHeaderKey)
	if publicKey == "" {
		c.httpService.HTTPError(ctx, http.StatusBadRequest,
			fmt.Sprintf(resources.HeaderIsMissing, constant.PublicKeyHeaderKey))

		return
	}

	pagination := c.httpService.Paginate(ctx)
	websites, _ := c.websiteService.GetAll(pagination)
	websitesResponse := dto.WebsitesOnSystemResponse{}

	for i := 0; i < len(websites); i++ {
		// @todo: Implement user key caching.
		user, _ := c.userService.Get(websites[i].UserID)

		rawToken, _ := websites[i].Token.Decrypt(user.KeyPair.PrivateKey)
		websites[i].Token = sqltype.WebsiteToken(rawToken)

		tokenEncrypted, err := websites[i].Token.Encrypt(publicKey)
		if err != nil {
			c.httpService.HTTPError(ctx, http.StatusBadRequest, err.Error())

			return
		}

		websites[i].Token = sqltype.WebsiteToken(tokenEncrypted)
	}

	_ = copier.Copy(&websitesResponse, &websites)

	c.httpService.HTTPPaginatedResponse(ctx, http.StatusOK, websitesResponse, *pagination)
}
