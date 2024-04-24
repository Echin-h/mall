package service

import (
	"context"
	"errors"
	conf "gin-mall/conf/sql"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/log"
	"gin-mall/pkg/util/upload"
	"gin-mall/respository/db/dao"
	"gin-mall/respository/db/model"
	"gin-mall/types"
	"mime/multipart"
	"strconv"
	"sync"
)

var ProductSrvIns *ProductSrv
var ProductSrvOnce sync.Once

type ProductSrv struct{}

func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}

func (s *ProductSrv) ProductList(ctx context.Context, req *types.ProductListReq) (resp interface{}, err error) {
	productDao := dao.NewProductDao(ctx)
	products, err := productDao.ListProductByCondition(req.CategoryID, req.BasePage)
	total, err := productDao.CountProductByCondition(req.CategoryID)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	pRespList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:            p.ID,
			Name:          p.Name,
			CategoryID:    p.CategoryID,
			Title:         p.Title,
			Info:          p.Info,
			ImgPath:       p.ImgPath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View(),
			CreatedAt:     p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			pResp.BossAvatar = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + pResp.BossAvatar
			pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
		}
		pRespList = append(pRespList, pResp)
	}

	resp = &types.DataListResp{
		Item:  pRespList,
		Total: total,
	}

	return
}

// 创建商品的流程
// 1. 从请求中获取用户信息
// 2. 从用户信息中获取商家信息
// 3. 获取商品的数据和图片文件
// 4. 上传图片到本地或者Oss
// 5. 创建商品
// 6. 创建商品图片

func (s *ProductSrv) ProductCreate(ctx context.Context, files []*multipart.FileHeader, req *types.ProductCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	boss, _ := dao.NewProductDao(ctx).GetBossById(uId)
	tmp, _ := files[0].Open()
	var path string
	if conf.Config.System.UploadModel == consts.UploadModelLocal {
		path, err = upload.ProductUploadToLocalStatic(tmp, uId, files[0].Filename)
		if err != nil {
			log.LogrusObj.Error(err)
			return nil, err
		}
	} else {
		// Oss upload
	}
	product := &model.Product{
		Name:          req.Name,
		CategoryID:    req.CategoryID,
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       path,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		tmp, _ := file.Open()
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			path, err = upload.ProductUploadToLocalStatic(tmp, uId, req.Name+num)
		} else {
			// Oss upload
		}
		productImg := &model.ProductImg{
			ProductID: product.ID, // product.ID is the primary key of the product
			ImgPath:   path,
		}
		err = dao.NewProductImgDaoByDB(productDao.DB).CreateProductImg(productImg)
		if err != nil {
			log.LogrusObj.Error(err)
			return nil, err
		}
		wg.Done()
	}
	wg.Wait()
	return
}

func (s *ProductSrv) ProductShow(ctx context.Context, req *types.ProductShowReq) (resp interface{}, err error) {
	p, err := dao.NewProductDao(ctx).ShowProduct(req.ID)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	pResp := &types.ProductResp{
		ID:            p.ID,
		Name:          p.Name,
		CategoryID:    p.CategoryID,
		Title:         p.Title,
		Info:          p.Info,
		ImgPath:       p.ImgPath,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		View:          p.View(),
		CreatedAt:     p.CreatedAt.Unix(),
		Num:           p.Num,
		OnSale:        p.OnSale,
		BossID:        p.BossID,
		BossName:      p.BossName,
		BossAvatar:    p.BossAvatar,
	}
	if conf.Config.System.UploadModel == consts.UploadModelLocal {
		pResp.BossAvatar = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + pResp.BossAvatar
		pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
	}
	resp = pResp
	return
}

func (s *ProductSrv) UpdateProduct(ctx context.Context, req *types.ProductUpdateReq) (resp interface{}, err error) {
	//u, err := ctl.GetUserInfo(ctx)
	//if err != nil {
	//	log.LogrusObj.Error(err)
	//	return nil, err
	//}
	//if u.Id != req.BossID {
	//	log.LogrusObj.Infoln("no auth")
	//	return nil, err
	//}
	u, _ := ctl.GetUserInfo(ctx)
	if u.Id != req.BossID {
		log.LogrusObj.Infoln("no auth")
		return nil, errors.New("no auth")
	}
	product := &model.Product{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Info:       req.Info,
		// ImgPath:       service.ImgPath,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		OnSale:        req.OnSale,
	}

	err = dao.NewProductDao(ctx).UpdateProduct(req.ID, product)
	if err != nil {
		log.LogrusObj.Error("update product error")
		return nil, err
	}

	return product, nil
}

func (s *ProductSrv) ProductDelete(ctx context.Context, req *types.ProductDeleteReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	err = dao.NewProductDao(ctx).DeleteProduct(req.ID, u.Id)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	return
}

// 搜索商品 TODO 后续用脚本同步数据MySQL到ES，用ES进行搜索

func (s *ProductSrv) ProductSearch(ctx context.Context, req *types.ProductSearchReq) (resp interface{}, err error) {
	productDao := dao.NewProductDao(ctx)
	products, err := productDao.SearchProduct(req.Name, req.BasePage)
	total, err := productDao.CountSearchProduct(req.Name)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	pRespList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:            p.ID,
			Name:          p.Name,
			CategoryID:    p.CategoryID,
			Title:         p.Title,
			Info:          p.Info,
			ImgPath:       p.ImgPath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View(),
			CreatedAt:     p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			pResp.BossAvatar = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + pResp.BossAvatar
			pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
		}
		pRespList = append(pRespList, pResp)
	}
	resp = &types.DataListResp{
		Item:  pRespList,
		Total: total,
	}
	return
}

func (s *ProductSrv) ProductImgList(ctx context.Context, req *types.ListProductImgReq) (resp interface{}, err error) {
	productImg, _ := dao.NewProductImgDao(ctx).ListProductImgByProductId(req.ID)
	for i := range productImg {
		if conf.Config.System.UploadModel == consts.UploadModelLocal {
			productImg[i].ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + productImg[i].ImgPath
		}
	}

	resp = &types.DataListResp{
		Item:  productImg,
		Total: int64(len(productImg)),
	}

	return
}
