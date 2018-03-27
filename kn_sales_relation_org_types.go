// DO NOT EDIT: generated by github.com/tim-online/go-afas-profit-rest/generate

package afas

import (
	"encoding/json"

	"github.com/aodin/date"
	"github.com/cockroachdb/apd"
	"github.com/cydev/zero"
)

// KnSalesRelationOrg
type KnSalesRelationOrg struct {
	NummerDebiteur                        string         `json:"DbId,omitempty"`                              // Nummer debiteur
	VoorkeurIbanNummer                    string         `json:"Iban,omitempty"`                              // Voorkeur Iban nummer
	VoorkeurBankgironummer                string         `json:"BaAc,omitempty"`                              // Voorkeur bank-/gironummer
	VoorkeurBankgironummerLandCode        string         `json:"CoId,omitempty"`                              // Voorkeur bank-/gironummer land code
	IsDebiteur                            bool           `json:"IsDb,omitempty"`                              // Is debiteur
	VoorkeurTegenrekening                 string         `json:"ToAc,omitempty"`                              // Voorkeur tegenrekening
	BtwNummer                             string         `json:"VaId,omitempty"`                              // Btw-nummer
	Betalingsvoorwaarde                   string         `json:"PaCd,omitempty"`                              // Betalingsvoorwaarde
	Vertegenwoordiger                     string         `json:"VeId,omitempty"`                              // Vertegenwoordiger
	Taal                                  string         `json:"LgId,omitempty"`                              // Taal
	Valuta                                string         `json:"CuId"`                                        // Valuta
	AfwijkendeAanmaningsset               int            `json:"DsId,omitempty"`                              // Afwijkende aanmaningsset
	Verantwoordelijke                     string         `json:"EmId,omitempty"`                              // Verantwoordelijke
	BtwPlicht                             string         `json:"VaDu,omitempty"`                              // Btw-plicht
	Profiel                               string         `json:"PfId,omitempty"`                              // Profiel
	Regelkorting                          *apd.Decimal   `json:"PrLi,omitempty"`                              // % Regelkorting
	Factuurkorting                        *apd.Decimal   `json:"PrFc,omitempty"`                              // Factuurkorting (%)
	Kredietbeperking                      *apd.Decimal   `json:"ClPc,omitempty"`                              // Kredietbeperking (%)
	Betalingskorting                      *apd.Decimal   `json:"PrPt,omitempty"`                              // Betalingskorting (%)
	Kredietlimiet                         *apd.Decimal   `json:"Krli,omitempty"`                              // Kredietlimiet
	FacturerenAan                         string         `json:"FaTo,omitempty"`                              // Factureren aan
	Vervoerder                            string         `json:"TrPt,omitempty"`                              // Vervoerder
	PrioriteitLevering                    int            `json:"PrDl,omitempty"`                              // Prioriteit levering
	PrijzenVan                            string         `json:"PrVn,omitempty"`                              // Prijzen van
	VoorkeurPrijslijst                    string         `json:"PrLs,omitempty"`                              // Voorkeur prijslijst
	VoorkeurMagazijn                      string         `json:"VkMa,omitempty"`                              // Voorkeur magazijn
	GeblokkeerdVoorLevering               bool           `json:"Bl,omitempty"`                                // Geblokkeerd voor levering
	VolledigBlokkerenNietMeerZichtbaar    bool           `json:"BlTl,omitempty"`                              // Volledig blokkeren, niet meer zichtbaar
	Hoofddeclarant                        string         `json:"LDId,omitempty"`                              // Hoofddeclarant
	AfwijkendeBtwTariefgroep              bool           `json:"VaYN,omitempty"`                              // Afwijkende btw-tariefgroep
	AanmaningVerzenden                    bool           `json:"DuYN,omitempty"`                              // Aanmaning verzenden
	CodeGroepsadministratie               string         `json:"VaIg,omitempty"`                              // Code groepsadministratie
	StatusBewaking                        string         `json:"VaGu,omitempty"`                              // Status bewaking
	Kortingsgroep                         string         `json:"DsGr,omitempty"`                              // Kortingsgroep
	Nettoprijs                            bool           `json:"NtPr,omitempty"`                              // Nettoprijs
	PrijsInclBtw                          bool           `json:"VtIn,omitempty"`                              // Prijs incl. btw
	Factuurtekst                          string         `json:"InTx,omitempty"`                              // Factuurtekst
	FactuurGeheelVerdichten               bool           `json:"CITo,omitempty"`                              // Factuur geheel verdichten
	TekstBijGeheelVerdichten              string         `json:"TxTc,omitempty"`                              // Tekst bij geheel verdichten
	StriktMaximum                         bool           `json:"StMa,omitempty"`                              // Strikt maximum
	MaximumFactuurbedrag                  *apd.Decimal   `json:"MaIn,omitempty"`                              // Maximum factuurbedrag
	StriktMinimum                         bool           `json:"StMi,omitempty"`                              // Strikt minimum
	Minimumfactuurbedrag                  *apd.Decimal   `json:"MiIn,omitempty"`                              // Minimumfactuurbedrag
	Afrondingsmethode                     string         `json:"RoOf,omitempty"`                              // Afrondingsmethode
	Declarant                             string         `json:"DeId,omitempty"`                              // Declarant
	Incassospecificatie                   bool           `json:"PaSp,omitempty"`                              // Incassospecificatie
	AutomatischIncasseren                 bool           `json:"AuPa,omitempty"`                              // Automatisch incasseren
	Verdichten                            bool           `json:"PaCo,omitempty"`                              // Verdichten
	EenmaligeIncassomachtigingVereist     bool           `json:"SiPA,omitempty"`                              // Eenmalige incassomachtiging vereist
	WaarschuwingBijOrder                  bool           `json:"WaOr,omitempty"`                              // Waarschuwing bij order
	TekstWaarschuwing                     string         `json:"WaTx,omitempty"`                              // Tekst waarschuwing
	CBSTypen                              string         `json:"CsTy,omitempty"`                              // CBS-typen
	Opmerking                             []byte         `json:"Rm,omitempty"`                                // Opmerking
	Contactpersoon                        int            `json:"CtP1,omitempty"`                              // Contactpersoon
	ExtraContactpersoon                   int            `json:"CtP2,omitempty"`                              // Extra contactpersoon
	Verzamelrekening                      string         `json:"ColA,omitempty"`                              // Verzamelrekening
	KlantSinds                            *date.Date     `json:"CsDa,omitempty"`                              // Klant sinds
	AangebrachtDoor                       string         `json:"BcBy,omitempty"`                              // Aangebracht door
	Leveringsconditie                     string         `json:"DeCo"`                                        // Leveringsconditie
	VoorkeurContact                       int            `json:"CtI1,omitempty"`                              // Voorkeur contact
	VoorkeurVerstrekkingswijze            string         `json:"InPv"`                                        // Voorkeur verstrekkingswijze
	Ordersortering                        string         `json:"SoId,omitempty"`                              // Ordersortering
	TypeBarcode                           string         `json:"VaBc,omitempty"`                              // Type barcode
	Barcode                               string         `json:"BaCo,omitempty"`                              // Barcode
	AdresseringEDIPakbonConformEDIFactuur bool           `json:"EDDn,omitempty"`                              // Adressering EDI-pakbon conform EDI-factuur
	VaA1                                  string         `json:"VaA1,omitempty"`                              // VaA1
	VaA2                                  string         `json:"VaA2,omitempty"`                              // VaA2
	VaA3                                  string         `json:"VaA3,omitempty"`                              // VaA3
	VaA4                                  string         `json:"VaA4,omitempty"`                              // VaA4
	VaA5                                  string         `json:"VaA5,omitempty"`                              // VaA5
	Wachtwoord                            string         `json:"Pwrd,omitempty"`                              // Wachtwoord
	Activeringscode                       string         `json:"AtCd,omitempty"`                              // Activeringscode
	Accounttype                           string         `json:"AcTp,omitempty"`                              // Accounttype
	VerwerkingOrder                       string         `json:"OrPr,omitempty"`                              // Verwerking order
	Assortiment                           string         `json:"AsGr,omitempty"`                              // Assortiment
	AfwijkenAssortimentToestaan           bool           `json:"AsYN,omitempty"`                              // Afwijken assortiment toestaan
	StudentNummer                         string         `json:"OINr,omitempty"`                              // Student nummer
	IncassowijzeSEPA                      string         `json:"VaDt,omitempty"`                              // Incassowijze SEPA
	BedrijfsIdEVerbinding                 string         `json:"EnId,omitempty"`                              // Bedrijfs-Id eVerbinding
	TypeVerkooprelatie                    string         `json:"VaTp,omitempty"`                              // Type verkooprelatie
	VerwijderingsbijdrageToepassen        bool           `json:"ReCo,omitempty"`                              // Verwijderingsbijdrage toepassen
	EenmaligeDebiteur                     bool           `json:"U36AC2FE34594A8F2D14DB88F867DEF52,omitempty"` // Eenmalige debiteur?
	Machtiging                            bool           `json:"U47C81A5D426695598163F6B1959DEEB1,omitempty"` // Machtiging?
	KnOrganisation                        KnOrganisation `json:"KnOrganisation"`                              // KnOrganisation

}

