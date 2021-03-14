/*

 Copyright 2021 The Crossplane Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

*/

package util

import (
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	ossapi "github.com/crossplane/provider-alibaba/apis/oss/v1alpha1"
)

// Domain is aliyun Domain
var Domain = "aliyuncs.com"

var (
	errRegionNotValid            = "region is not valid"
	errCloudResourceNotSupported = "cloud resource is not supported"
)

// GetEndpoint gets endpoints for all cloud resources
func GetEndpoint(res runtime.Object, region string) (string, error) {
	klog.Info("getting endpoint for resource")
	if region == "" {
		klog.Error(errRegionNotValid)
		return "", errors.New(errRegionNotValid)
	}

	err := errors.New(errCloudResourceNotSupported)
	if res == nil || res.GetObjectKind() == nil {
		klog.ErrorS(err, "ResourceObject", res)
		return "", err
	}

	var endpoint string
	switch res.GetObjectKind().GroupVersionKind().Kind {
	case ossapi.OSSKind:
		endpoint = fmt.Sprintf("http://oss-%s.%s", region, Domain)
	default:
		return "", err
	}
	return endpoint, nil
}
