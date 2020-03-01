// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build windows

package providers

import (
	"github.com/DataDog/datadog-agent/pkg/util/containers"
	"github.com/DataDog/datadog-agent/pkg/util/containers/providers/windows"
)

// ContainerImpl provides a ContainerImplementation for Windows builds
var ContainerImpl containers.ContainerImplementation

func init() {
	ContainerImpl = &windows.Provider{}
}
