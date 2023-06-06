package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type EnvError error

type Task interface {
	doWork() EnvError
	ID() int
}

type MyTask struct {
	id int
}

func (mt *MyTask) doWork() EnvError {
	fmt.Printf("Task %d work begins \n", mt.id)
	if mt.id == 5 {
		time.Sleep(time.Duration(1) * time.Second)
	} else {
		time.Sleep(time.Duration(10) * time.Second)
	}
	if mt.id == 5 {
		fmt.Printf("Task %d failed successfully\n", mt.id)
		return errors.New("Env: error thrown\n")
	}
	fmt.Printf("Task %d work ends\n", mt.id)
	return nil
}

func (mt *MyTask) ID() int {
	return mt.id
}

type Scheduler struct {
	tasks []MyTask
}

func (sc *Scheduler) execute(ctx context.Context) EnvError {

	if ctx == nil {
		ctx = context.Background()
	}
	parent_ctx, parent_cancel := context.WithCancel(ctx)
	// defer parent_cancel()

	n := len(sc.tasks)

	tracker := make(chan EnvError, n)
	// parent_ctx = context.WithValue(parent_ctx, 0, tracker)

	for _, t := range sc.tasks {
		// child_ctx, _ := context.WithCancel(parent_ctx)
		go func(ct context.Context, task MyTask) {

			// go func() {
			// 	ct.Value(0).(chan EnvError) <- task.doWork()
			// }()
			for {
				select {
				case <-ct.Done():
					fmt.Printf("Task %d cancelled abruptly\n", task.ID())
					return
				case tracker <- task.doWork():
					return
				}
			}

		}(parent_ctx, t)
	}

	var err EnvError = nil

	for i := 0; i < n; i++ {
		err = <-tracker
		if err != nil {
			// signal all goroutines to stop
			fmt.Println("Stopping remaining tasks")
			parent_cancel()
			time.Sleep(15 * time.Second)
			break

		}
	}
	close(tracker)
	if err == nil {
		parent_cancel()
		time.Sleep(5 * time.Second)

	}

	return err
}

func main() {

	tasks := make([]MyTask, 6)

	for i := 0; i < 6; i++ {
		tasks[i].id = i
	}

	scd := Scheduler{tasks: tasks}
	err := scd.execute(nil)
	fmt.Println(err)

}