func (k KnSalesRelationOrg) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnSalesRelationOrg

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				DBID:    k.DBID(),
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}
func (k KnSalesRelationOrg) DBIDField() string {
	return "NummerDebiteur"
}

func (k KnSalesRelationOrg) DBID() string {
	return k.NummerDebiteur
}

func (k KnSalesRelationOrg) JSONFields() []string {
	return []string{
		"DbId",
		"Iban",
		"BaAc",
		"CoId",
		"IsDb",
		"ToAc",
		"VaId",
		"PaCd",
		"VeId",
		"LgId",
		"CuId",
		"DsId",
		"EmId",
		"VaDu",
		"PfId",
		"PrLi",
		"PrFc",
		"ClPc",
		"PrPt",
		"Krli",
		"FaTo",
		"TrPt",
		"PrDl",
		"PrVn",
		"PrLs",
		"VkMa",
		"Bl",
		"BlTl",
		"LDId",
		"VaYN",
		"DuYN",
		"VaIg",
		"VaGu",
		"DsGr",
		"NtPr",
		"VtIn",
		"InTx",
		"CITo",
		"TxTc",
		"StMa",
		"MaIn",
		"StMi",
		"MiIn",
		"RoOf",
		"DeId",
		"PaSp",
		"AuPa",
		"PaCo",
		"SiPA",
		"WaOr",
		"WaTx",
		"CsTy",
		"Rm",
		"CtP1",
		"CtP2",
		"ColA",
		"CsDa",
		"BcBy",
		"DeCo",
		"CtI1",
		"InPv",
		"SoId",
		"VaBc",
		"BaCo",
		"EDDn",
		"VaA1",
		"VaA2",
		"VaA3",
		"VaA4",
		"VaA5",
		"Pwrd",
		"AtCd",
		"AcTp",
		"OrPr",
		"AsGr",
		"AsYN",
		"OINr",
		"VaDt",
		"EnId",
		"VaTp",
		"ReCo",
		"U36AC2FE34594A8F2D14DB88F867DEF52",
		"U47C81A5D426695598163F6B1959DEEB1",
		"",
	}
}

