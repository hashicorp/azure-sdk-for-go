// Copyright IBM Corp. 2014, 2021
// SPDX-License-Identifier: Apache-2.0

package foo

// GatewaysClient ...
type GatewaysClient struct {
	BaseClient
}

// NewGatewaysClient ...
func NewGatewaysClient(subscriptionID string) GatewaysClient {
	return NewGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGatewaysClientWithBaseURI ...
func NewGatewaysClientWithBaseURI(baseURI string, subscriptionID string) GatewaysClient {
	return GatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}
