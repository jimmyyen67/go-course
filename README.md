# Go - The Complete Guide

[Go - The Complete Guide on Udemy](https://www.udemy.com/course/go-the-complete-guide/)

This project stems from my journey through the "Go - The Complete Guide" Udemy course, where I aimed to master the Go programming language.

Throughout this course, I explored:
- **Go Fundamentals**: Learning the basic syntax, variables, types, and control structures.
- **Data Structures**: Working with arrays, slices, maps, and structs.
- **Advanced Features**: Delving into interfaces, error handling, and package management.
- **Concurrent Programming**: Mastering goroutines and channels for efficient concurrency.
- **Real-World Applications**: Building a REST API with user authentication and SQL database integration.
- **Optimization**: Implementing best practices for efficient Go code.

### Course Project: Build a REST API
 - **GET** `/events` Get a list of available events
 - **GET** `/events/{id}` Get a specific event
 - **POST** `/events` Create a new bookable event -> Auth Required
 - **PUT** `/events/{id}` Update an event -> Auth Required & Creator Only
 - **DELETE** `/events/{id}` Delete an event -> Auth Required & Creator Only
 - **POST** `/signup` Create new user
 - **POST** `/login` Authenticate user -> Auth Token (JWT)
 - **POST** `/events/{id}/register` Register user for event -> Auth Required
 - **DELETE** `/events/{id}/register` Cancel registration -> Auth Required