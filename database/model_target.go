package database

const (
	ModelTargetTable = "model_target"
)

type ModelTarget struct {
	ID                int64  `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	FrameWork         string `json:"frame_work" gorm:"column:id"`
	AlgorithmCategory string `json:"algorithm_category" gorm:"column:algorithm_category"`
	Algorithm         string `json:"algorithm" gorm:"column:algorithm"`
	ModelType         string `json:"model_type" gorm:"column:algorithm"`
}

func (m ModelTarget) TableName() string {
	return ModelTargetTable
}

type ModelTargetDao struct {
	database
}

func NewModelTargetDao() *ModelTargetDao {
	return &ModelTargetDao{
		database{
			instance: DB(),
		},
	}
}

type GetListFilter struct {
	framework          string
	algorithm_category string
	algorithm          string
	model_type         string
}

func (m *ModelTargetDao) GetCount(filter *GetListFilter) (int64, error) {
	q := m.instance.Where("model_type = ?", 0)

	var count int64 = 0
	q = q.Count(&count)

	return count, q.Error
	
}