package afas

// Verkooprelatie (Rapport)
type OMNIVerkooprelatie struct {
	NummerDebiteur         string `json:"Nummer_debiteur"`          // Nummer debiteur
	NummerOrgPers          string `json:"Nummer_org-pers"`          // Nummer org-pers
	OrganisatiePersoonCode string `json:"Organisatie_persoon_code"` // Organisatie/persoon code

}
