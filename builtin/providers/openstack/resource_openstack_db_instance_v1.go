package openstack

import (
	"fmt"
	"log"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/db/v1/instances"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatabaseInstanceV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseInstanceV1Create,
		Read:   resourceDatabaseInstanceV1Read,
		Delete: resourceDatabaseInstanceV1Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"flavor_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Computed:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_FLAVOR_ID", nil),
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDatabaseInstanceV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseInstanceClient, err := config.databaseInstanceV1Client(GetRegion(d))

	if err != nil {
		return fmt.Errorf("Error creating RS cloud database client: %s", err)
	}

	createOpts := &instances.CreateOpts{
		FlavorRef: d.Get("flavor_id").(string),
		Size: d.Get("size").(int),
		Name: d.Get("name").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	instance, err := instances.Create(databaseInstanceClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating RS cloud database instance: %s", err)
	}
	log.Printf("[INFO] instance ID: %s", instance.ID)

	// Wait for the volume to become available.
	log.Printf(
		"[DEBUG] Waiting for volume (%s) to become available",
		instance.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"BUILD"},
		Target:     []string{"ACTIVE"},
		Refresh:    InstanceV1StateRefreshFunc(databaseInstanceClient, instance.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for instance (%s) to become ready: %s",
			instance.ID, err)
	}

	// Store the ID now
	d.SetId(instance.ID)

	return resourceDatabaseInstanceV1Read(d, meta)
}

func resourceDatabaseInstanceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseInstanceClient, err := config.databaseInstanceV1Client(GetRegion(d))
	if err != nil {
		return fmt.Errorf("Error creating OpenStack cloud database client: %s", err)
	}

	instance, err := instances.Get(databaseInstanceClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "instance")
	}

	log.Printf("[DEBUG] Retrieved instance %s: %+v", d.Id(), instance)

	d.Set("size", instance.Size)
	d.Set("name", instance.Name)
	d.Set("flavor_id", instance.FlavorRef)
	d.Set("region", GetRegion(d))

	return nil
}

func resourceDatabaseInstanceV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	databaseInstanceClient, err := config.databaseInstanceV1Client(GetRegion(d))
	if err != nil {
		return fmt.Errorf("Error creating RS cloud instance client: %s", err)
	}

	instance, err := instances.Get(databaseInstanceClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "instance")
	}

	log.Printf("[DEBUG] Deleting cloud database instance %s", d.Id())
	err = instances.Delete(databaseInstanceClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting cloud database instance: %s", err)
	}

	// Wait for the volume to delete before moving on.
	log.Printf("[DEBUG] Waiting for volume (%s) to delete", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "SHUTOFF"},
		Target:     []string{"DELETED", "SOFT_DELETED"},
		Refresh:    InstanceV1StateRefreshFunc(databaseInstanceClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for instance (%s) to delete: %s",
			d.Id(), err)
	}

	d.SetId("")
	return nil
}

// InstanceV1StateRefreshFunc returns a resource.StateRefreshFunc that is used to watch
// an cloud database instance.
func InstanceV1StateRefreshFunc(client *gophercloud.ServiceClient, instanceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		i, err := instances.Get(client, instanceID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return i, "deleted", nil
			}
			return nil, "", err
		}

		if i.Status == "error" {
			return i, i.Status, fmt.Errorf("There was an error creating the instance.")
		}

		return i, i.Status, nil
	}
}