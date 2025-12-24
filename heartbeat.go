package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Heartbeat struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	TTL         string            `json:"ttl,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	RefreshedAt *time.Time        `json:"refreshed_at,omitempty"`
}

type CreateHeartbeatInput struct {
	Name        string            `json:"name" validate:"required,max=255,min=1"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty" validate:"dive,keys,max=255,min=1,endkeys,max=255,min=1"`
	TTL         string            `json:"ttl,omitempty"`
}

type UpdateHeartbeatInput struct {
	ID          string            `json:"id" validate:"required,uuid"`
	Name        string            `json:"name" validate:"required,max=255,min=1"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty" validate:"dive,keys,max=255,min=1,endkeys,max=255,min=1"`
	TTL         string            `json:"ttl,omitempty"`
}

type DeleteHeartbeatInput struct {
	ID string `param:"id" validate:"required,uuid"`
}

type GetHeartbeatInput struct {
	Identifier string `param:"identifier" description:"Heartbeat Name or ID"`
}

type RefreshHeartbeatInput struct {
	ID string `param:"id" validate:"required,uuid"`
}

type ListHeartbeatsOutput struct {
	Result []Heartbeat `json:"result"`
}

func (c *Client) CreateHeartbeat(ctx context.Context, input CreateHeartbeatInput) (Response, error) {
	var result Response
	_, err := c.sendRequest(ctx, "/api/v1/heartbeat", http.MethodPost, input, &result, nil)
	if err != nil {
		return Response{}, err
	}
	return result, nil
}

func (c *Client) UpdateHeartbeat(ctx context.Context, input UpdateHeartbeatInput) (Response, error) {
	var result Response
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/heartbeat/%s", input.ID), http.MethodPut, input, &result, nil)
	if err != nil {
		return Response{}, err
	}
	return result, nil
}

func (c *Client) DeleteHeartbeat(ctx context.Context, input DeleteHeartbeatInput) (Response, error) {
	var result Response
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/heartbeat/%s", input.ID), http.MethodDelete, nil, &result, nil)
	if err != nil {
		return Response{}, err
	}
	return result, nil
}

func (c *Client) GetHeartbeat(ctx context.Context, input GetHeartbeatInput) (Heartbeat, error) {
	var result Heartbeat
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/heartbeat/%s", input.Identifier), http.MethodGet, nil, &result, nil)
	if err != nil {
		return Heartbeat{}, err
	}
	return result, nil
}

func (c *Client) RefreshHeartbeat(ctx context.Context, input RefreshHeartbeatInput) (Response, error) {
	var result Response
	_, err := c.sendRequest(ctx, fmt.Sprintf("/api/v1/heartbeat/%s/refresh", input.ID), http.MethodPost, nil, &result, nil)
	if err != nil {
		return Response{}, err
	}
	return result, nil
}

func (c *Client) ListHeartbeats(ctx context.Context) (ListHeartbeatsOutput, error) {
	var result ListHeartbeatsOutput
	_, err := c.sendRequest(ctx, "/api/v1/heartbeat", http.MethodGet, nil, &result, nil)
	if err != nil {
		return ListHeartbeatsOutput{}, err
	}
	return result, nil
}