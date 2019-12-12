package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/storage"
)

// Storage executes a cdproto command under Storage domain.
type Storage struct {
	ctxWithExecutor context.Context
}

// ClearDataForOrigin clears storage for origin.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearDataForOrigin
//
// parameters:
//  - `origin`: Security origin.
//  - `storageTypes`: Comma separated list of StorageType to clear.
func (doStorage Storage) ClearDataForOrigin(origin string, storageTypes string) (err error) {
	b := storage.ClearDataForOrigin(origin, storageTypes)
	return b.Do(doStorage.ctxWithExecutor)
}

// GetUsageAndQuota returns usage and quota in bytes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getUsageAndQuota
//
// parameters:
//  - `origin`: Security origin.
//
// returns:
//  - `retUsage`: Storage usage (bytes).
//  - `retQuota`: Storage quota (bytes).
//  - `retUsageBreakdown`: Storage usage per type (bytes).
func (doStorage Storage) GetUsageAndQuota(origin string) (retUsage float64, retQuota float64, retUsageBreakdown []*storage.UsageForType, err error) {
	b := storage.GetUsageAndQuota(origin)
	return b.Do(doStorage.ctxWithExecutor)
}

// TrackCacheStorageForOrigin registers origin to be notified when an update
// occurs to its cache storage list.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackCacheStorageForOrigin
//
// parameters:
//  - `origin`: Security origin.
func (doStorage Storage) TrackCacheStorageForOrigin(origin string) (err error) {
	b := storage.TrackCacheStorageForOrigin(origin)
	return b.Do(doStorage.ctxWithExecutor)
}

// TrackIndexedDBForOrigin registers origin to be notified when an update
// occurs to its IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackIndexedDBForOrigin
//
// parameters:
//  - `origin`: Security origin.
func (doStorage Storage) TrackIndexedDBForOrigin(origin string) (err error) {
	b := storage.TrackIndexedDBForOrigin(origin)
	return b.Do(doStorage.ctxWithExecutor)
}

// UntrackCacheStorageForOrigin unregisters origin from receiving
// notifications for cache storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackCacheStorageForOrigin
//
// parameters:
//  - `origin`: Security origin.
func (doStorage Storage) UntrackCacheStorageForOrigin(origin string) (err error) {
	b := storage.UntrackCacheStorageForOrigin(origin)
	return b.Do(doStorage.ctxWithExecutor)
}

// UntrackIndexedDBForOrigin unregisters origin from receiving notifications
// for IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackIndexedDBForOrigin
//
// parameters:
//  - `origin`: Security origin.
func (doStorage Storage) UntrackIndexedDBForOrigin(origin string) (err error) {
	b := storage.UntrackIndexedDBForOrigin(origin)
	return b.Do(doStorage.ctxWithExecutor)
}
