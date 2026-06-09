# Snippetbox

Snippetbox is a web application for storing and sharing code snippets. It provides a simple interface for users to create, view, and manage their code snippets.

## Technologies Used

* Go
* HTML
* CSS
* JavaScript

## Project Structure

```
.
├── go.mod
├── cmd/
│   └── web/
│       ├── config.go
│       ├── handlers.go
│       ├── helpers.go
│       └── main.go
├── internal/
└── ui/
    ├── html/
    │   ├── base.tmpl
    │   ├── pages/
    │   │   └── home.tmpl
    │   └── partials/
    │       └── nav.tmpl.html
    └── static/
        ├── index.html
        ├── css/
        │   ├── index.html
        │   └── main.css
        ├── img/
        │   └── index.html
        └── js/
            ├── index.html
            └── main.js
```

## How to Run

To run the application, navigate to the project root directory in your terminal and execute the following command:

```bash
go run ./cmd/web/ -addr=":8080" -static-dir="./ui/static/"
```

The application will then be accessible in your browser at `http://localhost:8080`.