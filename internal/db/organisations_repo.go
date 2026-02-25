package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
)

// ═══════════════════════════════════════════════════════════════════════════════
// REPOSITORY — Organisations & Payeurs
// ═══════════════════════════════════════════════════════════════════════════════

// ─────────────────────────────────────────────────────────────────────────────
// ORGANISATIONS
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetAllOrganisations(s *crypto.CryptoService, typeOrg string, actifSeulement bool) ([]models.Organisation, error) {
	q := `SELECT * FROM organisations WHERE 1=1`
	args := []interface{}{}
	if typeOrg != "" {
		q += ` AND type_org = ?`
		args = append(args, typeOrg)
	}
	if actifSeulement {
		q += ` AND actif = 1`
	}
	q += ` ORDER BY type_org, nom`

	var list []models.Organisation
	if err := db.Select(&list, q, args...); err != nil {
		return nil, err
	}
	for i := range list {
		decryptOrganisation(&list[i], s)
	}
	return list, nil
}

// GetOrganisationsListItems — version légère pour les selects (déchiffre nom + contact seulement)
func (db *Database) GetOrganisationsListItems(s *crypto.CryptoService) ([]models.OrganisationListItem, error) {
	var orgs []models.Organisation
	if err := db.Select(&orgs, `SELECT * FROM organisations WHERE actif = 1 ORDER BY type_org, nom`); err != nil {
		return nil, err
	}

	items := make([]models.OrganisationListItem, 0, len(orgs))
	for _, o := range orgs {
		decryptOrganisation(&o, s)
		items = append(items, models.OrganisationListItem{
			ID:         o.ID,
			TypeOrg:    o.TypeOrg,
			Nom:        ptrStr(o.Nom),
			Acronyme:   ptrStr(o.Acronyme),
			ContactNom: ptrStr(o.ContactNom),
			Telephone:  ptrStr(o.Telephone),
			Actif:      o.Actif,
		})
	}
	return items, nil
}

func (db *Database) GetOrganisationByID(id int, s *crypto.CryptoService) (*models.Organisation, error) {
	var o models.Organisation
	if err := db.Get(&o, `SELECT * FROM organisations WHERE id = ?`, id); err != nil {
		return nil, err
	}
	decryptOrganisation(&o, s)
	return &o, nil
}

