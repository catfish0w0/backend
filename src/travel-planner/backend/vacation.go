package backend

import (
	"fmt"
	"travel-planner/model"
)

func (backend *MySQLBackend) GetVacations() ([]model.Vacation, error) {
	var vacations []model.Vacation
	result := backend.db.Table("Vacations").Find(&vacations)
	fmt.Println(vacations, result)
	if result.Error != nil {
		return nil, result.Error
	}
	return vacations, nil
}

func (backend *MySQLBackend) SaveVacation(vacation *model.Vacation) (bool, error) {
	result := backend.db.Table("Vacations").Create(&vacation)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (backend *MySQLBackend) GetSingleVacation(Id uint32) (*model.Vacation, error) {
	var vacation model.Vacation

    result := backend.db.Table("Vacations").Where("id = ?",Id).Find(&vacation)
	fmt.Println(vacation, result)
	if result.Error != nil{
		fmt.Println("Failed to get vacation from db")
		return nil, result.Error
	}
	if result.RowsAffected == 0{
		fmt.Printf("No vacation record in vacation %v\n", Id)
      return nil, nil
	}
	return &vacation, nil
}