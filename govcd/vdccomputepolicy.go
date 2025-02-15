package govcd

/*
 * Copyright 2022 VMware, Inc.  All rights reserved.  Licensed under the Apache v2 License.
 */

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/vmware/go-vcloud-director/v2/types/v56"
	"github.com/vmware/go-vcloud-director/v2/util"
)

// In UI called VM sizing policy. In API VDC compute policy
type VdcComputePolicy struct {
	VdcComputePolicy *types.VdcComputePolicy
	Href             string
	client           *Client
}

// GetVdcComputePolicyById retrieves VDC compute policy by given ID
func (client *Client) GetVdcComputePolicyById(id string) (*VdcComputePolicy, error) {
	return getVdcComputePolicyById(client, id)
}

// GetVdcComputePolicyById retrieves VDC compute policy by given ID
// Deprecated: use client.GetVdcComputePolicyById
func (org *AdminOrg) GetVdcComputePolicyById(id string) (*VdcComputePolicy, error) {
	return getVdcComputePolicyById(org.client, id)
}

// GetVdcComputePolicyById retrieves VDC compute policy by given ID
// Deprecated: use client.GetVdcComputePolicyById
func (org *Org) GetVdcComputePolicyById(id string) (*VdcComputePolicy, error) {
	return getVdcComputePolicyById(org.client, id)
}

// getVdcComputePolicyById retrieves VDC compute policy by given ID
func getVdcComputePolicyById(client *Client, id string) (*VdcComputePolicy, error) {
	endpoint := types.OpenApiPathVersion1_0_0 + types.OpenApiEndpointVdcComputePolicies
	minimumApiVersion, err := client.checkOpenApiEndpointCompatibility(endpoint)
	if err != nil {
		return nil, err
	}

	if id == "" {
		return nil, fmt.Errorf("empty VDC id")
	}

	urlRef, err := client.OpenApiBuildEndpoint(endpoint, id)

	if err != nil {
		return nil, err
	}

	vdcComputePolicy := &VdcComputePolicy{
		VdcComputePolicy: &types.VdcComputePolicy{},
		Href:             urlRef.String(),
		client:           client,
	}

	err = client.OpenApiGetItem(minimumApiVersion, urlRef, nil, vdcComputePolicy.VdcComputePolicy, nil)
	if err != nil {
		return nil, err
	}

	return vdcComputePolicy, nil
}

// GetAllVdcComputePolicies retrieves all VDC compute policies using OpenAPI endpoint. Query parameters can be supplied to perform additional
// filtering
func (client *Client) GetAllVdcComputePolicies(queryParameters url.Values) ([]*VdcComputePolicy, error) {
	return getAllVdcComputePolicies(client, queryParameters)
}

// GetAllVdcComputePolicies retrieves all VDC compute policies using OpenAPI endpoint. Query parameters can be supplied to perform additional
// filtering
// Deprecated: use client.GetAllVdcComputePolicies
func (org *AdminOrg) GetAllVdcComputePolicies(queryParameters url.Values) ([]*VdcComputePolicy, error) {
	return getAllVdcComputePolicies(org.client, queryParameters)
}

// GetAllVdcComputePolicies retrieves all VDC compute policies using OpenAPI endpoint. Query parameters can be supplied to perform additional
// filtering
// Deprecated: use client.GetAllVdcComputePolicies
func (org *Org) GetAllVdcComputePolicies(queryParameters url.Values) ([]*VdcComputePolicy, error) {
	return getAllVdcComputePolicies(org.client, queryParameters)
}

// getAllVdcComputePolicies retrieves all VDC compute policies using OpenAPI endpoint. Query parameters can be supplied to perform additional
// filtering
func getAllVdcComputePolicies(client *Client, queryParameters url.Values) ([]*VdcComputePolicy, error) {
	endpoint := types.OpenApiPathVersion1_0_0 + types.OpenApiEndpointVdcComputePolicies
	minimumApiVersion, err := client.checkOpenApiEndpointCompatibility(endpoint)
	if err != nil {
		return nil, err
	}

	urlRef, err := client.OpenApiBuildEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	responses := []*types.VdcComputePolicy{{}}

	err = client.OpenApiGetAllItems(minimumApiVersion, urlRef, queryParameters, &responses, nil)
	if err != nil {
		return nil, err
	}

	var wrappedVdcComputePolicies []*VdcComputePolicy
	for _, response := range responses {
		wrappedVdcComputePolicy := &VdcComputePolicy{
			client:           client,
			VdcComputePolicy: response,
		}
		wrappedVdcComputePolicies = append(wrappedVdcComputePolicies, wrappedVdcComputePolicy)
	}

	return wrappedVdcComputePolicies, nil
}

// CreateVdcComputePolicy creates a new VDC Compute Policy using OpenAPI endpoint
// Deprecated: use client.CreateVdcComputePolicy
func (org *AdminOrg) CreateVdcComputePolicy(newVdcComputePolicy *types.VdcComputePolicy) (*VdcComputePolicy, error) {
	return org.client.CreateVdcComputePolicy(newVdcComputePolicy)
}

// CreateVdcComputePolicy creates a new VDC Compute Policy using OpenAPI endpoint
func (client *Client) CreateVdcComputePolicy(newVdcComputePolicy *types.VdcComputePolicy) (*VdcComputePolicy, error) {
	endpoint := types.OpenApiPathVersion1_0_0 + types.OpenApiEndpointVdcComputePolicies
	minimumApiVersion, err := client.checkOpenApiEndpointCompatibility(endpoint)
	if err != nil {
		return nil, err
	}

	urlRef, err := client.OpenApiBuildEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	returnVdcComputePolicy := &VdcComputePolicy{
		VdcComputePolicy: &types.VdcComputePolicy{},
		client:           client,
	}

	err = client.OpenApiPostItem(minimumApiVersion, urlRef, nil, newVdcComputePolicy, returnVdcComputePolicy.VdcComputePolicy, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating VDC compute policy: %s", err)
	}

	return returnVdcComputePolicy, nil
}

