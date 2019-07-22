// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Comment defines model for Comment.
type Comment struct {
	Author    User      `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentsPage defines model for CommentsPage.
type CommentsPage struct {
	Comments  *[]Comment `json:"comments,omitempty"`
	NextToken *string    `json:"nextToken,omitempty"`
}

// Issue defines model for Issue.
type Issue struct {
	Assignee  *User         `json:"assignee,omitempty"`
	Author    User          `json:"author"`
	Category  *string       `json:"category,omitempty"`
	Comments  *CommentsPage `json:"comments,omitempty"`
	Content   *string       `json:"content,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	Id        string        `json:"id"`
	Labels    *[]string     `json:"labels,omitempty"`
	Severity  *string       `json:"severity,omitempty"`
	Subject   string        `json:"subject"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// IssuesPage defines model for IssuesPage.
type IssuesPage struct {
	Issues    *[]Issue `json:"issues,omitempty"`
	NextToken *string  `json:"nextToken,omitempty"`
}

// NewComment defines model for NewComment.
type NewComment struct {
	Content string `json:"content"`
}

// NewIssue defines model for NewIssue.
type NewIssue struct {
	Content *string   `json:"content,omitempty"`
	Labels  *[]string `json:"labels,omitempty"`
	Subject string    `json:"subject"`
}

// NewProject defines model for NewProject.
type NewProject struct {
	Description *string  `json:"description,omitempty"`
	Labels      []string `json:"labels"`
	Name        string   `json:"name"`
}

// Project defines model for Project.
type Project struct {
	CreatedAt   time.Time `json:"created_at"`
	Description *string   `json:"description,omitempty"`
	Id          string    `json:"id"`
	Labels      []string  `json:"labels"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProjectsPage defines model for ProjectsPage.
type ProjectsPage struct {
	NextToken *string    `json:"nextToken,omitempty"`
	Projects  *[]Project `json:"projects,omitempty"`
}

// UpdatedProject defines model for UpdatedProject.
type UpdatedProject struct {
	// Embedded struct due to allOf(#/components/schemas/NewProject)
	NewProject
	// Embedded fields due to inline allOf schema
	Version int64 `json:"version"`
}

// User defines model for User.
type User struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UsersPage defines model for UsersPage.
type UsersPage struct {
	NextToken *string `json:"nextToken,omitempty"`
	Users     *[]User `json:"users,omitempty"`
}

// FilterIssues defines model for filterIssues.
type FilterIssues []string

// Limit defines model for limit.
type Limit int64

// NextToken defines model for nextToken.
type NextToken string

// ProjectsParams defines parameters for Projects.
type ProjectsParams struct {
	NextToken *NextToken `json:"nextToken,omitempty"`
	Limit     *Limit     `json:"limit,omitempty"`
}

// IssuesParams defines parameters for Issues.
type IssuesParams struct {
	NextToken *NextToken    `json:"nextToken,omitempty"`
	Limit     *Limit        `json:"limit,omitempty"`
	Filter    *FilterIssues `json:"filter,omitempty"`
}

// CommentsParams defines parameters for Comments.
type CommentsParams struct {
	NextToken *NextToken `json:"nextToken,omitempty"`
	Limit     *Limit     `json:"limit,omitempty"`
}

// UsersParams defines parameters for Users.
type UsersParams struct {
	NextToken *NextToken `json:"nextToken,omitempty"`
	Limit     *Limit     `json:"limit,omitempty"`
}

// NewProjectRequestBody defines body for NewProject for application/json ContentType.
type NewProjectJSONRequestBody NewProject

// UpdateProjectRequestBody defines body for UpdateProject for application/json ContentType.
type UpdateProjectJSONRequestBody UpdatedProject

// NewIssueRequestBody defines body for NewIssue for application/json ContentType.
type NewIssueJSONRequestBody NewIssue

// NewCommentRequestBody defines body for NewComment for application/json ContentType.
type NewCommentJSONRequestBody NewComment

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of projects.// (GET /projects)
	Projects(ctx echo.Context, params ProjectsParams) error
	// Create a project.// (POST /projects)
	NewProject(ctx echo.Context) error
	// (GET /projects/{id})
	GetProject(ctx echo.Context, id string) error
	// Update a project.// (PUT /projects/{id})
	UpdateProject(ctx echo.Context, id string) error
	// Get a list of issues.// (GET /projects/{project_id}/issues)
	Issues(ctx echo.Context, projectId string, params IssuesParams) error
	// Create a issue.// (POST /projects/{project_id}/issues)
	NewIssue(ctx echo.Context, projectId string) error
	// (GET /projects/{project_id}/issues/{id})
	GetIssue(ctx echo.Context, projectId string, id string) error
	// Get a list of Comments.// (GET /projects/{project_id}/issues/{issue_id}/comments)
	Comments(ctx echo.Context, projectId string, issueId string, params CommentsParams) error
	// Create a comment on a issue.// (POST /projects/{project_id}/issues/{issue_id}/comments)
	NewComment(ctx echo.Context, projectId string, issueId string) error
	// (GET /projects/{project_id}/issues/{issue_id}/comments/{id})
	GetComment(ctx echo.Context, projectId string, issueId string, id string) error
	// Get a list of users.// (GET /users)
	Users(ctx echo.Context, params UsersParams) error
	// (GET /users/{id})
	GetUser(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Projects converts echo context to params.
func (w *ServerInterfaceWrapper) Projects(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ProjectsParams
	// ------------- Optional query parameter "nextToken" -------------
	if paramValue := ctx.QueryParam("nextToken"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "nextToken", ctx.QueryParams(), &params.NextToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nextToken: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Projects(ctx, params)
	return err
}

// NewProject converts echo context to params.
func (w *ServerInterfaceWrapper) NewProject(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.NewProject(ctx)
	return err
}

// GetProject converts echo context to params.
func (w *ServerInterfaceWrapper) GetProject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProject(ctx, id)
	return err
}

// UpdateProject converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateProject(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateProject(ctx, id)
	return err
}

// Issues converts echo context to params.
func (w *ServerInterfaceWrapper) Issues(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "project_id" -------------
	var projectId string

	err = runtime.BindStyledParameter("simple", false, "project_id", ctx.Param("project_id"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params IssuesParams
	// ------------- Optional query parameter "nextToken" -------------
	if paramValue := ctx.QueryParam("nextToken"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "nextToken", ctx.QueryParams(), &params.NextToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nextToken: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "filter" -------------
	if paramValue := ctx.QueryParam("filter"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("pipeDelimited", true, false, "filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter filter: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Issues(ctx, projectId, params)
	return err
}

// NewIssue converts echo context to params.
func (w *ServerInterfaceWrapper) NewIssue(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "project_id" -------------
	var projectId string

	err = runtime.BindStyledParameter("simple", false, "project_id", ctx.Param("project_id"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.NewIssue(ctx, projectId)
	return err
}

// GetIssue converts echo context to params.
func (w *ServerInterfaceWrapper) GetIssue(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "project_id" -------------
	var projectId string

	err = runtime.BindStyledParameter("simple", false, "project_id", ctx.Param("project_id"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_id: %s", err))
	}

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIssue(ctx, projectId, id)
	return err
}

// Comments converts echo context to params.
func (w *ServerInterfaceWrapper) Comments(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "project_id" -------------
	var projectId string

	err = runtime.BindStyledParameter("simple", false, "project_id", ctx.Param("project_id"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_id: %s", err))
	}

	// ------------- Path parameter "issue_id" -------------
	var issueId string

	err = runtime.BindStyledParameter("simple", false, "issue_id", ctx.Param("issue_id"), &issueId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter issue_id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params CommentsParams
	// ------------- Optional query parameter "nextToken" -------------
	if paramValue := ctx.QueryParam("nextToken"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "nextToken", ctx.QueryParams(), &params.NextToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nextToken: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Comments(ctx, projectId, issueId, params)
	return err
}

// NewComment converts echo context to params.
func (w *ServerInterfaceWrapper) NewComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "project_id" -------------
	var projectId string

	err = runtime.BindStyledParameter("simple", false, "project_id", ctx.Param("project_id"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_id: %s", err))
	}

	// ------------- Path parameter "issue_id" -------------
	var issueId string

	err = runtime.BindStyledParameter("simple", false, "issue_id", ctx.Param("issue_id"), &issueId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter issue_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.NewComment(ctx, projectId, issueId)
	return err
}

// GetComment converts echo context to params.
func (w *ServerInterfaceWrapper) GetComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "project_id" -------------
	var projectId string

	err = runtime.BindStyledParameter("simple", false, "project_id", ctx.Param("project_id"), &projectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_id: %s", err))
	}

	// ------------- Path parameter "issue_id" -------------
	var issueId string

	err = runtime.BindStyledParameter("simple", false, "issue_id", ctx.Param("issue_id"), &issueId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter issue_id: %s", err))
	}

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComment(ctx, projectId, issueId, id)
	return err
}

// Users converts echo context to params.
func (w *ServerInterfaceWrapper) Users(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params UsersParams
	// ------------- Optional query parameter "nextToken" -------------
	if paramValue := ctx.QueryParam("nextToken"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "nextToken", ctx.QueryParams(), &params.NextToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nextToken: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {
	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Users(ctx, params)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUser(ctx, id)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/projects", wrapper.Projects)
	router.POST("/projects", wrapper.NewProject)
	router.GET("/projects/:id", wrapper.GetProject)
	router.PUT("/projects/:id", wrapper.UpdateProject)
	router.GET("/projects/:project_id/issues", wrapper.Issues)
	router.POST("/projects/:project_id/issues", wrapper.NewIssue)
	router.GET("/projects/:project_id/issues/:id", wrapper.GetIssue)
	router.GET("/projects/:project_id/issues/:issue_id/comments", wrapper.Comments)
	router.POST("/projects/:project_id/issues/:issue_id/comments", wrapper.NewComment)
	router.GET("/projects/:project_id/issues/:issue_id/comments/:id", wrapper.GetComment)
	router.GET("/users", wrapper.Users)
	router.GET("/users/:id", wrapper.GetUser)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"9fn69wAAAP//+dK2XdQsAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}