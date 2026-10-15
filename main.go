package main

import (
	"errors"
	"fmt"
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
	Process(job) error
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

func (e EmailProcessor) Process(j job) error {
	if j.jobInfo == "" {
		return ErrEmptyEmail
	}
	fmt.Printf("Job Id:%s\n email:%s\n", j.id, j.jobInfo)
	return nil
}

func runJob(jp JobProcessor, j job) error {
	err := jp.Process(j)
	if err != nil {
		return fmt.Errorf("Job %s: %v", j.id, err)
	}
	return nil
}

func main() {
	emailJob := job{id: "001", jobInfo: "1", jobType: "email"}
	err := runJob(EmailProcessor{}, emailJob)
	if errors.Is(err, ErrEmptyEmail) {
		fmt.Print("Please provide email information")
		return
	}
	if err != nil {
		fmt.Print("Job failed:%s", err)
		return
	}
	reportJob := job{id: "002", jobInfo: "", jobType: "report"}
	err = runJob(ReportProcessor{}, reportJob)
	if errors.Is(err, ErrEmptyReport) {
		fmt.Print("Please provide report Information")
		return
	}
	if err != nil {
		fmt.Print("Job failed:%s", err)
		return
	}
	fmt.Println("Job completed")
}
