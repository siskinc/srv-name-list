package list_item

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"sync"
)

type ItemHitAllReq struct {
	Namespace string   `json:"namespace" binding:"required"`          // 命名空间
	CodeList  []string `json:"code_list"`                             // 名单类型code, 如果不传, 或者传入的长度为0, 默认命中所有code
	Data      string   `json:"data" binding:"required" minlength:"2"` // 数据
}

type ItemHitAllResp struct {
	HitListItemList []*models.ListItem `json:"hit_list_item_list"`
}

func (service *Service) queryAllListType(namespace string, codeList []string) (listTypeList []*models.ListType, err error) {
	filter := bson.D{
		bson.E{
			Key:   "namespace",
			Value: namespace,
		},
		bson.E{
			Key:   "is_valid",
			Value: true,
		},
	}
	if len(codeList) > 0 {
		filter = append(filter, bson.E{
			Key: "code",
			Value: bson.M{
				"$in": codeList,
			},
		})
	}
	listTypeServiceObj := listTypeService.NewService()
	listTypeList, err = listTypeServiceObj.QueryAll(filter)
	return
}

type itemHitInfo struct {
	wg            *sync.WaitGroup
	listType      *models.ListType
	listItem      *models.ListItem
	allMultiValue map[string]string
	err           error
}

func (service *Service) itemHitOne(namespace, code, value string) (listItem *models.ListItem, err error) {
	filter := bson.D{
		bson.E{
			Key:   "namespace",
			Value: namespace,
		},
		bson.E{
			Key:   "code",
			Value: code,
		},
		bson.E{
			Key:   "value",
			Value: value,
		},
		bson.E{
			Key:   "is_valid",
			Value: true,
		},
	}
	var result []*models.ListItem
	result, _, err = service.listItemRepoObj.Query(filter, 1, 1, "")
	if err != nil {
		return
	}
	if len(result) > 0 {
		listItem = result[0]
	}
	return
}

func (service *Service) itemHit(info *itemHitInfo) {
	if info.wg != nil {
		defer info.wg.Done()
	}
	if info.listType == nil {
		return
	}
	listType := info.listType
	namespace, code := listType.Namespace, listType.Code
	allMultiValue := info.allMultiValue
	fields := listType.Fields
	values := make([]string, len(fields))
	for i := range fields {
		field := fields[i]
		values[i] = allMultiValue[field]
	}
	value := strings.Join(values, ",")
	listItem, err := service.itemHitOne(namespace, code, value)
	if err != nil {
		logrus.Errorf("hit one item have an err: %v, namespace: %s, code: %s, value: %s",
			err, namespace, code, value)
		info.err = err
		return
	}
	info.listItem = listItem
	return
}

func (service *Service) ItemHitAll(req *ItemHitAllReq) (resp *ItemHitAllResp, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("hit all list item req is nil"))
		return
	}

	namespace, codeList := req.Namespace, req.CodeList

	resp = &ItemHitAllResp{}

	var exist bool
	namespaceServiceObj := namespaceService.NewService()
	exist, err = namespaceServiceObj.CheckExist(namespace)
	if err != nil {
		return
	}
	if !exist {
		err = errorx.NewError(error_code.CustomForbiddenNotFoundNamespace, fmt.Errorf("not found namespace %s", namespace))
		return
	}

	listTypeList, err := service.queryAllListType(namespace, codeList)
	if err != nil {
		logrus.Errorf("query all list type have an err: %v", err)
		return
	}

	if len(listTypeList) <= 0 {
		return
	}

	fields := make(map[string]struct{})
	var fieldList = make([]string, 0)
	for i := range listTypeList {
		for j := range listTypeList[i].Fields {
			field := listTypeList[i].Fields[j]
			if _, exist := fields[field]; !exist {
				fields[field] = struct{}{}
				fieldList = append(fieldList, field)
			}
		}
	}
	results := gjson.GetMany(req.Data, fieldList...)
	allMultiValue := make(map[string]string, len(results))
	for i := range results {
		value := results[i].Str
		field := fieldList[i]
		allMultiValue[field] = value
	}
	wg := &sync.WaitGroup{}
	itemHitInfoList := make([]*itemHitInfo, len(listTypeList))
	for i := range listTypeList {
		info := &itemHitInfo{
			wg:            wg,
			listType:      listTypeList[i],
			allMultiValue: allMultiValue,
		}
		itemHitInfoList[i] = info
		wg.Add(1)
		go service.itemHit(info)
	}
	wg.Wait()
	for i := range itemHitInfoList {
		info := itemHitInfoList[i]
		if info.listItem != nil {
			resp.HitListItemList = append(resp.HitListItemList, info.listItem)
		}
	}
	return
}
