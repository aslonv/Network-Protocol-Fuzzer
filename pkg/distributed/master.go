package distributed

import (
    "log"
    "sync"
)

type Master struct {
    workers []*Worker
    tasks   chan Task
    wg      sync.WaitGroup
}

func NewMaster(workerCount int) *Master {
    m := &Master{
        workers: make([]*Worker, workerCount),
        tasks:   make(chan Task, workerCount),
    }

    for i := 0; i < workerCount; i++ {
        m.workers[i] = NewWorker(i, m.tasks, &m.wg)
    }

    return m
}

func (m *Master) Start() {
    log.Println("Starting master node...")

    for _, worker := range m.workers {
        m.wg.Add(1)
        go worker.Start()
    }

    // Example task generation
    for i := 0; i < 1000; i++ {
        m.tasks <- Task{ID: i, Data: []byte{byte(i)}}
    }

    close(m.tasks) // Close the channel when tasks are done
    m.wg.Wait()    // Wait for all workers to finish
    log.Println("All tasks completed.")
}
