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

func PermissionList(c *gin.Context) {
	permissionList := model.GetPermissionList()
	c.JSON(http.StatusOK, util.BuildSuccessResult(permissionList))
}

func checkPermission(c *gin.Context, permission string) {
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
		for _, pn := range permissionList {
			if permission == pn {
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildNeedPermissionResult())
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