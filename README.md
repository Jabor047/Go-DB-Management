# Go Database Management
A project that shows the usage of Go for Database Management. The database used was sqlite.

## Commands

- go run main.go create : This Creates the table users and items
- go run main.go add_user <user_name> <password>: This adds a new user and his desired hashed password to the db
- go run main.go login <user_name> <password>: Checks if the user_name and password given match the records in the db
- go run main.go add_item <giver_id> <item_name> <description> <image_location>: adds an new item to the DB
- go run main.go get_image <image_id>: gets an image from the DB
- go run main.go search_items <search_term>: finds data in the DB that matches the search_term
- go run main.go set_receiver <item_id> <receiver_id>: updates the receiver column in the DB