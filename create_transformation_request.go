// This file was auto-generated by Fern from our API Definition.

package api

type CreateTransformationRequest struct {
	// A unique, human-friendly name for the transformation <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name,omitempty"`
	// JavaScript code to be executed
	Code string `json:"code,omitempty"`
	// Key-value environment variables to be passed to the transformation
	Env *map[string]*CreateTransformationRequestEnvValue `json:"env,omitempty"`
}