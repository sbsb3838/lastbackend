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

package ingress

import (
	"github.com/lastbackend/lastbackend/pkg/util/http"
	"github.com/lastbackend/lastbackend/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path: "/cluster/ingress", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressListH},
	{Path: "/cluster/ingress/{ingress}", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressInfoH},
	{Path: "/cluster/ingress/{ingress}/spec", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressGetSpecH},
	{Path: "/cluster/ingress/{ingress}", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressRemoveH},
	{Path: "/cluster/ingress/{ingress}", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressConnectH},
	{Path: "/cluster/ingress/{ingress}/meta", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressSetMetaH},
	{Path: "/cluster/ingress/{ingress}/status", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressSetStatusH},
	{Path: "/cluster/ingress/{ingress}/status/route/{pod}", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Authenticate}, Handler: IngressSetRouteStatusH},
}