func (k KnSalesRelationOrg) JSONObjects() []string {
	return []string{
		"KnOrganisation",
	}
}

// KnOrganisation
type KnOrganisation struct {
	PostadresIsAdres         bool              `json:"PadAdr"`               // Postadres is adres
	Autonummering            bool              `json:"AutoNum"`              // Autonummering
	OrganisatieVergelijkenOp string            `json:"MatchOga"`             // Organisatie vergelijken op
	Organisatiepersoonintern int               `json:"BcId"`                 // Organisatie/persoon (intern)
	Nummer                   string            `json:"BcCo"`                 // Nummer
	Zoeknaam                 string            `json:"SeNm,omitempty"`       // Zoeknaam
	Naam                     string            `json:"Nm"`                   // Naam
	Rechtsvorm               string            `json:"ViLe,omitempty"`       // Rechtsvorm
	Branche                  string            `json:"ViLb,omitempty"`       // Branche
	KvKNummer                string            `json:"CcNr,omitempty"`       // KvK-nummer
	DatumKvK                 *date.Date        `json:"CcDa,omitempty"`       // Datum KvK
	Naamstatutair            string            `json:"NmRg,omitempty"`       // Naam (statutair)
	Vestigingstatutair       string            `json:"RsRg,omitempty"`       // Vestiging (statutair)
	Titelaanhef              string            `json:"TtId,omitempty"`       // Titel/aanhef
	Briefaanhef              string            `json:"LeHe,omitempty"`       // Briefaanhef
	OrganisatorischeEenheid  string            `json:"OuId,omitempty"`       // Organisatorische eenheid
	TelefoonnrWerk           string            `json:"TeNr,omitempty"`       // Telefoonnr. werk
	FaxWerk                  string            `json:"FaNr,omitempty"`       // Fax werk
	MobielWerk               string            `json:"MbNr,omitempty"`       // Mobiel werk
	EMailWerk                string            `json:"EmAd,omitempty"`       // E-mail werk
	Website                  string            `json:"HoPa,omitempty"`       // Website
	Correspondentie          bool              `json:"Corr,omitempty"`       // Correspondentie
	Voorkeursmedium          string            `json:"ViMd,omitempty"`       // Voorkeursmedium
	Opmerking                []byte            `json:"Re,omitempty"`         // Opmerking
	Fiscaalnummer            string            `json:"FiNr,omitempty"`       // Fiscaalnummer
	Status                   string            `json:"StId,omitempty"`       // Status
	SocialeNetwerken         string            `json:"SocN,omitempty"`       // Sociale netwerken
	Facebook                 string            `json:"Face,omitempty"`       // Facebook
	LinkedIn                 string            `json:"Link,omitempty"`       // LinkedIn
	Twitter                  string            `json:"Twtr,omitempty"`       // Twitter
	OnderdeelVanOrganisatie  string            `json:"BcPa,omitempty"`       // Onderdeel van organisatie
	NaamBestand              string            `json:"FileName,omitempty"`   // Naam bestand
	Afbeelding               []byte            `json:"FileStream,omitempty"` // Afbeelding
	KnBankAccount            KnBankAccount     `json:"KnBankAccount"`        // KnBankAccount
	KnBasicAddressAdr        KnBasicAddressAdr `json:"KnBasicAddressAdr"`    // KnBasicAddressAdr
	KnBasicAddressPad        KnBasicAddressPad `json:"KnBasicAddressPad"`    // KnBasicAddressPad
	KnContact                KnContact         `json:"KnContact"`            // KnContact

}

