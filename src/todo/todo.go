package todo

import (
	"log"
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
}

func (s *Server) ListTasks(ctx context.Context, in *Empty) (*TasksListMessage, error) {
	log.Println("List tasks.")
	tasksListMessage := TasksListMessage{}
	for _,task := range s.tasks {
		tasksListMessage.List = append(tasksListMessage.List,
			&TaskMessage{
				Title:  task.title,
				Description: task.description,
				CreatedAt: task.createdAt,
				Deadline: task.deadline,
			},
		)
	}
	return &tasksListMessage, nil
}

func (s *Server) AddTask(ctx context.Context, in *TaskMessage) (*ReplyMessage, error) {
	s.tasks = append(s.tasks, &Task{
		title: in.Title,
		description: in.Description,
		createdAt: in.CreatedAt,
		deadline: in.Deadline,
	})
	log.Println("Task added.")
	return &ReplyMessage{Text: "OK."}, nil
}
