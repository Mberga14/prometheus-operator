// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
)

// MSTeamsV2ConfigApplyConfiguration represents a declarative configuration of the MSTeamsV2Config type for use
// with apply.
type MSTeamsV2ConfigApplyConfiguration struct {
	SendResolved *bool                         `json:"sendResolved,omitempty"`
	WebhookURL   *v1.SecretKeySelector         `json:"webhookURL,omitempty"`
	Title        *string                       `json:"title,omitempty"`
	Text         *string                       `json:"text,omitempty"`
	HTTPConfig   *HTTPConfigApplyConfiguration `json:"httpConfig,omitempty"`
}

// MSTeamsV2ConfigApplyConfiguration constructs a declarative configuration of the MSTeamsV2Config type for use with
// apply.
func MSTeamsV2Config() *MSTeamsV2ConfigApplyConfiguration {
	return &MSTeamsV2ConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *MSTeamsV2ConfigApplyConfiguration) WithSendResolved(value bool) *MSTeamsV2ConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithWebhookURL sets the WebhookURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookURL field is set to the value of the last call.
func (b *MSTeamsV2ConfigApplyConfiguration) WithWebhookURL(value v1.SecretKeySelector) *MSTeamsV2ConfigApplyConfiguration {
	b.WebhookURL = &value
	return b
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *MSTeamsV2ConfigApplyConfiguration) WithTitle(value string) *MSTeamsV2ConfigApplyConfiguration {
	b.Title = &value
	return b
}

// WithText sets the Text field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Text field is set to the value of the last call.
func (b *MSTeamsV2ConfigApplyConfiguration) WithText(value string) *MSTeamsV2ConfigApplyConfiguration {
	b.Text = &value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *MSTeamsV2ConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *MSTeamsV2ConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}
