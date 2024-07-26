package mysql

import (
	"log"
	"todoList/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlCardRepository struct {
	conn *gorm.DB
}

func NewMysqlCardRepository(conn *gorm.DB) entities.CardRepository {
	return &mysqlCardRepository{conn}
}

func (m *mysqlCardRepository) GetAllCards() (allCardContent []entities.Card, err error) {
	cards := make([]entities.Card, 0)
	err = m.conn.Order("location").Find(&cards).Error
	if err != nil {
		return
	}
	for _, card := range cards {
		tasks := make([]entities.Task, 0)
		m.conn.Where("card_id = ? ", card.ID).Find(&tasks)
		cardContent := entities.Card{
			ID:       card.ID,
			Name:     card.Name,
			Location: card.Location,
			Tasks:    tasks,
		}
		allCardContent = append(allCardContent, cardContent)
	}

	return
}

func (m *mysqlCardRepository) GetSpecificCardIDTask(cardID int) (tasks []entities.Task, err error) {
	tasks = make([]entities.Task, 0)
	err = m.conn.Where("card_id=?", cardID).Find(&tasks).Error
	if err != nil {
		return
	}

	return
}

func (m *mysqlCardRepository) DeleteCardAndTask(toDeleteCard entities.Card, task entities.Task) (err error) {
	// transaction
	tx := m.conn.Begin()
	err1 := tx.Table("tasks").Where("card_id=?", toDeleteCard.ID).Delete(&task).Error
	err2 := tx.Table("cards").Where("ID=?", toDeleteCard.ID).Delete(&toDeleteCard).Error
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")

	}

	return
}

func (m *mysqlCardRepository) CheckCardExist(card entities.Card) (int64, error) {
	result := m.conn.Where("id=?", card.ID).Find(&card)
	return result.RowsAffected, result.Error
}

func (m *mysqlCardRepository) CheckTaskExist(task entities.Task) (reTask entities.Task, err error) {
	err = m.conn.Table("tasks").Where("id=?", task.ID).Find(&reTask).Error
	if err != nil {
		return
	}

	return
}

func (m *mysqlCardRepository) CreateCard(card entities.Card) error {
	err := m.conn.Create(&card).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlCardRepository) CreateTask(task entities.Task) error {
	err := m.conn.Create(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlCardRepository) UpdateTask(task entities.Task, taskID int) error {
	err := m.conn.Model(task).Where("id=?", taskID).Updates(map[string]interface{}{"status": task.Status, "content": task.Content}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlCardRepository) UpdateCard(card entities.Card, cardID int) error {
	err := m.conn.Model(card).Where("id=?", cardID).Updates(map[string]interface{}{"name": card.Name}).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *mysqlCardRepository) DeleteTask(task entities.Task) error {
	err := m.conn.Where("id=?", task.ID).Delete(task).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlCardRepository) UpdateCardLocation(domainCard []map[string]interface{}) error {
	err := m.conn.Table("cards").Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"location",
		}),
	}).Create(&domainCard).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *mysqlCardRepository) UpdateTaskLocation(domainTask []map[string]interface{}) error {
	err := m.conn.Table("tasks").Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"location",
		}),
	}).Create(&domainTask).Error
	if err != nil {
		return err
	}
	return nil
}
