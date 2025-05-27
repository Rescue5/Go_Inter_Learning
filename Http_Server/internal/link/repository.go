package link

import (
	"HttpServer/db"
)

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{Database: database}
}

func (repo *Repository) Create(link *Link) (*Link, error) {
	result := repo.Database.Db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *Repository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.Db.First(&link, "hash=?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *Repository) Update(link *Link) (*Link, error) {
	result := repo.Database.Db.Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *Repository) Delete(id uint) error {
	check := repo.Database.Db.First(&Link{}, id)
	if check.Error != nil {
		return check.Error
	}
	result := repo.Database.Db.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
