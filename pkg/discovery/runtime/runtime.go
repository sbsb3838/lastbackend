//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package runtime

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/log"
	"github.com/lastbackend/lastbackend/pkg/discovery/runtime/endpoint"
)

const (
	logPrefix = "discovery:runtime"
)

type Runtime struct {
	ctx context.Context
}

func (r *Runtime) Restore() {
	log.Debugf("%s:restore:> restore init", logPrefix)
}

func (r *Runtime) Loop() {
	log.Debugf("%s:restore:> watch endpoint start", logPrefix)
	endpoint.Watch(r.ctx)
}

func NewRuntime(ctx context.Context) *Runtime {
	return &Runtime{ctx: ctx}
}
