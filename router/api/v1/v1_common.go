package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/pkg/codes"
	"github.com/olongfen/user_base/pkg/setting"
)

var (
	logCtl = log.NewLogFile("./log/log_ctrl", setting.Setting.IsProduct, setting.Setting.LogPatent)
)

// GetSession 获取会话信息
func GetSession(c *gin.Context) (ret *session.Session, err error) {
	var (
		ok bool
		s  interface{}
	)
	if s, ok = c.Get("sessionTag"); !ok {
		err = contrib.ErrSessionUndefined
		return
	}

	ret = s.(*session.Session)
	return
}

func GetSessionAndBindingForm(form interface{}, c *gin.Context) (ret *session.Session, code int, err error) {
	if ret, err = GetSession(c); err != nil {
		return nil, 401, err
	}
	if err = c.ShouldBind(form); err != nil {
		return nil, codes.CodeParamInvalid, err
	}
	return
}
