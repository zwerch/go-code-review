package api

import (
	"coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply(c *gin.Context) {
	apiReq := entity.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a *API) Create(c *gin.Context) {
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	_, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	// TODO: proper error handling
	// TODO: should be StatusCreated (201) and possibly return coupon
	c.Status(http.StatusOK)
}

func (a *API) Get(c *gin.Context) {
	apiReq := entity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		return
	}
	// TODO: proper error handling
	// TODO: return possible errors for each coupon
	c.JSON(http.StatusOK, coupons)
}
