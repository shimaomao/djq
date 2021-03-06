package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/constant"
	"mimi/djq/model"
	"mimi/djq/session"
	"mimi/djq/util"
	"net/http"
	"strings"
)

const (
	PermissionCheckModeAnd = iota
	PermissionCheckModeOr
)

func PermissionList(c *gin.Context) {
	permissionList := model.GetPermissionList()
	c.JSON(http.StatusOK, util.BuildSuccessResult(permissionList))
}

func checkPermission(c *gin.Context, permission string) {
	checkPermissionComplicated(c, PermissionCheckModeOr, permission)
	//sn, err := session.GetMi(c.Writer, c.Request)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//permissionStr, err := sn.Get(session.SessionNameMiPermission)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//permissionList := strings.Split(permissionStr, constant.Split4Permission)
	//if permissionList != nil || len(permissionList) != 0 {
	//	for _, pn := range permissionList {
	//		if permission == pn {
	//			return
	//		}
	//	}
	//}
	//c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
}

func checkPermissionComplicated(c *gin.Context, mode int, permissions ...string) {
	sn, err := session.GetMi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	permissionStr, err := sn.Get(session.SessionNameMiPermission)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	permissionList := strings.Split(permissionStr, constant.Split4Permission)
	if permissionList != nil || len(permissionList) != 0 {
		existCount := 0
		for i, permission := range permissions {
			for _, pn := range permissionList {
				if permission == pn {
					existCount++
					break
				}
			}
			if mode == PermissionCheckModeOr {
				if existCount > 0 {
					return
				}
			} else {
				if existCount != i+1 {
					c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
					return
				}
			}
		}
	}
	if mode == PermissionCheckModeOr {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedPermissionResult())
	}
}

func checkPermissionOr(c *gin.Context, permissions ...string) {
	checkPermissionComplicated(c, PermissionCheckModeOr, permissions...)
}

func PermissionAdminC(c *gin.Context) {
	checkPermission(c, "admin_c")
}

func PermissionAdminR(c *gin.Context) {
	checkPermission(c, "admin_r")
}

func PermissionAdminU(c *gin.Context) {
	checkPermission(c, "admin_u")
}

func PermissionAdminD(c *gin.Context) {
	checkPermission(c, "admin_d")
}

func PermissionRoleC(c *gin.Context) {
	checkPermission(c, "role_c")
}

func PermissionRoleR(c *gin.Context) {
	checkPermission(c, "role_r")
}

func PermissionRoleU(c *gin.Context) {
	checkPermission(c, "role_u")
}

func PermissionRoleD(c *gin.Context) {
	checkPermission(c, "role_d")
}

func PermissionAdvertisementC(c *gin.Context) {
	checkPermission(c, "advertisement_c")
}

func PermissionAdvertisementR(c *gin.Context) {
	checkPermission(c, "advertisement_r")
}

func PermissionAdvertisementU(c *gin.Context) {
	checkPermission(c, "advertisement_u")
}

func PermissionAdvertisementD(c *gin.Context) {
	checkPermission(c, "advertisement_d")
}

func PermissionAdvertisementCU(c *gin.Context) {
	checkPermissionOr(c, "advertisement_c", "advertisement_u")
}

func PermissionShopC(c *gin.Context) {
	checkPermission(c, "shop_c")
}

func PermissionShopR(c *gin.Context) {
	checkPermission(c, "shop_r")
}

func PermissionShopU(c *gin.Context) {
	checkPermission(c, "shop_u")
}

func PermissionShopD(c *gin.Context) {
	checkPermission(c, "shop_d")
}

func PermissionShopCU(c *gin.Context) {
	checkPermissionOr(c, "shop_c", "shop_u")
}

func PermissionShopClassificationC(c *gin.Context) {
	checkPermission(c, "shopClassification_c")
}

func PermissionShopClassificationR(c *gin.Context) {
	checkPermission(c, "shopClassification_r")
}

func PermissionShopClassificationU(c *gin.Context) {
	checkPermission(c, "shopClassification_u")
}

func PermissionShopClassificationD(c *gin.Context) {
	checkPermission(c, "shopClassification_d")
}

