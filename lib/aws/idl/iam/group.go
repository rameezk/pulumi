// Copyright 2017 Pulumi, Inc. All rights reserved.

package iam

import (
	"github.com/pulumi/coconut/pkg/resource/idl"
)

// Group is an AWS Identity and Access Management (IAM) group.
type Group struct {
	idl.NamedResource
	// groupName is a name for the IAM group.  If you don't specify a name, a unique physical ID will be generated.
	//
	// Important: if you specify a name, you cannot perform updates that require replacement of this resource.  You can
	// perform updates that require no or some interruption.  If you must replace this resource, specify a new name.
	//
	// If you specify a new name, you must specify the `CAPABILITY_NAMED_IAM` value to acknowledge your capabilities.
	//
	// Warning: Naming an IAM resource can cause an unrecoverable error if you reuse the same code in multiple regions.
	// To prevent this, create a name that includes the region name itself, to create a region-specific name.
	GroupName *string `coco:"groupName,replaces,optional"`
	// managedPolicies is one or more managed policies to attach to this role.
	ManagedPolicies *[]*Policy `coco:"managedPolicies,optional"`
	// path is the path associated with this role.  For more information about paths, see
	// http://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html#Identifiers_FriendlyNames.
	Path *string `coco:"path,optional"`
	// policies are the policies to associate with this role.
	Policies *InlinePolicy `coco:"policies,optional"`
}