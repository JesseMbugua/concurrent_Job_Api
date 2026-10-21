package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type job struct {
	id      string
	jobInfo string
	jobType string
}

type EmailProcessor struct {
}

type ReportProcessor struct {
}

type JobProcessor interface {
	Process(context.Context, job) error
}

var ErrEmptyEmail = errors.New("email information is empty")
var ErrEmptyReport = errors.New("report information is empty")

func (r ReportProcessor) Process(j job) error {
	if j.jobInfo == "" {
		return ErrEmptyReport
	}
	fmt.Printf("Job Id:%s\n report:%s\n", j.id, j.jobInfo)
	return nil
}

func (e EmailProcessor) Process(ctx context.Context, j job) error {
	err := ctx.Err()
	if err != nil {
		return err
	}
	if j.jobInfo == "" {
		return ErrEmptyEmail
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second):
		fmt.Printf("Job Id:%s\n email:%s\n", j.id, j.jobInfo)
		return nil
	}
}

func runJob(ctx context.Context, jp JobProcessor, j job) error {
	err := jp.Process(ctx, j)
	if err != nil {
		return fmt.Errorf("Job %s: %w", j.id, err)
	}
	return nil
}

func main() {
	emailJob := job{id: "001", jobInfo: "1", jobType: "email"}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := runJob(ctx, EmailProcessor{}, emailJob)

	if errors.Is(err, ErrEmptyEmail) {
		fmt.Print("Please provide email information")
		return
	}
	if err != nil {
		fmt.Printf("Job failed:%s", err)
		return
	}
	fmt.Println("Job completed")
}
