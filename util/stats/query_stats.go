// Copyright 2013 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stats

// QueryTiming identifies the code area or functionality in which time is spent
// during a query.
type QueryTiming int

// Query timings.
const (
	EvalTotalTime QueryTiming = iota
	ResultSortTime
	QueryPreparationTime
	InnerEvalTime
	ResultAppendTime
	ExecQueueTime
	ExecTotalTime
)

// Return a string representation of a QueryTiming identifier.
func (s QueryTiming) String() string {
	switch s {
	case EvalTotalTime:
		return "Eval total time"
	case ResultSortTime:
		return "Result sorting time"
	case QueryPreparationTime:
		return "Query preparation time"
	case InnerEvalTime:
		return "Inner eval time"
	case ResultAppendTime:
		return "Result append time"
	case ExecQueueTime:
		return "Exec queue wait time"
	case ExecTotalTime:
		return "Exec total time"
	default:
		return "Unknown query timing"
	}
}

// QueryStats with all query timers mapped to durations.
type QueryStats struct {
	EvalTotalTime        float64 `json:"evalTotalTime"`
	ResultSortTime       float64 `json:"resultSortTime"`
	QueryPreparationTime float64 `json:"queryPreparationTime"`
	InnerEvalTime        float64 `json:"innerEvalTime"`
	ResultAppendTime     float64 `json:"resultAppendTime"`
	ExecQueueTime        float64 `json:"execQueueTime"`
	ExecTotalTime        float64 `json:"execTotalTime"`
}

// NewQueryStats makes a QueryStats struct with all QueryTimings found in the
// given TimerGroup.
func NewQueryStats(tg *TimerGroup) *QueryStats {
	var qs QueryStats

	for s, timer := range tg.timers {
		switch s {
		case EvalTotalTime:
			qs.EvalTotalTime = timer.Duration()
		case ResultSortTime:
			qs.ResultSortTime = timer.Duration()
		case QueryPreparationTime:
			qs.QueryPreparationTime = timer.Duration()
		case InnerEvalTime:
			qs.InnerEvalTime = timer.Duration()
		case ResultAppendTime:
			qs.ResultAppendTime = timer.Duration()
		case ExecQueueTime:
			qs.ExecQueueTime = timer.Duration()
		case ExecTotalTime:
			qs.ExecTotalTime = timer.Duration()
		}
	}

	return &qs
}
