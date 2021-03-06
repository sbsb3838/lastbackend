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

package pod

import (
	"github.com/gorilla/mux"
	"github.com/lastbackend/lastbackend/pkg/distribution/errors"
	"github.com/lastbackend/lastbackend/pkg/log"
	"github.com/lastbackend/lastbackend/pkg/node/envs"
	"github.com/lastbackend/lastbackend/pkg/node/runtime/pod"
	"net/http"
)

const logLevel = 2

func PodGetH(w http.ResponseWriter, _ *http.Request) {

	log.V(logLevel).Debug("node:http:pod:get:> get pod info")
}

// PodLogsH handler streams pod logs into response writer
func PodLogsH(w http.ResponseWriter, r *http.Request) {

	log.V(logLevel).Debug("node:http:pod:get:> get pod logs")

	var (
		c      = mux.Vars(r)["container"]
		p      = envs.Get().GetState().Pods().GetPod(mux.Vars(r)["pod"])
		notify = w.(http.CloseNotifier).CloseNotify()
		done   = make(chan bool, 1)
	)

	go func() {
		<-notify
		log.Debug("HTTP connection just closed.")
		done <- true
	}()

	if p == nil {
		log.Errorf("node:http:pod:get:> pod not found")
		errors.New("pod").NotFound().Http(w)
		return
	}

	if _, ok := p.Containers[c]; !ok {
		log.Errorf("node:http:pod:get:> container not found")
		errors.New("pod").NotFound().Http(w)
		return
	}

	if err := pod.Logs(r.Context(), c, true, w, done); err != nil {
		log.Errorf("node:http:pod:get:> get pod logs err: %s", err.Error())
	}

	return
}
