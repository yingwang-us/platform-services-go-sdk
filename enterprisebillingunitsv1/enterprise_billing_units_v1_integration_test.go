// +build integration

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
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/enterprisebillingunitsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the enterprisebillingunitsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`EnterpriseBillingUnitsV1 Integration Tests`, func() {

	const externalConfigFile = "../enterprise_billing_units.env"

	var (
		err                           error
		enterpriseBillingUnitsService *enterprisebillingunitsv1.EnterpriseBillingUnitsV1
		serviceURL                    string
		config                        map[string]string

		enterpriseID   string
		accountID      string
		accountGroupID string
		billingUnitID  string
		billingMonth   string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(enterprisebillingunitsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			enterpriseID = config["ENTERPRISE_ID"]
			Expect(enterpriseID).ToNot(BeEmpty())

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			accountGroupID = config["ACCOUNT_GROUP_ID"]
			Expect(accountGroupID).ToNot(BeEmpty())

			billingUnitID = config["BILLING_UNIT_ID"]
			Expect(billingUnitID).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			enterpriseBillingUnitsServiceOptions := &enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{}

			enterpriseBillingUnitsService, err = enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(enterpriseBillingUnitsServiceOptions)

			Expect(err).To(BeNil())
			Expect(enterpriseBillingUnitsService).ToNot(BeNil())
			Expect(enterpriseBillingUnitsService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`GetBillingUnit - Get billing unit by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingUnit(getBillingUnitOptions *GetBillingUnitOptions)`, func() {

			getBillingUnitOptions := &enterprisebillingunitsv1.GetBillingUnitOptions{
				BillingUnitID: &billingUnitID,
			}

			billingUnit, response, err := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingUnit).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetBillingUnit response:\n%s\n", common.ToJSON(billingUnit))
		})
	})

	Describe(`ListBillingUnits - List billing units`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBillingUnits(enterpriseID)`, func() {

			listBillingUnitsOptions := &enterprisebillingunitsv1.ListBillingUnitsOptions{
				EnterpriseID: &enterpriseID,
			}

			billingUnitsList, response, err := enterpriseBillingUnitsService.ListBillingUnits(listBillingUnitsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingUnitsList).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "ListBillingUnits(enterpriseID) response:\n%s\n", common.ToJSON(billingUnitsList))
		})
	})

	Describe(`ListBillingOptions - List billing options`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBillingOptions(listBillingOptionsOptions *ListBillingOptionsOptions)`, func() {

			listBillingOptionsOptions := &enterprisebillingunitsv1.ListBillingOptionsOptions{
				BillingUnitID: &billingUnitID,
			}

			billingOptionsList, response, err := enterpriseBillingUnitsService.ListBillingOptions(listBillingOptionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsList).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "ListBillingOptions() response:\n%s\n", common.ToJSON(billingOptionsList))
		})
	})

	Describe(`GetCreditPools - Get credit pools`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPools(getCreditPoolsOptions *GetCreditPoolsOptions)`, func() {

			getCreditPoolsOptions := &enterprisebillingunitsv1.GetCreditPoolsOptions{
				BillingUnitID: &billingUnitID,
				//			Date:          &billingMonth,
				Type: core.StringPtr("PLATFORM"),
			}

			creditPoolsList, response, err := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsList).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetCreditPools() response:\n%s\n", common.ToJSON(creditPoolsList))
		})
	})
})
