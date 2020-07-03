package srv_user

import (
	"encoding/json"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/query"
	"github.com/olongfen/user_base/utils"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// AddUser 添加哟个用户
func AddUser(form *utils.AddUserForm) (ret *models.UserBase, err error) {
	var (
		u = new(models.UserBase)
	)
	if len(utils.RegPhoneNum.FindString(form.Phone)) == 0 {
		err = utils.ErrPhoneInvalid
		return
	}
	u.Phone = form.Phone
	u.Username = form.Phone
	if _d, _err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		u.LoginPwd = string(_d)
	}
	u.Uid = uuid.NewV4().String()
	if err = u.InsertUserData(); err != nil {
		return
	}
	return u, nil
}

// EditUser 修改用户信息
func EditUser(uid string, form *utils.FormEditUser) (ret *models.UserBase, err error) {
	var (
		dataMap map[string]interface{}
		data    = &models.UserBase{}
	)
	if dataMap, err = form.Valid(); err != nil {
		return
	}
	//
	if _d, _err := json.Marshal(dataMap); _err != nil {
		return
	} else {
		if err = json.Unmarshal(_d, data); err != nil {
			return
		}
	}
	if err = data.UpdateUser(uid); err != nil {
		return nil, err
	}

	//
	ret = data
	return
}

// ChangePwd 修改密码
func ChangePwd(uid string, oldPasswd, newPasswd string) (err error) {
	var (
		data = &models.UserBase{}
	)
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		err = utils.ErrFormParamInvalid
		return
	}
	if err = data.GetUserByUId(uid); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPwd), []byte(oldPasswd)); err != nil {
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return data.UpdateUserOneColumn(uid, "login_pwd", string(_d))
	}
}

// ChangePayPwd 修改密码
func ChangePayPwd(uid string, oldPasswd, newPasswd string) (err error) {
	var (
		data = &models.UserBase{}
	)
	if len(oldPasswd) == 0 || len(newPasswd) == 0 {
		err = utils.ErrFormParamInvalid
		return
	}
	if err = data.GetUserByUId(uid); err != nil {
		return
	}
	if len(data.PayPwd) == 0 {
		err = utils.ErrPayPwdNotSet
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(data.PayPwd), []byte(oldPasswd)); err != nil {
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return data.UpdateUserOneColumn(uid, "pay_pwd", string(_d))
	}
}

// SetUserPayPwd
func SetUserPayPwd(uid string, pwd string) (err error) {
	u := new(models.UserBase)
	if err = u.GetUserByUId(uid); err != nil {
		return
	}
	if len(u.PayPwd) > 0 {
		err = utils.ErrActionNotAllow
		return
	}
	if _d, _err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); _err != nil {
		err = _err
		return
	} else {
		return u.UpdateUserOneColumn(uid, "pay_pwd", string(_d))
	}
}

// GetUserList
func GetUserList(uid string, form *utils.FormUserList) (ret []*models.UserBase, err error) {
	cond := map[string]interface{}{}
	if form.Username != "" {
		cond["username"] = form.Username
	}
	if form.Status != "" {
		cond["status"] = form.Status
	}
	if form.ID != "" {
		cond["id"] = form.ID
	}
	if form.CreatedTime != "" {
		cond["created_at"] = form.CreatedTime
	}
	var (
		q *query.Query
	)
	if q, err = query.NewQuery(form.PageNum, form.PageSize).ValidCond(cond); err != nil {
		return
	}
	return models.GetUserList(q)
}
