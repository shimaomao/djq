package model

type Permission struct {
	Code        string `json:"code" desc:"编码"`
	Name        string `json:"name" desc:"名称"`
	Description string `json:"description" desc:"描述"`
}

func GetPermissionList() []*Permission {
	return []*Permission{
		{"admin_c", "添加管理员信息", "添加管理员信息"},
		{"admin_r", "查看管理员信息", "查看管理员信息"},
		{"admin_u", "更新管理员信息", "更新管理员信息"},
		{"admin_d", "删除管理员信息", "删除管理员信息"},

		{"role_c", "添加角色信息", "添加角色信息"},
		{"role_r", "查看角色信息", "查看角色信息"},
		{"role_u", "更新角色信息", "更新角色信息"},
		{"role_d", "删除角色信息", "删除角色信息"},

		{"advertisement_c", "添加广告信息", "添加广告信息"},
		{"advertisement_r", "查看广告信息", "查看广告信息"},
		{"advertisement_u", "更新广告信息", "更新广告信息"},
		{"advertisement_d", "删除广告信息", "删除广告信息"}}
}