package finvoice

import (
	"encoding/xml"
	"log"
	"os"

	"golang.org/x/net/html/charset"
)

func FromFile(filePath string) (*Finvoice, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer xmlFile.Close()

	// Parse the XML file
	var finvoice Finvoice
	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&finvoice)
	if err != nil {
		return nil, err
	}

	return &finvoice, nil
}

type Finvoice struct {
	SellerPartyDetails struct {
		SellerOrganisationName    string `xml:"SellerOrganisationName"`
		SellerOrganisationTaxCode string `xml:"SellerOrganisationTaxCode"`
		SellerCode                struct {
			IdentifierType string `xml:"IdentifierType,attr"`
			Value          string `xml:",chardata"`
		} `xml:"SellerCode"`
		SellerPostalAddressDetails struct {
			SellerStreetName         string `xml:"SellerStreetName"`
			SellerTownName           string `xml:"SellerTownName"`
			SellerPostCodeIdentifier string `xml:"SellerPostCodeIdentifier"`
			CountryCode              string `xml:"CountryCode"`
		} `xml:"SellerPostalAddressDetails"`
	} `xml:"SellerPartyDetails"`
	SellerOrganisationUnitNumber string `xml:"SellerOrganisationUnitNumber"`
	SellerInformationDetails     struct {
		SellerAccountDetails struct {
			SellerAccountID struct {
				IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
				Value                    string `xml:",chardata"`
			} `xml:"SellerAccountID"`
			SellerBic struct {
				IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
				Value                    string `xml:",chardata"`
			} `xml:"SellerBic"`
		} `xml:"SellerAccountDetails"`
	} `xml:"SellerInformationDetails"`
	InvoiceRecipientLanguageCode string `xml:"InvoiceRecipientLanguageCode"`
	BuyerPartyDetails            struct {
		BuyerPartyIdentifier      string `xml:"BuyerPartyIdentifier"`
		BuyerOrganisationName     string `xml:"BuyerOrganisationName"`
		BuyerPostalAddressDetails struct {
			BuyerStreetName         string `xml:"BuyerStreetName"`
			BuyerTownName           string `xml:"BuyerTownName"`
			BuyerPostCodeIdentifier string `xml:"BuyerPostCodeIdentifier"`
			CountryCode             string `xml:"CountryCode"`
			CountryName             string `xml:"CountryName"`
		} `xml:"BuyerPostalAddressDetails"`
	} `xml:"BuyerPartyDetails"`
	BuyerCommunicationDetails struct {
		BuyerPhoneNumberIdentifier string `xml:"BuyerPhoneNumberIdentifier"`
	} `xml:"BuyerCommunicationDetails"`
	DeliveryDetails struct {
		DeliveryPeriodDetails struct {
			StartDate struct {
				Format string `xml:"Format,attr"`
				Value  string `xml:",chardata"`
			} `xml:"StartDate"`
			EndDate struct {
				Format string `xml:"Format,attr"`
				Value  string `xml:",chardata"`
			} `xml:"EndDate"`
		} `xml:"DeliveryPeriodDetails"`
	} `xml:"DeliveryDetails"`
	InvoiceRows []struct {
		ArticleName      string `xml:"ArticleName"`
		InvoicedQuantity struct {
			QuantityUnitCode string `xml:"QuantityUnitCode,attr"`
			Value            string `xml:",chardata"`
		} `xml:"InvoicedQuantity"`
		UnitPriceAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"UnitPriceAmount"`
		UnitPriceVatIncludedAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"UnitPriceVatIncludedAmount"`
		RowDeliveryDate struct {
			Format string `xml:"Format,attr"`
			Value  string `xml:",chardata"`
		} `xml:"RowDeliveryDate"`
		RowVatRatePercent    string `xml:"RowVatRatePercent"`
		RowVatExcludedAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"RowVatExcludedAmount"`
		RowAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"RowAmount"`
	} `xml:"InvoiceRow"`
	InvoiceDetails struct {
		InvoiceTypeCode struct {
			CodeListAgencyIdentifier string `xml:"CodeListAgencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"InvoiceTypeCode"`
		InvoiceTypeText string `xml:"InvoiceTypeText"`
		OriginCode      string `xml:"OriginCode"`
		InvoiceNumber   string `xml:"InvoiceNumber"`
		InvoiceDate     struct {
			Format string `xml:"Format,attr"`
			Value  string `xml:",chardata"`
		} `xml:"InvoiceDate"`
		InvoicingPeriodStartDate struct {
			Format string `xml:"Format,attr"`
			Value  string `xml:",chardata"`
		} `xml:"InvoicingPeriodStartDate"`
		InvoicingPeriodEndDate struct {
			Format string `xml:"Format,attr"`
			Value  string `xml:",chardata"`
		} `xml:"InvoicingPeriodEndDate"`
		SellersBuyerIdentifier        string `xml:"SellersBuyerIdentifier"`
		InvoiceTotalVatExcludedAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"InvoiceTotalVatExcludedAmount"`
		InvoiceTotalVatAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"InvoiceTotalVatAmount"`
		InvoiceTotalVatIncludedAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"InvoiceTotalVatIncludedAmount"`
		CreditLimitAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"CreditLimitAmount"`
		CreditInterestPercent string `xml:"CreditInterestPercent"`
		OperationLimitAmount  struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"OperationLimitAmount"`
		MonthlyAmount struct {
			AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
			Value                    string `xml:",chardata"`
		} `xml:"MonthlyAmount"`
		VatSpecificationDetails struct {
			VatBaseAmount struct {
				AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
				Value                    string `xml:",chardata"`
			} `xml:"VatBaseAmount"`
			VatRatePercent string `xml:"VatRatePercent"`
			VatRateAmount  struct {
				AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
				Value                    string `xml:",chardata"`
			} `xml:"VatRateAmount"`
		} `xml:"VatSpecificationDetails"`
		InvoiceFreeText     string `xml:"InvoiceFreeText"`
		PaymentTermsDetails struct {
			InvoiceDueDate struct {
				Format string `xml:"Format,attr"`
				Value  string `xml:",chardata"`
			} `xml:"InvoiceDueDate"`
			PaymentOverDueFineDetails struct {
				// Add your fields here
			} `xml:"PaymentOverDueFineDetails"`
		} `xml:"PaymentTermsDetails"`
	} `xml:"InvoiceDetails"`
}
