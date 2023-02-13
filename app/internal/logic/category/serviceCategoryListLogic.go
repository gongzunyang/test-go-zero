package category

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gov-cms/app/internal/svc"
	"gov-cms/app/internal/types"
	"gov-cms/model"
	"math"
	"strconv"
	"strings"
)

type ServiceCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServiceCategoryListLogic {
	return &ServiceCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func IsContain(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func (l *ServiceCategoryListLogic) ServiceCategoryList(req *types.CategoryReq) (pageResp types.PageResp, err error) {
	// todo: add your logic here and delete this line
	if req.Name != "" || req.Status != 0 {
		var wherePar string
		if req.Name != "" {
			wherePar += "name like " + "'" + "%" + req.Name + "%'"
		}
		if req.Status != 0 {
			if req.Name != "" {
				wherePar += " and "
			}
			wherePar += " status = " + string(req.Status)
		}
		whereBuilderSearch := l.svcCtx.GovServiceCategoryModel.RowBuilder().Where(wherePar)
		searchData, error := l.svcCtx.GovServiceCategoryModel.FindData(l.ctx, whereBuilderSearch)
		if error != nil {
			return
		}
		var idsArr []int64

		for i := range searchData {
			idsArr = append(idsArr, searchData[i].Id, searchData[i].Pid)

		}
		allWhereBuilder := l.svcCtx.GovServiceCategoryModel.RowBuilder().Where("")
		allData, error := l.svcCtx.GovServiceCategoryModel.FindData(l.ctx, allWhereBuilder)
		var categoryArr []types.CategoryResp
		var category types.CategoryResp
		var cateArr []types.CategoryRespSon
		for _, value := range allData {
			isExists := IsContain(idsArr, value.Id)
			if isExists == true {
				if value.Pid == 0 {
					category.Id = value.Id
					category.Pid = value.Pid
					category.Name = value.Name.String
					category.Status = value.Status
					category.Remark = value.Remark.String
					category.Get_second_service = cateArr
					categoryArr = append(categoryArr, category)
				}
			}
		}
		var cate types.CategoryRespSon
		for i := range categoryArr {
			for ii := range searchData {
				if searchData[ii].Pid != 0 {
					if categoryArr[i].Id == searchData[ii].Pid {
						cate.Id = searchData[ii].Id
						cate.Pid = searchData[ii].Pid
						cate.Name = searchData[ii].Name.String
						cate.Status = searchData[ii].Status
						cate.Remark = searchData[ii].Remark.String
						categoryArr[i].Get_second_service = append(categoryArr[i].Get_second_service, cate)
					}

				}
			}
		}
		total_item := int64(len(categoryArr))
		pageResp.Current_page = req.Page
		pageResp.Page_size = req.PageSize
		pageResp.Total_page = int64(math.Ceil(float64(total_item) / float64(req.PageSize)))
		pageResp.Total_item = total_item
		start := req.PageSize * (req.Page - 1)
		end := req.PageSize * req.Page
		if end > total_item {
			end = total_item
		}
		var categoryReturnArr []types.CategoryResp
		for i := range categoryArr {
			if (int64(i) >= start) && (int64(i) < end) {
				categoryReturnArr = append(categoryReturnArr, categoryArr[i])
			}
		}
		pageResp.List = categoryReturnArr
	} else {
		//total number data
		whereCountBuilder := l.svcCtx.GovServiceCategoryModel.CountBuilder("id").Where("pid = 0")
		total_item, cErr := l.svcCtx.GovServiceCategoryModel.FindCount(l.ctx, whereCountBuilder)
		if cErr != nil {
			return
		}
		pageResp.Current_page = req.Page
		pageResp.Page_size = req.PageSize
		pageResp.Total_page = int64(math.Ceil(float64(total_item) / float64(req.PageSize)))
		pageResp.Total_item = total_item
		whereBuilder := l.svcCtx.GovServiceCategoryModel.RowBuilder().Where("pid = 0")
		homestayActivityList, error := l.svcCtx.GovServiceCategoryModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "id desc")
		if error != nil && error != model.ErrNotFound {
			logx.WithContext(l.ctx).Errorf("ActivityHomestayListLogic ActivityHomestayList 获取活动数据失败 id : %d ,err : %v", 453962077333278721, error)
			return pageResp, nil
		} else {
			var pidArr []string
			var categoryArr []types.CategoryResp
			var cateArr []types.CategoryRespSon
			var category types.CategoryResp
			for _, value := range homestayActivityList {
				pidArr = append(pidArr, strconv.FormatInt(value.Id, 10))
				category.Id = value.Id
				category.Pid = value.Pid
				category.Name = value.Name.String
				category.Remark = value.Remark.String
				category.Status = value.Status
				category.Get_second_service = cateArr
				categoryArr = append(categoryArr, category)
			}
			if len(pidArr) > 0 {
				str := strings.Join(pidArr, ",")
				whereBuilder = l.svcCtx.GovServiceCategoryModel.RowBuilder().Where("pid in (" + str + ")")
				sonData, error := l.svcCtx.GovServiceCategoryModel.FindData(l.ctx, whereBuilder)
				if error != nil {
					return
				}
				for key, value := range categoryArr {
					var sonson []types.CategoryRespSon
					var threeSon types.CategoryRespSon
					for _, v := range sonData {
						if value.Id == v.Pid {
							threeSon.Id = v.Id
							threeSon.Pid = v.Pid
							threeSon.Name = v.Name.String
							threeSon.Remark = v.Remark.String
							threeSon.Status = v.Status
							sonson = append(sonson, threeSon)
						}
					}
					categoryArr[key].Get_second_service = sonson
				}
				pageResp.List = categoryArr
			}
		}
	}
	return pageResp, nil
}
