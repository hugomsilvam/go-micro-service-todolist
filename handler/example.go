package handler

import (
	"context"

	"github.com/micro/go-log"

	todo "github.com/hugomsilvam/todo.list.micro/proto/example"
)

// Todo data struct
type Todo struct {
	id          int32
	title       string
	description string
	completed   bool
}

var todoList = []Todo{}

// Hello is a simple handler function to test
func (t *Todo) Hello(ctx context.Context, req *todo.Message, rsp *todo.Message) error {
	log.Log("Hello")
	log.Log("hello request message ", req.Msg)
	log.Log("t todo ", t)
	rsp.Msg = "Hello " + req.Msg
	return nil
}

// Create a new Todo
func (t *Todo) Create(ctx context.Context, req *todo.CreateRequest, rsp *todo.CreateResponse) error {
	log.Log("CreateTodo")
	var todoObj = Todo{id: req.Todo.Id, title: req.Todo.Title, description: req.Todo.Description, completed: req.Todo.Completed}
	log.Log("todoObj ", todoObj)
	todoList = append(todoList, todoObj)
	log.Log("Created todo ", todoObj.id, " and add to Todolist")
	rsp.Id = todoObj.id
	return nil
}

// Get a Todo by Id
func (t *Todo) Get(ctx context.Context, req *todo.GetRequest, rsp *todo.GetResponse) error {
	log.Log("GetTodo")
	for _, todoObj := range todoList {
		if todoObj.id == req.Id {
			log.Log("get todo", todoObj, "from array")
			rsp.Todo = &todo.Todo{Id: todoObj.id, Title: todoObj.title, Description: todoObj.description, Completed: todoObj.completed}
			break
		}
	}
	log.Log("Todo id ", req.Id, " does not exist!")
	return nil
}

// GetAll Todos from array
func (t *Todo) GetAll(ctx context.Context, req *todo.GetAllRequest, rsp *todo.GetAllResponse) error {
	log.Log("GetAllTodos")
	for _, todoObj := range todoList {
		rsp.Todos = append(rsp.Todos, &todo.Todo{Id: todoObj.id, Title: todoObj.title, Description: todoObj.description, Completed: todoObj.completed})
	}

	return nil
}

// Update a Todo by Id
func (t *Todo) Update(ctx context.Context, req *todo.UpdateRequest, rsp *todo.UpdateResponse) error {
	log.Log("Update Todo")
	var id = req.Todo.Id
	var newTodo = req.Todo
	log.Log("Update todo ", id, " with new todo content ", newTodo)
	for i, todoObj := range todoList {
		if todoObj.id == id {
			todoList[i] = Todo{id: id, title: newTodo.Title, description: newTodo.Description, completed: newTodo.Completed}
			log.Log("Todo updated!")
			rsp.Updated = 1
			break
		}
	}
	return nil
}

// Delete a Todo by Id
func (t *Todo) Delete(ctx context.Context, req *todo.DeleteRequest, rsp *todo.DeleteResponse) error {
	log.Log("Delete Todo")
	for i, todoObj := range todoList {
		if todoObj.id == req.Id {
			log.Log("Prepare for delete todo", todoObj)
			todoList = append(todoList[:i], todoList[i+1:]...)
			log.Log("Deleted todo", req.Id, "!")
			rsp.Deleted = 1
			break
		}
	}
	return nil
}
