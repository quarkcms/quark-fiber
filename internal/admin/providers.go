package admin

import (
	"github.com/quarkcms/quark-fiber/internal/admin/dashboards"
	"github.com/quarkcms/quark-fiber/internal/admin/resources"
)

// 注册服务
var Providers = []interface{}{
	&dashboards.Index{},
	&resources.Admin{},
	&resources.Role{},
	&resources.Permission{},
	&resources.Menu{},
	&resources.ActionLog{},
	&resources.Config{},
	&resources.File{},
	&resources.Picture{},
	&resources.WebConfig{},
	&resources.Account{},
}
