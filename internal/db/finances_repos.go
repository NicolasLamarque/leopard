package database

import (
	"fmt"
	"leopard/internal/crypto"
	models "leopard/internal/model"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// REPOSITORY FINANCIER — Leopard — version finale
//
// Champs chiffrés restants (4 seulement) :
//   ParametresFinance : nom_complet, titre, no_membre, email, téléphone,
//                       adresse, no_tps, no_tvq, interac_email
//   Revenu            : reference_paiement
//   Facture           : tiers_payant_nom
//   Paiement          : reference
//
// Tout le reste est en clair → SQL natif, zéro parseFloat, zéro goroutine crypto.
// Les triggers SQLite maintiennent stats_mensuelles automatiquement.
// ═══════════════════════════════════════════════════════════════════════════════

// ─────────────────────────────────────────────────────────────────────────────
// PARAMÈTRES FINANCE
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetParametresFinance(s *crypto.CryptoService) (*models.ParametresFinance, error) {
	var p models.ParametresFinance
	if err := db.Get(&p, `SELECT * FROM parametres_finance WHERE id = 1`); err != nil {
		return &models.ParametresFinance{
			ID: 1, SousSeuilTPS: 1, TauxTPS: 0.05, TauxTVQ: 0.09975,
			TauxKM: 0.72, PrefixeFacture: "FAC",
			AnneeFacture: time.Now().Year(), ProchainNoFacture: 1,
			JoursPaiementDefaut: 30,
		}, nil
	}
	decryptParametresFinance(&p, s)
	return &p, nil
}

func (db *Database) SaveParametresFinance(p models.ParametresFinance, s *crypto.CryptoService) error {
	enc := p
	encryptParametresFinance(&enc, s)
	_, err := db.Exec(`
		INSERT INTO parametres_finance (
			id, nom_complet, titre_professionnel, no_membre_ordre,
			email_professionnel, telephone_professionnel, adresse_pratique,
			no_tps, no_tvq, interac_email,
			taux_horaire_defaut, taux_forfait_evaluation, taux_rapport_expertise, taux_km,
			sous_seuil_tps, inscrite_tps, inscrite_tvq, taux_tps, taux_tvq,
			prefixe_facture, annee_facture, prochain_no_facture,
			jours_paiement_defaut, infos_paiement, politique_annulation, mentions_legales_pdf,
			updated_at
		) VALUES (1,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP)
		ON CONFLICT(id) DO UPDATE SET
			nom_complet=excluded.nom_complet,
			titre_professionnel=excluded.titre_professionnel,
			no_membre_ordre=excluded.no_membre_ordre,
			email_professionnel=excluded.email_professionnel,
			telephone_professionnel=excluded.telephone_professionnel,
			adresse_pratique=excluded.adresse_pratique,
			no_tps=excluded.no_tps, no_tvq=excluded.no_tvq,
			interac_email=excluded.interac_email,
			taux_horaire_defaut=excluded.taux_horaire_defaut,
			taux_forfait_evaluation=excluded.taux_forfait_evaluation,
			taux_rapport_expertise=excluded.taux_rapport_expertise,
			taux_km=excluded.taux_km,
			sous_seuil_tps=excluded.sous_seuil_tps,
			inscrite_tps=excluded.inscrite_tps, inscrite_tvq=excluded.inscrite_tvq,
			prefixe_facture=excluded.prefixe_facture,
			interac_email=excluded.interac_email,
			jours_paiement_defaut=excluded.jours_paiement_defaut,
			infos_paiement=excluded.infos_paiement,
			politique_annulation=excluded.politique_annulation,
			mentions_legales_pdf=excluded.mentions_legales_pdf,
			updated_at=CURRENT_TIMESTAMP`,
		enc.NomComplet, enc.TitreProfessionnel, enc.NoMembreOrdre,
		enc.EmailProfessionnel, enc.TelephoneProfessionnel, enc.AdressePratique,
		enc.NoTPS, enc.NoTVQ, enc.InteracEmail,
		enc.TauxHoraireDefaut, enc.TauxForfaitEvaluation, enc.TauxRapportExpertise, enc.TauxKM,
		enc.SousSeuilTPS, enc.InscriteTPS, enc.InscriteTVQ, enc.TauxTPS, enc.TauxTVQ,
		enc.PrefixeFacture, enc.AnneeFacture, enc.ProchainNoFacture,
		enc.JoursPaiementDefaut, enc.InfosPaiement, enc.PolitiqueAnnulation, enc.MentionsLegalesPDF)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// SERVICES — aucun chiffrement, CRUD direct
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetAllServices(categorie string, actifSeulement bool) ([]models.Service, error) {
	q := `SELECT * FROM services WHERE 1=1`
	args := []interface{}{}
	if categorie != "" {
		q += ` AND categorie = ?`
		args = append(args, categorie)
	}
	if actifSeulement {
		q += ` AND actif = 1`
	}
	q += ` ORDER BY ordre_affichage, categorie, id`
	var list []models.Service
	return list, db.Select(&list, q, args...)
}

func (db *Database) GetServiceByID(id int) (*models.Service, error) {
	var svc models.Service
	return &svc, db.Get(&svc, `SELECT * FROM services WHERE id = ?`, id)
}

func (db *Database) CreateService(req models.CreateServiceRequest, createdBy int) (int64, error) {
	res, err := db.Exec(`
		INSERT INTO services (
			code, categorie, type_livraison, actif, ordre_affichage,
			nom, description_courte, description_longue, public_cible, notes_internes,
			mode_tarification, duree_minutes, taux_horaire, tarif_unitaire, tarif_min, tarif_max,
			exempte_taxes, tps_applicable, tvq_applicable,
			nb_places_max, nb_seances_forfait, duree_programme_semaines,
			format_tract, couleur_tract, duree_video_minutes, url_ressource, tags, icone,
			created_by
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		req.Code, req.Categorie, req.TypeLivraison, req.Actif, req.OrdreAffichage,
		req.Nom, req.DescriptionCourte, req.DescriptionLongue, req.PublicCible, req.NotesInternes,
		req.ModeTarification, req.DureeMinutes, req.TauxHoraire, req.TarifUnitaire, req.TarifMin, req.TarifMax,
		req.ExempteTaxes, req.TPSApplicable, req.TVQApplicable,
		req.NbPlacesMax, req.NbSeancesForfait, req.DureeProgrammeSemaines,
		req.FormatTract, req.CouleurTract, req.DureeVideoMinutes, req.URLRessource, req.Tags, req.Icone,
		createdBy)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) UpdateService(req models.UpdateServiceRequest) error {
	_, err := db.Exec(`
		UPDATE services SET
			code=?, categorie=?, type_livraison=?, actif=?, ordre_affichage=?,
			nom=?, description_courte=?, description_longue=?, public_cible=?, notes_internes=?,
			mode_tarification=?, duree_minutes=?, taux_horaire=?, tarif_unitaire=?,
			tarif_min=?, tarif_max=?, exempte_taxes=?, tps_applicable=?, tvq_applicable=?,
			nb_places_max=?, nb_seances_forfait=?, duree_programme_semaines=?,
			format_tract=?, couleur_tract=?, duree_video_minutes=?,
			url_ressource=?, tags=?, icone=?,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=?`,
		req.Code, req.Categorie, req.TypeLivraison, req.Actif, req.OrdreAffichage,
		req.Nom, req.DescriptionCourte, req.DescriptionLongue, req.PublicCible, req.NotesInternes,
		req.ModeTarification, req.DureeMinutes, req.TauxHoraire, req.TarifUnitaire,
		req.TarifMin, req.TarifMax, req.ExempteTaxes, req.TPSApplicable, req.TVQApplicable,
		req.NbPlacesMax, req.NbSeancesForfait, req.DureeProgrammeSemaines,
		req.FormatTract, req.CouleurTract, req.DureeVideoMinutes,
		req.URLRessource, req.Tags, req.Icone, req.ID)
	return err
}

func (db *Database) ArchiverService(id int) error {
	_, err := db.Exec(`UPDATE services SET actif=0, updated_at=CURRENT_TIMESTAMP WHERE id=?`, id)
	return err
}

func (db *Database) DeleteService(id int) error {
	var count int
	db.Get(&count, `SELECT COUNT(*) FROM revenus WHERE service_id=?`, id)
	if count > 0 {
		return fmt.Errorf("ce service est utilisé dans %d revenu(s) — archivez-le plutôt que de le supprimer", count)
	}
	_, err := db.Exec(`DELETE FROM services WHERE id=?`, id)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// REVENUS — un seul champ chiffré : reference_paiement
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetAllRevenus(s *crypto.CryptoService, dateDebut, dateFin, statut, categorie string) ([]models.Revenu, error) {
	q := `SELECT * FROM revenus WHERE 1=1`
	args := []interface{}{}
	if dateDebut != "" {
		q += ` AND date_service >= ?`
		args = append(args, dateDebut)
	}
	if dateFin != "" {
		q += ` AND date_service <= ?`
		args = append(args, dateFin)
	}
	if statut != "" {
		q += ` AND statut_paiement = ?`
		args = append(args, statut)
	}
	if categorie != "" {
		q += ` AND categorie = ?`
		args = append(args, categorie)
	}
	q += ` ORDER BY date_service DESC`

	var list []models.Revenu
	if err := db.Select(&list, q, args...); err != nil {
		return nil, err
	}
	for i := range list {
		list[i].ReferencePaiement, _ = s.DecryptStringPtr(list[i].ReferencePaiement)
	}
	return list, nil
}

func (db *Database) GetRevenusByClient(clientID int, s *crypto.CryptoService) ([]models.Revenu, error) {
	var list []models.Revenu
	if err := db.Select(&list, `SELECT * FROM revenus WHERE client_id=? ORDER BY date_service DESC`, clientID); err != nil {
		return nil, err
	}
	for i := range list {
		list[i].ReferencePaiement, _ = s.DecryptStringPtr(list[i].ReferencePaiement)
	}
	return list, nil
}

func (db *Database) CreateRevenu(req models.CreateRevenuRequest, createdBy int, s *crypto.CryptoService) (int64, error) {
	ref, _ := s.EncryptStringPtr(req.ReferencePaiement)
	res, err := db.Exec(`
		INSERT INTO revenus (
			client_id, service_id, facture_id,
			date_service, categorie, mode_facturation, statut_paiement,
			mode_paiement, date_paiement,
			duree_heures, taux_horaire_applique,
			montant_brut, montant_tps, montant_tvq, montant_total,
			description, notes, reference_paiement, created_by
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		req.ClientID, req.ServiceID, req.FactureID,
		req.DateService, req.Categorie, req.ModeFacturation, req.StatutPaiement,
		req.ModePaiement, req.DatePaiement,
		req.DureeHeures, req.TauxHoraireApplique,
		req.MontantBrut, req.MontantTPS, req.MontantTVQ, req.MontantTotal,
		req.Description, req.Notes, ref, createdBy)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) UpdateRevenu(req models.UpdateRevenuRequest, s *crypto.CryptoService) error {
	ref, _ := s.EncryptStringPtr(req.ReferencePaiement)
	_, err := db.Exec(`
		UPDATE revenus SET
			client_id=?, service_id=?, facture_id=?,
			date_service=?, categorie=?, mode_facturation=?, statut_paiement=?,
			mode_paiement=?, date_paiement=?,
			duree_heures=?, taux_horaire_applique=?,
			montant_brut=?, montant_tps=?, montant_tvq=?, montant_total=?,
			description=?, notes=?, reference_paiement=?,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=?`,
		req.ClientID, req.ServiceID, req.FactureID,
		req.DateService, req.Categorie, req.ModeFacturation, req.StatutPaiement,
		req.ModePaiement, req.DatePaiement,
		req.DureeHeures, req.TauxHoraireApplique,
		req.MontantBrut, req.MontantTPS, req.MontantTVQ, req.MontantTotal,
		req.Description, req.Notes, ref, req.ID)
	return err
}

func (db *Database) DeleteRevenu(id int) error {
	_, err := db.Exec(`DELETE FROM revenus WHERE id=?`, id)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// DÉPENSES — aucun chiffrement
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetAllDepenses(dateDebut, dateFin, categorie string, deductibleSeulement bool) ([]models.Depense, error) {
	q := `SELECT * FROM depenses WHERE 1=1`
	args := []interface{}{}
	if dateDebut != "" {
		q += ` AND date_depense >= ?`
		args = append(args, dateDebut)
	}
	if dateFin != "" {
		q += ` AND date_depense <= ?`
		args = append(args, dateFin)
	}
	if categorie != "" {
		q += ` AND categorie = ?`
		args = append(args, categorie)
	}
	if deductibleSeulement {
		q += ` AND deductible = 1`
	}
	q += ` ORDER BY date_depense DESC`
	var list []models.Depense
	return list, db.Select(&list, q, args...)
}

func (db *Database) CreateDepense(req models.CreateDepenseRequest, createdBy int) (int64, error) {
	res, err := db.Exec(`
		INSERT INTO depenses (
			date_depense, categorie, deductible, pct_utilisation_affaires,
			est_kilometrage, sous_total, montant_tps, montant_tvq,
			montant_total, montant_deductible, km_parcourus, taux_km_utilise,
			fournisseur, no_recu, description, notes, created_by
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		req.DateDepense, req.Categorie, req.Deductible, req.PctUtilisationAffaires,
		req.EstKilometrage, req.SousTotal, req.MontantTPS, req.MontantTVQ,
		req.MontantTotal, req.MontantDeductible, req.KMParcourus, req.TauxKMUtilise,
		req.Fournisseur, req.NoRecu, req.Description, req.Notes, createdBy)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) UpdateDepense(req models.UpdateDepenseRequest) error {
	_, err := db.Exec(`
		UPDATE depenses SET
			date_depense=?, categorie=?, deductible=?, pct_utilisation_affaires=?,
			est_kilometrage=?, sous_total=?, montant_tps=?, montant_tvq=?,
			montant_total=?, montant_deductible=?, km_parcourus=?, taux_km_utilise=?,
			fournisseur=?, no_recu=?, description=?, notes=?,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=?`,
		req.DateDepense, req.Categorie, req.Deductible, req.PctUtilisationAffaires,
		req.EstKilometrage, req.SousTotal, req.MontantTPS, req.MontantTVQ,
		req.MontantTotal, req.MontantDeductible, req.KMParcourus, req.TauxKMUtilise,
		req.Fournisseur, req.NoRecu, req.Description, req.Notes, req.ID)
	return err
}

func (db *Database) DeleteDepense(id int) error {
	_, err := db.Exec(`DELETE FROM depenses WHERE id=?`, id)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// FACTURES — seul tiers_payant_nom est chiffré
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetAllFactures(s *crypto.CryptoService, statut, dateDebut, dateFin string) ([]models.Facture, error) {
	q := `SELECT * FROM factures WHERE 1=1`
	args := []interface{}{}
	if statut != "" {
		q += ` AND statut = ?`
		args = append(args, statut)
	}
	if dateDebut != "" {
		q += ` AND date_emission >= ?`
		args = append(args, dateDebut)
	}
	if dateFin != "" {
		q += ` AND date_emission <= ?`
		args = append(args, dateFin)
	}
	q += ` ORDER BY date_emission DESC`

	var list []models.Facture
	if err := db.Select(&list, q, args...); err != nil {
		return nil, err
	}
	for i := range list {
		list[i].TiersPayantNom, _ = s.DecryptStringPtr(list[i].TiersPayantNom)
		list[i].Lignes, _ = db.getFactureLignes(list[i].ID)
	}
	return list, nil
}

func (db *Database) GetFactureByID(id int, s *crypto.CryptoService) (*models.Facture, error) {
	var f models.Facture
	if err := db.Get(&f, `SELECT * FROM factures WHERE id=?`, id); err != nil {
		return nil, err
	}
	f.TiersPayantNom, _ = s.DecryptStringPtr(f.TiersPayantNom)
	f.Lignes, _ = db.getFactureLignes(id)
	return &f, nil
}

func (db *Database) CreateFacture(req models.CreateFactureRequest, createdBy int, s *crypto.CryptoService) (int64, error) {
	numero, err := db.genererNumeroFacture()
	if err != nil {
		return 0, err
	}
	tiersNom, _ := s.EncryptStringPtr(req.TiersPayantNom)

	res, err := db.Exec(`
		INSERT INTO factures (
			numero, client_id, date_emission, date_echeance, statut,
			exempte_taxes, avec_tps, avec_tvq, jours_paiement,
			montant_sous_total, montant_tps, montant_tvq, montant_total,
			montant_paye, montant_solde,
			montant_tiers_payant, montant_du_client, tiers_payant_nom,
			titre_facture, notes_client, notes_internes, created_by
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,0,?,?,?,?,?,?,?,?)`,
		numero, req.ClientID, req.DateEmission, req.DateEcheance, "brouillon",
		req.ExempteTaxes, req.AvecTPS, req.AvecTVQ, req.JoursPaiement,
		req.MontantSousTotal, req.MontantTPS, req.MontantTVQ, req.MontantTotal,
		req.MontantTotal, // solde initial = total
		req.MontantTiersPayant, req.MontantDuClient, tiersNom,
		req.TitreFacture, req.NotesClient, req.NotesInternes, createdBy)
	if err != nil {
		return 0, err
	}
	factureID, _ := res.LastInsertId()
	for i, ligne := range req.Lignes {
		if err := db.createFactureLigne(int(factureID), i, ligne); err != nil {
			return factureID, err
		}
	}
	db.Exec(`UPDATE parametres_finance SET prochain_no_facture = prochain_no_facture + 1 WHERE id=1`)
	return factureID, nil
}

func (db *Database) UpdateFactureStatut(id int, statut string) error {
	_, err := db.Exec(`UPDATE factures SET statut=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`, statut, id)
	return err
}

// EnregistrerPaiement — transaction atomique, SUM direct (montants en clair)
func (db *Database) EnregistrerPaiement(req models.CreatePaiementRequest, s *crypto.CryptoService) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ref, _ := s.EncryptStringPtr(req.Reference)

	_, err = tx.Exec(`
		INSERT INTO paiements (facture_id, date_paiement, source, mode_paiement, montant, notes, reference)
		VALUES (?,?,?,?,?,?,?)`,
		req.FactureID, req.DatePaiement, req.Source, req.ModePaiement,
		req.Montant, req.Notes, ref)
	if err != nil {
		return err
	}

	// SUM direct — montants en clair → pas de déchiffrement
	var totalPaye float64
	tx.Get(&totalPaye, `SELECT COALESCE(SUM(montant), 0) FROM paiements WHERE facture_id=?`, req.FactureID)

	var montantTotal float64
	tx.Get(&montantTotal, `SELECT montant_total FROM factures WHERE id=?`, req.FactureID)

	solde := montantTotal - totalPaye

	statut := "envoyee"
	var datePaiement interface{} = nil
	if solde <= 0.005 {
		statut = "payee"
		datePaiement = req.DatePaiement
	} else if totalPaye > 0 {
		statut = "partielle"
	}

	_, err = tx.Exec(`
		UPDATE factures SET
			montant_paye=?, montant_solde=?, statut=?,
			date_paiement=?, updated_at=CURRENT_TIMESTAMP
		WHERE id=?`,
		totalPaye, solde, statut, datePaiement, req.FactureID)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// VerifierFacturesEnRetard — pure SQL, aucun déchiffrement
func (db *Database) VerifierFacturesEnRetard() error {
	_, err := db.Exec(`
		UPDATE factures SET statut='en_retard', updated_at=CURRENT_TIMESTAMP
		WHERE statut='envoyee' AND date_echeance < date('now')`)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// CONTRATS — aucun chiffrement
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetAllContrats(clientID int, statut string) ([]models.Contrat, error) {
	q := `SELECT * FROM contrats WHERE 1=1`
	args := []interface{}{}
	if clientID > 0 {
		q += ` AND client_id=?`
		args = append(args, clientID)
	}
	if statut != "" {
		q += ` AND statut=?`
		args = append(args, statut)
	}
	q += ` ORDER BY date_debut DESC`
	var list []models.Contrat
	return list, db.Select(&list, q, args...)
}

func (db *Database) GetContratByID(id int) (*models.Contrat, error) {
	var c models.Contrat
	return &c, db.Get(&c, `SELECT * FROM contrats WHERE id=?`, id)
}

func (db *Database) CreateContrat(req models.CreateContratRequest, createdBy int) (int64, error) {
	res, err := db.Exec(`
		INSERT INTO contrats (
			client_id, service_id, type_contrat, statut,
			date_debut, date_fin, renouvellement_auto,
			mode_facturation, taux_horaire, taux_forfait, duree_seance_min,
			frequence_seances,
			clause_objet, clause_objectifs,
			clause_services_inclus, clause_services_exclus,
			clause_tarification, clause_paiement_retard,
			clause_annulation, clause_absence_prolongee,
			clause_confidentialite, clause_limites_confidentialite, clause_dossier_client,
			clause_communication, clause_urgences,
			clause_resiliation_client, clause_resiliation_ts,
			clauses_specifiques, notes_internes, created_by
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		req.ClientID, req.ServiceID, req.TypeContrat, "actif",
		req.DateDebut, req.DateFin, req.RenouvellementAuto,
		req.ModeFacturation, req.TauxHoraire, req.TauxForfait, req.DureeSeanceMin,
		req.FrequenceSeances,
		req.ClauseObjet, req.ClauseObjectifs,
		req.ClauseServicesInclus, req.ClauseServicesExclus,
		req.ClauseTarification, req.ClausePaiementRetard,
		req.ClauseAnnulation, req.ClauseAbsenceProlongee,
		req.ClauseConfidentialite, req.ClauseLimitesConfidentialite, req.ClauseDossierClient,
		req.ClauseCommunication, req.ClauseUrgences,
		req.ClauseResiliationClient, req.ClauseResiliationTS,
		req.ClausesSpecifiques, req.NotesInternes, createdBy)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) UpdateContrat(req models.UpdateContratRequest) error {
	_, err := db.Exec(`
		UPDATE contrats SET
			type_contrat=?, date_fin=?, renouvellement_auto=?,
			mode_facturation=?, taux_horaire=?, taux_forfait=?, duree_seance_min=?,
			frequence_seances=?,
			clause_objet=?, clause_objectifs=?,
			clause_services_inclus=?, clause_services_exclus=?,
			clause_tarification=?, clause_paiement_retard=?,
			clause_annulation=?, clause_absence_prolongee=?,
			clause_confidentialite=?, clause_limites_confidentialite=?, clause_dossier_client=?,
			clause_communication=?, clause_urgences=?,
			clause_resiliation_client=?, clause_resiliation_ts=?,
			clauses_specifiques=?, notes_internes=?,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=?`,
		req.TypeContrat, req.DateFin, req.RenouvellementAuto,
		req.ModeFacturation, req.TauxHoraire, req.TauxForfait, req.DureeSeanceMin,
		req.FrequenceSeances,
		req.ClauseObjet, req.ClauseObjectifs,
		req.ClauseServicesInclus, req.ClauseServicesExclus,
		req.ClauseTarification, req.ClausePaiementRetard,
		req.ClauseAnnulation, req.ClauseAbsenceProlongee,
		req.ClauseConfidentialite, req.ClauseLimitesConfidentialite, req.ClauseDossierClient,
		req.ClauseCommunication, req.ClauseUrgences,
		req.ClauseResiliationClient, req.ClauseResiliationTS,
		req.ClausesSpecifiques, req.NotesInternes, req.ID)
	return err
}

func (db *Database) SignerContrat(id int, dateSignature string) error {
	_, err := db.Exec(`
		UPDATE contrats SET consentement_signe=1, date_signature=?, updated_at=CURRENT_TIMESTAMP
		WHERE id=?`, dateSignature, id)
	return err
}

// ─────────────────────────────────────────────────────────────────────────────
// STATS & DASHBOARD — pure SQL, zéro crypto
// ─────────────────────────────────────────────────────────────────────────────

func (db *Database) GetStatsMensuelles(annee, mois int) (*models.StatsMensuelles, error) {
	var s models.StatsMensuelles
	if err := db.Get(&s, `SELECT * FROM stats_mensuelles WHERE annee=? AND mois=?`, annee, mois); err != nil {
		return &models.StatsMensuelles{Annee: annee, Mois: mois}, nil
	}
	return &s, nil
}

func (db *Database) GetDerniersMois() ([]models.StatsMensuelles, error) {
	var list []models.StatsMensuelles
	if err := db.Select(&list, `SELECT * FROM stats_mensuelles ORDER BY annee DESC, mois DESC LIMIT 12`); err != nil {
		return nil, err
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list, nil
}

// GetDashboardData — zéro déchiffrement, tout en SQL
func (db *Database) GetDashboardData() (*models.DashboardData, error) {
	now := time.Now()
	statsMois, _ := db.GetStatsMensuelles(now.Year(), int(now.Month()))
	derniersMois, _ := db.GetDerniersMois()

	var annee models.StatsAnnee
	annee.Annee = now.Year()
	db.Get(&annee.RevenuBrutTotal, `SELECT COALESCE(SUM(rev_total),0)      FROM stats_mensuelles WHERE annee=?`, now.Year())
	db.Get(&annee.TotalDepenses, `SELECT COALESCE(SUM(dep_total),0)      FROM stats_mensuelles WHERE annee=?`, now.Year())
	db.Get(&annee.TotalDeductible, `SELECT COALESCE(SUM(dep_deductible),0) FROM stats_mensuelles WHERE annee=?`, now.Year())
	db.Get(&annee.MontantAPercevoir, `SELECT COALESCE(SUM(montant_solde),0)  FROM factures WHERE statut NOT IN ('payee','annulee')`)
	db.Get(&annee.FacturesEnRetard, `SELECT COUNT(*) FROM factures WHERE statut='en_retard'`)
	annee.BeneficeNet = annee.RevenuBrutTotal - annee.TotalDepenses
	if annee.RevenuBrutTotal > 0 {
		annee.PctVersSeuilTPS = (annee.RevenuBrutTotal / 30000) * 100
	}

	return &models.DashboardData{
		StatsMois:    *statsMois,
		StatsAnnee:   annee,
		DerniersMois: derniersMois,
		Alertes:      genererAlertes(annee),
	}, nil
}

// GetRapportFiscalAnnuel — GROUP BY natif sur colonnes en clair
func (db *Database) GetRapportFiscalAnnuel(annee int) (*models.RapportFiscalAnnuel, error) {
	var mois []models.StatsMensuelles
	db.Select(&mois, `SELECT * FROM stats_mensuelles WHERE annee=? ORDER BY mois`, annee)

	r := &models.RapportFiscalAnnuel{Annee: annee, ParMois: mois}

	type catRow struct {
		Categorie string  `db:"categorie"`
		Total     float64 `db:"total"`
	}

	var revsParCat []catRow
	db.Select(&revsParCat, `
		SELECT categorie, COALESCE(SUM(montant_brut), 0) as total
		FROM revenus
		WHERE strftime('%Y', date_service) = ? AND statut_paiement != 'annule'
		GROUP BY categorie`, fmt.Sprintf("%d", annee))

	for _, c := range revsParCat {
		r.RevenuBrutTotal += c.Total
		switch c.Categorie {
		case "consultation":
			r.RevConsultations = c.Total
		case "evaluation":
			r.RevEvaluations = c.Total
		case "rapport":
			r.RevRapports = c.Total
		case "forfait":
			r.RevForfaits = c.Total
		case "atelier", "groupe":
			r.RevAteliers += c.Total
		default:
			r.RevAutres += c.Total
		}
	}

	var depsParCat []catRow
	db.Select(&depsParCat, `
		SELECT categorie, COALESCE(SUM(montant_deductible), 0) as total
		FROM depenses
		WHERE strftime('%Y', date_depense) = ? AND deductible = 1
		GROUP BY categorie`, fmt.Sprintf("%d", annee))

	for _, d := range depsParCat {
		r.TotalDepensesDeduct += d.Total
		switch d.Categorie {
		case "bureau":
			r.DepBureau = d.Total
		case "formation":
			r.DepFormation = d.Total
		case "assurance":
			r.DepAssurance = d.Total
		case "logiciel":
			r.DepLogiciel = d.Total
		case "materiel":
			r.DepMateriel = d.Total
		case "deplacement":
			r.DepDeplacement = d.Total
		case "honoraires":
			r.DepHonoraires = d.Total
		case "marketing":
			r.DepMarketing = d.Total
		}
	}

	r.RevenuNetEntreprise = r.RevenuBrutTotal - r.TotalDepensesDeduct
	if r.RevenuNetEntreprise > 0 {
		r.CotisationsRRQ = r.RevenuNetEntreprise * 0.119
		r.DeductionRRQ = r.CotisationsRRQ / 2
		r.RevenuImposable = r.RevenuNetEntreprise - r.DeductionRRQ
	}
	r.SeuilTPSAtteint = r.RevenuBrutTotal >= 30000
	return r, nil
}

// GetRentabiliteParClient — GROUP BY client_id, SUM direct
func (db *Database) GetRentabiliteParClient(s *crypto.CryptoService, annee int) ([]models.RentabiliteParClient, error) {
	type row struct {
		ClientID         int     `db:"client_id"`
		NbSeances        int     `db:"nb_seances"`
		RevenuTotal      float64 `db:"revenu_total"`
		DernierContact   string  `db:"dernier_contact"`
		MontantEnAttente float64 `db:"montant_en_attente"`
	}
	var rows []row
	if err := db.Select(&rows, `
		SELECT
			client_id,
			COUNT(*) as nb_seances,
			COALESCE(SUM(montant_total), 0) as revenu_total,
			MAX(date_service) as dernier_contact,
			COALESCE(SUM(CASE WHEN statut_paiement IN ('en_attente','partiel') THEN montant_total ELSE 0 END), 0) as montant_en_attente
		FROM revenus
		WHERE strftime('%Y', date_service) = ? AND statut_paiement != 'annule'
		GROUP BY client_id
		ORDER BY revenu_total DESC`, fmt.Sprintf("%d", annee)); err != nil {
		return nil, err
	}

	result := make([]models.RentabiliteParClient, 0, len(rows))
	for _, r := range rows {
		client, err := db.GetClientByID(r.ClientID, s)
		if err != nil {
			continue
		}
		revMoyen := 0.0
		if r.NbSeances > 0 {
			revMoyen = r.RevenuTotal / float64(r.NbSeances)
		}
		result = append(result, models.RentabiliteParClient{
			ClientID:         r.ClientID,
			ClientNom:        client.Nom,
			ClientPrenom:     client.Prenom,
			NbSeances:        r.NbSeances,
			RevenuTotal:      r.RevenuTotal,
			RevenuMoyen:      revMoyen,
			DernierContact:   r.DernierContact,
			MontantEnAttente: r.MontantEnAttente,
		})
	}
	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// ALERTES
// ─────────────────────────────────────────────────────────────────────────────

func genererAlertes(annee models.StatsAnnee) []models.AlerteFinance {
	var alertes []models.AlerteFinance
	if annee.RevenuBrutTotal >= 25000 {
		niveau := "warn"
		if annee.RevenuBrutTotal >= 28000 {
			niveau = "danger"
		}
		alertes = append(alertes, models.AlerteFinance{
			Type:    "seuil_tps",
			Niveau:  niveau,
			Message: fmt.Sprintf("%.0f%% du seuil TPS/TVQ atteint — consultez votre comptable.", annee.PctVersSeuilTPS),
			Valeur:  annee.RevenuBrutTotal,
		})
	}
	if annee.FacturesEnRetard > 0 {
		alertes = append(alertes, models.AlerteFinance{
			Type:    "facture_retard",
			Niveau:  "warn",
			Message: fmt.Sprintf("%d facture(s) en retard — %.2f $ à percevoir.", annee.FacturesEnRetard, annee.MontantAPercevoir),
			Valeur:  annee.MontantAPercevoir,
		})
	}
	return alertes
}

// ─────────────────────────────────────────────────────────────────────────────
// HELPERS PRIVÉS — Crypto (paramètres seulement) + utilitaires
// ─────────────────────────────────────────────────────────────────────────────

func decryptParametresFinance(p *models.ParametresFinance, s *crypto.CryptoService) {
	p.NomComplet, _ = s.DecryptStringPtr(p.NomComplet)
	p.TitreProfessionnel, _ = s.DecryptStringPtr(p.TitreProfessionnel)
	p.NoMembreOrdre, _ = s.DecryptStringPtr(p.NoMembreOrdre)
	p.EmailProfessionnel, _ = s.DecryptStringPtr(p.EmailProfessionnel)
	p.TelephoneProfessionnel, _ = s.DecryptStringPtr(p.TelephoneProfessionnel)
	p.AdressePratique, _ = s.DecryptStringPtr(p.AdressePratique)
	p.NoTPS, _ = s.DecryptStringPtr(p.NoTPS)
	p.NoTVQ, _ = s.DecryptStringPtr(p.NoTVQ)
	p.InteracEmail, _ = s.DecryptStringPtr(p.InteracEmail)
}

func encryptParametresFinance(p *models.ParametresFinance, s *crypto.CryptoService) {
	p.NomComplet, _ = s.EncryptStringPtr(p.NomComplet)
	p.TitreProfessionnel, _ = s.EncryptStringPtr(p.TitreProfessionnel)
	p.NoMembreOrdre, _ = s.EncryptStringPtr(p.NoMembreOrdre)
	p.EmailProfessionnel, _ = s.EncryptStringPtr(p.EmailProfessionnel)
	p.TelephoneProfessionnel, _ = s.EncryptStringPtr(p.TelephoneProfessionnel)
	p.AdressePratique, _ = s.EncryptStringPtr(p.AdressePratique)
	p.NoTPS, _ = s.EncryptStringPtr(p.NoTPS)
	p.NoTVQ, _ = s.EncryptStringPtr(p.NoTVQ)
	p.InteracEmail, _ = s.EncryptStringPtr(p.InteracEmail)
}

func (db *Database) getFactureLignes(factureID int) ([]models.FactureLigne, error) {
	var lignes []models.FactureLigne
	return lignes, db.Select(&lignes,
		`SELECT * FROM facture_lignes WHERE facture_id=? ORDER BY ordre`, factureID)
}

func (db *Database) createFactureLigne(factureID, ordre int, req models.CreateFactureLigneRequest) error {
	_, err := db.Exec(`
		INSERT INTO facture_lignes
			(facture_id, service_id, revenu_id, ordre, description, date_service, quantite, unite, tarif_unitaire, montant_ligne)
		VALUES (?,?,?,?,?,?,?,?,?,?)`,
		factureID, req.ServiceID, req.RevenuID, ordre,
		req.Description, req.DateService, req.Quantite, req.Unite,
		req.TarifUnitaire, req.MontantLigne)
	return err
}

func (db *Database) genererNumeroFacture() (string, error) {
	var prefixe string
	var annee, no int
	err := db.QueryRow(
		`SELECT prefixe_facture, annee_facture, prochain_no_facture FROM parametres_finance WHERE id=1`).
		Scan(&prefixe, &annee, &no)
	if err != nil || prefixe == "" {
		prefixe, annee, no = "FAC", time.Now().Year(), 1
	}
	if annee != time.Now().Year() {
		annee = time.Now().Year()
		no = 1
		db.Exec(`UPDATE parametres_finance SET annee_facture=?, prochain_no_facture=1 WHERE id=1`, annee)
	}
	return fmt.Sprintf("%s-%d-%04d", prefixe, annee, no), nil
}
