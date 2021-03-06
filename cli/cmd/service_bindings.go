package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	scmodel "github.com/servicebroker/servicebroker/model/service_controller"
	"github.com/spf13/cobra"
)

const (
	SERVICE_BINDINGS_URL     = "/v2/service_bindings"
	SERVICE_BINDINGS_FMT_STR = "/v2/service_bindings/%s"
)

var (
	bindingParameters string
)

func init() {
	RootCmd.AddCommand(serviceBindingsCmd)
	serviceBindingsCmd.AddCommand(createServiceBindingsCmd)
	serviceBindingsCmd.AddCommand(listServiceBindingsCmd)
	serviceBindingsCmd.AddCommand(describeServiceBindingsCmd)
	createServiceBindingsCmd.Flags().StringVarP(&bindingParameters, "parameters", "p", "", "Parameters to pass to the service broker for binding creation (must be JSON object)")

}

var serviceBindingsCmd = &cobra.Command{
	Use:   "service-bindings",
	Short: "Manage service bindings",
	Long:  "Manage service bindings",
}

var createServiceBindingsCmd = &cobra.Command{
	Use:   "create <APP_NAME> <SERVICE_NAME>",
	Short: "Create a service binding",
	Long:  "Create a service binding",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("need APP_NAME and SERVICE_NAME")
		}
		app := args[0]
		service := args[1]
		ServiceInstanceGUID, err := fetchServiceInstanceGUID(service)
		if err != nil {
			return err
		}
		req := scmodel.CreateServiceBindingRequest{
			AppName:             app,
			ServiceInstanceName: service,
			ServiceInstanceGUID: ServiceInstanceGUID,
			Parameters:          nil,
		}
		if len(bindingParameters) > 0 {
			var m interface{}
			err := json.Unmarshal([]byte(bindingParameters), &m)
			if err != nil {
				return err
			}
			req.Parameters = m.(map[string]interface{})
		}
		body, err := json.Marshal(req)
		if err != nil {
			return err
		}
		return callService(SERVICE_BINDINGS_URL, "POST", "create service binding", ioutil.NopCloser(bytes.NewReader(body)))
	},
}

var listServiceBindingsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all service bindings",
	Long:  "List all service bindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		foo, err := fetchPrettyBindings()
		if err != nil {
			return err
		}
		fmt.Printf(foo)
		return nil
	},
}

var describeServiceBindingsCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe a service binding",
	Long:  "Describe all service binding",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("need GUID of the service instance to describe")
		}
		foo := fmt.Sprintf(SERVICE_BINDINGS_FMT_STR, args[0])
		u := fmt.Sprintf("%s%s", controller, foo)
		i, err := callHttp(u, "GET", "describe binding", nil)
		if err != nil {
			return err
		}
		var sb scmodel.ServiceBinding
		err = json.Unmarshal([]byte(i), &sb)
		if err != nil {
			return err
		}

		bar := fmt.Sprintf(SERVICE_INSTANCES_FMT_STR, sb.ServiceInstanceGUID)
		u = fmt.Sprintf("%s%s", controller, bar)
		i, err = callHttp(u, "GET", "describe service instance", nil)
		if err != nil {
			return err
		}
		var si scmodel.ServiceInstance
		err = json.Unmarshal([]byte(i), &si)
		if err != nil {
			return err
		}

		fmt.Printf("%s -> %s\n\t%+v\n", sb.ServiceInstanceName, si.Name, sb.Parameters)
		return nil
	},
}
