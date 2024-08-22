package distributed

import (
    "log"
)

type Worker struct {
    id    int
    tasks chan Task
    wg    *sync.WaitGroup
}

type Task struct {
    ID   int
    Data []byte
}

func NewWorker(id int, tasks chan Task, wg *sync.WaitGroup) *Worker {
    return &Worker{
        id:    id,
        tasks: tasks,
        wg:    wg,
    }
}

func (w *Worker) Start() {
    log.Printf("Worker %d started.", w.id)
    defer w.wg.Done()

    for task
