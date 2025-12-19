package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AuthsSearch struct {
	Id             *int        `json:"id" form:"id"`
	AuthKeyId      *int        `json:"authKeyId" form:"authKeyId"`
	LangPack       *string     `json:"langPack" form:"langPack"`
	ApiId          *int        `json:"apiId" form:"apiId"`
	ClientIp       *string     `json:"clientIp" form:"clientIp"`
	DeviceModel    *string     `json:"deviceModel" form:"deviceModel"`
	SystemLangCode *string     `json:"systemLangCode" form:"systemLangCode"`
	SystemVersion  *string     `json:"systemVersion" form:"systemVersion"`
	AppVersion     *string     `json:"appVersion" form:"appVersion"`
	Proxy          *string     `json:"proxy" form:"proxy"`
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UpdatedAtRange []time.Time `json:"updatedAtRange" form:"updatedAtRange[]"`
	request.PageInfo
}
