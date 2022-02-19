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

package controller

import (
	"net/http"

	"github.com/dupman/server/dto"
	"github.com/dupman/server/model"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// WebsiteController data type.
type WebsiteController struct {
	httpService    service.HTTPService
	websiteService service.WebsiteService
}

// NewWebsiteController creates a new WebsiteController.
func NewWebsiteController(
	httpService service.HTTPService,
	websiteService service.WebsiteService,
) WebsiteController {
	return WebsiteController{
		httpService:    httpService,
		websiteService: websiteService,
	}
}

func (c WebsiteController) Create(ctx *gin.Context) {
	var (
		websiteModel = model.Website{}

		websitePayload  *dto.WebsiteOnCreate
		websiteResponse dto.WebsiteOnResponse
		err             error
	)

	if err = ctx.ShouldBind(&websitePayload); err != nil {
		c.httpService.HTTPValidationError(ctx, err)

		return
	}

	user, _ := c.httpService.CurrentUser(ctx)
	_ = copier.Copy(&websiteModel, &websitePayload)
	websiteModel.UserID = user.ID

	if err = c.websiteService.Create(&websiteModel, user.KeyPair.PublicKey); err != nil {
		c.httpService.HTTPError(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	_ = copier.Copy(&websiteResponse, &websiteModel)
	c.httpService.HTTPResponse(ctx, http.StatusCreated, websiteResponse)
}

func (c WebsiteController) All(ctx *gin.Context) {
	websitesResponse := dto.WebsitesOnResponse{}

	uid := c.httpService.CurrentUserID(ctx)
	websites, _ := c.websiteService.GetByUser(uid)
	_ = copier.Copy(&websitesResponse, &websites)

	c.httpService.HTTPResponse(ctx, http.StatusOK, websitesResponse)
}