func (k KnOrganisation) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnOrganisation

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}

func (k KnOrganisation) JSONFields() []string {
	return []string{
		"PadAdr",
		"AutoNum",
		"MatchOga",
		"BcId",
		"BcCo",
		"SeNm",
		"Nm",
		"ViLe",
		"ViLb",
		"CcNr",
		"CcDa",
		"NmRg",
		"RsRg",
		"TtId",
		"LeHe",
		"OuId",
		"TeNr",
		"FaNr",
		"MbNr",
		"EmAd",
		"HoPa",
		"Corr",
		"ViMd",
		"Re",
		"FiNr",
		"StId",
		"SocN",
		"Face",
		"Link",
		"Twtr",
		"BcPa",
		"FileName",
		"FileStream",
		"",
		"",
		"",
		"",
	}
}

func (k KnOrganisation) JSONObjects() []string {
	return []string{
		"KnBankAccount",
		"KnBasicAddressAdr",
		"KnBasicAddressPad",
		"KnContact",
	}
}

// KnBasicAddressAdr
type KnBasicAddressAdr struct {
	Land                                                   string    `json:"CoId"`           // Land
	Postbusadres                                           bool      `json:"PbAd"`           // Postbusadres
	ToevVoorStraat                                         string    `json:"StAd,omitempty"` // Toev. voor straat
	Straat                                                 string    `json:"Ad,omitempty"`   // Straat
	Huisnummer                                             int       `json:"HmNr,omitempty"` // Huisnummer
	ToevAanHuisnr                                          string    `json:"HmAd,omitempty"` // Toev. aan huisnr.
	Postcode                                               string    `json:"ZpCd,omitempty"` // Postcode
	Woonplaats                                             string    `json:"Rs,omitempty"`   // Woonplaats
	AdresToevoeging                                        string    `json:"AdAd,omitempty"` // Adres toevoeging
	IngangsdatumAdreswijzigingwordtGenegeerdBijEersteDatum date.Date `json:"BeginDate"`      // Ingangsdatum adreswijziging (wordt genegeerd bij eerste datum)
	ZoekWoonplaatsBijPostcode                              bool      `json:"ResZip"`         // Zoek woonplaats bij postcode

}

