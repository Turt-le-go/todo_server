package todo

import (
	"log"
	"todo_server/src/db"
	"todo_server/src/utils"

	"golang.org/x/net/context"
)

type Task struct {
	title string
	description string
	createdAt int64
	deadline int64
}

type Server struct {
	tasks []*Task
	DBConn db.Connection
}


func (s *Server) ListTasks(ctx context.Context, in *Empty) (*TasksListMessage, error) {
	log.Println("List tasks.")
	db := s.DBConn.Open()
	defer db.Close()
	rows, err := db.Query("select createdAt,title,description,deadline from tasks")
	utils.Check(err)
	log.Println("OK")
	defer rows.Close()
	tasksListMessage := TasksListMessage{}
	for rows.Next() {
		task := TaskMessage{}
		err := rows.Scan(&task.CreatedAt, &task.Title, &task.Description, &task.Deadline)
		utils.Check(err)
		tasksListMessage.List = append(tasksListMessage.List, &task)
	}
	return &tasksListMessage, nil
}

func (s *Server) AddTask(ctx context.Context, in *TaskMessage) (*ReplyMessage, error) {
	db := s.DBConn.Open()
	defer db.Close()
	_, err := db.Exec(
		"insert into tasks values (null, $1, $2, $3, $4, 0)",
		in.CreatedAt, in.Title,in.Description, in.Deadline,
	)
	utils.Check(err)
	log.Println("Task added.")
	return &ReplyMessage{Text: "OK."}, nil
}

