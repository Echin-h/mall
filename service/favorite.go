package service

import (
	"context"
	"errors"
	conf "gin-mall/conf/sql"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/respository/db/dao"
	"gin-mall/respository/db/model"
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

func (s *FavoriteSrv) FavoriteCreate(ctx context.Context, req *types.FavoriteCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	uid := u.Id

	favorDao := dao.NewFavoritesDao(ctx)
	exist, _ := favorDao.FavoriteExistOrNot(uid, req.ProductId)
	if exist {
		log.LogrusObj.Error(`favorite exist`)
		return nil, errors.New("favorite exist")
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(req.BossId)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ProductId)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	favorite := &model.Favorite{
		UserID:    uid,
		ProductID: req.ProductId,
		BossID:    req.BossId,
		User:      *user,
		Boss:      *boss,
		Product:   *product,
	}

	err = favorDao.CreateFavorite(favorite)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}

func (s *FavoriteSrv) FavoriteDelete(ctx context.Context, req *types.FavoriteDeleteReq) (resp interface{}, err error) {
	favorDao := dao.NewFavoritesDao(ctx)
	err = favorDao.DeleteFavoriteById(req.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}

	return
}
