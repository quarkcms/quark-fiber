package actions

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/ui/admin/actions"
)

type DetailLink struct {
	actions.Link
}

// 初始化
func (p *DetailLink) Init(name string) *DetailLink {
	// 初始化父结构
	p.ParentInit()

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 文字
	p.Name = name

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 跳转链接
func (p *DetailLink) GetHref(c *fiber.Ctx) string {
	return "#/index?api=" + strings.Replace(strings.Replace(c.Path(), "/api/", "", -1), "/index", "/detail&id=${id}", -1)
}
