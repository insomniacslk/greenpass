package greenpass

import (
	// used to embed the JSON schema
	_ "embed"
	"fmt"
	"strings"

	"github.com/kr/pretty"
	"github.com/xeipuuv/gojsonschema"
)

// Generate types from the updated JSON specifications by running `go generate`.
//
//go:generate go run tools/generate_types/main.go -o types.go
//go:generate gofmt -w types.go
//go:generate go run tools/generate_jsonschema/main.go -o DCC.combined-schema.json

//go:embed DCC.combined-schema.json
var digitalCertificateJSONSchema string

// CovidCertificate represents an EU Digital Covid Certificate, see
// specification at
// https://ec.europa.eu/health/sites/default/files/ehealth/docs/covid-certificate_json_specification_en.pdf
// and schema validator at
// https://github.com/ehn-dcc-development/ehn-dcc-schema/blob/release/1.3.0/DCC.schema.json
type CovidCertificate struct {
	// version string information, semver
	Ver string `json:"ver"`
	// person name information
	Nam PersonName `json:"nam"`
	// date of birth
	DoB string `json:"dob"`
	// vaccination dose, test, and recovery information. Only one of them must
	// be present.
	V []VaccinationDose   `json:"v,omitempty"`
	T []TestResult        `json:"t,omitempty"`
	R []RecoveryStatement `json:"r,omitempty"`
}

// Validate validates the COVID certificate against the official JSON schema.
func (c *CovidCertificate) Validate() error {
	// validate the digital COVID certificate using the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(digitalCertificateJSONSchema)
	documentLoader := gojsonschema.NewGoLoader(c)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if result.Valid() {
		return nil
	}
	errors := make([]string, 0, len(result.Errors()))
	for idx, e := range result.Errors() {
		errors = append(errors, fmt.Sprintf("%d) %s", idx+1, e.String()))
	}
	pretty.Println(c)
	return fmt.Errorf("found %d errors: %v", len(result.Errors()), strings.Join(errors, "; "))
}

// Summary returns a multiline string that summarizes the content of the
//CovidCertificate object.
func (c *CovidCertificate) Summary() string {
	var status string
	vaccineFound := false
	if c.V != nil && len(c.V) == 1 {
		vaccineFound = true
	}
	if vaccineFound {
		dn := int(c.V[0].Dn)
		sd := int(c.V[0].Sd)
		if dn == 0 {
			status = "not vaccinated"
		} else if dn < sd {
			status = fmt.Sprintf("partially vaccinated (%d of %d doses)", dn, sd)
		} else if dn == sd {
			status = "fully vaccinated"
		} else {
			status = fmt.Sprintf("fully vaccinated with booster dose (%d of %d doses)", dn, sd)
		}
	}
	ret := fmt.Sprintf("Name                    : %s %s\n", c.Nam.Gn, c.Nam.Fn)
	ret += fmt.Sprintf("Date of Birth           : %s\n", c.DoB)
	ret += fmt.Sprintf("Status                  : %s\n", status)
	if vaccineFound {
		ret += fmt.Sprintf("Vaccination date        : %s\n", c.V[0].Dt)
		ret += fmt.Sprintf("Country                 : %s\n", c.V[0].Co)
		ret += fmt.Sprintf("Vaccine Prophylaxis     : %s\n", c.V[0].Vp)
		ret += fmt.Sprintf("Vaccine Product         : %s\n", c.V[0].Mp)
		ret += fmt.Sprintf("Marketing Authorization : %s\n", c.V[0].Ma)
	}
	return ret
}

// PersonName describes a person name according to the DCC spec.
type PersonName struct {
	// Surname
	Fn string `json:"fn"`
	// Standardised surname
	Fnt string `json:"fnt"`
	// Forename
	Gn string `json:"gn"`
	// Standardised given name
	Gnt string `json:"gnt"`
}

// VaccinationDose describes a vaccination dose according to the DCC spec.
type VaccinationDose struct {
	// Disease or agent targeted. COVID-19 is 840539006.
	Tg DiseaseAgentTargeted `json:"tg"`
	// Vaccine or prophylaxis. Example: "1119349007" is a SARS-COV-2 mRNA
	// vaccine.
	Vp VaccineProphylaxis `json:"vp"`
	// Vaccine medicinal product. See vaccine-medicinal-product.json in the
	// spec. Example: "EU/1/20/1528" is Comirnaty (Pfizer).
	Mp VaccineMedicinalProduct `json:"mp"`
	// Marketing authorization holder or manufacturer. See vaccine-mah-manf.json
	// in the spec. Example "ORG-100030215" is "Biontech Manufacturing GmbH"
	Ma VaccineMahManf `json:"ma"`
	// Dose sequence number. 1 for first dose, 2 for second dose, etc.
	Dn float64 `json:"dn"`
	// Overall number of doses for this vaccine. 1 for 1-dose vaccine, 2 for
	// 2-dose vaccine, 3 for booster.
	Sd float64 `json:"sd"`
	// Date of vaccination of the described dose.
	Dt string `json:"dt"`
	// Country where the vaccine was administered, ISO3166 (2-letter) coded.
	// Example: "IE" is Ireland. See country-2-codes.json from the spec.
	Co Country2Codes `json:"co"`
	// Certificate issuer, name of the organization that issued the certificate.
	Is string `json:"is"`
	// Unique Certificate Identifier, prefixed with "URN:UVCI:". See UVCI in
	// https://ec.europa.eu/health/sites/default/files/ehealth/docs/vaccination-proof_interoperability-guidelines_en.pdf
	Ci string `json:"ci"`
}

type TestResult struct {
}

type RecoveryStatement struct {
}
