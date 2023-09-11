package query

import "gorm.io/gorm"

func paginate(p PaginationInput) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.Page
		if page <= 0 {
			page = 1
		}
		pageSize := p.PageSize
		switch {
		case pageSize > 30:
			pageSize = 30
		case pageSize <= 0:
			pageSize = 30
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
