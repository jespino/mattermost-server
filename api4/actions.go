// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/actions"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

// TODO: Add AuditLog/AuditRecords

func (api *API) InitActions() {
	api.BaseRoutes.Actions.Handle("/events", api.APISessionRequired(getSystemBusEvents)).Methods("GET")
	api.BaseRoutes.Actions.Handle("/actions", api.APISessionRequired(getActions)).Methods("GET")
	api.BaseRoutes.Actions.Handle("/graphs", api.APISessionRequired(getActionsGraphs)).Methods("GET")
	api.BaseRoutes.Actions.Handle("/graphs", api.APISessionRequired(createActionsGraph)).Methods("POST")
	api.BaseRoutes.Actions.Handle("/graphs/{graph_id:[A-Za-z0-9]+}", api.APISessionRequired(deleteActionsGraph)).Methods("DELETE")
	api.BaseRoutes.Actions.Handle("/graphs/{graph_id:[A-Za-z0-9]+}", api.APISessionRequired(updateActionsGraph)).Methods("POST")
	api.BaseRoutes.Actions.Handle("/webhook/{hook_id:[A-Za-z0-9]+}", api.APIHandler(triggerActionsWebhook)).Methods("POST")
}

func getSystemBusEvents(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	events, err := c.App.Srv().SystemBus.ListEvents()
	if err != nil {
		c.Err = model.NewAppError("Api4.getSystemBusEvents", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(events); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}
}

func getActions(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	actions, err := c.App.Srv().Actions.ListActions()
	if err != nil {
		c.Err = model.NewAppError("Api4.getActions", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(actions); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}
}

func getActionsGraphs(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	graphs, err := c.App.Srv().Actions.ListGraphs()
	if err != nil {
		c.Err = model.NewAppError("Api4.getSystemBusGraphs", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	graphsData := []*actions.GraphData{}
	for _, graph := range graphs {
		graphsData = append(graphsData, graph.ToGraphData())
	}

	if err := json.NewEncoder(w).Encode(graphsData); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}
}

func createActionsGraph(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	graphData := actions.GraphData{}
	err := json.NewDecoder(r.Body).Decode(&graphData)
	if err != nil {
		c.Err = model.NewAppError("Api4.createSystemBusLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	graphData.ID = model.NewId()

	c.App.Srv().Actions.AddGraphData(&graphData)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(graphData); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}

	ReturnStatusOK(w)
}

func deleteActionsGraph(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	c.RequireGraphId()
	if c.Err != nil {
		return
	}

	if err := c.App.Srv().Actions.DeleteGraph(c.Params.GraphId); err != nil {
		c.Err = model.NewAppError("Api4.createActionsLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnStatusOK(w)
}

func updateActionsGraph(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	c.RequireGraphId()
	if c.Err != nil {
		return
	}

	graphData := actions.GraphData{}
	err := json.NewDecoder(r.Body).Decode(&graphData)
	if err != nil {
		c.Err = model.NewAppError("Api4.createSystemBusLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.App.Srv().Actions.DeleteGraph(c.Params.GraphId); err != nil {
		c.Err = model.NewAppError("Api4.createActionsLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	graphData.ID = c.Params.GraphId

	c.App.Srv().Actions.AddGraphData(&graphData)

	ReturnStatusOK(w)
}

func triggerActionsWebhook(c *Context, w http.ResponseWriter, r *http.Request) {
	c.RequireHookId()
	if c.Err != nil {
		return
	}

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		c.Err = model.NewAppError("Api4.createSystemBusLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.App.Srv().Actions.RunHook(c.Params.HookId, data); err != nil {
		c.Err = model.NewAppError("Api4.createActionsLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnStatusOK(w)
}
