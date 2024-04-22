package types

// 页面展示
type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
