// +build examples

/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package enterprisebillingunitsv1_test

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/enterprisebillingunitsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

const externalConfigFile = "../enterprise_billing_units_v1.env"

var (
	enterpriseBillingUnitsService *enterprisebillingunitsv1.EnterpriseBillingUnitsV1
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`EnterpriseBillingUnitsV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(enterprisebillingunitsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			enterpriseBillingUnitsServiceOptions := &enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{}

			enterpriseBillingUnitsService, err = enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(enterpriseBillingUnitsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(enterpriseBillingUnitsService).ToNot(BeNil())
		})
	})

	Describe(`EnterpriseBillingUnitsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingUnit request example`, func() {
			// begin-get_billing_unit

			getBillingUnitOptions := enterpriseBillingUnitsService.NewGetBillingUnitOptions(
				"testString",
			)

			billingUnit, response, err := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(billingUnit, "", "  ")
			fmt.Println(string(b))

			// end-get_billing_unit

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingUnit).ToNot(BeNil())

		})
		It(`ListBillingUnits request example`, func() {
			// begin-list_billing_units

			listBillingUnitsOptions := enterpriseBillingUnitsService.NewListBillingUnitsOptions()

			billingUnitsList, response, err := enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(billingUnitsList, "", "  ")
			fmt.Println(string(b))

			// end-list_billing_units

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingUnitsList).ToNot(BeNil())

		})
		It(`ListBillingOptions request example`, func() {
			// begin-list_billing_options

			listBillingOptionsOptions := enterpriseBillingUnitsService.NewListBillingOptionsOptions(
				"testString",
			)

			billingOption, response, err := enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(billingOption, "", "  ")
			fmt.Println(string(b))

			// end-list_billing_options

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOption).ToNot(BeNil())

		})
		It(`GetCreditPools request example`, func() {
			// begin-get_credit_pools

			getCreditPoolsOptions := enterpriseBillingUnitsService.NewGetCreditPoolsOptions(
				"testString",
			)

			creditPoolsList, response, err := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(creditPoolsList, "", "  ")
			fmt.Println(string(b))

			// end-get_credit_pools

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsList).ToNot(BeNil())

		})
	})
})