func (k KnBasicAddressAdr) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnBasicAddressAdr

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}

func (k KnBasicAddressAdr) JSONFields() []string {
	return []string{
		"CoId",
		"PbAd",
		"StAd",
		"Ad",
		"HmNr",
		"HmAd",
		"ZpCd",
		"Rs",
		"AdAd",
		"BeginDate",
		"ResZip",
	}
}

func (k KnBasicAddressAdr) JSONObjects() []string {
	return []string{}
}

// KnContact
type KnContact struct {
	PostadresIsAdres     bool              `json:"PadAdr"`            // Postadres is adres
	SoortContact         string            `json:"ViKc,omitempty"`    // Soort Contact
	Afdeling             string            `json:"ExAd,omitempty"`    // Afdeling
	Functie              string            `json:"ViFu,omitempty"`    // Functie
	FunctieOpVisitekaart string            `json:"FuDs,omitempty"`    // Functie op visitekaart
	Correspondentie      bool              `json:"Corr,omitempty"`    // Correspondentie
	Voorkeursmedium      string            `json:"ViMd,omitempty"`    // Voorkeursmedium
	TelefoonnrWerk       string            `json:"TeNr,omitempty"`    // Telefoonnr. werk
	FaxWerk              string            `json:"FaNr,omitempty"`    // Fax werk
	MobielWerk           string            `json:"MbNr,omitempty"`    // Mobiel werk
	EMailWerk            string            `json:"EmAd,omitempty"`    // E-mail werk
	Website              string            `json:"HoPa,omitempty"`    // Website
	Toelichting          []byte            `json:"Re,omitempty"`      // Toelichting
	Geblokkeerd          bool              `json:"Bl,omitempty"`      // Geblokkeerd
	TavRegel             string            `json:"AtLn,omitempty"`    // T.a.v. regel
	Briefaanhef          string            `json:"LeHe,omitempty"`    // Briefaanhef
	SocialeNetwerken     string            `json:"SocN,omitempty"`    // Sociale netwerken
	Facebook             string            `json:"Face,omitempty"`    // Facebook
	LinkedIn             string            `json:"Link,omitempty"`    // LinkedIn
	Twitter              string            `json:"Twtr,omitempty"`    // Twitter
	Kerstkaart           bool              `json:"U006,omitempty"`    // Kerstkaart
	MailingOntvangen     bool              `json:"U007,omitempty"`    // Mailing ontvangen
	KnBasicAddressAdr    KnBasicAddressAdr `json:"KnBasicAddressAdr"` // KnBasicAddressAdr
	KnBasicAddressPad    KnBasicAddressPad `json:"KnBasicAddressPad"` // KnBasicAddressPad
	KnPerson             KnPerson          `json:"KnPerson"`          // KnPerson

}

func (k KnContact) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnContact

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}

