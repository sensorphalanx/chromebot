package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/io"
	"github.com/chromedp/cdproto/runtime"
)

// IO executes a cdproto command under IO domain.
type IO struct {
	ctxWithExecutor context.Context
}

// Close close the stream, discard any temporary backing storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IO#method-close
//
// parameters:
//  - `handle`: Handle of the stream to close.
func (doIO IO) Close(handle io.StreamHandle) (err error) {
	b := io.Close(handle)
	return b.Do(doIO.ctxWithExecutor)
}

// Read read a chunk of the stream.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IO#method-read
//
// parameters:
//  - `handle`: Handle of the stream to read.
//  - `offset`: This can be nil. (Optional) Seek to the specified offset before reading (if not specificed, proceed with offset following the last read). Some types of streams may only support sequential reads.
//  - `size`: This can be nil. (Optional) Maximum number of bytes to read (left upon the agent discretion if not specified).
//
// returns:
//  - `retData`: Data that were read.
//  - `retEOF`: Set if the end-of-file condition occurred while reading.
func (doIO IO) Read(handle io.StreamHandle, offset *int64, size *int64) (retData string, retEOF bool, err error) {
	b := io.Read(handle)
	if offset != nil {
		b = b.WithOffset(*offset)
	}
	if size != nil {
		b = b.WithSize(*size)
	}
	return b.Do(doIO.ctxWithExecutor)
}

// ResolveBlob return UUID of Blob object specified by a remote object id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/IO#method-resolveBlob
//
// parameters:
//  - `objectID`: Object id of a Blob object wrapper.
//
// returns:
//  - `retUUID`: UUID of the specified Blob.
func (doIO IO) ResolveBlob(objectID runtime.RemoteObjectID) (retUUID string, err error) {
	b := io.ResolveBlob(objectID)
	return b.Do(doIO.ctxWithExecutor)
}