func PermissionUserC(c *gin.Context) {
	checkPermission(c, "user_c")
}

func PermissionUserR(c *gin.Context) {
	checkPermission(c, "user_r")
}

func PermissionUserU(c *gin.Context) {
	checkPermission(c, "user_u")
}

func PermissionUserD(c *gin.Context) {
	checkPermission(c, "user_d")
}

func PermissionPromotionalPartnerC(c *gin.Context) {
	checkPermission(c, "promotionalPartner_c")
}

func PermissionPromotionalPartnerR(c *gin.Context) {
	checkPermission(c, "promotionalPartner_r")
}

func PermissionPromotionalPartnerU(c *gin.Context) {
	checkPermission(c, "promotionalPartner_u")
}

func PermissionPromotionalPartnerD(c *gin.Context) {
	checkPermission(c, "promotionalPartner_d")
}

func PermissionPresentC(c *gin.Context) {
	checkPermission(c, "present_c")
}

func PermissionPresentR(c *gin.Context) {
	checkPermission(c, "present_r")
}

func PermissionPresentU(c *gin.Context) {
	checkPermission(c, "present_u")
}

func PermissionPresentD(c *gin.Context) {
	checkPermission(c, "present_d")
}

func PermissionPresentOrderC(c *gin.Context) {
	checkPermission(c, "presentOrder_c")
}

func PermissionPresentOrderR(c *gin.Context) {
	checkPermission(c, "presentOrder_r")
}

func PermissionPresentOrderU(c *gin.Context) {
	checkPermission(c, "presentOrder_u")
}

func PermissionPresentOrderD(c *gin.Context) {
	checkPermission(c, "presentOrder_d")
}

func PermissionCashCouponOrderC(c *gin.Context) {
	checkPermission(c, "cashCouponOrder_c")
}

func PermissionCashCouponOrderR(c *gin.Context) {
	checkPermission(c, "cashCouponOrder_r")
}

func PermissionCashCouponOrderU(c *gin.Context) {
	checkPermission(c, "cashCouponOrder_u")
}

func PermissionCashCouponOrderD(c *gin.Context) {
	checkPermission(c, "cashCouponOrder_d")
}

func PermissionRefundCU(c *gin.Context) {
	checkPermissionOr(c, "refund_c", "refund_u")
}

func PermissionRefundC(c *gin.Context) {
	checkPermission(c, "refund_c")
}

func PermissionRefundR(c *gin.Context) {
	checkPermission(c, "refund_r")
}

func PermissionRefundU(c *gin.Context) {
	checkPermission(c, "refund_u")
}

func PermissionRefundD(c *gin.Context) {
	checkPermission(c, "refund_d")
}

func PermissionRefundReasonC(c *gin.Context) {
	checkPermission(c, "refundReason_c")
}

func PermissionRefundReasonR(c *gin.Context) {
	checkPermission(c, "refundReason_r")
}

func PermissionRefundReasonU(c *gin.Context) {
	checkPermission(c, "refundReason_u")
}

func PermissionRefundReasonD(c *gin.Context) {
	checkPermission(c, "refundReason_d")
}

func PermissionIndexContactWayManage(c *gin.Context) {
	checkPermission(c, "indexContactWay_manage")
}

func PermissionPromotionalPartnerManage(c *gin.Context) {
	checkPermission(c, "promotionalPartner_manage")
}

func PermissionShopAccountRedPackManage(c *gin.Context) {
	checkPermission(c, "shopAccountRedPack_manage")
}

func PermissionCashCouponOrderCountManage(c *gin.Context) {
	checkPermission(c, "cashCouponOrderCount_manage")
}

//func PermissionCashCouponC(c *gin.Context) {
//	checkPermission(c, "cashCoupon_c")
//}
//
//func PermissionCashCouponR(c *gin.Context) {
//	checkPermission(c, "cashCoupon_r")
//}
//
//func PermissionCashCouponU(c *gin.Context) {
//	checkPermission(c, "cashCoupon_u")
//}
//
//func PermissionCashCouponD(c *gin.Context) {
//	checkPermission(c, "cashCoupon_d")
//}
