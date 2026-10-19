package model

type TbBlog struct {
	Id uint64 `gorm:"column:id;type:bigint unsigned;not null;primaryKey;autoIncrement;commont:主键" json:"id"`
}

func (t *TbBlog) TableName() string {
	return "tb_blog"
}
