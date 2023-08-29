// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/token)
	PostAuthToken(ctx echo.Context, params PostAuthTokenParams) error

	// (GET /daemon/dns/dump)
	GetDaemonDNSDump(ctx echo.Context) error

	// (GET /daemon/events)
	GetDaemonEvents(ctx echo.Context, params GetDaemonEventsParams) error

	// (POST /daemon/join)
	PostDaemonJoin(ctx echo.Context, params PostDaemonJoinParams) error

	// (POST /daemon/leave)
	PostDaemonLeave(ctx echo.Context, params PostDaemonLeaveParams) error

	// (POST /daemon/logs/control)
	PostDaemonLogsControl(ctx echo.Context) error

	// (GET /daemon/running)
	GetDaemonRunning(ctx echo.Context) error

	// (GET /daemon/status)
	GetDaemonStatus(ctx echo.Context, params GetDaemonStatusParams) error

	// (POST /daemon/stop)
	PostDaemonStop(ctx echo.Context) error

	// (POST /daemon/sub/action)
	PostDaemonSubAction(ctx echo.Context) error

	// (GET /instance/status)
	GetInstanceStatus(ctx echo.Context) error

	// (POST /instance/status)
	PostInstanceStatus(ctx echo.Context) error

	// (GET /networks)
	GetNetworks(ctx echo.Context, params GetNetworksParams) error

	// (GET /node/backlogs)
	GetNodeBacklogs(ctx echo.Context, params GetNodeBacklogsParams) error

	// (POST /node/clear)
	PostNodeClear(ctx echo.Context) error

	// (GET /node/drbd/allocation)
	GetNodeDRBDAllocation(ctx echo.Context) error

	// (GET /node/drbd/config)
	GetNodeDRBDConfig(ctx echo.Context, params GetNodeDRBDConfigParams) error

	// (POST /node/drbd/config)
	PostNodeDRBDConfig(ctx echo.Context, params PostNodeDRBDConfigParams) error

	// (GET /node/logs)
	GetNodeLogs(ctx echo.Context, params GetNodeLogsParams) error

	// (POST /node/monitor)
	PostNodeMonitor(ctx echo.Context) error

	// (GET /nodes/info)
	GetNodesInfo(ctx echo.Context) error

	// (POST /object/abort)
	PostObjectAbort(ctx echo.Context) error

	// (GET /object/backlogs)
	GetObjectBacklogs(ctx echo.Context, params GetObjectBacklogsParams) error

	// (POST /object/clear)
	PostObjectClear(ctx echo.Context) error

	// (GET /object/config)
	GetObjectConfig(ctx echo.Context, params GetObjectConfigParams) error

	// (GET /object/file)
	GetObjectFile(ctx echo.Context, params GetObjectFileParams) error

	// (GET /object/logs)
	GetObjectLogs(ctx echo.Context, params GetObjectLogsParams) error

	// (POST /object/monitor)
	PostObjectMonitor(ctx echo.Context) error

	// (POST /object/progress)
	PostObjectProgress(ctx echo.Context) error

	// (GET /object/selector)
	GetObjectSelector(ctx echo.Context, params GetObjectSelectorParams) error

	// (POST /object/switchTo)
	PostObjectSwitchTo(ctx echo.Context) error

	// (GET /pools)
	GetPools(ctx echo.Context, params GetPoolsParams) error

	// (GET /public/openapi)
	GetSwagger(ctx echo.Context) error

	// (GET /relay/message)
	GetRelayMessage(ctx echo.Context, params GetRelayMessageParams) error

	// (POST /relay/message)
	PostRelayMessage(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostAuthToken converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthToken(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostAuthTokenParams
	// ------------- Optional query parameter "role" -------------

	err = runtime.BindQueryParameter("form", true, false, "role", ctx.QueryParams(), &params.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter role: %s", err))
	}

	// ------------- Optional query parameter "duration" -------------

	err = runtime.BindQueryParameter("form", true, false, "duration", ctx.QueryParams(), &params.Duration)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter duration: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostAuthToken(ctx, params)
	return err
}

// GetDaemonDNSDump converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonDNSDump(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonDNSDump(ctx)
	return err
}

// GetDaemonEvents converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonEvents(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonEventsParams
	// ------------- Optional query parameter "duration" -------------

	err = runtime.BindQueryParameter("form", true, false, "duration", ctx.QueryParams(), &params.Duration)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter duration: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonEvents(ctx, params)
	return err
}

