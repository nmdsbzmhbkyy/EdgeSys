package services

import (
	"EdgeSys/apps/job/entity"
	"EdgeSys/pkg/global/model"
	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	JobLogModel interface {
		Insert(data entity.JobLog) *entity.JobLog
		FindListPage(page, pageSize int, data entity.JobLog) (*[]entity.JobLog, int64)
		Delete(infoId []int64)
		DeleteAll()
	}

	JobLogModelImpl struct {
		table string
	}
)

var JobLogModelDao JobLogModel = &JobLogModelImpl{
	table: `job_logs`,
}

func (m *JobLogModelImpl) Insert(data entity.JobLog) *entity.JobLog {
	pkg.Db.Table(m.table).Create(&data)
	return &data
}

func (m *JobLogModelImpl) FindListPage(page, pageSize int, data entity.JobLog) (*[]entity.JobLog, int64) {
	list := make([]entity.JobLog, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}

	// 组织数据访问权限
	model.OrgAuthSet(db, data.RoleId, data.Owner)

	err := db.Count(&total).Error
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error

	biz.ErrIsNil(err, "查询登录分页日志信息失败")
	return &list, total
}

func (m *JobLogModelImpl) Delete(logIds []int64) {
	err := pkg.Db.Table(m.table).Delete(&entity.JobLog{}, "id in (?)", logIds).Error
	biz.ErrIsNil(err, "删除登录日志信息失败")
	return
}

func (m *JobLogModelImpl) DeleteAll() {
	pkg.Db.Exec("DELETE FROM job_logs")
}
