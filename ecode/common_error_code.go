package ecode

var (
	// 00 共用
	ErrInvalidParameter = add("211000000", "帶入參數格式錯誤.")

	// 01 domain - card
	ErrGetTodoList         = add("211003002", "取得Todo清單時發生錯誤.")
	ErrGetTasks            = add("211003003", "取得task時發生錯誤")
	ErrCreateCard          = add("211003004", "新增Card時發生錯誤")
	ErrUpdateCard          = add("211003005", "更新Card時發生錯誤")
	ErrCreateTask          = add("211003006", "新增Task時發生錯誤.")
	ErrUpdateTask          = add("211003007", "更新Task時發生錯誤.")
	ErrDeleteCard          = add("211003008-1", "刪除Card時發生錯誤.")
	ErrDeleteTask          = add("211003008-2", "刪除Task時發生錯誤")
	ErrReOrderCardLocation = add("211003012", "重新排序Card時發生錯誤.")
	ErrReOrderTaskLocation = add("2110030013", "重新排序Task時發生錯誤")
)
