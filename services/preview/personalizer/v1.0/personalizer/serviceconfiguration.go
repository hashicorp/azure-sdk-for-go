package personalizer

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ServiceConfigurationClient is the personalizer Service is an Azure Cognitive Service that makes it easy to target
// content and experiences without complex pre-analysis or cleanup of past data. Given a context and featurized
// content, the Personalizer Service returns which content item to show to users in rewardActionId. As rewards are sent
// in response to the use of rewardActionId, the reinforcement learning algorithm will improve the model and improve
// performance of future rank calls.
type ServiceConfigurationClient struct {
	BaseClient
}

// NewServiceConfigurationClient creates an instance of the ServiceConfigurationClient client.
func NewServiceConfigurationClient(endpoint string) ServiceConfigurationClient {
	return ServiceConfigurationClient{New(endpoint)}
}

// Get get the Personalizer service configuration.
func (client ServiceConfigurationClient) Get(ctx context.Context) (result ServiceConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceConfigurationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.ServiceConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "personalizer.ServiceConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.ServiceConfigurationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceConfigurationClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/personalizer/v1.0", urlParameters),
		autorest.WithPath("/configurations/service"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceConfigurationClient) GetResponder(resp *http.Response) (result ServiceConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update the Personalizer service configuration.
// Parameters:
// config - the personalizer service configuration.
func (client ServiceConfigurationClient) Update(ctx context.Context, config ServiceConfiguration) (result ServiceConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceConfigurationClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: config,
			Constraints: []validation.Constraint{{Target: "config.RewardWaitTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "config.DefaultReward", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "config.DefaultReward", Name: validation.InclusiveMaximum, Rule: int64(1), Chain: nil},
						{Target: "config.DefaultReward", Name: validation.InclusiveMinimum, Rule: -1, Chain: nil},
					}},
				{Target: "config.RewardAggregation", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "config.RewardAggregation", Name: validation.MaxLength, Rule: 256, Chain: nil}}},
				{Target: "config.ExplorationPercentage", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "config.ExplorationPercentage", Name: validation.InclusiveMaximum, Rule: int64(1), Chain: nil},
						{Target: "config.ExplorationPercentage", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
					}},
				{Target: "config.ModelExportFrequency", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "config.LogMirrorSasURI", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "config.LogMirrorSasURI", Name: validation.MaxLength, Rule: 2048, Chain: nil}}},
				{Target: "config.LogRetentionDays", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "config.LogRetentionDays", Name: validation.InclusiveMaximum, Rule: int64(2147483647), Chain: nil},
						{Target: "config.LogRetentionDays", Name: validation.InclusiveMinimum, Rule: -1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("personalizer.ServiceConfigurationClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, config)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.ServiceConfigurationClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "personalizer.ServiceConfigurationClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.ServiceConfigurationClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServiceConfigurationClient) UpdatePreparer(ctx context.Context, config ServiceConfiguration) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}/personalizer/v1.0", urlParameters),
		autorest.WithPath("/configurations/service"),
		autorest.WithJSON(config))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceConfigurationClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServiceConfigurationClient) UpdateResponder(resp *http.Response) (result ServiceConfiguration, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
