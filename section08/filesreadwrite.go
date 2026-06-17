package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Todo represents a single todo item stored as one JSON object per line.
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var fileName = filepath.Join("section08", "todos.json")

func writeToFile(todo Todo, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(todo)
}

func readFromFile(fileName string) ([]Todo, error) {
	file, err := os.Open(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return []Todo{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var todos []Todo
	decoder := json.NewDecoder(file)
	for {
		var todo Todo
		err := decoder.Decode(&todo)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func updateTodoInFile(updatedTodo Todo, fileName string) error {
	existingTodos, err := readFromFile(fileName)
	if err != nil {
		return err
	}

	for i, todo := range existingTodos {
		if todo.ID == updatedTodo.ID {
			existingTodos[i] = updatedTodo
			return rewriteTodos(fileName, existingTodos)
		}
	}

	return fmt.Errorf("todo with ID %d not found", updatedTodo.ID)
}

func deleteTodoFromFile(todoID int, fileName string) error {
	existingTodos, err := readFromFile(fileName)
	if err != nil {
		return err
	}

	updatedTodos := make([]Todo, 0, len(existingTodos))
	found := false
	for _, todo := range existingTodos {
		if todo.ID == todoID {
			found = true
			continue
		}
		updatedTodos = append(updatedTodos, todo)
	}

	if !found {
		return fmt.Errorf("todo with ID %d not found", todoID)
	}

	return rewriteTodos(fileName, updatedTodos)
}

func rewriteTodos(fileName string, todos []Todo) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, todo := range todos {
		if err := encoder.Encode(todo); err != nil {
			return err
		}
	}

	return nil
}

func nextTodoID(todos []Todo) int {
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}

func prompt(reader *bufio.Reader, message string) (string, error) {
	fmt.Print(message)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value), nil
}

func promptInt(reader *bufio.Reader, message string) (int, error) {
	value, err := prompt(reader, message)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

func main() {
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 8: Reading and Writing from a File in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Todo App Menu:")
		fmt.Println("1. Add Todo")
		fmt.Println("2. View Todos")
		fmt.Println("3. Edit Todo")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Println()

		choice, err := promptInt(reader, "Enter your choice: ")
		if err != nil {
			fmt.Println("Invalid input. Please enter a number from 1 to 5.")
			continue
		}

		switch choice {
		case 1:
			title, err := prompt(reader, "Enter Todo Title: ")
			if err != nil {
				fmt.Println("Error reading title:", err)
				continue
			}

			content, err := prompt(reader, "Enter Todo Content: ")
			if err != nil {
				fmt.Println("Error reading content:", err)
				continue
			}

			existingTodos, err := readFromFile(fileName)
			if err != nil {
				fmt.Println("Error reading todos:", err)
				continue
			}

			todo := Todo{ID: nextTodoID(existingTodos), Title: title, Content: content}
			if err := writeToFile(todo, fileName); err != nil {
				fmt.Println("Error adding todo:", err)
			} else {
				fmt.Println("Todo added successfully!")
			}

		case 2:
			todos, err := readFromFile(fileName)
			if err != nil {
				fmt.Println("Error reading todos:", err)
			} else if len(todos) == 0 {
				fmt.Println("No todos found.")
			} else {
				fmt.Println("Todos:")
				for _, todo := range todos {
					fmt.Printf("ID: %d, Title: %s, Content: %s\n", todo.ID, todo.Title, todo.Content)
				}
			}

		case 3:
			id, err := promptInt(reader, "Enter Todo ID to edit: ")
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}
			title, err := prompt(reader, "Enter Updated Todo Title: ")
			if err != nil {
				fmt.Println("Error reading title:", err)
				continue
			}
			content, err := prompt(reader, "Enter Updated Todo Content: ")
			if err != nil {
				fmt.Println("Error reading content:", err)
				continue
			}

			updatedTodo := Todo{ID: id, Title: title, Content: content}
			if err := updateTodoInFile(updatedTodo, fileName); err != nil {
				fmt.Println("Error editing todo:", err)
			} else {
				fmt.Println("Todo edited successfully!")
			}

		case 4:
			id, err := promptInt(reader, "Enter Todo ID to delete: ")
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}

			if err := deleteTodoFromFile(id, fileName); err != nil {
				fmt.Println("Error deleting todo:", err)
			} else {
				fmt.Println("Todo deleted successfully!")
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number from 1 to 5.")
		}
	}
}
