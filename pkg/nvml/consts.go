/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nvml

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// Return constants
const (
	SUCCESS                       = Return(nvml.SUCCESS)
	ERROR_UNINITIALIZED           = Return(nvml.ERROR_UNINITIALIZED)
	ERROR_INVALID_ARGUMENT        = Return(nvml.ERROR_INVALID_ARGUMENT)
	ERROR_NOT_SUPPORTED           = Return(nvml.ERROR_NOT_SUPPORTED)
	ERROR_NO_PERMISSION           = Return(nvml.ERROR_NO_PERMISSION)
	ERROR_ALREADY_INITIALIZED     = Return(nvml.ERROR_ALREADY_INITIALIZED)
	ERROR_NOT_FOUND               = Return(nvml.ERROR_NOT_FOUND)
	ERROR_INSUFFICIENT_SIZE       = Return(nvml.ERROR_INSUFFICIENT_SIZE)
	ERROR_INSUFFICIENT_POWER      = Return(nvml.ERROR_INSUFFICIENT_POWER)
	ERROR_DRIVER_NOT_LOADED       = Return(nvml.ERROR_DRIVER_NOT_LOADED)
	ERROR_TIMEOUT                 = Return(nvml.ERROR_TIMEOUT)
	ERROR_IRQ_ISSUE               = Return(nvml.ERROR_IRQ_ISSUE)
	ERROR_LIBRARY_NOT_FOUND       = Return(nvml.ERROR_LIBRARY_NOT_FOUND)
	ERROR_FUNCTION_NOT_FOUND      = Return(nvml.ERROR_FUNCTION_NOT_FOUND)
	ERROR_CORRUPTED_INFOROM       = Return(nvml.ERROR_CORRUPTED_INFOROM)
	ERROR_GPU_IS_LOST             = Return(nvml.ERROR_GPU_IS_LOST)
	ERROR_RESET_REQUIRED          = Return(nvml.ERROR_RESET_REQUIRED)
	ERROR_OPERATING_SYSTEM        = Return(nvml.ERROR_OPERATING_SYSTEM)
	ERROR_LIB_RM_VERSION_MISMATCH = Return(nvml.ERROR_LIB_RM_VERSION_MISMATCH)
	ERROR_IN_USE                  = Return(nvml.ERROR_IN_USE)
	ERROR_MEMORY                  = Return(nvml.ERROR_MEMORY)
	ERROR_NO_DATA                 = Return(nvml.ERROR_NO_DATA)
	ERROR_VGPU_ECC_NOT_SUPPORTED  = Return(nvml.ERROR_VGPU_ECC_NOT_SUPPORTED)
	ERROR_INSUFFICIENT_RESOURCES  = Return(nvml.ERROR_INSUFFICIENT_RESOURCES)
	ERROR_UNKNOWN                 = Return(nvml.ERROR_UNKNOWN)
)

// MIG Mode constants
const (
	DEVICE_MIG_ENABLE  = nvml.DEVICE_MIG_ENABLE
	DEVICE_MIG_DISABLE = nvml.DEVICE_MIG_DISABLE
)

// GPU Instance Profiles
const (
	GPU_INSTANCE_PROFILE_1_SLICE      = nvml.GPU_INSTANCE_PROFILE_1_SLICE
	GPU_INSTANCE_PROFILE_2_SLICE      = nvml.GPU_INSTANCE_PROFILE_2_SLICE
	GPU_INSTANCE_PROFILE_3_SLICE      = nvml.GPU_INSTANCE_PROFILE_3_SLICE
	GPU_INSTANCE_PROFILE_4_SLICE      = nvml.GPU_INSTANCE_PROFILE_4_SLICE
	GPU_INSTANCE_PROFILE_6_SLICE      = nvml.GPU_INSTANCE_PROFILE_6_SLICE
	GPU_INSTANCE_PROFILE_7_SLICE      = nvml.GPU_INSTANCE_PROFILE_7_SLICE
	GPU_INSTANCE_PROFILE_8_SLICE      = nvml.GPU_INSTANCE_PROFILE_8_SLICE
	GPU_INSTANCE_PROFILE_1_SLICE_REV1 = nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV1
	GPU_INSTANCE_PROFILE_1_SLICE_REV2 = nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2
	GPU_INSTANCE_PROFILE_2_SLICE_REV1 = nvml.GPU_INSTANCE_PROFILE_2_SLICE_REV1
	GPU_INSTANCE_PROFILE_COUNT        = nvml.GPU_INSTANCE_PROFILE_COUNT
)

// Compute Instance Profiles
const (
	COMPUTE_INSTANCE_PROFILE_1_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE
	COMPUTE_INSTANCE_PROFILE_2_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE
	COMPUTE_INSTANCE_PROFILE_3_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE
	COMPUTE_INSTANCE_PROFILE_4_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE
	COMPUTE_INSTANCE_PROFILE_6_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_6_SLICE
	COMPUTE_INSTANCE_PROFILE_7_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_7_SLICE
	COMPUTE_INSTANCE_PROFILE_8_SLICE      = nvml.COMPUTE_INSTANCE_PROFILE_8_SLICE
	COMPUTE_INSTANCE_PROFILE_1_SLICE_REV1 = nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE_REV1
	COMPUTE_INSTANCE_PROFILE_COUNT        = nvml.COMPUTE_INSTANCE_PROFILE_COUNT
)

// Compute Instance Engine Profiles
const (
	COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED = nvml.COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED
	COMPUTE_INSTANCE_ENGINE_PROFILE_COUNT  = nvml.COMPUTE_INSTANCE_ENGINE_PROFILE_COUNT
)

// Event Types
const (
	EventTypeXidCriticalError  = nvml.EventTypeXidCriticalError
	EventTypeSingleBitEccError = nvml.EventTypeSingleBitEccError
	EventTypeDoubleBitEccError = nvml.EventTypeDoubleBitEccError
)
