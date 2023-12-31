package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Todo struct represents a single todo item
type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

const fileName = "section08\\todos.json"

// Function to write a Todo to a file
func writeToFile(todo Todo, fileName string) error {
	// Check if the file exists, create it if not
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}

	// Open the file for writing
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode and write the Todo to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(todo)
	if err != nil {
		return err
	}

	return nil
}

// Function to read all Todos from a file
func readFromFile(fileName string) ([]Todo, error) {
	var todos []Todo

	// Check if the file exists, return an empty list if not
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return todos, nil
	} else if err != nil {
		return nil, err
	}

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read each line from the file and decode into Todo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var todo Todo
		err := json.Unmarshal([]byte(scanner.Text()), &todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Function to update a Todo in the file
func updateTodoInFile(updatedTodo Todo, fileName string) error {
	// Read existing todos
	existingTodos, err := readFromFile(fileName)
	if err != nil {
		return err
	}

	// Find and replace the todo with the updated one
	var found bool
	for i, todo := range existingTodos {
		if todo.ID == updatedTodo.ID {
			existingTodos[i] = updatedTodo
			found = true
			break
		}
	}

	// If todo not found, return an error
	if !found {
		return fmt.Errorf("Todo with ID %d not found", updatedTodo.ID)
	}

	// Write the updated todos back to the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, todo := range existingTodos {
		encoder := json.NewEncoder(file)
		err := encoder.Encode(todo)
		if err != nil {
			return err
		}
	}

	return nil
}

// Function to delete a Todo from the file
func deleteTodoFromFile(todoID int, fileName string) error {
	// Read existing todos
	existingTodos, err := readFromFile(fileName)
	if err != nil {
		return err
	}

	// Find and remove the todo with the specified ID
	var found bool
	var updatedTodos []Todo
	for _, todo := range existingTodos {
		if todo.ID == todoID {
			found = true
		} else {
			updatedTodos = append(updatedTodos, todo)
		}
	}

	// If todo not found, return an error
	if !found {
		return fmt.Errorf("Todo with ID %d not found", todoID)
	}

	// Write the updated todos back to the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, todo := range updatedTodos {
		encoder := json.NewEncoder(file)
		err := encoder.Encode(todo)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	// SECTION 8: Reading and Writing from a File in Go
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 8: Reading and Writing from a File in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()
	for {
		fmt.Println("Todo App Menu:")
		fmt.Println("1. Add Todo")
		fmt.Println("2. View Todos")
		fmt.Println("3. Edit Todo")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Println()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title, content string
			fmt.Print("Enter Todo Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Todo Content: ")
			fmt.Scan(&content)

			// Read existing todos
			existingTodos, err := readFromFile(fileName)
			if err != nil {
				fmt.Println("Error reading Todos:", err)
				break
			}

			// Increment ID for simplicity, in a real application, use a unique ID generator
			todo := Todo{ID: len(existingTodos) + 1, Title: title, Content: content}
			err = writeToFile(todo, fileName)
			if err != nil {
				fmt.Println("Error adding Todo:", err)
			} else {
				fmt.Println("Todo added successfully!")
			}

		case 2:
			todos, err := readFromFile(fileName)
			if err != nil {
				fmt.Println("Error reading Todos:", err)
			} else if len(todos) == 0 {
				fmt.Println("No Todos found.")
			} else {
				fmt.Println("Todos:")
				for _, todo := range todos {
					fmt.Printf("ID: %d, Title: %s, Content: %s\n", todo.ID, todo.Title, todo.Content)
				}
			}

		case 3:
			var updatedTodo Todo
			fmt.Print("Enter Todo ID to edit: ")
			fmt.Scan(&updatedTodo.ID)
			fmt.Print("Enter Updated Todo Title: ")
			fmt.Scan(&updatedTodo.Title)
			fmt.Print("Enter Updated Todo Content: ")
			fmt.Scan(&updatedTodo.Content)

			err := updateTodoInFile(updatedTodo, fileName)
			if err != nil {
				fmt.Println("Error editing Todo:", err)
			} else {
				fmt.Println("Todo edited successfully!")
			}
		case 4:
			var todoID int
			fmt.Print("Enter Todo ID to delete: ")
			fmt.Scan(&todoID)

			err := deleteTodoFromFile(todoID, fileName)
			if err != nil {
				fmt.Println("Error deleting Todo:", err)
			} else {
				fmt.Println("Todo deleted successfully!")
			}
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
