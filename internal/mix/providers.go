package mix

import (
	"github.com/quarkcms/quark-go/internal/mix/resources/forms"
	"github.com/quarkcms/quark-go/internal/mix/resources/pages"
)

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
	&pages.Test{},
	&forms.Index{},
}
