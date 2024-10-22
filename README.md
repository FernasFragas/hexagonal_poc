# hexagonal_poc

This is a proof of concept for a hexagonal architecture in golang.
Is a simple task manager that allows you to create a new task and store it in a SQLite database.

- In task.go you can find the business logic of the application.
- In the task_repository.go and task_http_handler.go you can find the respective Adapters for the business logic.

Read the comments in the code to understand who's the primary and secondary ports.

![Screenshot](https://github.com/FernasFragas/hexagonal_poc/blob/main/Screenshot.png?raw=true)

For a more detailed explanation you can read this [article](https://medium.com/@patronfragas/hexagonal-architecture-walk-the-talk-a1c036de6f8d).