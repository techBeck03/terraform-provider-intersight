package intersight

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexClusterRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"capacity_runway": {
				Description: "The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%.\nDefault value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_name": {
				Description: "The name of this HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"cluster_purpose": {
				Description: "This can be a Storage or Compute cluster. A storage cluster contains storage nodes that are used to persist data. A compute cluster contains compute nodes that are used for executing business logic.\n* `Storage` - Cluster of storage nodes used to persist data.\n* `Compute` - Cluster of compute nodes used to execute business logic.\n* `Unknown` - This cluster type is Unknown. Expect Compute or Storage as valid values.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_type": {
				Description: "The storage type of this cluster (All Flash or Hybrid).",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"cluster_uuid": {
				Description: "The unique identifier for this HyperFlex cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"compute_node_count": {
				Description: "The number of compute nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"converged_node_count": {
				Description: "The number of converged nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"deployment_type": {
				Description: "The deployment type of the HyperFlex cluster.\nThe cluster can have one of the following configurations:\n1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.\n2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.\n3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.\nIf the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,\nthe deployment type is set as 'NA' (not available).\n* `NA` - The deployment type of the HyperFlex cluster is not available.\n* `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.\n* `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.\n* `Edge` - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_id": {
				Description: "The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"drive_type": {
				Description: "The type of the drives used for storage in this cluster.\n* `NA` - The drive type of the HyperFlex cluster is not available.\n* `All-Flash` - Indicates that this HyperFlex cluster contains flash drives only.\n* `Hybrid` - Indicates that this HyperFlex cluster contains both flash and hard disk drives.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"flt_aggr": {
				Description: "The number of yellow (warning) and red (critical) alarms stored as an aggregate.\nThe first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"hx_version": {
				Description: "The HyperFlex Data or Application Platform version of this cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hxdp_build_version": {
				Description: "The version and build number of the HyperFlex Data Platform for this cluster.\nAfter a cluster upgrade, this version string will be updated on the next inventory cycle to reflect\nthe newly installed version.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"hypervisor_type": {
				Description: "Identifies the broad type of the underlying hypervisor.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hypervisor_version": {
				Description: "The version of hypervisor running on this cluster.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"identity": {
				Description: "The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference).",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "The user-provided name for this cluster to facilitate identification.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Description: "Cluster health status as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"storage_capacity": {
				Description: "The storage capacity in this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"storage_node_count": {
				Description: "The number of storage nodes that belong to this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"storage_utilization": {
				Description: "The storage utilization is computed based on total capacity and current capacity utilization.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"total_cores": {
				Description: "Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"utilization_percentage": {
				Description: "The storage utilization percentage is computed based on total capacity and current capacity utilization.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"utilization_trend_percentage": {
				Description: "The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.",
				Type:        schema.TypeFloat,
				Optional:    true,
				Computed:    true,
			},
			"vm_count": {
				Description: "The number of virtual machines present on this cluster.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"account_moid": {
					Description: "The Account ID for this managed object.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"alarm": {
						Description: "An array of relationships to hyperflexAlarm resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"alarm_summary": {
						Description: "The summary of alarm counts based on the severity types (Critical or Warning).",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"critical": {
									Description: "The count of alarms that have severity type Critical.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"warning": {
									Description: "The count of alarms that have severity type Warning.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
					},
					"ancestors": {
						Description: "An array of relationships to moBaseMo resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"associated_profile": {
						Description: "A reference to a policyAbstractProfile resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"capacity_runway": {
						Description: "The number of days remaining before the cluster's storage utilization reaches the recommended capacity limit of 76%.\nDefault value is math.MaxInt32 to indicate that the capacity runway is \"Unknown\" for a cluster that is not connected or with not sufficient data.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"child_clusters": {
						Description: "An array of relationships to hyperflexBaseCluster resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster_name": {
						Description: "The name of this HyperFlex cluster.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"cluster_purpose": {
						Description: "This can be a Storage or Compute cluster. A storage cluster contains storage nodes that are used to persist data. A compute cluster contains compute nodes that are used for executing business logic.\n* `Storage` - Cluster of storage nodes used to persist data.\n* `Compute` - Cluster of compute nodes used to execute business logic.\n* `Unknown` - This cluster type is Unknown. Expect Compute or Storage as valid values.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster_type": {
						Description: "The storage type of this cluster (All Flash or Hybrid).",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"cluster_uuid": {
						Description: "The unique identifier for this HyperFlex cluster.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"compute_node_count": {
						Description: "The number of compute nodes that belong to this cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"converged_node_count": {
						Description: "The number of converged nodes that belong to this cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"create_time": {
						Description: "The time when this managed object was created.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"deployment_type": {
						Description: "The deployment type of the HyperFlex cluster.\nThe cluster can have one of the following configurations:\n1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.\n2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.\n3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.\nIf the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,\nthe deployment type is set as 'NA' (not available).\n* `NA` - The deployment type of the HyperFlex cluster is not available.\n* `Datacenter` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.\n* `Stretched Cluster` - The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.\n* `Edge` - The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device_id": {
						Description: "The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"domain_group_moid": {
						Description: "The DomainGroup ID for this managed object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"drive_type": {
						Description: "The type of the drives used for storage in this cluster.\n* `NA` - The drive type of the HyperFlex cluster is not available.\n* `All-Flash` - Indicates that this HyperFlex cluster contains flash drives only.\n* `Hybrid` - Indicates that this HyperFlex cluster contains both flash and hard disk drives.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"flt_aggr": {
						Description: "The number of yellow (warning) and red (critical) alarms stored as an aggregate.\nThe first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"health": {
						Description: "A reference to a hyperflexHealth resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"hx_version": {
						Description: "The HyperFlex Data or Application Platform version of this cluster.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"hxdp_build_version": {
						Description: "The version and build number of the HyperFlex Data Platform for this cluster.\nAfter a cluster upgrade, this version string will be updated on the next inventory cycle to reflect\nthe newly installed version.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"hypervisor_type": {
						Description: "Identifies the broad type of the underlying hypervisor.\n* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.\n* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.\n* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.\n* `Unknown` - The hypervisor running on the HyperFlex cluster is not known.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"hypervisor_version": {
						Description: "The version of hypervisor running on this cluster.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"identity": {
						Description: "The internally generated identity of this cluster. This entity is not manipulated by users. It aids in uniquely identifying the cluster object. In case of VMware, this is a MOR (managed object reference).",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"license": {
						Description: "A reference to a hyperflexLicense resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"memory_capacity": {
						Description: "The capacity and usage information for memory on this cluster.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"capacity": {
									Description: "The total memory capacity of the entity in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Free memory (bytes) that is unused and available for allocation, as a point-in-time snapshot. The available memory capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used memory capacity is also reported.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Memory (bytes) that has been already used up, as a point-in-time snapshot.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"mod_time": {
						Description: "The time when this managed object was last modified.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"name": {
						Description: "The user-provided name for this cluster to facilitate identification.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"nodes": {
						Description: "An array of relationships to hyperflexNode resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"owners": {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Schema{
							Type: schema.TypeString}},
					"parent": {
						Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"parent_cluster": {
						Description: "A reference to a hyperflexBaseCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"permission_resources": {
						Description: "An array of relationships to moBaseMo resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"processor_capacity": {
						Description: "The capacity and usage information for CPU power on this cluster.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"capacity": {
									Description: "Total capacity of the entity in MHz.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"free": {
									Description: "Free CPU capacity in MHz, as a point-in-time snapshot. The available CPU capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used CPU capacity is also reported.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"used": {
									Description: "Used CPU capacity of the entity in MHz, as a point-in-time snapshot.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
						Computed: true,
					},
					"registered_device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"shared_scope": {
						Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"status": {
						Description: "Cluster health status as reported by the hypervisor platform.\n* `Unknown` - Entity status is unknown.\n* `Degraded` - State is degraded, and might impact normal operation of the entity.\n* `Critical` - Entity is in a critical state, impacting operations.\n* `Ok` - Entity status is in a stable state, operating normally.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"storage_capacity": {
						Description: "The storage capacity in this cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"storage_containers": {
						Description: "An array of relationships to storageHyperFlexStorageContainer resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"storage_node_count": {
						Description: "The number of storage nodes that belong to this cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"storage_utilization": {
						Description: "The storage utilization is computed based on total capacity and current capacity utilization.",
						Type:        schema.TypeFloat,
						Optional:    true,
						Computed:    true,
					},
					"summary": {
						Description: "The summary of HyperFlex cluster health, storage, and number of nodes.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"active_nodes": {
									Description: "The number of nodes currently participating in the storage cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"address": {
									Description: "The data IP address of the HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"boottime": {
									Description: "The time taken during last cluster startup in seconds.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"cluster_access_policy": {
									Description: "The cluster access policy for the HyperFlex cluster. An access policy of 'STRICT' means that the cluster becomes readonly once any fragment of data is reduced to one copy. 'LENIENT' means that the cluster stays in read-write mode even if any fragment of data is reduced to one copy.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"compression_savings": {
									Description: "The percentage of storage space saved using data compression.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
								"data_replication_compliance": {
									Description: "The compliance with the data replication factor set for the HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"data_replication_factor": {
									Description: "The number of data copies retained by the HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"deduplication_savings": {
									Description: "The percentage of storage space saved using data deduplication.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
								"downtime": {
									Description: "The amount of time the HyperFlex cluster has been offline.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"free_capacity": {
									Description: "The amount of storage capacity currently not in use, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"healing_info": {
									Description: "The information about the HyperFlex cluster auto-healing process.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"estimated_completion_time_in_seconds": {
												Description: "The estimated time in seconds it will take to complete the auto-healing process.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"in_progress": {
												Description: "The status of the cluster's auto-healing process. If 'true', auto-healing is in progress for the cluster.",
												Type:        schema.TypeBool,
												Optional:    true,
												Computed:    true,
											},
											"messages": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Schema{
													Type: schema.TypeString}},
											"messages_iterator": {
												Description: "The current message describing the auto-healing process of the cluster.",
												Type:        schema.TypeMap,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												}, Optional: true,
												Computed: true,
											},
											"messages_size": {
												Description: "The number of elements in the messages collection.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"percent_complete": {
												Description: "The progress of the auto-healing process as a percentage.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"name": {
									Description: "The name of the HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"resiliency_details": {
									Description: "The details about the resiliency health of the cluster. Includes information about the cluster healing status and the storage cluster health.",
									Type:        schema.TypeMap,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									}, Optional: true,
									Computed: true,
								},
								"resiliency_details_size": {
									Description: "The number of elements in the resiliency details property.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"resiliency_info": {
									Description: "The information about the storage cluster resiliency. The resiliency info has details about the health status and number of device failures tolerable.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"hdd_failures_tolerable": {
												Description: "The number of persistent storage device failures tolerable before the storage cluster becomes offline.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"messages": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Schema{
													Type: schema.TypeString}},
											"messages_iterator": {
												Description: "The current message describing the auto-healing process of the cluster.",
												Type:        schema.TypeMap,
												Elem: &schema.Schema{
													Type: schema.TypeString,
												}, Optional: true,
												Computed: true,
											},
											"messages_size": {
												Description: "The number of elements in the messages collection.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"node_failures_tolerable": {
												Description: "The number of node failures tolerable before the storage cluster becomes offline.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"ssd_failures_tolerable": {
												Description: "The number of caching device failures tolerable before the storage cluster becomes offline.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"state": {
												Description: "The resiliency state of the cluster. The resiliency status is 'HEALTHY' if there are no failures and the storage cluster is fully operational. The resiliency status is 'WARNING' when the cluster has experienced failures that may adversely affect the cluster. It is 'UNKNOWN' if the cluster is offline or if the state cannot be determined.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"space_status": {
									Description: "The space utilization status of the HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"state": {
									Description: "The operational state of the HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"total_capacity": {
									Description: "The total amount of storage capacity available for the HyperFlex cluster, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"total_savings": {
									Description: "The percentage of storage space saved in total.",
									Type:        schema.TypeFloat,
									Optional:    true,
									Computed:    true,
								},
								"uptime": {
									Description: "The amount of time the HyperFlex cluster has been running since last startup.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"used_capacity": {
									Description: "The amount of storage capacity in use, represented in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"total_cores": {
						Description: "Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"utilization_percentage": {
						Description: "The storage utilization percentage is computed based on total capacity and current capacity utilization.",
						Type:        schema.TypeFloat,
						Optional:    true,
						Computed:    true,
					},
					"utilization_trend_percentage": {
						Description: "The storage utilization trend percentage represents the trend in percentage computed using the first and last point from historical data.",
						Type:        schema.TypeFloat,
						Optional:    true,
						Computed:    true,
					},
					"version_context": {
						Description: "The versioning info for this managed object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"interested_mos": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"moid": {
												Description: "The Moid of the referenced REST resource.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the remote type referred by this relationship.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"selector": {
												Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
									Computed: true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"ref_mo": {
									Description: "A reference to the original Managed Object.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"moid": {
												Description: "The Moid of the referenced REST resource.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the remote type referred by this relationship.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"selector": {
												Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
								"timestamp": {
									Description: "The time this versioned Managed Object was created.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"nr_version": {
									Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"version_type": {
									Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"vm_count": {
						Description: "The number of virtual machines present on this cluster.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"volumes": {
						Description: "An array of relationships to storageHyperFlexVolume resources.",
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexClusterRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexCluster{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("capacity_runway"); ok {
		x := int64(v.(int))
		o.SetCapacityRunway(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("cluster_name"); ok {
		x := (v.(string))
		o.SetClusterName(x)
	}
	if v, ok := d.GetOk("cluster_purpose"); ok {
		x := (v.(string))
		o.SetClusterPurpose(x)
	}
	if v, ok := d.GetOk("cluster_type"); ok {
		x := int64(v.(int))
		o.SetClusterType(x)
	}
	if v, ok := d.GetOk("cluster_uuid"); ok {
		x := (v.(string))
		o.SetClusterUuid(x)
	}
	if v, ok := d.GetOk("compute_node_count"); ok {
		x := int64(v.(int))
		o.SetComputeNodeCount(x)
	}
	if v, ok := d.GetOk("converged_node_count"); ok {
		x := int64(v.(int))
		o.SetConvergedNodeCount(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("deployment_type"); ok {
		x := (v.(string))
		o.SetDeploymentType(x)
	}
	if v, ok := d.GetOk("device_id"); ok {
		x := (v.(string))
		o.SetDeviceId(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("drive_type"); ok {
		x := (v.(string))
		o.SetDriveType(x)
	}
	if v, ok := d.GetOk("flt_aggr"); ok {
		x := int64(v.(int))
		o.SetFltAggr(x)
	}
	if v, ok := d.GetOk("hx_version"); ok {
		x := (v.(string))
		o.SetHxVersion(x)
	}
	if v, ok := d.GetOk("hxdp_build_version"); ok {
		x := (v.(string))
		o.SetHxdpBuildVersion(x)
	}
	if v, ok := d.GetOk("hypervisor_type"); ok {
		x := (v.(string))
		o.SetHypervisorType(x)
	}
	if v, ok := d.GetOk("hypervisor_version"); ok {
		x := (v.(string))
		o.SetHypervisorVersion(x)
	}
	if v, ok := d.GetOk("identity"); ok {
		x := (v.(string))
		o.SetIdentity(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("storage_capacity"); ok {
		x := int64(v.(int))
		o.SetStorageCapacity(x)
	}
	if v, ok := d.GetOk("storage_node_count"); ok {
		x := int64(v.(int))
		o.SetStorageNodeCount(x)
	}
	if v, ok := d.GetOk("storage_utilization"); ok {
		x := v.(float32)
		o.SetStorageUtilization(x)
	}
	if v, ok := d.GetOk("total_cores"); ok {
		x := int64(v.(int))
		o.SetTotalCores(x)
	}
	if v, ok := d.GetOk("utilization_percentage"); ok {
		x := v.(float32)
		o.SetUtilizationPercentage(x)
	}
	if v, ok := d.GetOk("utilization_trend_percentage"); ok {
		x := v.(float32)
		o.SetUtilizationTrendPercentage(x)
	}
	if v, ok := d.GetOk("vm_count"); ok {
		x := int64(v.(int))
		o.SetVmCount(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexCluster object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexCluster: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexCluster: %s", responseErr.Error())
	}
	count := countResponse.HyperflexClusterList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexCluster data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexClusterResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexClusterList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexCluster: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexCluster: %s", responseErr.Error())
		}
		results := resMo.HyperflexClusterList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["alarm"] = flattenListHyperflexAlarmRelationship(s.GetAlarm(), d)

				temp["alarm_summary"] = flattenMapHyperflexAlarmSummary(s.GetAlarmSummary(), d)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["associated_profile"] = flattenMapPolicyAbstractProfileRelationship(s.GetAssociatedProfile(), d)
				temp["capacity_runway"] = (s.GetCapacityRunway())

				temp["child_clusters"] = flattenListHyperflexBaseClusterRelationship(s.GetChildClusters(), d)
				temp["class_id"] = (s.GetClassId())
				temp["cluster_name"] = (s.GetClusterName())
				temp["cluster_purpose"] = (s.GetClusterPurpose())
				temp["cluster_type"] = (s.GetClusterType())
				temp["cluster_uuid"] = (s.GetClusterUuid())
				temp["compute_node_count"] = (s.GetComputeNodeCount())
				temp["converged_node_count"] = (s.GetConvergedNodeCount())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["deployment_type"] = (s.GetDeploymentType())
				temp["device_id"] = (s.GetDeviceId())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["drive_type"] = (s.GetDriveType())
				temp["flt_aggr"] = (s.GetFltAggr())

				temp["health"] = flattenMapHyperflexHealthRelationship(s.GetHealth(), d)
				temp["hx_version"] = (s.GetHxVersion())
				temp["hxdp_build_version"] = (s.GetHxdpBuildVersion())
				temp["hypervisor_type"] = (s.GetHypervisorType())
				temp["hypervisor_version"] = (s.GetHypervisorVersion())
				temp["identity"] = (s.GetIdentity())

				temp["license"] = flattenMapHyperflexLicenseRelationship(s.GetLicense(), d)

				temp["memory_capacity"] = flattenMapVirtualizationMemoryCapacity(s.GetMemoryCapacity(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["nodes"] = flattenListHyperflexNodeRelationship(s.GetNodes(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["parent_cluster"] = flattenMapHyperflexBaseClusterRelationship(s.GetParentCluster(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["processor_capacity"] = flattenMapVirtualizationComputeCapacity(s.GetProcessorCapacity(), d)

				temp["registered_device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetRegisteredDevice(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["status"] = (s.GetStatus())
				temp["storage_capacity"] = (s.GetStorageCapacity())

				temp["storage_containers"] = flattenListStorageHyperFlexStorageContainerRelationship(s.GetStorageContainers(), d)
				temp["storage_node_count"] = (s.GetStorageNodeCount())
				temp["storage_utilization"] = (s.GetStorageUtilization())

				temp["summary"] = flattenMapHyperflexSummary(s.GetSummary(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["total_cores"] = (s.GetTotalCores())
				temp["utilization_percentage"] = (s.GetUtilizationPercentage())
				temp["utilization_trend_percentage"] = (s.GetUtilizationTrendPercentage())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vm_count"] = (s.GetVmCount())

				temp["volumes"] = flattenListStorageHyperFlexVolumeRelationship(s.GetVolumes(), d)
				hyperflexClusterResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexClusterResults))
	if err := d.Set("results", hyperflexClusterResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexClusterResults[0]["moid"].(string))
	return de
}
