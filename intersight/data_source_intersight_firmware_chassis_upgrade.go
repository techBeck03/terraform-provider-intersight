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

func dataSourceFirmwareChassisUpgrade() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFirmwareChassisUpgradeRead,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"create_time": {
				Description: "The time when this managed object was created.",
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
			"skip_estimate_impact": {
				Description: "User has the option to skip the estimate impact calculation.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"status": {
				Description: "Status of the upgrade operation.\n* `NONE` - Upgrade status is not populated.\n* `IN_PROGRESS` - The upgrade is in progress.\n* `SUCCESSFUL` - The upgrade successfully completed.\n* `FAILED` - The upgrade shows failed status.\n* `TERMINATED` - The upgrade has been terminated.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"upgrade_type": {
				Description: "Desired upgrade mode to choose either direct download based upgrade or network share upgrade.\n* `direct_upgrade` - Upgrade mode is direct download.\n* `network_upgrade` - Upgrade mode is network upgrade.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"results": {
				Type:     schema.TypeList,
				Elem:     &schema.Resource{Schema: resourceFirmwareChassisUpgrade().Schema},
				Computed: true,
			}},
	}
}

func dataSourceFirmwareChassisUpgradeRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.FirmwareChassisUpgrade{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("skip_estimate_impact"); ok {
		x := (v.(bool))
		o.SetSkipEstimateImpact(x)
	}
	if v, ok := d.GetOk("status"); ok {
		x := (v.(string))
		o.SetStatus(x)
	}
	if v, ok := d.GetOk("upgrade_type"); ok {
		x := (v.(string))
		o.SetUpgradeType(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of FirmwareChassisUpgrade object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareChassisUpgradeList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of FirmwareChassisUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of FirmwareChassisUpgrade: %s", responseErr.Error())
	}
	count := countResponse.FirmwareChassisUpgradeList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for FirmwareChassisUpgrade data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var firmwareChassisUpgradeResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.FirmwareApi.GetFirmwareChassisUpgradeList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching FirmwareChassisUpgrade: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching FirmwareChassisUpgrade: %s", responseErr.Error())
		}
		results := resMo.FirmwareChassisUpgradeList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["chassis"] = flattenMapEquipmentChassisRelationship(s.GetChassis(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)

				temp["direct_download"] = flattenMapFirmwareDirectDownload(s.GetDirectDownload(), d)

				temp["distributable"] = flattenMapFirmwareDistributableRelationship(s.GetDistributable(), d)
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["exclude_component_list"] = (s.GetExcludeComponentList())

				temp["file_server"] = flattenMapSoftwarerepositoryFileServer(s.GetFileServer(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())

				temp["network_share"] = flattenMapFirmwareNetworkShare(s.GetNetworkShare(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)

				temp["release"] = flattenMapSoftwarerepositoryReleaseRelationship(s.GetRelease(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["skip_estimate_impact"] = (s.GetSkipEstimateImpact())
				temp["status"] = (s.GetStatus())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["upgrade_impact"] = flattenMapFirmwareUpgradeImpactStatusRelationship(s.GetUpgradeImpact(), d)

				temp["upgrade_status"] = flattenMapFirmwareUpgradeStatusRelationship(s.GetUpgradeStatus(), d)
				temp["upgrade_type"] = (s.GetUpgradeType())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				firmwareChassisUpgradeResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(firmwareChassisUpgradeResults))
	if err := d.Set("results", firmwareChassisUpgradeResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(firmwareChassisUpgradeResults[0]["moid"].(string))
	return de
}
