package domain

// Code generated by chromebot-domain-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/runtime"
)

// DOM executes a cdproto command under DOM domain.
type DOM struct {
	ctxWithExecutor context.Context
}

// CollectClassNamesFromSubtree collects class names for the node with given
// id and all of it's child nodes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-collectClassNamesFromSubtree
//
// parameters:
//  - `nodeID`: Id of the node to collect class names.
//
// returns:
//  - `retClassNames`: Class name list.
func (doDOM DOM) CollectClassNamesFromSubtree(nodeID cdp.NodeID) (retClassNames []string, err error) {
	b := dom.CollectClassNamesFromSubtree(nodeID)
	return b.Do(doDOM.ctxWithExecutor)
}

// CopyTo creates a deep copy of the specified node and places it into the
// target container before the given anchor.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-copyTo
//
// parameters:
//  - `nodeID`: Id of the node to copy.
//  - `targetNodeID`: Id of the element to drop the copy into.
//  - `insertBeforeNodeID`: This can be nil. (Optional) Drop the copy before this node (if absent, the copy becomes the last child of targetNodeId).
//
// returns:
//  - `retNodeID`: Id of the node clone.
func (doDOM DOM) CopyTo(nodeID cdp.NodeID, targetNodeID cdp.NodeID, insertBeforeNodeID *cdp.NodeID) (retNodeID cdp.NodeID, err error) {
	b := dom.CopyTo(nodeID, targetNodeID)
	if insertBeforeNodeID != nil {
		b = b.WithInsertBeforeNodeID(*insertBeforeNodeID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// DescribeNode describes node given its id, does not require domain to be
// enabled. Does not start tracking any objects, can be used for automation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-describeNode
//
// parameters:
//  - `nodeID`: This can be nil. (Optional) Identifier of the node.
//  - `backendNodeID`: This can be nil. (Optional) Identifier of the backend node.
//  - `objectID`: This can be nil. (Optional) JavaScript object id of the node wrapper.
//  - `depth`: This can be nil. (Optional) The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
//  - `pierce`: This can be nil. (Optional) Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
//
// returns:
//  - `retNode`: Node description.
func (doDOM DOM) DescribeNode(nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectID *runtime.RemoteObjectID, depth *int64, pierce *bool) (retNode *cdp.Node, err error) {
	b := dom.DescribeNode()
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectID != nil {
		b = b.WithObjectID(*objectID)
	}
	if depth != nil {
		b = b.WithDepth(*depth)
	}
	if pierce != nil {
		b = b.WithPierce(*pierce)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// Disable disables DOM agent for the given page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-disable
func (doDOM DOM) Disable() (err error) {
	b := dom.Disable()
	return b.Do(doDOM.ctxWithExecutor)
}

// DiscardSearchResults discards search results from the session with the
// given id. getSearchResults should no longer be called for that search.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-discardSearchResults
//
// parameters:
//  - `searchID`: Unique search session identifier.
func (doDOM DOM) DiscardSearchResults(searchID string) (err error) {
	b := dom.DiscardSearchResults(searchID)
	return b.Do(doDOM.ctxWithExecutor)
}

// Enable enables DOM agent for the given page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-enable
func (doDOM DOM) Enable() (err error) {
	b := dom.Enable()
	return b.Do(doDOM.ctxWithExecutor)
}

// Focus focuses the given element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-focus
//
// parameters:
//  - `nodeID`: This can be nil. (Optional) Identifier of the node.
//  - `backendNodeID`: This can be nil. (Optional) Identifier of the backend node.
//  - `objectID`: This can be nil. (Optional) JavaScript object id of the node wrapper.
func (doDOM DOM) Focus(nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectID *runtime.RemoteObjectID) (err error) {
	b := dom.Focus()
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectID != nil {
		b = b.WithObjectID(*objectID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetAttributes returns attributes for the specified node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getAttributes
//
// parameters:
//  - `nodeID`: Id of the node to retrieve attibutes for.
//
// returns:
//  - `retAttributes`: An interleaved array of node attribute names and values.
func (doDOM DOM) GetAttributes(nodeID cdp.NodeID) (retAttributes []string, err error) {
	b := dom.GetAttributes(nodeID)
	return b.Do(doDOM.ctxWithExecutor)
}

// GetBoxModel returns boxes for the given node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getBoxModel
//
// parameters:
//  - `nodeID`: This can be nil. (Optional) Identifier of the node.
//  - `backendNodeID`: This can be nil. (Optional) Identifier of the backend node.
//  - `objectID`: This can be nil. (Optional) JavaScript object id of the node wrapper.
//
// returns:
//  - `retModel`: Box model for the node.
func (doDOM DOM) GetBoxModel(nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectID *runtime.RemoteObjectID) (retModel *dom.BoxModel, err error) {
	b := dom.GetBoxModel()
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectID != nil {
		b = b.WithObjectID(*objectID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetContentQuads returns quads that describe node position on the page.
// This method might return multiple quads for inline nodes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getContentQuads
//
// parameters:
//  - `nodeID`: This can be nil. (Optional) Identifier of the node.
//  - `backendNodeID`: This can be nil. (Optional) Identifier of the backend node.
//  - `objectID`: This can be nil. (Optional) JavaScript object id of the node wrapper.
//
// returns:
//  - `retQuads`: Quads that describe node layout relative to viewport.
func (doDOM DOM) GetContentQuads(nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectID *runtime.RemoteObjectID) (retQuads []dom.Quad, err error) {
	b := dom.GetContentQuads()
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectID != nil {
		b = b.WithObjectID(*objectID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetDocument returns the root DOM node (and optionally the subtree) to the
// caller.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getDocument
//
// parameters:
//  - `depth`: This can be nil. (Optional) The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
//  - `pierce`: This can be nil. (Optional) Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
//
// returns:
//  - `retRoot`: Resulting node.
func (doDOM DOM) GetDocument(depth *int64, pierce *bool) (retRoot *cdp.Node, err error) {
	b := dom.GetDocument()
	if depth != nil {
		b = b.WithDepth(*depth)
	}
	if pierce != nil {
		b = b.WithPierce(*pierce)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetFlattenedDocument returns the root DOM node (and optionally the
// subtree) to the caller.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getFlattenedDocument
//
// parameters:
//  - `depth`: This can be nil. (Optional) The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
//  - `pierce`: This can be nil. (Optional) Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
//
// returns:
//  - `retNodes`: Resulting node.
func (doDOM DOM) GetFlattenedDocument(depth *int64, pierce *bool) (retNodes []*cdp.Node, err error) {
	b := dom.GetFlattenedDocument()
	if depth != nil {
		b = b.WithDepth(*depth)
	}
	if pierce != nil {
		b = b.WithPierce(*pierce)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetNodeForLocation returns node id at given location. Depending on whether
// DOM domain is enabled, nodeId is either returned or not.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getNodeForLocation
//
// parameters:
//  - `x`: X coordinate.
//  - `y`: Y coordinate.
//  - `includeUserAgentShadowDOM`: This can be nil. (Optional) False to skip to the nearest non-UA shadow root ancestor (default: false).
//  - `ignorePointerEventsNone`: This can be nil. (Optional) Whether to ignore pointer-events: none on elements and hit test them.
//
// returns:
//  - `retBackendNodeID`: Resulting node.
//  - `retFrameID`: Frame this node belongs to.
//  - `retNodeID`: Id of the node at given coordinates, only when enabled and requested document.
func (doDOM DOM) GetNodeForLocation(x int64, y int64, includeUserAgentShadowDOM *bool, ignorePointerEventsNone *bool) (retBackendNodeID cdp.BackendNodeID, retFrameID cdp.FrameID, retNodeID cdp.NodeID, err error) {
	b := dom.GetNodeForLocation(x, y)
	if includeUserAgentShadowDOM != nil {
		b = b.WithIncludeUserAgentShadowDOM(*includeUserAgentShadowDOM)
	}
	if ignorePointerEventsNone != nil {
		b = b.WithIgnorePointerEventsNone(*ignorePointerEventsNone)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetOuterHTML returns node's HTML markup.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getOuterHTML
//
// parameters:
//  - `nodeID`: This can be nil. (Optional) Identifier of the node.
//  - `backendNodeID`: This can be nil. (Optional) Identifier of the backend node.
//  - `objectID`: This can be nil. (Optional) JavaScript object id of the node wrapper.
//
// returns:
//  - `retOuterHTML`: Outer HTML markup.
func (doDOM DOM) GetOuterHTML(nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectID *runtime.RemoteObjectID) (retOuterHTML string, err error) {
	b := dom.GetOuterHTML()
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectID != nil {
		b = b.WithObjectID(*objectID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// GetRelayoutBoundary returns the id of the nearest ancestor that is a
// relayout boundary.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getRelayoutBoundary
//
// parameters:
//  - `nodeID`: Id of the node.
//
// returns:
//  - `retNodeID`: Relayout boundary node id for the given node.
func (doDOM DOM) GetRelayoutBoundary(nodeID cdp.NodeID) (retNodeID cdp.NodeID, err error) {
	b := dom.GetRelayoutBoundary(nodeID)
	return b.Do(doDOM.ctxWithExecutor)
}

// GetSearchResults returns search results from given fromIndex to given
// toIndex from the search with the given identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getSearchResults
//
// parameters:
//  - `searchID`: Unique search session identifier.
//  - `fromIndex`: Start index of the search result to be returned.
//  - `toIndex`: End index of the search result to be returned.
//
// returns:
//  - `retNodeIds`: Ids of the search result nodes.
func (doDOM DOM) GetSearchResults(searchID string, fromIndex int64, toIndex int64) (retNodeIds []cdp.NodeID, err error) {
	b := dom.GetSearchResults(searchID, fromIndex, toIndex)
	return b.Do(doDOM.ctxWithExecutor)
}

// MarkUndoableState marks last undoable state.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-markUndoableState
func (doDOM DOM) MarkUndoableState() (err error) {
	b := dom.MarkUndoableState()
	return b.Do(doDOM.ctxWithExecutor)
}

// MoveTo moves node into the new container, places it before the given
// anchor.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-moveTo
//
// parameters:
//  - `nodeID`: Id of the node to move.
//  - `targetNodeID`: Id of the element to drop the moved node into.
//  - `insertBeforeNodeID`: This can be nil. (Optional) Drop node before this one (if absent, the moved node becomes the last child of targetNodeId).
//
// returns:
//  - `retNodeID`: New id of the moved node.
func (doDOM DOM) MoveTo(nodeID cdp.NodeID, targetNodeID cdp.NodeID, insertBeforeNodeID *cdp.NodeID) (retNodeID cdp.NodeID, err error) {
	b := dom.MoveTo(nodeID, targetNodeID)
	if insertBeforeNodeID != nil {
		b = b.WithInsertBeforeNodeID(*insertBeforeNodeID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// PerformSearch searches for a given string in the DOM tree. Use
// getSearchResults to access search results or cancelSearch to end this search
// session.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-performSearch
//
// parameters:
//  - `query`: Plain text or query selector or XPath search query.
//  - `includeUserAgentShadowDOM`: This can be nil. (Optional) True to search in user agent shadow DOM.
//
// returns:
//  - `retSearchID`: Unique search session identifier.
//  - `retResultCount`: Number of search results.
func (doDOM DOM) PerformSearch(query string, includeUserAgentShadowDOM *bool) (retSearchID string, retResultCount int64, err error) {
	b := dom.PerformSearch(query)
	if includeUserAgentShadowDOM != nil {
		b = b.WithIncludeUserAgentShadowDOM(*includeUserAgentShadowDOM)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// PushNodeByPathToFrontend requests that the node is sent to the caller
// given its path. // FIXME, use XPath.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-pushNodeByPathToFrontend
//
// parameters:
//  - `path`: Path to node in the proprietary format.
//
// returns:
//  - `retNodeID`: Id of the node for given path.
func (doDOM DOM) PushNodeByPathToFrontend(path string) (retNodeID cdp.NodeID, err error) {
	b := dom.PushNodeByPathToFrontend(path)
	return b.Do(doDOM.ctxWithExecutor)
}

// PushNodesByBackendIdsToFrontend requests that a batch of nodes is sent to
// the caller given their backend node ids.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-pushNodesByBackendIdsToFrontend
//
// parameters:
//  - `backendNodeIds`: The array of backend node ids.
//
// returns:
//  - `retNodeIds`: The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
func (doDOM DOM) PushNodesByBackendIdsToFrontend(backendNodeIds []cdp.BackendNodeID) (retNodeIds []cdp.NodeID, err error) {
	b := dom.PushNodesByBackendIdsToFrontend(backendNodeIds)
	return b.Do(doDOM.ctxWithExecutor)
}

// QuerySelector executes querySelector on a given node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-querySelector
//
// parameters:
//  - `nodeID`: Id of the node to query upon.
//  - `selector`: Selector string.
//
// returns:
//  - `retNodeID`: Query selector result.
func (doDOM DOM) QuerySelector(nodeID cdp.NodeID, selector string) (retNodeID cdp.NodeID, err error) {
	b := dom.QuerySelector(nodeID, selector)
	return b.Do(doDOM.ctxWithExecutor)
}

// QuerySelectorAll executes querySelectorAll on a given node.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-querySelectorAll
//
// parameters:
//  - `nodeID`: Id of the node to query upon.
//  - `selector`: Selector string.
//
// returns:
//  - `retNodeIds`: Query selector result.
func (doDOM DOM) QuerySelectorAll(nodeID cdp.NodeID, selector string) (retNodeIds []cdp.NodeID, err error) {
	b := dom.QuerySelectorAll(nodeID, selector)
	return b.Do(doDOM.ctxWithExecutor)
}

// Redo re-does the last undone action.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-redo
func (doDOM DOM) Redo() (err error) {
	b := dom.Redo()
	return b.Do(doDOM.ctxWithExecutor)
}

// RemoveAttribute removes attribute with given name from an element with
// given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-removeAttribute
//
// parameters:
//  - `nodeID`: Id of the element to remove attribute from.
//  - `name`: Name of the attribute to remove.
func (doDOM DOM) RemoveAttribute(nodeID cdp.NodeID, name string) (err error) {
	b := dom.RemoveAttribute(nodeID, name)
	return b.Do(doDOM.ctxWithExecutor)
}

// RemoveNode removes node with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-removeNode
//
// parameters:
//  - `nodeID`: Id of the node to remove.
func (doDOM DOM) RemoveNode(nodeID cdp.NodeID) (err error) {
	b := dom.RemoveNode(nodeID)
	return b.Do(doDOM.ctxWithExecutor)
}

// RequestChildNodes requests that children of the node with given id are
// returned to the caller in form of setChildNodes events where not only
// immediate children are retrieved, but all children down to the specified
// depth.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-requestChildNodes
//
// parameters:
//  - `nodeID`: Id of the node to get children for.
//  - `depth`: This can be nil. (Optional) The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
//  - `pierce`: This can be nil. (Optional) Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
func (doDOM DOM) RequestChildNodes(nodeID cdp.NodeID, depth *int64, pierce *bool) (err error) {
	b := dom.RequestChildNodes(nodeID)
	if depth != nil {
		b = b.WithDepth(*depth)
	}
	if pierce != nil {
		b = b.WithPierce(*pierce)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// RequestNode requests that the node is sent to the caller given the
// JavaScript node object reference. All nodes that form the path from the node
// to the root are also sent to the client as a series of setChildNodes
// notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-requestNode
//
// parameters:
//  - `objectID`: JavaScript object id to convert into node.
//
// returns:
//  - `retNodeID`: Node id for given object.
func (doDOM DOM) RequestNode(objectID runtime.RemoteObjectID) (retNodeID cdp.NodeID, err error) {
	b := dom.RequestNode(objectID)
	return b.Do(doDOM.ctxWithExecutor)
}

// ResolveNode resolves the JavaScript node object for a given NodeId or
// BackendNodeId.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-resolveNode
//
// parameters:
//  - `nodeID`: This can be nil. (Optional) Id of the node to resolve.
//  - `backendNodeID`: This can be nil. (Optional) Backend identifier of the node to resolve.
//  - `objectGroup`: This can be nil. (Optional) Symbolic group name that can be used to release multiple objects.
//  - `executionContextID`: This can be nil. (Optional) Execution context in which to resolve the node.
//
// returns:
//  - `retObject`: JavaScript object wrapper for given node.
func (doDOM DOM) ResolveNode(nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectGroup *string, executionContextID *runtime.ExecutionContextID) (retObject *runtime.RemoteObject, err error) {
	b := dom.ResolveNode()
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectGroup != nil {
		b = b.WithObjectGroup(*objectGroup)
	}
	if executionContextID != nil {
		b = b.WithExecutionContextID(*executionContextID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// SetAttributeValue sets attribute for an element with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setAttributeValue
//
// parameters:
//  - `nodeID`: Id of the element to set attribute for.
//  - `name`: Attribute name.
//  - `value`: Attribute value.
func (doDOM DOM) SetAttributeValue(nodeID cdp.NodeID, name string, value string) (err error) {
	b := dom.SetAttributeValue(nodeID, name, value)
	return b.Do(doDOM.ctxWithExecutor)
}

// SetAttributesAsText sets attributes on element with given id. This method
// is useful when user edits some existing attribute value and types in several
// attribute name/value pairs.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setAttributesAsText
//
// parameters:
//  - `nodeID`: Id of the element to set attributes for.
//  - `text`: Text with a number of attributes. Will parse this text using HTML parser.
//  - `name`: This can be nil. (Optional) Attribute name to replace with new attributes derived from text in case text parsed successfully.
func (doDOM DOM) SetAttributesAsText(nodeID cdp.NodeID, text string, name *string) (err error) {
	b := dom.SetAttributesAsText(nodeID, text)
	if name != nil {
		b = b.WithName(*name)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// SetFileInputFiles sets files for the given file input element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setFileInputFiles
//
// parameters:
//  - `files`: Array of file paths to set.
//  - `nodeID`: This can be nil. (Optional) Identifier of the node.
//  - `backendNodeID`: This can be nil. (Optional) Identifier of the backend node.
//  - `objectID`: This can be nil. (Optional) JavaScript object id of the node wrapper.
func (doDOM DOM) SetFileInputFiles(files []string, nodeID *cdp.NodeID, backendNodeID *cdp.BackendNodeID, objectID *runtime.RemoteObjectID) (err error) {
	b := dom.SetFileInputFiles(files)
	if nodeID != nil {
		b = b.WithNodeID(*nodeID)
	}
	if backendNodeID != nil {
		b = b.WithBackendNodeID(*backendNodeID)
	}
	if objectID != nil {
		b = b.WithObjectID(*objectID)
	}
	return b.Do(doDOM.ctxWithExecutor)
}

// SetNodeStackTracesEnabled sets if stack traces should be captured for
// Nodes. See Node.getNodeStackTraces. Default is disabled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setNodeStackTracesEnabled
//
// parameters:
//  - `enable`: Enable or disable.
func (doDOM DOM) SetNodeStackTracesEnabled(enable bool) (err error) {
	b := dom.SetNodeStackTracesEnabled(enable)
	return b.Do(doDOM.ctxWithExecutor)
}

// GetNodeStackTraces gets stack traces associated with a Node. As of now,
// only provides stack trace for Node creation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getNodeStackTraces
//
// parameters:
//  - `nodeID`: Id of the node to get stack traces for.
//
// returns:
//  - `retCreation`: Creation stack trace, if available.
func (doDOM DOM) GetNodeStackTraces(nodeID cdp.NodeID) (retCreation *runtime.StackTrace, err error) {
	b := dom.GetNodeStackTraces(nodeID)
	return b.Do(doDOM.ctxWithExecutor)
}

// GetFileInfo returns file information for the given File wrapper.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getFileInfo
//
// parameters:
//  - `objectID`: JavaScript object id of the node wrapper.
//
// returns:
//  - `retPath`
func (doDOM DOM) GetFileInfo(objectID runtime.RemoteObjectID) (retPath string, err error) {
	b := dom.GetFileInfo(objectID)
	return b.Do(doDOM.ctxWithExecutor)
}

// SetInspectedNode enables console to refer to the node with given id via $x
// (see Command Line API for more details $x functions).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setInspectedNode
//
// parameters:
//  - `nodeID`: DOM node id to be accessible by means of $x command line API.
func (doDOM DOM) SetInspectedNode(nodeID cdp.NodeID) (err error) {
	b := dom.SetInspectedNode(nodeID)
	return b.Do(doDOM.ctxWithExecutor)
}

// SetNodeName sets node name for a node with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setNodeName
//
// parameters:
//  - `nodeID`: Id of the node to set name for.
//  - `name`: New node's name.
//
// returns:
//  - `retNodeID`: New node's id.
func (doDOM DOM) SetNodeName(nodeID cdp.NodeID, name string) (retNodeID cdp.NodeID, err error) {
	b := dom.SetNodeName(nodeID, name)
	return b.Do(doDOM.ctxWithExecutor)
}

// SetNodeValue sets node value for a node with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setNodeValue
//
// parameters:
//  - `nodeID`: Id of the node to set value for.
//  - `value`: New node's value.
func (doDOM DOM) SetNodeValue(nodeID cdp.NodeID, value string) (err error) {
	b := dom.SetNodeValue(nodeID, value)
	return b.Do(doDOM.ctxWithExecutor)
}

// SetOuterHTML sets node HTML markup, returns new node id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-setOuterHTML
//
// parameters:
//  - `nodeID`: Id of the node to set markup for.
//  - `outerHTML`: Outer HTML markup to set.
func (doDOM DOM) SetOuterHTML(nodeID cdp.NodeID, outerHTML string) (err error) {
	b := dom.SetOuterHTML(nodeID, outerHTML)
	return b.Do(doDOM.ctxWithExecutor)
}

// Undo undoes the last performed action.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-undo
func (doDOM DOM) Undo() (err error) {
	b := dom.Undo()
	return b.Do(doDOM.ctxWithExecutor)
}

// GetFrameOwner returns iframe node that owns iframe with the given domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getFrameOwner
//
// parameters:
//  - `frameID`
//
// returns:
//  - `retBackendNodeID`: Resulting node.
//  - `retNodeID`: Id of the node at given coordinates, only when enabled and requested document.
func (doDOM DOM) GetFrameOwner(frameID cdp.FrameID) (retBackendNodeID cdp.BackendNodeID, retNodeID cdp.NodeID, err error) {
	b := dom.GetFrameOwner(frameID)
	return b.Do(doDOM.ctxWithExecutor)
}
