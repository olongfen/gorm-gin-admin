package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/pkg/codes"
	"github.com/olongfen/gorm-gin-admin/src/service"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"strconv"
)

// @tags 管理员
// @Summary 添加角色
// @Description 添加角色
// @Accept json
// @Produce json
// @Param {} body utils.FormRole true "添加角色form"
// @Success 200  {object}
// @Failure 500  {object}
// @router /api/v1/admin/addRole [post]
func AddRole(c *gin.Context)  {
	var(
		err error
		code = codes.CodeProcessingFailed
		form =new(utils.FormRole)
		data *models.Role
	)
	defer func() {
		if err!=nil{
			app.NewGinResponse(c).Fail(code,err.Error()).Response()
		}
	}()
	if _,code,err = GetSessionAndBindingForm(form,c);err!=nil{
		return
	}
	if data,err = service.AddRole(form);err!=nil{
		return
	}
	app.NewGinResponse(c).Success(data).Response()
}

// @tags 管理员
// @Summary 更新角色
// @Description 更新角色
// @Accept json
// @Produce json
// @Param {} body utils.FormUpdateRole true "更新角色form"
// @Success 200  {object}
// @Failure 500  {object}
// @router /api/v1/admin/editRole [put]
func EditRole(c *gin.Context)  {
	var(
		err error
		code = codes.CodeProcessingFailed
		form =new(utils.FormUpdateRole)
		s *session.Session
	)
	defer func() {
		if err!=nil{
			app.NewGinResponse(c).Fail(code,err.Error()).Response()
		}
	}()
	if s,code,err = GetSessionAndBindingForm(form,c);err!=nil{
		return
	}
	if err = service.UpdateRole(s.UID,form);err!=nil{
		return
	}
	app.NewGinResponse(c).Success(nil).Response()
}

// @tags 管理员
// @Summary 删除角色
// @Description 删除角色
// @Accept json
// @Produce json
// @Param id query string true "角色id"
// @Success 200  {object}
// @Failure 500  {object}
// @router /api/v1/admin/removeRole [delete]
func RemoveRole(c *gin.Context)  {
	var(
		err error
		code = codes.CodeProcessingFailed
		id int
		s *session.Session
	)
	defer func() {
		if err!=nil{
			app.NewGinResponse(c).Fail(code,err.Error()).Response()
		}
	}()
	if _id,ok:=c.GetQuery("id");!ok{
		err = utils.ErrParamInvalid
		code = codes.CodeParamInvalid
		return
	}else {
		if id,err = strconv.Atoi(_id);err!=nil{
			code = codes.CodeParamInvalid
			return
		}
	}
	if s,code,err = GetSession(c);err!=nil{
		return
	}
	if err = service.DelRole(s.UID,id);err!=nil{
		return
	}
	app.NewGinResponse(c).Success(nil).Response()
}

// @tags 管理员
// @Summary 删除角色
// @Description 删除角色
// @Accept json
// @Produce json
// @Param id query string true "角色id"
// @Success 200  {object}
// @Failure 500  {object}
// @router /api/v1/admin/getRoleList [get]
func GetRoleList(c *gin.Context)  {
	var(
		err error
		code = codes.CodeProcessingFailed
		data []*models.Role
	)
	defer func() {
		if err!=nil{
			app.NewGinResponse(c).Fail(code,err.Error()).Response()
		}
	}()

	if data,err = models.GetRoleList();err!=nil{
		return
	}
	app.NewGinResponse(c).Success(data).Response()
}