func (k KnContact) JSONFields() []string {
	return []string{
		"PadAdr",
		"ViKc",
		"ExAd",
		"ViFu",
		"FuDs",
		"Corr",
		"ViMd",
		"TeNr",
		"FaNr",
		"MbNr",
		"EmAd",
		"HoPa",
		"Re",
		"Bl",
		"AtLn",
		"LeHe",
		"SocN",
		"Face",
		"Link",
		"Twtr",
		"U006",
		"U007",
		"",
		"",
		"",
	}
}

func (k KnContact) JSONObjects() []string {
	return []string{
		"KnBasicAddressAdr",
		"KnBasicAddressPad",
		"KnPerson",
	}
}

// KnBasicAddressPad
type KnBasicAddressPad struct {
	Land                                                   string    `json:"CoId"`           // Land
	Postbusadres                                           bool      `json:"PbAd"`           // Postbusadres
	ToevVoorStraat                                         string    `json:"StAd,omitempty"` // Toev. voor straat
	Straat                                                 string    `json:"Ad,omitempty"`   // Straat
	Huisnummer                                             int       `json:"HmNr,omitempty"` // Huisnummer
	ToevAanHuisnr                                          string    `json:"HmAd,omitempty"` // Toev. aan huisnr.
	Postcode                                               string    `json:"ZpCd,omitempty"` // Postcode
	Woonplaats                                             string    `json:"Rs,omitempty"`   // Woonplaats
	AdresToevoeging                                        string    `json:"AdAd,omitempty"` // Adres toevoeging
	IngangsdatumAdreswijzigingwordtGenegeerdBijEersteDatum date.Date `json:"BeginDate"`      // Ingangsdatum adreswijziging (wordt genegeerd bij eerste datum)
	ZoekWoonplaatsBijPostcode                              bool      `json:"ResZip"`         // Zoek woonplaats bij postcode

}

func (k KnBasicAddressPad) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnBasicAddressPad

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}

func (k KnBasicAddressPad) JSONFields() []string {
	return []string{
		"CoId",
		"PbAd",
		"StAd",
		"Ad",
		"HmNr",
		"HmAd",
		"ZpCd",
		"Rs",
		"AdAd",
		"BeginDate",
		"ResZip",
	}
}

func (k KnBasicAddressPad) JSONObjects() []string {
	return []string{}
}

