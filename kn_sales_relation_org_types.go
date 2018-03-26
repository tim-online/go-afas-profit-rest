package afas

import (
	"encoding/json"
	"time"

	"github.com/cockroachdb/apd"
)

// Verkooprelatie organisatie
type KnSalesRelationOrg struct {
	NummerDebiteur                        string         `json:"DbId,omitempty"`                    // Nummer debiteur
	VoorkeurIbanNummer                    string         `json:"Iban"`                              // Voorkeur Iban nummer
	VoorkeurBankgironummer                string         `json:"BaAc"`                              // Voorkeur bank-/gironummer
	VoorkeurBankgironummerLandCode        string         `json:"CoId"`                              // Voorkeur bank-/gironummer land code
	IsDebiteur                            bool           `json:"IsDb,omitempty"`                    // Is debiteur
	VoorkeurTegenrekening                 string         `json:"ToAc,omitempty"`                    // Voorkeur tegenrekening
	BtwNummer                             string         `json:"VaId,omitempty"`                    // Btw-nummer
	Betalingsvoorwaarde                   string         `json:"PaCd,omitempty"`                    // Betalingsvoorwaarde
	Vertegenwoordiger                     string         `json:"VeId,omitempty"`                    // Vertegenwoordiger
	Taal                                  string         `json:"LgId,omitempty"`                    // Taal
	Valuta                                string         `json:"CuId,omitempty"`                    // Valuta
	AfwijkendeAanmaningsset               int            `json:"DsId,omitempty"`                    // Afwijkende aanmaningsset
	Verantwoordelijke                     string         `json:"EmId,omitempty"`                    // Verantwoordelijke
	BtwPlicht                             string         `json:"VaDu,omitempty"`                    // Btw-plicht
	Profiel                               string         `json:"PfId,omitempty"`                    // Profiel
	Regelkorting                          *apd.Decimal   `json:"PrLi"`                              // % Regelkorting
	Factuurkorting                        *apd.Decimal   `json:"PrFc"`                              // Factuurkorting (%)
	Kredietbeperking                      *apd.Decimal   `json:"ClPc"`                              // Kredietbeperking (%)
	Betalingskorting                      *apd.Decimal   `json:"PrPt"`                              // Betalingskorting (%)
	Kredietlimiet                         *apd.Decimal   `json:"Krli"`                              // Kredietlimiet
	FacturerenAan                         string         `json:"FaTo,omitempty"`                    // Factureren aan
	Vervoerder                            string         `json:"TrPt,omitempty"`                    // Vervoerder
	PrioriteitLevering                    int            `json:"PrDl"`                              // Prioriteit levering
	PrijzenVan                            string         `json:"PrVn,omitempty"`                    // Prijzen van
	VoorkeurPrijslijst                    string         `json:"PrLs,omitempty"`                    // Voorkeur prijslijst
	VoorkeurMagazijn                      string         `json:"VkMa,omitempty"`                    // Voorkeur magazijn
	GeblokkeerdVoorLevering               bool           `json:"Bl,omitempty"`                      // Geblokkeerd voor levering
	VolledigBlokkerenNietMeerZichtbaar    bool           `json:"BlTl,omitempty"`                    // Volledig blokkeren, niet meer zichtbaar
	Hoofddeclarant                        string         `json:"LDId,omitempty"`                    // Hoofddeclarant
	AfwijkendeBtwTariefgroep              bool           `json:"VaYN,omitempty"`                    // Afwijkende btw-tariefgroep
	AanmaningVerzenden                    bool           `json:"DuYN,omitempty"`                    // Aanmaning verzenden
	CodeGroepsadministratie               string         `json:"VaIg,omitempty"`                    // Code groepsadministratie
	StatusBewaking                        string         `json:"VaGu,omitempty"`                    // Status bewaking
	Kortingsgroep                         string         `json:"DsGr,omitempty"`                    // Kortingsgroep
	Nettoprijs                            bool           `json:"NtPr,omitempty"`                    // Nettoprijs
	PrijsInclBtw                          bool           `json:"VtIn,omitempty"`                    // Prijs incl. btw
	Factuurtekst                          string         `json:"InTx"`                              // Factuurtekst
	FactuurGeheelVerdichten               bool           `json:"CITo,omitempty"`                    // Factuur geheel verdichten
	TekstBijGeheelVerdichten              string         `json:"TxTc"`                              // Tekst bij geheel verdichten
	StriktMaximum                         bool           `json:"StMa,omitempty"`                    // Strikt maximum
	MaximumFactuurbedrag                  *apd.Decimal   `json:"MaIn"`                              // Maximum factuurbedrag
	StriktMinimum                         bool           `json:"StMi,omitempty"`                    // Strikt minimum
	Minimumfactuurbedrag                  *apd.Decimal   `json:"MiIn"`                              // Minimumfactuurbedrag
	Afrondingsmethode                     string         `json:"RoOf,omitempty"`                    // Afrondingsmethode
	Declarant                             string         `json:"DeId,omitempty"`                    // Declarant
	Incassospecificatie                   bool           `json:"PaSp,omitempty"`                    // Incassospecificatie
	AutomatischIncasseren                 bool           `json:"AuPa,omitempty"`                    // Automatisch incasseren
	Verdichten                            bool           `json:"PaCo,omitempty"`                    // Verdichten
	EenmaligeIncassomachtigingVereist     bool           `json:"SiPA,omitempty"`                    // Eenmalige incassomachtiging vereist
	WaarschuwingBijOrder                  bool           `json:"WaOr,omitempty"`                    // Waarschuwing bij order
	TekstWaarschuwing                     string         `json:"WaTx,omitempty"`                    // Tekst waarschuwing
	CBSTypen                              string         `json:"CsTy,omitempty"`                    // CBS-typen
	Opmerking                             []byte         `json:"Rm"`                                // Opmerking
	Contactpersoon                        int            `json:"CtP1"`                              // Contactpersoon
	ExtraContactpersoon                   int            `json:"CtP2"`                              // Extra contactpersoon
	Verzamelrekening                      string         `json:"ColA,omitempty"`                    // Verzamelrekening
	KlantSinds                            time.Time      `json:"CsDa,omitempty"`                    // Klant sinds
	AangebrachtDoor                       string         `json:"BcBy,omitempty"`                    // Aangebracht door
	Leveringsconditie                     string         `json:"DeCo,omitempty"`                    // Leveringsconditie
	VoorkeurContact                       int            `json:"CtI1"`                              // Voorkeur contact
	VoorkeurVerstrekkingswijze            string         `json:"InPv,omitempty"`                    // Voorkeur verstrekkingswijze
	Ordersortering                        string         `json:"SoId"`                              // Ordersortering
	TypeBarcode                           string         `json:"VaBc,omitempty"`                    // Type barcode
	Barcode                               string         `json:"BaCo,omitempty"`                    // Barcode
	AdresseringEDIPakbonConformEDIFactuur bool           `json:"EDDn,omitempty"`                    // Adressering EDI-pakbon conform EDI-factuur
	VaA1                                  string         `json:"VaA1"`                              // VaA1
	VaA2                                  string         `json:"VaA2"`                              // VaA2
	VaA3                                  string         `json:"VaA3"`                              // VaA3
	VaA4                                  string         `json:"VaA4"`                              // VaA4
	VaA5                                  string         `json:"VaA5"`                              // VaA5
	Wachtwoord                            string         `json:"Pwrd,omitempty"`                    // Wachtwoord
	Activeringscode                       string         `json:"AtCd,omitempty"`                    // Activeringscode
	Accounttype                           string         `json:"AcTp,omitempty"`                    // Accounttype
	VerwerkingOrder                       string         `json:"OrPr,omitempty"`                    // Verwerking order
	Assortiment                           string         `json:"AsGr,omitempty"`                    // Assortiment
	AfwijkenAssortimentToestaan           bool           `json:"AsYN,omitempty"`                    // Afwijken assortiment toestaan
	StudentNummer                         string         `json:"OINr,omitempty"`                    // Student nummer
	IncassowijzeSEPA                      string         `json:"VaDt,omitempty"`                    // Incassowijze SEPA
	BedrijfsIdEVerbinding                 string         `json:"EnId"`                              // Bedrijfs-Id eVerbinding
	TypeVerkooprelatie                    string         `json:"VaTp,omitempty"`                    // Type verkooprelatie
	VerwijderingsbijdrageToepassen        bool           `json:"ReCo,omitempty"`                    // Verwijderingsbijdrage toepassen
	EenmaligeDebiteur                     bool           `json:"U36AC2FE34594A8F2D14DB88F867DEF52"` // Eenmalige debiteur?
	Machtiging                            bool           `json:"U47C81A5D426695598163F6B1959DEEB1"` // Machtiging?
	KnOrganisation                        KnOrganisation `json:"KnOrganisation"`
}

