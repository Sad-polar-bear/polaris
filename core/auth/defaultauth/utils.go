/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package defaultauth

import (
	api "github.com/polarismesh/polaris-server/common/api/v1"
	"github.com/polarismesh/polaris-server/store"
)

var storeCodeAPICodeMap = map[store.StatusCode]uint32{
	store.EmptyParamsErr:             api.InvalidParameter,
	store.OutOfRangeErr:              api.InvalidParameter,
	store.DataConflictErr:            api.DataConflict,
	store.NotFoundNamespace:          api.NotFoundNamespace,
	store.NotFoundService:            api.NotFoundService,
	store.NotFoundMasterConfig:       api.NotFoundMasterConfig,
	store.NotFoundTagConfigOrService: api.NotFoundTagConfigOrService,
	store.ExistReleasedConfig:        api.ExistReleasedConfig,
	store.DuplicateEntryErr:          api.ExistedResource,
}

// StoreCode2APICode store code to api code
func StoreCode2APICode(err error) uint32 {
	code := store.Code(err)
	apiCode, ok := storeCodeAPICodeMap[code]
	if ok {
		return apiCode
	}

	return api.StoreLayerException
}