// KnPerson
type KnPerson struct {
	PostadresIsAdres                                   bool              `json:"PadAdr"`                // Postadres is adres
	Autonummering                                      bool              `json:"AutoNum"`               // Autonummering
	PersoonVergelijkenOp                               string            `json:"MatchPer"`              // Persoon vergelijken op
	Organisatiepersoonintern                           int               `json:"BcId"`                  // Organisatie/persoon (intern)
	Nummer                                             string            `json:"BcCo"`                  // Nummer
	Zoeknaam                                           string            `json:"SeNm,omitempty"`        // Zoeknaam
	Roepnaam                                           string            `json:"CaNm,omitempty"`        // Roepnaam
	Voornaam                                           string            `json:"FiNm,omitempty"`        // Voornaam
	Voorletters                                        string            `json:"In,omitempty"`          // Voorletters
	Voorvoegsel                                        string            `json:"Is,omitempty"`          // Voorvoegsel
	Achternaam                                         string            `json:"LaNm,omitempty"`        // Achternaam
	GeboortenaamApartVastleggen                        bool              `json:"SpNm"`                  // Geboortenaam apart vastleggen
	VoorvGebnaam                                       string            `json:"IsBi,omitempty"`        // Voorv. geb.naam
	Geboortenaam                                       string            `json:"NmBi,omitempty"`        // Geboortenaam
	VoorvoegselPartner                                 string            `json:"IsPa,omitempty"`        // Voorvoegsel partner
	GebnaamPartner                                     string            `json:"NmPa,omitempty"`        // Geb.naam partner
	Naamgebruik                                        string            `json:"ViUs,omitempty"`        // Naamgebruik
	Geslacht                                           string            `json:"ViGe,omitempty"`        // Geslacht
	Nationaliteit                                      string            `json:"PsNa,omitempty"`        // Nationaliteit
	Geboortedatum                                      *date.Date        `json:"DaBi,omitempty"`        // Geboortedatum
	Geboorteland                                       string            `json:"CoBi,omitempty"`        // Geboorteland
	Geboorteplaats                                     string            `json:"RsBi,omitempty"`        // Geboorteplaats
	BSN                                                string            `json:"SoSe,omitempty"`        // BSN
	BurgerlijkeStaat                                   string            `json:"ViCs,omitempty"`        // Burgerlijke staat
	Huwelijksdatum                                     *date.Date        `json:"DaMa,omitempty"`        // Huwelijksdatum
	DatumScheiding                                     *date.Date        `json:"DaDi,omitempty"`        // Datum scheiding
	Overlijdensdatum                                   *date.Date        `json:"DaDe,omitempty"`        // Overlijdensdatum
	Titelaanhef                                        string            `json:"TtId,omitempty"`        // Titel/aanhef
	TweedeTitel                                        string            `json:"TtEx,omitempty"`        // Tweede titel
	Briefaanhef                                        string            `json:"LeHe,omitempty"`        // Briefaanhef
	TelefoonnrWerk                                     string            `json:"TeNr,omitempty"`        // Telefoonnr. werk
	TelefoonnrPrivé                                    string            `json:"TeN2,omitempty"`        // Telefoonnr. privé
	FaxWerk                                            string            `json:"FaNr,omitempty"`        // Fax werk
	MobielWerk                                         string            `json:"MbNr,omitempty"`        // Mobiel werk
	MobielPrivé                                        string            `json:"MbN2,omitempty"`        // Mobiel privé
	EMailWerk                                          string            `json:"EmAd,omitempty"`        // E-mail werk
	EMailPrivé                                         string            `json:"EmA2,omitempty"`        // E-mail privé
	Website                                            string            `json:"HoPa,omitempty"`        // Website
	Correspondentie                                    bool              `json:"Corr"`                  // Correspondentie
	Voorkeursmedium                                    string            `json:"ViMd,omitempty"`        // Voorkeursmedium
	Opmerking                                          []byte            `json:"Re,omitempty"`          // Opmerking
	Status                                             string            `json:"StId,omitempty"`        // Status
	SocialeNetwerken                                   string            `json:"SocN,omitempty"`        // Sociale netwerken
	Facebook                                           string            `json:"Face,omitempty"`        // Facebook
	LinkedIn                                           string            `json:"Link,omitempty"`        // LinkedIn
	Twitter                                            string            `json:"Twtr,omitempty"`        // Twitter
	LandWetgeving                                      string            `json:"CoLw,omitempty"`        // Land wetgeving
	NaamBestand                                        string            `json:"FileName,omitempty"`    // Naam bestand
	Afbeelding                                         []byte            `json:"FileStream,omitempty"`  // Afbeelding
	PersoonToegangGevenTotAfgeschermdeDeelVanDePortals bool              `json:"AddToPortal"`           // Persoon toegang geven tot afgeschermde deel van de portal(s)
	EMailToegang                                       string            `json:"EmailPortal,omitempty"` // E-mail toegang
	KnBankAccount                                      KnBankAccount     `json:"KnBankAccount"`         // KnBankAccount
	KnBasicAddressAdr                                  KnBasicAddressAdr `json:"KnBasicAddressAdr"`     // KnBasicAddressAdr
	KnBasicAddressPad                                  KnBasicAddressPad `json:"KnBasicAddressPad"`     // KnBasicAddressPad

}

func (k KnPerson) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnPerson

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}

