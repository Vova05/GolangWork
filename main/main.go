package main

import (
	"Projects/awesomeProject/GolangWork/main/taskstore"
	"Projects/awesomeProject/GolangWork/gin"
	"log"
	"net/http"
	"path/filepath"
	_ "time"
)

//Rest server
type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func main() {

	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Rest server
	router := gin.Default()

	server := NewTaskServer()
	router.POST("/tag/", server.createTaskHandler)
	//mux.HandleFunc("/task/", server.taskHandler)
	mux.HandleFunc("/tag/", server.createTaskHandler)
	mux.HandleFunc("/getTask/", server.getTaskHandler)


	fileServer := http.FileServer(neuteredFileSystem{http.Dir("GolangWork/ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}