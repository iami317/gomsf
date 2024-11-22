package rpc

type job struct {
	rpc *RPC
}

type JobInfoReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
	JobID  string `json:"job_id"`
}

type JobInfoRes struct {
	JobID     int                    `msgpack:"jid" json:"job_id"`
	Name      string                 `msgpack:"name" json:"name"`
	StartTime int                    `msgpack:"start_time" json:"start_time"`
	URIPath   interface{}            `msgpack:"uripath,omitempty" json:"uripath,omitempty"`
	Datastore map[string]interface{} `msgpack:"datastore,omitempty" json:"datastore,omitempty"`
}

// Info returns information about a job
func (j *job) Info(jobID string) (*JobInfoRes, error) {
	req := &JobInfoReq{
		Method: "job.info",
		Token:  j.rpc.Token(),
		JobID:  jobID,
	}

	var res *JobInfoRes
	if err := j.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type JobListReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

type JobListRes map[string]string

// List returns a list of jobs
func (j *job) List() (*JobListRes, error) {
	req := &JobListReq{
		Method: "job.list",
		Token:  j.rpc.Token(),
	}

	var res *JobListRes
	if err := j.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type JobStopReq struct {
	Method string `json:"method"`
	Token  string `json:"token"`
	JobID  string `json:"job_id"`
}

type JobStopRes struct {
	Result Result `msgpack:"result"`
}

// Stop stops a job
func (j *job) Stop(jobID string) (*JobStopRes, error) {
	req := &JobStopReq{
		Method: "job.stop",
		Token:  j.rpc.Token(),
		JobID:  jobID,
	}

	var res *JobStopRes
	if err := j.rpc.Call(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
