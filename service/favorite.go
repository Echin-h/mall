package service

import (
	"context"
	conf "gin-mall/conf/sql"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/db/dao"
	"gin-mall/types"
	"sync"
)

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

type FavoriteSrv struct{}

func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

func (s *FavoriteSrv) FavoritesList(ctx context.Context, req *types.FavoritesServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	uid := u.Id

	favorDao := dao.NewFavoritesDao(ctx)
	favorites, total, err := favorDao.ListFavoritesByUserId(uid, req.PageNum, req.PageSize)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	for i := range favorites {
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			favorites[i].ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + favorites[i].ImgPath
		}
	}

	resp = &types.DataListResp{
		Item:  favorites,
		Total: total,
	}

	return

}
