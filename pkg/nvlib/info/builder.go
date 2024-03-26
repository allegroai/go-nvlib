/**
# Copyright 2024 NVIDIA CORPORATION
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package info

type options struct {
	root root
}

// New creates a new instance of the 'info' Interface.
func New(opts ...Option) Interface {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	if o.root == "" {
		o.root = "/"
	}
	return &propertyExtractor{
		root: o.root,
	}
}
