package gitlab

import "strings"

const (
	PipelineStatusFailed = "failed"
)

// has_failed_jobs
func (p *ContextPipeline) HasFailedJobs() bool {
	if p == nil {
		return false
	}

	for _, job := range p.Jobs {
		status := strings.ToLower(*job.Status)
		if status == PipelineStatusFailed {
			return true
		}
	}

	return false
}
