package entities

type Task struct {
	ID       int    `json:"id"`
	Content  string `json:"content"` // 內容
	Status   bool   `json:"status"`  // 狀態(true:已完成、false:未完成)
	Location int64  `json:"location"`
	CardID   int    `json:"card_id"`
}

type Card struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location int64  `json:"location"` // card位置(1、2、3)排序用
	// Tasks    []Task `json:"tasks"`
}

type CardRepository interface {
	UpdateTask(task Task, taskID int) error
	GetAllCards() (allCards []Card, err error)
	DeleteCardAndTask(card Card, task Task) error
	CreateCard(card Card) error
	CheckCardExist(card Card) (int64, error)
	CreateTask(task Task) error
	DeleteTask(task Task) error
	UpdateCardLocation(domainCard []map[string]interface{}) error
	CheckTaskExist(task Task) (reTask Task, err error)
	UpdateTaskLocation(domainTask []map[string]interface{}) error
	GetSpecificCardIDTask(cardID int) (allTask []Task, err error)
	UpdateCard(card Card, cardID int) error
}

type CardUsecase interface {
	UpdateTask(task Task, taskID int) error
	UpdateCard(card Card, cardID int) error
	GetAllCards() (allCards []Card, err error)
	DeleteCardAndTask(card Card, task Task) error
	DeleteTask(task Task) error
	SortCard(changeCard []Card) error
	SortTask(domainTask []Task) error
	GetSpecificCardIDTask(cardID int) (allTask []Task, err error)
	CreateCard(newCard Card) error
	CreateTask(newTask Task) error
}
