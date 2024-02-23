// Code generated by ifacemaker; DO NOT EDIT.

package hcloud

import (
	"context"
)

// IFirewallClient ...
type IFirewallClient interface {
	// GetByID retrieves a Firewall by its ID. If the Firewall does not exist, nil is returned.
	GetByID(ctx context.Context, id int64) (*Firewall, *Response, error)
	// GetByName retrieves a Firewall by its name. If the Firewall does not exist, nil is returned.
	GetByName(ctx context.Context, name string) (*Firewall, *Response, error)
	// Get retrieves a Firewall by its ID if the input can be parsed as an integer, otherwise it
	// retrieves a Firewall by its name. If the Firewall does not exist, nil is returned.
	Get(ctx context.Context, idOrName string) (*Firewall, *Response, error)
	// List returns a list of Firewalls for a specific page.
	//
	// Please note that filters specified in opts are not taken into account
	// when their value corresponds to their zero value or when they are empty.
	List(ctx context.Context, opts FirewallListOpts) ([]*Firewall, *Response, error)
	// All returns all Firewalls.
	All(ctx context.Context) ([]*Firewall, error)
	// AllWithOpts returns all Firewalls for the given options.
	AllWithOpts(ctx context.Context, opts FirewallListOpts) ([]*Firewall, error)
	// Create creates a new Firewall.
	Create(ctx context.Context, opts FirewallCreateOpts) (FirewallCreateResult, *Response, error)
	// Update updates a Firewall.
	Update(ctx context.Context, firewall *Firewall, opts FirewallUpdateOpts) (*Firewall, *Response, error)
	// Delete deletes a Firewall.
	Delete(ctx context.Context, firewall *Firewall) (*Response, error)
	// SetRules sets the rules of a Firewall.
	SetRules(ctx context.Context, firewall *Firewall, opts FirewallSetRulesOpts) ([]*Action, *Response, error)
	ApplyResources(ctx context.Context, firewall *Firewall, resources []FirewallResource) ([]*Action, *Response, error)
	RemoveResources(ctx context.Context, firewall *Firewall, resources []FirewallResource) ([]*Action, *Response, error)
}
