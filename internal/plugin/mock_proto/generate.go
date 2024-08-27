// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run github.com/golang/mock/mockgen -destination mock.go github.com/terramate-io/opentofulib/internal/tfplugin5 ProviderClient,ProvisionerClient,Provisioner_ProvisionResourceClient,Provisioner_ProvisionResourceServer

package mock_tfplugin5