func (db *Database) CreateOrganisation(req models.CreateOrganisationRequest, createdBy int, s *crypto.CryptoService) (int64, error) {
	enc := req
	encryptOrganisationRequest(&enc, s)

	res, err := db.Exec(`
		INSERT INTO organisations (
			type_org, no_fournisseur, conditions_paiement, taux_tps_exempt, mode_paiement_pref,
			nom, acronyme, no_organisme,
			adresse, ville, code_postal, province, pays,
			telephone, telecopieur, email_general, site_web,
			contact_nom, contact_titre, contact_telephone, contact_email,
			notes, conditions_speciales, created_by
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		enc.TypeOrg, enc.NoFournisseur, enc.ConditionsPaiement, enc.TauxTPSExempt, enc.ModePaiementPref,
		enc.Nom, enc.Acronyme, enc.NoOrganisme,
		enc.Adresse, enc.Ville, enc.CodePostal, enc.Province, enc.Pays,
		enc.Telephone, enc.Telecopieur, enc.EmailGeneral, enc.SiteWeb,
		enc.ContactNom, enc.ContactTitre, enc.ContactTelephone, enc.ContactEmail,
		enc.Notes, enc.ConditionsSpeciales, createdBy)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) UpdateOrganisation(req models.UpdateOrganisationRequest, s *crypto.CryptoService) error {
	enc := req.CreateOrganisationRequest
	encryptOrganisationRequest(&enc, s)

	_, err := db.Exec(`
		UPDATE organisations SET
			type_org=?, no_fournisseur=?, conditions_paiement=?, taux_tps_exempt=?, mode_paiement_pref=?,
			nom=?, acronyme=?, no_organisme=?,
			adresse=?, ville=?, code_postal=?, province=?, pays=?,
			telephone=?, telecopieur=?, email_general=?, site_web=?,
			contact_nom=?, contact_titre=?, contact_telephone=?, contact_email=?,
			notes=?, conditions_speciales=?, updated_at=CURRENT_TIMESTAMP
		WHERE id=?`,
		enc.TypeOrg, enc.NoFournisseur, enc.ConditionsPaiement, enc.TauxTPSExempt, enc.ModePaiementPref,
		enc.Nom, enc.Acronyme, enc.NoOrganisme,
		enc.Adresse, enc.Ville, enc.CodePostal, enc.Province, enc.Pays,
		enc.Telephone, enc.Telecopieur, enc.EmailGeneral, enc.SiteWeb,
		enc.ContactNom, enc.ContactTitre, enc.ContactTelephone, enc.ContactEmail,
		enc.Notes, enc.ConditionsSpeciales, req.ID)
	return err
}

func (db *Database) ArchiverOrganisation(id int) error {
	_, err := db.Exec(`UPDATE organisations SET actif=0, updated_at=CURRENT_TIMESTAMP WHERE id=?`, id)
	return err
}

func (db *Database) DeleteOrganisation(id int) error {
	// Vérifier qu'aucun payeur ne pointe vers cette org
	var count int
	db.Get(&count, `SELECT COUNT(*) FROM payeurs WHERE organisation_id=?`, id)
	if count > 0 {
		return fmt.Errorf("cette organisation est liée à %d payeur(s) — archivez-la plutôt", count)
	}
	_, err := db.Exec(`DELETE FROM organisations WHERE id=?`, id)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// PAYEURS
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) CreatePayeur(req models.CreatePayeurRequest, createdBy int, s *crypto.CryptoService) (int64, error) {
	nomFact, _ := s.EncryptStringPtr(req.NomFacturation)
	adrFact, _ := s.EncryptStringPtr(req.AdresseFacturation)
	emailFact, _ := s.EncryptStringPtr(req.EmailFacturation)
	notes, _ := s.EncryptStringPtr(req.Notes)

	res, err := db.Exec(`
		INSERT INTO payeurs (type_payeur, client_id, contact_id, organisation_id, est_defaut,
			nom_facturation, adresse_facturation, email_facturation, notes, created_by)
		VALUES (?,?,?,?,?,?,?,?,?,?)`,
		req.TypePayeur, req.ClientID, req.ContactID, req.OrganisationID, req.EstDefaut,
		nomFact, adrFact, emailFact, notes, createdBy)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) GetPayeursByClient(clientID int, s *crypto.CryptoService) ([]models.PayeurResolu, error) {
	// Récupère tous les payeurs liés directement au client
	// (type client) OU dont le contact appartient au client
	var payeurs []models.Payeur
	if err := db.Select(&payeurs, `
		SELECT p.* FROM payeurs p
		LEFT JOIN contacts c ON p.contact_id = c.id
		WHERE p.client_id = ? OR c.client_id = ?
		ORDER BY p.est_defaut DESC, p.id`,
		clientID, clientID); err != nil {
		return nil, err
	}
	return db.resoudrePayeurs(payeurs, s)
}

// ResoudrePayeur — résout un payeur_id en PayeurResolu (pour affichage sur facture/contrat)
func (db *Database) ResoudrePayeur(payeurID int, s *crypto.CryptoService) (*models.PayeurResolu, error) {
	var p models.Payeur
	if err := db.Get(&p, `SELECT * FROM payeurs WHERE id=?`, payeurID); err != nil {
		return nil, err
	}
	resolved, err := db.resoudrePayeurs([]models.Payeur{p}, s)
	if err != nil || len(resolved) == 0 {
		return nil, err
	}
	return &resolved[0], nil
}

// resoudrePayeurs — enrichit chaque payeur avec les infos déchiffrées de sa source
func (db *Database) resoudrePayeurs(payeurs []models.Payeur, s *crypto.CryptoService) ([]models.PayeurResolu, error) {
	result := make([]models.PayeurResolu, 0, len(payeurs))

	for _, p := range payeurs {
		// Déchiffrer les surcharges optionnelles
		nomFact, _ := s.DecryptStringPtr(p.NomFacturation)
		emailFact, _ := s.DecryptStringPtr(p.EmailFacturation)
		adrFact, _ := s.DecryptStringPtr(p.AdresseFacturation)

		resolu := models.PayeurResolu{
			PayeurID:   p.ID,
			TypePayeur: p.TypePayeur,
		}

		switch p.TypePayeur {

		case "client":
			if p.ClientID != nil {
				client, err := db.GetClientByID(*p.ClientID, s)
				if err == nil {
					resolu.Nom = fmt.Sprintf("%s %s", client.Prenom, client.Nom)
					resolu.Email = ptrStr(emailFact)
					resolu.Detail = "Client"
					resolu.JoursPaiement = 30
				}
			}

		case "contact":
			if p.ContactID != nil {
				var contact struct {
					Nom       *string `db:"nom"`
					Prenom    *string `db:"prenom"`
					Telephone *string `db:"telephone"`
					Email     *string `db:"email"`
				}
				if err := db.Get(&contact, `SELECT nom, prenom, telephone, email FROM contacts WHERE id=?`, *p.ContactID); err == nil {
					nom, _ := s.DecryptStringPtr(contact.Nom)
					prenom, _ := s.DecryptStringPtr(contact.Prenom)
					tel, _ := s.DecryptStringPtr(contact.Telephone)
					email, _ := s.DecryptStringPtr(contact.Email)
					resolu.Nom = fmt.Sprintf("%s %s", ptrStr(prenom), ptrStr(nom))
					resolu.Telephone = ptrStr(tel)
					resolu.Email = ptrStr(email)
					resolu.Detail = "Contact / Mandataire"
					resolu.JoursPaiement = 30
				}
			}

		case "organisation":
			if p.OrganisationID != nil {
				org, err := db.GetOrganisationByID(*p.OrganisationID, s)
				if err == nil {
					resolu.Nom = ptrStr(org.Nom)
					resolu.Detail = org.TypeOrg
					resolu.Telephone = ptrStr(org.Telephone)
					resolu.Email = ptrStr(org.EmailGeneral)
					resolu.Adresse = ptrStr(org.Adresse)
					resolu.ModePaiement = org.ModePaiementPref
					resolu.JoursPaiement = org.ConditionsPaiement
					resolu.ExempteTPSTVQ = org.TauxTPSExempt == 1
				}
			}
		}

		// Surcharges optionnelles — écrasent les valeurs de la source
		if ptrStr(nomFact) != "" {
			resolu.Nom = ptrStr(nomFact)
		}
		if ptrStr(emailFact) != "" {
			resolu.Email = ptrStr(emailFact)
		}
		if ptrStr(adrFact) != "" {
			resolu.Adresse = ptrStr(adrFact)
		}

		result = append(result, resolu)
	}
	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// HELPERS — Chiffrement organisations
// ─────────────────────────────────────────────────────────────────────────────

func decryptOrganisation(o *models.Organisation, s *crypto.CryptoService) {
	o.Nom, _ = s.DecryptStringPtr(o.Nom)
	o.Acronyme, _ = s.DecryptStringPtr(o.Acronyme)
	o.NoOrganisme, _ = s.DecryptStringPtr(o.NoOrganisme)
	o.Adresse, _ = s.DecryptStringPtr(o.Adresse)
	o.Ville, _ = s.DecryptStringPtr(o.Ville)
	o.CodePostal, _ = s.DecryptStringPtr(o.CodePostal)
	o.Province, _ = s.DecryptStringPtr(o.Province)
	o.Pays, _ = s.DecryptStringPtr(o.Pays)
	o.Telephone, _ = s.DecryptStringPtr(o.Telephone)
	o.Telecopieur, _ = s.DecryptStringPtr(o.Telecopieur)
	o.EmailGeneral, _ = s.DecryptStringPtr(o.EmailGeneral)
	o.SiteWeb, _ = s.DecryptStringPtr(o.SiteWeb)
	o.ContactNom, _ = s.DecryptStringPtr(o.ContactNom)
	o.ContactTitre, _ = s.DecryptStringPtr(o.ContactTitre)
	o.ContactTelephone, _ = s.DecryptStringPtr(o.ContactTelephone)
	o.ContactEmail, _ = s.DecryptStringPtr(o.ContactEmail)
	o.Notes, _ = s.DecryptStringPtr(o.Notes)
	o.ConditionsSpeciales, _ = s.DecryptStringPtr(o.ConditionsSpeciales)
}

func encryptOrganisationRequest(enc *models.CreateOrganisationRequest, s *crypto.CryptoService) {
	enc.Nom, _ = s.EncryptStringPtr(enc.Nom)
	enc.Acronyme, _ = s.EncryptStringPtr(enc.Acronyme)
	enc.NoOrganisme, _ = s.EncryptStringPtr(enc.NoOrganisme)
	enc.Adresse, _ = s.EncryptStringPtr(enc.Adresse)
	enc.Ville, _ = s.EncryptStringPtr(enc.Ville)
	enc.CodePostal, _ = s.EncryptStringPtr(enc.CodePostal)
	enc.Province, _ = s.EncryptStringPtr(enc.Province)
	enc.Pays, _ = s.EncryptStringPtr(enc.Pays)
	enc.Telephone, _ = s.EncryptStringPtr(enc.Telephone)
	enc.Telecopieur, _ = s.EncryptStringPtr(enc.Telecopieur)
	enc.EmailGeneral, _ = s.EncryptStringPtr(enc.EmailGeneral)
	enc.SiteWeb, _ = s.EncryptStringPtr(enc.SiteWeb)
	enc.ContactNom, _ = s.EncryptStringPtr(enc.ContactNom)
	enc.ContactTitre, _ = s.EncryptStringPtr(enc.ContactTitre)
	enc.ContactTelephone, _ = s.EncryptStringPtr(enc.ContactTelephone)
	enc.ContactEmail, _ = s.EncryptStringPtr(enc.ContactEmail)
	enc.Notes, _ = s.EncryptStringPtr(enc.Notes)
	enc.ConditionsSpeciales, _ = s.EncryptStringPtr(enc.ConditionsSpeciales)
}

// ptrStr — retourne "" si le pointeur est nil
func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
