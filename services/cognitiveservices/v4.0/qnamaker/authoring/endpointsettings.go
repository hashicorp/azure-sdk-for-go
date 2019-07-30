package authoring

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

// EndpointSettingsClient is the an API for QnAMaker Service
type EndpointSettingsClient struct {
	BaseClient
}

// NewEndpointSettingsClient creates an instance of the EndpointSettingsClient client.
func NewEndpointSettingsClient(endpoint string) EndpointSettingsClient {
	return EndpointSettingsClient{New(endpoint)}
}

// GetSettings sends the get settings request.
func (client EndpointSettingsClient) GetSettings(ctx context.Context) (result EndpointSettingsDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointSettingsClient.GetSettings")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSettingsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.EndpointSettingsClient", "GetSettings", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSettingsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.EndpointSettingsClient", "GetSettings", resp, "Failure sending request")
		return
	}

	result, err = client.GetSettingsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.EndpointSettingsClient", "GetSettings", resp, "Failure responding to request")
	}

	return
}

// GetSettingsPreparer prepares the GetSettings request.
func (client EndpointSettingsClient) GetSettingsPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v4.0", urlParameters),
		autorest.WithPath("/endpointSettings"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSettingsSender sends the GetSettings request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointSettingsClient) GetSettingsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetSettingsResponder handles the response to the GetSettings request. The method always
// closes the http.Response Body.
func (client EndpointSettingsClient) GetSettingsResponder(resp *http.Response) (result EndpointSettingsDTO, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateSettings sends the update settings request.
// Parameters:
// endpointSettingsPayload - post body of the request.
func (client EndpointSettingsClient) UpdateSettings(ctx context.Context, endpointSettingsPayload EndpointSettingsDTO) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointSettingsClient.UpdateSettings")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateSettingsPreparer(ctx, endpointSettingsPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.EndpointSettingsClient", "UpdateSettings", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSettingsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.EndpointSettingsClient", "UpdateSettings", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateSettingsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.EndpointSettingsClient", "UpdateSettings", resp, "Failure responding to request")
	}

	return
}

// UpdateSettingsPreparer prepares the UpdateSettings request.
func (client EndpointSettingsClient) UpdateSettingsPreparer(ctx context.Context, endpointSettingsPayload EndpointSettingsDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v4.0", urlParameters),
		autorest.WithPath("/endpointSettings"),
		autorest.WithJSON(endpointSettingsPayload))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSettingsSender sends the UpdateSettings request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointSettingsClient) UpdateSettingsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateSettingsResponder handles the response to the UpdateSettings request. The method always
// closes the http.Response Body.
func (client EndpointSettingsClient) UpdateSettingsResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
