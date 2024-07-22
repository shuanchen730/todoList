package usecase

import (
	"errors"
	"todoList/entities"
)

type cardUsecase struct {
	cardRepo entities.CardRepository
}

func NewCardUsecase(cardrepo entities.CardRepository) *cardUsecase {
	return &cardUsecase{
		cardRepo: cardrepo,
	}
}

func (ca *cardUsecase) GetAllCards() (allCardContent []entities.Card, err error) {
	allCardContent, err = ca.cardRepo.GetAllCards()
	if err != nil {
		return
	}
	return
}

func (ca *cardUsecase) GetSpecificCardIDTask(cardID int) (allTasks []entities.Task, err error) {
	allTasks, err = ca.cardRepo.GetSpecificCardIDTask(cardID)
	if err != nil {
		return
	}
	return
}

func (ca *cardUsecase) DeleteCardAndTask(toDeleteCard entities.Card, task entities.Task) error {
	cardExist, err := ca.cardRepo.CheckCardExist(toDeleteCard)
	if cardExist != 1 {
		err = errors.New("查無此卡片")
		return err
	}
	if err != nil {
		return err
	}
	err = ca.cardRepo.DeleteCardAndTask(toDeleteCard, task)
	if err != nil {
		return err
	}

	return nil
}

func (ca *cardUsecase) CreateCard(newCard entities.Card) error {
	err := ca.cardRepo.CreateCard(newCard)
	if err != nil {
		return err
	}
	return nil
}
func (ca *cardUsecase) CreateTask(newTask entities.Task) error {
	err := ca.cardRepo.CreateTask(newTask)
	if err != nil {
		return err
	}
	return nil
}

func (ca *cardUsecase) UpdateTask(task entities.Task, taskID int) error {
	taskExist, err := ca.cardRepo.CheckTaskExist(task)
	if err != nil {
		return err
	}
	if taskExist.Status != true {
		err = errors.New("查無此task")
		return err
	}
	err = ca.cardRepo.UpdateTask(task, taskID)
	if err != nil {
		return err
	}
	return nil
}

func (ca *cardUsecase) UpdateCard(card entities.Card, cardID int) error {
	cardExist, err := ca.cardRepo.CheckCardExist(card)
	if err != nil {
		return err
	}
	if cardExist != 1 {
		err = errors.New("查無此card")
		return err
	}
	err = ca.cardRepo.UpdateCard(card, cardID)
	if err != nil {
		return err
	}
	return nil
}

func (ca *cardUsecase) DeleteTask(task entities.Task) error {
	taskExist, err := ca.cardRepo.CheckTaskExist(task)
	if err != nil {
		return err
	}
	if taskExist.Status != true {
		err = errors.New("查無此task")
		return err
	}
	err = ca.cardRepo.DeleteTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (ca *cardUsecase) SortCard(changeCard []entities.Card) (err error) {
	domainCard := make([]map[string]interface{}, 0)
	cardLen := len(changeCard)
	for i := 0; i < cardLen; i++ {
		changeCard[i].Location = int64(i + 1)
		for i, card := range changeCard {
			cardColum := map[string]interface{}{
				"id":       card.ID,
				"name":     card.Name,
				"location": int64(i + 1),
			}
			domainCard = append(domainCard, cardColum)
		}
		err = ca.cardRepo.UpdateCardLocation(domainCard)
		if err != nil {
			return
		}
	}
	return
}
func (ca *cardUsecase) SortTask(changeTask []entities.Task) (err error) {
	domainTask := make([]map[string]interface{}, 0)
	taskLen := len(changeTask)
	for i := 0; i < taskLen; i++ {
		changeTask[i].Location = int64(i + 1)
		for i, task := range changeTask {
			taskColum := map[string]interface{}{
				"id":       task.ID,
				"content":  task.Content,
				"status":   task.Status,
				"card_id":  task.CardID,
				"location": int64(i + 1),
			}
			domainTask = append(domainTask, taskColum)
		}
		err = ca.cardRepo.UpdateTaskLocation(domainTask)
		if err != nil {
			return
		}
	}
	return
}
