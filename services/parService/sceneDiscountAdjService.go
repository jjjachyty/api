package parService

import (
	"fmt"

	"pccqcpa.com.cn/app/rpm/api/models/par"
	"pccqcpa.com.cn/app/rpm/api/util"
)

var sceneDiscountAdjModle par.SceneDiscountAdj

type SceneDiscountAdjService struct{}

func (s SceneDiscountAdjService) List(param ...map[string]interface{}) (*util.PageData, error) {
	return sceneDiscountAdjModle.List(param...)
}

func (d SceneDiscountAdjService) Find(param ...map[string]interface{}) ([]*par.SceneDiscountAdj, error) {
	paramMap := make(map[string]interface{})
	if 0 != len(param) {
		paramMap = param[0]
		if GapProportion, ok := paramMap["gap_proportion"]; ok {
			delete(paramMap, "gap_proportion")
			var whereInSql = `
				select max(gap_proportion)  gap_proportion  from rpm_par_scene_discount_adj t  where gap_proportion <=` + fmt.Sprint(GapProportion)
			paramMap["searchLike"] = []map[string]interface{}{
				map[string]interface{}{
					"type":  "in",
					"key":   "gap_proportion",
					"value": whereInSql,
				},
			}
		}
	}
	return sceneDiscountAdjModle.Find(paramMap)
}

func (s SceneDiscountAdjService) Add(sceneDiscountAdj *par.SceneDiscountAdj) error {
	return sceneDiscountAdj.Add()
}

func (s SceneDiscountAdjService) Update(sceneDiscountAdj *par.SceneDiscountAdj) error {
	return sceneDiscountAdj.Update()
}

func (s SceneDiscountAdjService) Delete(sceneDiscountAdj *par.SceneDiscountAdj) error {
	return sceneDiscountAdj.Delete()
}