// Update existing VDC compute policy
func (vdcComputePolicy *VdcComputePolicy) Update() (*VdcComputePolicy, error) {
	endpoint := types.OpenApiPathVersion1_0_0 + types.OpenApiEndpointVdcComputePolicies
	minimumApiVersion, err := vdcComputePolicy.client.checkOpenApiEndpointCompatibility(endpoint)
	if err != nil {
		return nil, err
	}

	if vdcComputePolicy.VdcComputePolicy.ID == "" {
		return nil, fmt.Errorf("cannot update VDC compute policy without ID")
	}

	urlRef, err := vdcComputePolicy.client.OpenApiBuildEndpoint(endpoint, vdcComputePolicy.VdcComputePolicy.ID)
	if err != nil {
		return nil, err
	}

	returnVdcComputePolicy := &VdcComputePolicy{
		VdcComputePolicy: &types.VdcComputePolicy{},
		client:           vdcComputePolicy.client,
	}

	err = vdcComputePolicy.client.OpenApiPutItem(minimumApiVersion, urlRef, nil, vdcComputePolicy.VdcComputePolicy, returnVdcComputePolicy.VdcComputePolicy, nil)
	if err != nil {
		return nil, fmt.Errorf("error updating VDC compute policy: %s", err)
	}

	return returnVdcComputePolicy, nil
}

// Delete deletes VDC compute policy
func (vdcComputePolicy *VdcComputePolicy) Delete() error {
	endpoint := types.OpenApiPathVersion1_0_0 + types.OpenApiEndpointVdcComputePolicies
	minimumApiVersion, err := vdcComputePolicy.client.checkOpenApiEndpointCompatibility(endpoint)
	if err != nil {
		return err
	}

	if vdcComputePolicy.VdcComputePolicy.ID == "" {
		return fmt.Errorf("cannot delete VDC compute policy without id")
	}

	urlRef, err := vdcComputePolicy.client.OpenApiBuildEndpoint(endpoint, vdcComputePolicy.VdcComputePolicy.ID)
	if err != nil {
		return err
	}

	err = vdcComputePolicy.client.OpenApiDeleteItem(minimumApiVersion, urlRef, nil, nil)

	if err != nil {
		return fmt.Errorf("error deleting VDC compute policy: %s", err)
	}

	return nil
}

// GetAllAssignedVdcComputePolicies retrieves all VDC assigned compute policies using OpenAPI endpoint. Query parameters can be supplied to perform additional
// filtering
func (vdc *AdminVdc) GetAllAssignedVdcComputePolicies(queryParameters url.Values) ([]*VdcComputePolicy, error) {
	endpoint := types.OpenApiPathVersion1_0_0 + types.OpenApiEndpointVdcAssignedComputePolicies
	minimumApiVersion, err := vdc.client.checkOpenApiEndpointCompatibility(endpoint)
	if err != nil {
		return nil, err
	}

	urlRef, err := vdc.client.OpenApiBuildEndpoint(fmt.Sprintf(endpoint, vdc.AdminVdc.ID))
	if err != nil {
		return nil, err
	}

	responses := []*types.VdcComputePolicy{{}}

	err = vdc.client.OpenApiGetAllItems(minimumApiVersion, urlRef, queryParameters, &responses, nil)
	if err != nil {
		return nil, err
	}

	var wrappedVdcComputePolicies []*VdcComputePolicy
	for _, response := range responses {
		wrappedVdcComputePolicy := &VdcComputePolicy{
			client:           vdc.client,
			VdcComputePolicy: response,
		}
		wrappedVdcComputePolicies = append(wrappedVdcComputePolicies, wrappedVdcComputePolicy)
	}

	return wrappedVdcComputePolicies, nil
}

// SetAssignedComputePolicies assign(set) compute policies.
func (vdc *AdminVdc) SetAssignedComputePolicies(computePolicyReferences types.VdcComputePolicyReferences) (*types.VdcComputePolicyReferences, error) {
	util.Logger.Printf("[TRACE] Set Compute Policies started")

	if !vdc.client.IsSysAdmin {
		return nil, fmt.Errorf("functionality requires System Administrator privileges")
	}

	adminVdcPolicyHREF, err := url.ParseRequestURI(vdc.AdminVdc.HREF)
	if err != nil {
		return nil, fmt.Errorf("error parsing VDC URL: %s", err)
	}

	vdcId, err := GetUuidFromHref(vdc.AdminVdc.HREF, true)
	if err != nil {
		return nil, fmt.Errorf("unable to get vdc ID from HREF: %s", err)
	}
	adminVdcPolicyHREF.Path = "/api/admin/vdc/" + vdcId + "/computePolicies"

	returnedVdcComputePolicies := &types.VdcComputePolicyReferences{}
	computePolicyReferences.Xmlns = types.XMLNamespaceVCloud

	_, err = vdc.client.ExecuteRequest(adminVdcPolicyHREF.String(), http.MethodPut,
		types.MimeVdcComputePolicyReferences, "error setting compute policies for VDC: %s", computePolicyReferences, returnedVdcComputePolicies)
	if err != nil {
		return nil, err
	}

	return returnedVdcComputePolicies, nil
}
