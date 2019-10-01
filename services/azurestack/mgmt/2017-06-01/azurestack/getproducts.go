package azurestack

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// GetProductsClient is the azure Stack
type GetProductsClient struct {
	BaseClient
}

// NewGetProductsClient creates an instance of the GetProductsClient client.
func NewGetProductsClient(subscriptionID string) GetProductsClient {
	return NewGetProductsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGetProductsClientWithBaseURI creates an instance of the GetProductsClient client.
func NewGetProductsClientWithBaseURI(baseURI string, subscriptionID string) GetProductsClient {
	return GetProductsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List returns a list of products.
// Parameters:
// resourceGroup - name of the resource group.
// registrationName - name of the Azure Stack registration.
// deviceConfiguration - device configuration.
func (client GetProductsClient) List(ctx context.Context, resourceGroup string, registrationName string, deviceConfiguration *DeviceConfiguration) (result ProductList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GetProductsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroup, registrationName, deviceConfiguration)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.GetProductsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azurestack.GetProductsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azurestack.GetProductsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GetProductsClient) ListPreparer(ctx context.Context, resourceGroup string, registrationName string, deviceConfiguration *DeviceConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registrationName": autorest.Encode("path", registrationName),
		"resourceGroup":    autorest.Encode("path", resourceGroup),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	deviceConfiguration.DeviceVersion = nil
	deviceConfiguration.IdentitySystem = ""
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/_all/GetProducts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if deviceConfiguration != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(deviceConfiguration))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GetProductsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GetProductsClient) ListResponder(resp *http.Response) (result ProductList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
