// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepositoriesService_ListHookDeliveries(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks/1/deliveries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"cursor": "v1_12077215967"})
		fmt.Fprint(w, `[{"id":1}, {"id":2}]`)
	})

	opt := &ListCursorOptions{Cursor: "v1_12077215967"}

	ctx := context.Background()
	hooks, _, err := client.Repositories.ListHookDeliveries(ctx, "o", "r", 1, opt)
	if err != nil {
		t.Errorf("Repositories.ListHookDeliveries returned error: %v", err)
	}

	want := []*HookDelivery{{ID: Int64(1)}, {ID: Int64(2)}}
	if d := cmp.Diff(hooks, want); d != "" {
		t.Errorf("Repositories.ListHooks want (-), got (+):\n%s", d)
	}

	const methodName = "ListHookDeliveries"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Repositories.ListHookDeliveries(ctx, "\n", "\n", -1, opt)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Repositories.ListHookDeliveries(ctx, "o", "r", 1, opt)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestRepositoriesService_ListHookDeliveries_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Repositories.ListHookDeliveries(ctx, "%", "%", 1, nil)
	testURLParseError(t, err)
}

func TestRepositoriesService_GetHookDelivery(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/hooks/1/deliveries/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()
	hook, _, err := client.Repositories.GetHookDelivery(ctx, "o", "r", 1, 1)
	if err != nil {
		t.Errorf("Repositories.GetHookDelivery returned error: %v", err)
	}

	want := &HookDelivery{ID: Int64(1)}
	if !cmp.Equal(hook, want) {
		t.Errorf("Repositories.GetHookDelivery returned %+v, want %+v", hook, want)
	}

	const methodName = "GetHookDelivery"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Repositories.GetHookDelivery(ctx, "\n", "\n", -1, -1)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Repositories.GetHookDelivery(ctx, "o", "r", 1, 1)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestRepositoriesService_GetHookDelivery_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Repositories.GetHookDelivery(ctx, "%", "%", 1, 1)
	testURLParseError(t, err)
}
