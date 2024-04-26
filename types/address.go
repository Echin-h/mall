package types

type AddressCreateReq struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type AddressGetReq struct {
	Id uint `form:"id" json:"id"`
}

type AddressListReq struct {
	BasePage
}

type AddressUpdateReq struct{}

type AddressServiceReq struct {
	Id      uint   `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type AddressDeleteReq struct {
	Id uint `form:"id" json:"id"`
}
