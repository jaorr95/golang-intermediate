package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name string
	Delay time.Duration
	Number int
}

type Worker struct {
	Id int
	JobQueue chan Job
	WorkerPool chan chan Job // channel in channel pattern
	QuitChan chan bool

}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id: id,
		JobQueue: make(chan Job),
		WorkerPool: workerPool,
		QuitChan: make(chan bool),
	}
}

func (w Worker) Start() {
  go func() {
  	for {

  		// Envio el canal (no el valor) por el worker pool
  		// se bloquea hasta que sea leido en el dispatch
  		// se bloquea porque workerPool es unbuffered channel
  		w.WorkerPool <- w.JobQueue // unbuffered channel, se bloquea
		select {
  		case job := <-w.JobQueue:
  			fmt.Printf("Worker con id %d ha iniciado\n", w.Id)
  			fib := Fibonacci(job.Number)
  			time.Sleep(job.Delay)
  			fmt.Printf("Worker con id %d ha finalizado con resultado %d", w.Id, fib)

		case <-w.QuitChan:
			fmt.Printf("Worker con id %d ha sido detenido", w.Id)
		}
	}
  }()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job // channel in channel pattern
	MaxWorkers int
	JobQueue chan Job
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher{
	worker := make(chan chan Job)
	return &Dispatcher{
		WorkerPool: worker,
		MaxWorkers: maxWorkers,
		JobQueue: jobQueue,
	}
}


func (d *Dispatcher) Dispatch() {
	for  {
		select {
		// espero a que llegue una peticion
		case job := <- d.JobQueue:
			go func() {
				// recibo el canal del worker atraves del cual se envia la peticion
				workerJobQueue := <- d.WorkerPool
				fmt.Printf("Dispatch workerpool %v\n", workerJobQueue)
				fmt.Printf("Dispatch job del request %v\n", job)
				workerJobQueue <- job
			}()

		}
	}
}

func (d *Dispatcher) Run() {
	for i:=0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}


func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Delay invalido", http.StatusBadRequest)
		log.Printf("Error en validacion: %v\n", err)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Value invalido", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name invalido", http.StatusBadRequest)
		return
	}

	job := Job{Delay: delay, Name: name, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers = 4
		maxQueueSize = 20
		port = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Printf("Iniciado servidor en el puerdo %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
