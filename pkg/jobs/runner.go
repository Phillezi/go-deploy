package jobs

import (
	"errors"
	"fmt"
	jobModel "go-deploy/models/sys/job"
	"go-deploy/utils"
	"strings"
)

type Runner struct {
	Job *jobModel.Job
}

func NewRunner(job *jobModel.Job) *Runner {
	return &Runner{Job: job}
}

func (runner *Runner) Run() {
	if jobDef := GetJobDef(runner.Job); jobDef != nil {
		if jobDef.TerminateFunc != nil {
			shouldTerminate, err := jobDef.TerminateFunc(runner.Job)
			if err != nil {
				utils.PrettyPrintError(fmt.Errorf("error executing job (%s) terminate function, terminating the job instead. details: %w", runner.Job.Type, err))

				err = jobModel.New().MarkTerminated(runner.Job.ID, err.Error())
				if err != nil {
					utils.PrettyPrintError(fmt.Errorf("error marking job as terminated. details: %w", err))
					return
				}
				return
			}

			if shouldTerminate {
				err = jobModel.New().MarkTerminated(runner.Job.ID, "gracefully terminated by system")
				utils.PrettyPrintError(fmt.Errorf("job (%s) gracefully terminated by system", runner.Job.Type))
				if err != nil {
					utils.PrettyPrintError(fmt.Errorf("error marking job as terminated. details: %w", err))
					return
				}
				return
			}
		}

		go wrapper(jobDef)
	} else {
		utils.PrettyPrintError(fmt.Errorf("unknown job type: %s", runner.Job.Type))

		err := jobModel.New().MarkTerminated(runner.Job.ID, "unknown job type")
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("error marking unknown job as terminated. details: %w", err))
			return
		}
	}
}

func wrapper(def *JobDefinition) {
	if def.EntryFunc != nil {
		err := def.EntryFunc(def.Job)
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("error executing job (%s) entry function. details: %w", def.Job.Type, err))

			err = jobModel.New().MarkFailed(def.Job.ID, err.Error())
			if err != nil {
				utils.PrettyPrintError(fmt.Errorf("error marking job as failed. details: %w", err))
				return
			}
			return
		}
	}

	defer func() {
		if def.ExitFunc != nil {
			err := def.ExitFunc(def.Job)
			if err != nil {
				utils.PrettyPrintError(fmt.Errorf("error executing job (%s) exit function. details: %w", def.Job.Type, err))

				err = jobModel.New().MarkFailed(def.Job.ID, err.Error())
				if err != nil {
					utils.PrettyPrintError(fmt.Errorf("error marking job as failed. details: %w", err))
					return
				}
				return
			}
		}
	}()

	err := def.JobFunc(def.Job)

	if err != nil {
		if strings.HasPrefix(err.Error(), "failed") {
			err = errors.Unwrap(err)
			utils.PrettyPrintError(fmt.Errorf("failed job (%s). details: %w", def.Job.Type, err))

			err = jobModel.New().MarkFailed(def.Job.ID, err.Error())
			if err != nil {
				utils.PrettyPrintError(fmt.Errorf("error marking job as failed. details: %w", err))
				return
			}
		} else if strings.HasPrefix(err.Error(), "terminated") {
			err = errors.Unwrap(err)
			utils.PrettyPrintError(fmt.Errorf("terminated job (%s). details: %w", def.Job.Type, err))

			err = jobModel.New().MarkTerminated(def.Job.ID, err.Error())
			if err != nil {
				utils.PrettyPrintError(fmt.Errorf("error marking job as terminated. details: %w", err))
				return
			}
		} else {
			utils.PrettyPrintError(fmt.Errorf("error executing job (%s). details: %w", def.Job.Type, err))

			err = jobModel.New().MarkFailed(def.Job.ID, err.Error())
			if err != nil {
				utils.PrettyPrintError(fmt.Errorf("error marking job as failed. details: %w", err))
				return
			}
		}
	} else {
		err = jobModel.New().MarkCompleted(def.Job.ID)
		if err != nil {
			utils.PrettyPrintError(fmt.Errorf("error marking job as completed. details: %w", err))
			return
		}
	}
}
