/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
type V1QuobyteVolumeSource struct {

	// Group to map volume access to Default is no group
	Group string `json:"group,omitempty"`

	// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty"`

	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	Registry string `json:"registry"`

	// User to map volume access to Defaults to serivceaccount user
	User string `json:"user,omitempty"`

	// Volume is a string that references an already created Quobyte volume by name.
	Volume string `json:"volume"`
}