// PostDaemonJoin converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonJoin(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostDaemonJoinParams
	// ------------- Required query parameter "node" -------------

	err = runtime.BindQueryParameter("form", true, true, "node", ctx.QueryParams(), &params.Node)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter node: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonJoin(ctx, params)
	return err
}

// PostDaemonLeave converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonLeave(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostDaemonLeaveParams
	// ------------- Required query parameter "node" -------------

	err = runtime.BindQueryParameter("form", true, true, "node", ctx.QueryParams(), &params.Node)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter node: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonLeave(ctx, params)
	return err
}

// PostDaemonLogsControl converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonLogsControl(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonLogsControl(ctx)
	return err
}

// GetDaemonRunning converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonRunning(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonRunning(ctx)
	return err
}

// GetDaemonStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetDaemonStatus(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonStatusParams
	// ------------- Optional query parameter "namespace" -------------

	err = runtime.BindQueryParameter("form", true, false, "namespace", ctx.QueryParams(), &params.Namespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter namespace: %s", err))
	}

	// ------------- Optional query parameter "selector" -------------

	err = runtime.BindQueryParameter("form", true, false, "selector", ctx.QueryParams(), &params.Selector)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter selector: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDaemonStatus(ctx, params)
	return err
}

// PostDaemonStop converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonStop(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonStop(ctx)
	return err
}

// PostDaemonSubAction converts echo context to params.
func (w *ServerInterfaceWrapper) PostDaemonSubAction(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostDaemonSubAction(ctx)
	return err
}

// GetInstanceStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetInstanceStatus(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetInstanceStatus(ctx)
	return err
}

// PostInstanceStatus converts echo context to params.
func (w *ServerInterfaceWrapper) PostInstanceStatus(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostInstanceStatus(ctx)
	return err
}

// GetNetworks converts echo context to params.
func (w *ServerInterfaceWrapper) GetNetworks(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNetworksParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNetworks(ctx, params)
	return err
}

// GetNodeBacklogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeBacklogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeBacklogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeBacklogs(ctx, params)
	return err
}

// PostNodeClear converts echo context to params.
func (w *ServerInterfaceWrapper) PostNodeClear(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNodeClear(ctx)
	return err
}

// GetNodeDRBDAllocation converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeDRBDAllocation(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeDRBDAllocation(ctx)
	return err
}

// GetNodeDRBDConfig converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeDRBDConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeDRBDConfigParams
	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeDRBDConfig(ctx, params)
	return err
}

// PostNodeDRBDConfig converts echo context to params.
func (w *ServerInterfaceWrapper) PostNodeDRBDConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostNodeDRBDConfigParams
	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNodeDRBDConfig(ctx, params)
	return err
}

// GetNodeLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodeLogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeLogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodeLogs(ctx, params)
	return err
}

// PostNodeMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) PostNodeMonitor(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNodeMonitor(ctx)
	return err
}

// GetNodesInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetNodesInfo(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNodesInfo(ctx)
	return err
}

// PostObjectAbort converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectAbort(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectAbort(ctx)
	return err
}

// GetObjectBacklogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectBacklogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectBacklogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectBacklogs(ctx, params)
	return err
}

// PostObjectClear converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectClear(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectClear(ctx)
	return err
}

// GetObjectConfig converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectConfig(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectConfigParams
	// ------------- Required query parameter "path" -------------

	err = runtime.BindQueryParameter("form", true, true, "path", ctx.QueryParams(), &params.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	// ------------- Optional query parameter "evaluate" -------------

	err = runtime.BindQueryParameter("form", true, false, "evaluate", ctx.QueryParams(), &params.Evaluate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter evaluate: %s", err))
	}

	// ------------- Optional query parameter "impersonate" -------------

	err = runtime.BindQueryParameter("form", true, false, "impersonate", ctx.QueryParams(), &params.Impersonate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter impersonate: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectConfig(ctx, params)
	return err
}

// GetObjectFile converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectFile(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectFileParams
	// ------------- Required query parameter "path" -------------

	err = runtime.BindQueryParameter("form", true, true, "path", ctx.QueryParams(), &params.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectFile(ctx, params)
	return err
}

// GetObjectLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectLogs(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectLogsParams
	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// ------------- Required query parameter "paths" -------------

	err = runtime.BindQueryParameter("form", true, true, "paths", ctx.QueryParams(), &params.Paths)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter paths: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectLogs(ctx, params)
	return err
}

// PostObjectMonitor converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectMonitor(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectMonitor(ctx)
	return err
}

