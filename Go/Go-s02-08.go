// package main

// import "fmt"

// func main() {
// 	// In controllers/user.go
// 	if err := db.CreateUser(user); err != nil {
// 		return fmt.Errorf("could not create user: %w", err)
// 	}
// }

// // In database/user.go
// func (db *Database) CreateUser(user *User) error {
// 	ok, err := db.DoesUserExists(user)
// 	if err != nil {
// 		return fmt.Errorf("could not check if user already exists in db: %w", err)
// 	}
// 	// ...
// }

// func (db *Database) DoesUserExists(user *User) error {
// 	if err := db.Connected(); err != nil {
// 		return fmt.Errorf("could not establish db connection: %w",
// 			err)
// 	}
// 	// ...
// }

// func (db *Database) Connected() error {
// 	if !hasInternetConnection() {
// 		return errors.New("no internet connection")
// 	}
// 	// ...
// }