func (k KnSalesRelationOrg) MarshalJSON() ([]byte, error) {
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
				// value is an object
				objects[k] = v
			}
		}
	}

	requestBody := map[string]interface{}{
		"KnSalesRelationOrg": map[string]interface{}{
			"Element": map[string]interface{}{
				"@DbId":   k.DBID(),
				"Fields":  fields,
				"Objects": objects,
			},
		},
	}

	return json.Marshal(requestBody)
}

func (k KnSalesRelationOrg) DBIDField() string {
	return "NummerDebiteur"
}

func (k KnSalesRelationOrg) DBID() string {
	return k.NummerDebiteur
}

func (k KnSalesRelationOrg) JSONFields() []string {
	return []string{
		"Iban",
	}
}

func (k KnSalesRelationOrg) JSONObjects() []string {
	return []string{
		"KnOrganisation",
	}
}

type KnOrganisation struct {
	PostadresIsAdres         bool      `json:"PadAdr"`         // Postadres is adres
	Autonummering            bool      `json:"AutoNum"`        // Autonummering
	OrganisatieVergelijkenOp string    `json:"MatchOga"`       // Organisatie vergelijken op
	Organisatiepersoonintern int       `json:"BcId,omitempty"` // Organisatie/persoon (intern)
	Nummer                   string    `json:"BcCo,omitempty"` // Nummer
	Zoeknaam                 string    `json:"SeNm,omitempty"` // Zoeknaam
	Naam                     string    `json:"Nm,omitempty"`   // Naam
	Rechtsvorm               string    `json:"ViLe,omitempty"` // Rechtsvorm
	Branche                  string    `json:"ViLb,omitempty"` // Branche
	KvKNummer                string    `json:"CcNr,omitempty"` // KvK-nummer
	DatumKvK                 time.Time `json:"CcDa,omitempty"` // Datum KvK
	Naamstatutair            string    `json:"NmRg,omitempty"` // Naam (statutair)
	Vestigingstatutair       string    `json:"RsRg,omitempty"` // Vestiging (statutair)
	Titelaanhef              string    `json:"TtId,omitempty"` // Titel/aanhef
	Briefaanhef              string    `json:"LeHe,omitempty"` // Briefaanhef
	OrganisatorischeEenheid  string    `json:"OuId,omitempty"` // Organisatorische eenheid
	TelefoonnrWerk           string    `json:"TeNr,omitempty"` // Telefoonnr. werk
	FaxWerk                  string    `json:"FaNr,omitempty"` // Fax werk
	MobielWerk               string    `json:"MbNr,omitempty"` // Mobiel werk
	EMailWerk                string    `json:"EmAd,omitempty"` // E-mail werk
	Website                  string    `json:"HoPa,omitempty"` // Website
	Correspondentie          bool      `json:"Corr,omitempty"` // Correspondentie
	Voorkeursmedium          string    `json:"ViMd,omitempty"` // Voorkeursmedium
	Opmerking                []byte    `json:"Re"`             // Opmerking
	Fiscaalnummer            string    `json:"FiNr,omitempty"` // Fiscaalnummer
	Status                   string    `json:"StId,omitempty"` // Status
	SocialeNetwerken         string    `json:"SocN,omitempty"` // Sociale netwerken
	Facebook                 string    `json:"Face,omitempty"` // Facebook
	LinkedIn                 string    `json:"Link,omitempty"` // LinkedIn
	Twitter                  string    `json:"Twtr,omitempty"` // Twitter
	OnderdeelVanOrganisatie  string    `json:"BcPa,omitempty"` // Onderdeel van organisatie
	NaamBestand              string    `json:"FileName"`       // Naam bestand
	Afbeelding               []byte    `json:"FileStream"`     // Afbeelding
}

func (k KnOrganisation) MarshalJSON() ([]byte, error) {
	type alias KnOrganisation

	// type to json
	b, err := json.Marshal(alias(k))
	if err != nil {
		return b, err
	}

	// json to map
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
				// value is an object
				objects[k] = v
			}
		}
	}

	body := map[string]interface{}{
		"Element": map[string]interface{}{
			"Fields":  fields,
			"Objects": objects,
		},
	}
	return json.Marshal(body)
}

func (k KnOrganisation) JSONFields() []string {
	return []string{
		"MatchOga",
		"Nm",
	}
}

func (k KnOrganisation) JSONObjects() []string {
	return []string{}
}
