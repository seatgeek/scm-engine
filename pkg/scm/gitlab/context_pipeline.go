package gitlab

const (
	PipelineStatusFailed = "failed"
)

// HasFailedJobs returns true if the pipeline has any jobs with status "failed"
func (p *ContextPipeline) HasFailedJobs() bool {
	if p == nil {
		return false
	}

	for _, job := range p.Jobs {
		if job.Status != nil && *job.Status == PipelineStatusFailed {
			return true
		}
	}

	return false
}