func (k KnPerson) JSONFields() []string {
	return []string{
		"PadAdr",
		"AutoNum",
		"MatchPer",
		"BcId",
		"BcCo",
		"SeNm",
		"CaNm",
		"FiNm",
		"In",
		"Is",
		"LaNm",
		"SpNm",
		"IsBi",
		"NmBi",
		"IsPa",
		"NmPa",
		"ViUs",
		"ViGe",
		"PsNa",
		"DaBi",
		"CoBi",
		"RsBi",
		"SoSe",
		"ViCs",
		"DaMa",
		"DaDi",
		"DaDe",
		"TtId",
		"TtEx",
		"LeHe",
		"TeNr",
		"TeN2",
		"FaNr",
		"MbNr",
		"MbN2",
		"EmAd",
		"EmA2",
		"HoPa",
		"Corr",
		"ViMd",
		"Re",
		"StId",
		"SocN",
		"Face",
		"Link",
		"Twtr",
		"CoLw",
		"FileName",
		"FileStream",
		"AddToPortal",
		"EmailPortal",
		"",
		"",
		"",
	}
}

func (k KnPerson) JSONObjects() []string {
	return []string{
		"KnBankAccount",
		"KnBasicAddressAdr",
		"KnBasicAddressPad",
	}
}

// KnBankAccount
type KnBankAccount struct {
	LandVanDeBank        string `json:"CoId"`           // Land van de bank
	IBANControle         bool   `json:"IbCk"`           // IBAN-controle
	IBANNummer           string `json:"Iban,omitempty"` // IBAN-nummer
	Bankrekening         string `json:"BaAc,omitempty"` // Bankrekening
	TypeBank             int    `json:"BkTp,omitempty"` // Type bank
	Bank                 string `json:"BkIc,omitempty"` // Bank
	GRekening            bool   `json:"AcGa,omitempty"` // G-rekening
	Cheque               bool   `json:"AcCk,omitempty"` // Cheque
	AfwijkendeNaam       string `json:"DiNm,omitempty"` // Afwijkende naam
	AfwijkendeWoonplaats string `json:"DiPl,omitempty"` // Afwijkende woonplaats
	BICCode              string `json:"Bic,omitempty"`  // BIC-code
	NaamBank             string `json:"BaNm,omitempty"` // Naam bank
	FiliaalVanDeBank     string `json:"BaFi,omitempty"` // Filiaal van de bank
	AdresVanDeBank       string `json:"BaAd,omitempty"` // Adres van de bank
	VestigingsplaatsBank string `json:"BaPl,omitempty"` // Vestigingsplaats bank
	CodeDoorberekening   string `json:"CalM,omitempty"` // Code doorberekening

}

func (k KnBankAccount) MarshalJSON() ([]byte, error) {
	// If struct is empty: do nothing
	if zero.IsZero(k) {
		return []byte("null"), nil
	}

	type alias KnBankAccount

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map with preservation of json struct tags
	m := map[string]interface{}{}
	json.Unmarshal(b, &m)

	jsonFields := k.JSONFields()
	fields := map[string]interface{}{}
	jsonObjects := k.JSONObjects()
	objects := map[string]interface{}{}
	for k, v := range m {
		for _, f := range jsonFields {
			if k == f {
				// value is a field
				fields[k] = v
			}
		}

		for _, f := range jsonObjects {
			if k == f {
				// skip empty objects
				// @TODO: move this logic to an Objects struct aliasing
				// map[string]interface{}
				if v == nil || zero.IsZero(v) {
					continue
				}

				// value is an object and not zero
				objects[k] = v
			}
		}
	}

	type Element struct {
		DBID    string                 `json:"@DbId,omitempty"`
		Fields  map[string]interface{} `json:"Fields,omitempty"`
		Objects map[string]interface{} `json:"Objects,omitempty"`
	}

	type Elements struct {
		Element []Element `json:"Element"`
	}

	structure := Elements{
		[]Element{
			Element{
				Fields:  fields,
				Objects: objects,
			},
		},
	}

	return json.Marshal(structure)
}

func (k KnBankAccount) JSONFields() []string {
	return []string{
		"CoId",
		"IbCk",
		"Iban",
		"BaAc",
		"BkTp",
		"BkIc",
		"AcGa",
		"AcCk",
		"DiNm",
		"DiPl",
		"Bic",
		"BaNm",
		"BaFi",
		"BaAd",
		"BaPl",
		"CalM",
	}
}

func (k KnBankAccount) JSONObjects() []string {
	return []string{}
}