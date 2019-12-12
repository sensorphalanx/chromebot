package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/tethering"
)

// Tethering executes a cdproto command under Tethering domain.
type Tethering struct {
	ctxWithExecutor context.Context
}

// Bind request browser port binding.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tethering#method-bind
//
// parameters:
//  - `port`: Port number to bind.
func (doTethering Tethering) Bind(port int64) (err error) {
	b := tethering.Bind(port)
	return b.Do(doTethering.ctxWithExecutor)
}

// Unbind request browser port unbinding.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tethering#method-unbind
//
// parameters:
//  - `port`: Port number to unbind.
func (doTethering Tethering) Unbind(port int64) (err error) {
	b := tethering.Unbind(port)
	return b.Do(doTethering.ctxWithExecutor)
}
