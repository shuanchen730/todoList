package http

import (
	"net/http"
	"todoList/ecode"
	"todoList/entities"
	"todoList/entities/delivery"
	"todoList/utils"

	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	CardUsecase entities.CardUsecase
}

func NewCardHandler(apiServer *gin.Engine, cardUsecase entities.CardUsecase) {
	handler := &CardHandler{
		CardUsecase: cardUsecase,
	}

	apiServer.GET("/card", handler.GetAllCards)          // 查詢card
	apiServer.DELETE("/card", handler.DeleteCardAndTask) // 刪除card&task
	apiServer.DELETE("/task", handler.DeleteTask)        // 刪除task
	apiServer.PUT("/task", handler.UpdateTask)           // 更新task
	apiServer.PUT("/card", handler.UpdateCard)           // 更新card
	apiServer.PUT("/card/location", handler.SortCard)    // 排序card
	apiServer.PUT("/task/location", handler.SortTask)    // 排序task
	apiServer.POST("/card", handler.CreateCard)          // 新增card
	apiServer.POST("/task", handler.CreateTask)          // 新增task
}

// GetAllCards 查詢所有card資料，含tasks
func (ca *CardHandler) GetAllCards(c *gin.Context) {
	allCardContent, err := ca.CardUsecase.GetAllCards()
	if err != nil {
		eCode := ecode.Cause(ecode.ErrGetTodoList)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, allCardContent)
}

// DeleteCardAndTask 刪除card&tasks
func (ca *CardHandler) DeleteCardAndTask(c *gin.Context) {
	var deleteCard entities.Card
	err := c.ShouldBindJSON(&deleteCard)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	toDeleteCard := entities.Card{
		ID: deleteCard.ID,
	}
	toDeleteTask := entities.Task{
		CardID: deleteCard.ID,
	}
	err = ca.CardUsecase.DeleteCardAndTask(toDeleteCard, toDeleteTask)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrDeleteCard)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}

	allTodo, err := ca.CardUsecase.GetAllCards()
	if err != nil {
		eCode := ecode.Cause(ecode.ErrGetTodoList)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	err = ca.CardUsecase.SortCard(allTodo)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrReOrderCardLocation)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}

	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}

// UpdateTask 更新task(內容、狀態)
func (ca *CardHandler) UpdateTask(c *gin.Context) {
	var updateTask entities.Task
	err := c.ShouldBind(&updateTask)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	task := entities.Task{
		ID:      updateTask.ID,
		Content: updateTask.Content,
		Status:  updateTask.Status,
	}
	err = ca.CardUsecase.UpdateTask(task, updateTask.ID)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrUpdateTask)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}

// UpdateCard 更新card
func (ca *CardHandler) UpdateCard(c *gin.Context) {
	var updateCard entities.Card
	err := c.ShouldBind(&updateCard)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	card := entities.Card{
		ID:   updateCard.ID,
		Name: updateCard.Name,
	}
	err = ca.CardUsecase.UpdateCard(card, updateCard.ID)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrUpdateCard)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}

// DeleteTask 單純刪除task
func (ca *CardHandler) DeleteTask(c *gin.Context) {
	var deleteTask entities.Task
	err := c.ShouldBindJSON(&deleteTask)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	task := entities.Task{
		ID: deleteTask.ID,
	}
	err = ca.CardUsecase.DeleteTask(task)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrDeleteTask)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}

	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}

// SortCard 排序card
func (ca *CardHandler) SortCard(c *gin.Context) {
	var changeCard []entities.Card
	if err := c.ShouldBindJSON(&changeCard); err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	err := ca.CardUsecase.SortCard(changeCard)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrReOrderCardLocation)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}

// SortTask 排序Task
func (ca *CardHandler) SortTask(c *gin.Context) {
	var changeTask []entities.Task
	if err := c.ShouldBindJSON(&changeTask); err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	err := ca.CardUsecase.SortTask(changeTask)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrReOrderTaskLocation)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}

// CreateCard 新增card
func (ca *CardHandler) CreateCard(c *gin.Context) {
	var newCard entities.Card
	if err := c.ShouldBindJSON(&newCard); err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	newCard = entities.Card{
		Name: newCard.Name,
	}
	err := ca.CardUsecase.CreateCard(newCard)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrCreateCard)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}

	allTodo, err := ca.CardUsecase.GetAllCards()
	if err != nil {
		eCode := ecode.Cause(ecode.ErrGetTodoList)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	err = ca.CardUsecase.SortCard(allTodo)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrReOrderCardLocation)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)

}

// CreateTask 新增task
func (ca *CardHandler) CreateTask(c *gin.Context) {
	var newTask entities.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		eCode := ecode.Cause(ecode.ErrInvalidParameter)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	newTask = entities.Task{
		Content: newTask.Content,
		CardID:  newTask.CardID,
	}
	err := ca.CardUsecase.CreateTask(newTask)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrCreateTask)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	allTasks, err := ca.CardUsecase.GetSpecificCardIDTask(newTask.CardID)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrGetTasks)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	err = ca.CardUsecase.SortTask(allTasks)
	if err != nil {
		eCode := ecode.Cause(ecode.ErrReOrderTaskLocation)
		c.JSON(http.StatusOK, utils.MakeECodeResponse(eCode, err.Error()))
		return
	}
	var responseData delivery.HttpResponse
	responseData.Result = 1
	c.JSON(http.StatusOK, responseData)
}
