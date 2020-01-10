package alicloud

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAlicloudAcmConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudAcmConfigurationCreate,
		Read:   resourceAlicloudAcmConfigurationRead,
		Update: resourceAlicloudAcmConfigurationUpdate,
		Delete: resourceAlicloudAcmConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"data_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 256),
				ForceNew:     true,
			},
			"group": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 128),
				ForceNew:     true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceAlicloudAcmConfigurationCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	tenant := d.Get("namespace").(string)
	dataID := d.Get("data_id").(string)
	group := d.Get("group").(string)
	content := d.Get("content").(string)

	status, err := acmPublishConfigurations(tenant, dataID, group, content, client.AccessKey, client.SecretKey, client.SecurityToken)
	if err != nil {
		return WrapError(err)
	}

	if status != 200 {
		return WrapError(errors.New("Error"))
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", tenant, group, dataID))

	return resourceAlicloudAcmConfigurationUpdate(d, meta)
}

func resourceAlicloudAcmConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("content") {
		client := meta.(*connectivity.AliyunClient)
		tenant := d.Get("namespace").(string)
		dataID := d.Get("data_id").(string)
		group := d.Get("group").(string)
		content := d.Get("content").(string)
		status, err := acmPublishConfigurations(tenant, dataID, group, content, client.AccessKey, client.SecretKey, client.SecurityToken)
		if err != nil {
			return WrapError(err)
		}
		if status != 200 {
			return WrapError(errors.New("Error"))
		}
	}

	return resourceAlicloudAcmConfigurationRead(d, meta)
}

func resourceAlicloudAcmConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	tenant := d.Get("namespace").(string)
	dataID := d.Get("data_id").(string)
	group := d.Get("group").(string)
	content, status, err := acmGetConfiguration(tenant, dataID, group, client.AccessKey, client.SecretKey, client.SecurityToken)
	if err != nil {
		return WrapError(err)
	}

	if status == 200 {
		d.Set("content", content)
	} else if status == 404 {
		d.SetId("")
	} else {
		return errors.New("Error")
	}

	return nil
}

func resourceAlicloudAcmConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	tenant := d.Get("namespace").(string)
	dataID := d.Get("data_id").(string)
	group := d.Get("group").(string)
	status, err := acmDeleteConfiguration(tenant, dataID, group, client.AccessKey, client.SecretKey, client.SecurityToken)
	if err != nil {
		return WrapError(err)
	}

	if status == 400 {
		return WrapError(errors.New("Not found"))
	} else if status != 200 {
		return WrapError(errors.New("Error"))
	}

	return nil
}
