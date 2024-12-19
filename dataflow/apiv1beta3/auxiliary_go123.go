// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

//go:build go1.23

package dataflow

import (
	"iter"

	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	"github.com/googleapis/gax-go/v2/iterator"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *JobIterator) All() iter.Seq2[*dataflowpb.Job, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *JobMessageIterator) All() iter.Seq2[*dataflowpb.JobMessage, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StageSummaryIterator) All() iter.Seq2[*dataflowpb.StageSummary, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *WorkerDetailsIterator) All() iter.Seq2[*dataflowpb.WorkerDetails, error] {
	return iterator.RangeAdapter(it.Next)
}
