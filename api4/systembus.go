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

func (api *API) InitSystemBus() {
	api.BaseRoutes.SystemBus.Handle("/events", api.APISessionRequired(getSystemBusEvents)).Methods("GET")
	api.BaseRoutes.SystemBus.Handle("/actions", api.APISessionRequired(getSystemBusActions)).Methods("GET")
	api.BaseRoutes.SystemBus.Handle("/links", api.APISessionRequired(getSystemBusLinks)).Methods("GET")
	api.BaseRoutes.SystemBus.Handle("/links", api.APISessionRequired(createSystemBusLink)).Methods("POST")
	api.BaseRoutes.SystemBus.Handle("/links/{link_id:[A-Za-z0-9]+}", api.APISessionRequired(deleteSystemBusLink)).Methods("DELETE")
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

func getSystemBusActions(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	actions, err := c.App.Srv().Actions.ListActions()
	if err != nil {
		c.Err = model.NewAppError("Api4.getSystemBusActions", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(actions); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}
}

func getSystemBusLinks(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	links, err := c.App.Srv().Actions.ListLinks()
	if err != nil {
		c.Err = model.NewAppError("Api4.getSystemBusLinks", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(links); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}
}

func createSystemBusLink(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	link := actions.LinkEventAction{}

	json.NewDecoder(r.Body).Decode(&link)
	if link.EventID == "" {
		c.SetInvalidParam("event_id")
		return
	}

	if link.ActionID == "" {
		c.SetInvalidParam("action_id")
		return
	}

	newLink, err := c.App.Srv().Actions.LinkEventAction(link.EventID, link.ActionID, link.Config)
	if err != nil {
		c.Err = model.NewAppError("Api4.createSystemBusLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newLink); err != nil {
		mlog.Warn("Error while writing response", mlog.Err(err))
	}
}

func deleteSystemBusLink(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.SessionHasPermissionTo(*c.AppContext.Session(), model.PermissionManageSystem) {
		c.SetPermissionError(model.PermissionManageSystem)
		return
	}

	c.RequireLinkId()
	if c.Err != nil {
		return
	}

	if err := c.App.Srv().Actions.UnlinkEventAction(c.Params.LinkId); err != nil {
		c.Err = model.NewAppError("Api4.createSystemBusLink", "api.systembus.request_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnStatusOK(w)
}
