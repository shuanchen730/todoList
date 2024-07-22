package mysql

import (
	"reflect"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"

	// "github.com/jinzhu/gorm"
	"testing"
	"todoList/entities"

	"gorm.io/gorm"
)

func ConnDB(t *testing.T) (mockSQL sqlmock.Sqlmock, gormDB *gorm.DB) {
	mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open() 發生錯誤：%v", err)
	}

	return
}

func Test_mysqlCardRepository_CheckTaskExist(t *testing.T) {
	mockSQL, gormDB := ConnDB(t)

	mockTask := entities.Task{
		ID:       888,
		Content:  "8887",
		Status:   false,
		Location: 0,
		CardID:   99,
	}

	rows2 := sqlmock.NewRows([]string{"id", "content", "status", "location", "card_id"}).AddRow(mockTask.ID, mockTask.Content, mockTask.Status, mockTask.Location, mockTask.CardID)

	query2 := "SELECT * FROM `tasks` WHERE id=? "

	mockSQL.ExpectQuery(query2).WithArgs(888).WillReturnRows(rows2)
	repo := NewMysqlCardRepository(gormDB)

	task, err := repo.CheckTaskExist(mockTask)

	if err = mockSQL.ExpectationsWereMet(); err != nil {
		t.Errorf("有預期執行的SQL語法未執行：%v", err)
	}

	if !reflect.DeepEqual(task, mockTask) {
		t.Errorf("got '%+v' but want '%+v'", task, mockTask)
	}
}

func Test_mysqlCardRepository_CreateCard(t *testing.T) {

	// mockSQL, gormDB := ConnDB(t)
	mockDB, mockSQL, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true}),
		&gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open() 發生錯誤：%v", err)
	}

	mockCard := entities.Card{
		ID:       1,
		Name:     "name1",
		Location: 0,
		// Tasks:    nil,
	}

	rows2 := sqlmock.NewRows([]string{"id", "name", "location"}).AddRow(mockCard.ID, mockCard.Name, mockCard.Location)

	query2 := `INSERT INTO "tasks"("id","name",location)VALUES (1,"name1",0)`

	mockSQL.ExpectQuery(query2).WithArgs().WillReturnRows(rows2)

	repo := NewMysqlCardRepository(gormDB)

	err = repo.CreateCard(mockCard)

	if err = mockSQL.ExpectationsWereMet(); err != nil {
		t.Errorf("有預期執行的SQL語法未執行：%v", err)
	}

}

// func Test_mysqlCardRepository_DeleteCardAndTask(t *testing.T) {
//
//   // type fields struct {
//   //   conn *gorm.DB
//   // }
//   //
//   // type args struct {
//   //   toDeleteCard entities.Card
//   //   task         entities.Task
//   // }
//
//   mockDB, mockSQL, err := sqlmock.New()
//
//   conn, err := gorm.Open(mysql.New(mysql.Config{
//     Conn:                      mockDB,
//     SkipInitializeWithVersion: true}),
//     &gorm.Config{})
//   if err != nil {
//     t.Fatalf("gorm.Open() 發生錯誤：%v", err)
//   }
//
//   mockCard := entities.Card{
//     ID:       99,
//     Name:     "work",
//     Location: 0,
//   }
//
//   mockTask := entities.Task{
//     ID:       888,
//     Content:  "8887",
//     Status:   false,
//     Location: 0,
//     CardID:   99,
//   }
//
//   rows := sqlmock.NewRows([]string{"id", "name", "location"}).AddRow(mockCard.ID, mockCard.Name, mockCard.Location)
//   rows2 := sqlmock.NewRows([]string{"id", "content", "status", "location", "card_id"}).AddRow(mockTask.ID, mockTask.Content, mockTask.Status, mockTask.Location, mockTask.CardID)
//
//   query1 := "DELETE FROM `cards` WHERE id = ?"
//   query2 := "DELETE FROM `tasks` WHERE  card_id =? AND tasks.id= ?"
//
//   mockSQL.ExpectBegin()
//   mockSQL.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(99, 888).WillReturnRows(rows2)
//   mockSQL.ExpectQuery(regexp.QuoteMeta(query1)).WithArgs(99, 888).WillReturnRows(rows)
//
//   repo := NewMysqlCardRepository(conn)
//
//   err = repo.DeleteCardAndTask(mockCard, mockTask)
//   if err != nil {
//     mockSQL.ExpectRollback()
//   } else {
//     mockSQL.ExpectCommit()
//
//   }
//
//   if err = mockSQL.ExpectationsWereMet(); err != nil {
//     t.Errorf("有預期執行的SQL語法未執行：%v", err)
//   }
//
//   // tests := []struct {
//   //   name    string
//   //   fields  fields
//   //   args    args
//   //   wantErr bool
//   // }{
//   //   {
//   //     name: "successful",
//   //     fields: fields{
//   //       conn: conn,
//   //     },
//   //     args: args{
//   //       toDeleteCard: entities.Card{
//   //         ID:       1,
//   //         Name:     "card1",
//   //         Location: 0,
//   //         Tasks: []entities.Task{
//   //           {
//   //             ID:       11,
//   //             Content:  "1111",
//   //             Status:   false,
//   //             Location: 1,
//   //             CardID:   1,
//   //           },
//   //           {
//   //             ID:       12,
//   //             Content:  "1212",
//   //             Status:   false,
//   //             Location: 0,
//   //             CardID:   1,
//   //           },
//   //         },
//   //       },
//   //       task: entities.Task{
//   //         ID:       11,
//   //         Content:  "1111",
//   //         Status:   false,
//   //         Location: 0,
//   //         CardID:   1,
//   //       },
//   //     },
//   //     wantErr: false,
//   //   },
//   // }
//   // for _, tt := range tests {
//   //   t.Run(tt.name, func(t *testing.T) {
//   //     m := &mysqlCardRepository{
//   //       conn: tt.fields.conn,
//   //     }
//   //     if err := m.DeleteCardAndTask(tt.args.toDeleteCard, tt.args.task); (err != nil) != tt.wantErr {
//   //       t.Errorf("DeleteCardAndTask() error = %v, wantErr %v", err, tt.wantErr)
//   //     }
//   //   })
//   // }
// }