// PostObjectProgress converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectProgress(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectProgress(ctx)
	return err
}

// GetObjectSelector converts echo context to params.
func (w *ServerInterfaceWrapper) GetObjectSelector(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectSelectorParams
	// ------------- Required query parameter "selector" -------------

	err = runtime.BindQueryParameter("form", true, true, "selector", ctx.QueryParams(), &params.Selector)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter selector: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetObjectSelector(ctx, params)
	return err
}

// PostObjectSwitchTo converts echo context to params.
func (w *ServerInterfaceWrapper) PostObjectSwitchTo(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostObjectSwitchTo(ctx)
	return err
}

// GetPools converts echo context to params.
func (w *ServerInterfaceWrapper) GetPools(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPoolsParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPools(ctx, params)
	return err
}

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetRelayMessage converts echo context to params.
func (w *ServerInterfaceWrapper) GetRelayMessage(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRelayMessageParams
	// ------------- Optional query parameter "nodename" -------------

	err = runtime.BindQueryParameter("form", true, false, "nodename", ctx.QueryParams(), &params.Nodename)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nodename: %s", err))
	}

	// ------------- Optional query parameter "cluster_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "cluster_id", ctx.QueryParams(), &params.ClusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter cluster_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRelayMessage(ctx, params)
	return err
}

// PostRelayMessage converts echo context to params.
func (w *ServerInterfaceWrapper) PostRelayMessage(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostRelayMessage(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/auth/token", wrapper.PostAuthToken)
	router.GET(baseURL+"/daemon/dns/dump", wrapper.GetDaemonDNSDump)
	router.GET(baseURL+"/daemon/events", wrapper.GetDaemonEvents)
	router.POST(baseURL+"/daemon/join", wrapper.PostDaemonJoin)
	router.POST(baseURL+"/daemon/leave", wrapper.PostDaemonLeave)
	router.POST(baseURL+"/daemon/logs/control", wrapper.PostDaemonLogsControl)
	router.GET(baseURL+"/daemon/running", wrapper.GetDaemonRunning)
	router.GET(baseURL+"/daemon/status", wrapper.GetDaemonStatus)
	router.POST(baseURL+"/daemon/stop", wrapper.PostDaemonStop)
	router.POST(baseURL+"/daemon/sub/action", wrapper.PostDaemonSubAction)
	router.GET(baseURL+"/instance/status", wrapper.GetInstanceStatus)
	router.POST(baseURL+"/instance/status", wrapper.PostInstanceStatus)
	router.GET(baseURL+"/networks", wrapper.GetNetworks)
	router.GET(baseURL+"/node/backlogs", wrapper.GetNodeBacklogs)
	router.POST(baseURL+"/node/clear", wrapper.PostNodeClear)
	router.GET(baseURL+"/node/drbd/allocation", wrapper.GetNodeDRBDAllocation)
	router.GET(baseURL+"/node/drbd/config", wrapper.GetNodeDRBDConfig)
	router.POST(baseURL+"/node/drbd/config", wrapper.PostNodeDRBDConfig)
	router.GET(baseURL+"/node/logs", wrapper.GetNodeLogs)
	router.POST(baseURL+"/node/monitor", wrapper.PostNodeMonitor)
	router.GET(baseURL+"/nodes/info", wrapper.GetNodesInfo)
	router.POST(baseURL+"/object/abort", wrapper.PostObjectAbort)
	router.GET(baseURL+"/object/backlogs", wrapper.GetObjectBacklogs)
	router.POST(baseURL+"/object/clear", wrapper.PostObjectClear)
	router.GET(baseURL+"/object/config", wrapper.GetObjectConfig)
	router.GET(baseURL+"/object/file", wrapper.GetObjectFile)
	router.GET(baseURL+"/object/logs", wrapper.GetObjectLogs)
	router.POST(baseURL+"/object/monitor", wrapper.PostObjectMonitor)
	router.POST(baseURL+"/object/progress", wrapper.PostObjectProgress)
	router.GET(baseURL+"/object/selector", wrapper.GetObjectSelector)
	router.POST(baseURL+"/object/switchTo", wrapper.PostObjectSwitchTo)
	router.GET(baseURL+"/pools", wrapper.GetPools)
	router.GET(baseURL+"/public/openapi", wrapper.GetSwagger)
	router.GET(baseURL+"/relay/message", wrapper.GetRelayMessage)
	router.POST(baseURL+"/relay/message", wrapper.PostRelayMessage)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
