package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/tracing"
)

// Tracing executes a cdproto command under Tracing domain.
type Tracing struct {
	ctxWithExecutor context.Context
}

// End stop trace events collection.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#method-end
func (doTracing Tracing) End() (err error) {
	b := tracing.End()
	return b.Do(doTracing.ctxWithExecutor)
}

// GetCategories gets supported tracing categories.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#method-getCategories
//
// returns:
//  - `retCategories`: A list of supported tracing categories.
func (doTracing Tracing) GetCategories() (retCategories []string, err error) {
	b := tracing.GetCategories()
	return b.Do(doTracing.ctxWithExecutor)
}

// RecordClockSyncMarker record a clock sync marker in the trace.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#method-recordClockSyncMarker
//
// parameters:
//  - `syncID`: The ID of this clock sync marker
func (doTracing Tracing) RecordClockSyncMarker(syncID string) (err error) {
	b := tracing.RecordClockSyncMarker(syncID)
	return b.Do(doTracing.ctxWithExecutor)
}

// RequestMemoryDump request a global memory dump.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#method-requestMemoryDump
//
// parameters:
//  - `deterministic`: This can be nil. (Optional) Enables more deterministic results by forcing garbage collection
//
// returns:
//  - `retDumpGUID`: GUID of the resulting global memory dump.
//  - `retSuccess`: True iff the global memory dump succeeded.
func (doTracing Tracing) RequestMemoryDump(deterministic *bool) (retDumpGUID string, retSuccess bool, err error) {
	b := tracing.RequestMemoryDump()
	if deterministic != nil {
		b = b.WithDeterministic(*deterministic)
	}
	return b.Do(doTracing.ctxWithExecutor)
}

// Start start trace events collection.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#method-start
//
// parameters:
//  - `bufferUsageReportingInterval`: This can be nil. (Optional) If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
//  - `transferMode`: This can be nil. (Optional) Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to ReportEvents).
//  - `streamFormat`: This can be nil. (Optional) Trace data format to use. This only applies when using ReturnAsStream transfer mode (defaults to json).
//  - `streamCompression`: This can be nil. (Optional) Compression format to use. This only applies when using ReturnAsStream transfer mode (defaults to none)
//  - `traceConfig`
func (doTracing Tracing) Start(bufferUsageReportingInterval *float64, transferMode *tracing.TransferMode, streamFormat *tracing.StreamFormat, streamCompression *tracing.StreamCompression, traceConfig *tracing.TraceConfig) (err error) {
	b := tracing.Start()
	if bufferUsageReportingInterval != nil {
		b = b.WithBufferUsageReportingInterval(*bufferUsageReportingInterval)
	}
	if transferMode != nil {
		b = b.WithTransferMode(*transferMode)
	}
	if streamFormat != nil {
		b = b.WithStreamFormat(*streamFormat)
	}
	if streamCompression != nil {
		b = b.WithStreamCompression(*streamCompression)
	}
	if traceConfig != nil {
		b = b.WithTraceConfig(traceConfig)
	}
	return b.Do(doTracing.ctxWithExecutor)
}
