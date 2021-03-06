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

package distribution

import (
	"context"
	"fmt"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/log"
	"github.com/lastbackend/lastbackend/pkg/storage"
	"github.com/lastbackend/lastbackend/pkg/storage/store"
	"strings"
)

const (
	logRoutePrefix = "distribution:route"
)

type IRoute interface {
	Get(namespace, name string) (*types.Route, error)
	ListSpec() (map[string]*types.RouteSpec, error)
	ListByNamespace(namespace string) (map[string]*types.Route, error)
	Create(namespace *types.Namespace, opts *types.RouteCreateOptions) (*types.Route, error)
	Update(route *types.Route, opts *types.RouteUpdateOptions) (*types.Route, error)
	SetStatus(route *types.Route, status *types.RouteStatus) error
	Remove(route *types.Route) error
}

type Route struct {
	context context.Context
	storage storage.Storage
}

func (n *Route) Get(namespace, name string) (*types.Route, error) {

	log.V(logLevel).Debug("%s:get:> get route by id %s/%s", logRoutePrefix, namespace, name)

	item, err := n.storage.Route().Get(n.context, namespace, name)
	if err != nil {

		if err.Error() == store.ErrEntityNotFound {
			log.V(logLevel).Warnf("%s:get:> in namespace %s by name %s not found", logRoutePrefix, namespace, name)
			return nil, nil
		}

		log.V(logLevel).Errorf("%s:get:> in namespace %s by name %s error: %s", logRoutePrefix, namespace, name, err.Error())
		return nil, err
	}

	return item, nil
}

func (n *Route) ListSpec() (map[string]*types.RouteSpec, error) {

	log.V(logLevel).Debugf("%s:listspec:> list specs", logRoutePrefix)

	item, err := n.storage.Route().ListSpec(n.context)
	if err != nil {
		log.V(logLevel).Errorf("%s:listspec:> error: %s", logRoutePrefix, err.Error())
		return nil, err
	}

	return item, nil
}

func (n *Route) ListByNamespace(namespace string) (map[string]*types.Route, error) {

	log.V(logLevel).Debug("%s:listbynamespace:> list route", logRoutePrefix)

	items, err := n.storage.Route().ListByNamespace(n.context, namespace)
	if err != nil {
		log.V(logLevel).Error("%s:listbynamespace:> list route err: %s", logRoutePrefix, err.Error())
		return items, err
	}

	log.V(logLevel).Debugf("%s:listbynamespace:> list route result: %d", logRoutePrefix, len(items))

	return items, nil
}

func (n *Route) Create(namespace *types.Namespace, opts *types.RouteCreateOptions) (*types.Route, error) {

	log.V(logLevel).Debugf("%s:create:> create route %#v", logRoutePrefix, opts)

	route := new(types.Route)
	route.Meta.SetDefault()
	route.Meta.Name = opts.Name
	route.Meta.Namespace = namespace.Meta.Name
	route.Meta.Security = opts.Security
	route.SelfLink()

	route.Status.State = types.StateInitialized

	route.Spec.Domain = fmt.Sprintf("%s.%s", strings.ToLower(opts.Name), strings.ToLower(opts.Domain))
	route.Spec.Rules = make([]*types.RouteRule, 0)
	for _, rule := range opts.Rules {
		route.Spec.Rules = append(route.Spec.Rules, &types.RouteRule{
			Service:  rule.Service,
			Endpoint: rule.Endpoint,
			Port:     rule.Port,
			Path:     rule.Path,
		})
	}

	if err := n.storage.Route().Insert(n.context, route); err != nil {
		log.V(logLevel).Errorf("%s:create:> insert route err: %s", logRoutePrefix, err.Error())
		return nil, err
	}

	return route, nil
}

func (n *Route) Update(route *types.Route, opts *types.RouteUpdateOptions) (*types.Route, error) {

	log.V(logLevel).Debugf("%s:update:> update route %s", logRoutePrefix, route.Meta.Name)

	route.Meta.Security = opts.Security
	route.Status.State = types.StateProvision
	route.Spec.Rules = make([]*types.RouteRule, 0)
	for _, rule := range opts.Rules {
		route.Spec.Rules = append(route.Spec.Rules, &types.RouteRule{
			Endpoint: rule.Endpoint,
			Port:     rule.Port,
			Path:     rule.Path,
		})
	}

	if err := n.storage.Route().Update(n.context, route); err != nil {
		log.V(logLevel).Errorf("%s:update:> update route err: %s", logRoutePrefix, err.Error())
		return nil, err
	}

	return route, nil
}

func (n *Route) SetStatus(route *types.Route, status *types.RouteStatus) error {

	if route == nil {
		log.V(logLevel).Warnf("%s:setstatus:> invalid argument %v", logRoutePrefix, route)
		return nil
	}

	log.V(logLevel).Debugf("%s:setstate:> set state route %s -> %#v", logRoutePrefix, route.Meta.Name, status)

	route.Status = *status
	if err := n.storage.Route().SetStatus(n.context, route); err != nil {
		log.Errorf("%s:setstatus:> pod set status err: %s", err.Error())
		return err
	}

	return nil
}

func (n *Route) Remove(route *types.Route) error {

	log.V(logLevel).Debugf("%s:remove:> remove route %#v", logRoutePrefix, route)

	if err := n.storage.Route().Remove(n.context, route); err != nil {
		log.V(logLevel).Errorf("%s:remove:> remove route  err: %s", logRoutePrefix, err.Error())
		return err
	}

	return nil
}

func NewRouteModel(ctx context.Context, stg storage.Storage) IRoute {
	return &Route{ctx, stg}
}
