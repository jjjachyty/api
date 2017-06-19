package loanService

import (
	"pccqcpa.com.cn/app/rpm/api/models/biz/ln"
)

var sid ln.SceneItd

type SceneItdService struct {
}

// Find func SceneItdService 多参数查询存款派生
func (SceneItdService) Find(param ...map[string]interface{}) ([]*ln.SceneItd, error) {
	return sid.Find(param...)
}

// AddSceneDp func SceneItdService 新增存款派生
func (SceneItdService) Add(sid *ln.SceneItd) error {
	// 判断业务定价业务
	return sid.Add()
}

// Update func SceneItdService 更新存款派生
func (SceneItdService) Update(sid *ln.SceneItd) error {
	return sid.Update()
}

// Delete func SceneItdService 删除存款派生
func (SceneItdService) Delete(sid *ln.SceneItd) error {
	return sid.Delete()
}

func (SceneItdService) DeleteByBusinessCode(businessCode string) error {
	return sid.DeleteByBusinessCode(businessCode)
}
