package services

import (
	"errors"
	"fmt"
	"leopard/internal/crypto"
	database "leopard/internal/db"
	models "leopard/internal/model"
)

type IntervenantService struct {
	db        *database.Database
	cryptoSvc *crypto.CryptoService
}

// NewIntervenantService initialise le service avec ses dépendances
func NewIntervenantService(db *database.Database, cryptoSvc *crypto.CryptoService) *IntervenantService {
	return &IntervenantService{
		db:        db,
		cryptoSvc: cryptoSvc,
	}
}

// GetList récupère tous les intervenants et les déchiffre pour l'affichage
func (s *IntervenantService) GetList() ([]models.Intervenant, error) {
	// 1. Appel au repo (SQL pur)
	intervenants, err := s.db.GetAllIntervenants()
	if err != nil {
		return nil, fmt.Errorf("erreur service GetList: %w", err)
	}

	// 2. Loi 25 : Déchiffrement de la liste
	for i := range intervenants {
		s.decrypt(&intervenants[i])
	}

	return intervenants, nil
}

// GetByID récupère un intervenant spécifique par son ID
func (s *IntervenantService) GetByID(id int) (*models.Intervenant, error) {
	intervenant, err := s.db.GetIntervenantByID(id)
	if err != nil {
		return nil, err
	}

	// Déchiffrement avant envoi au frontend
	s.decrypt(intervenant)
	return intervenant, nil
}

// Save gère la création ou la mise à jour avec chiffrement automatique
func (s *IntervenantService) Save(inter models.Intervenant) error {
	if inter.NomComplet == "" {
		return errors.New("le nom complet est obligatoire")
	}

	// --- CHIFFREMENT SYSTÉMATIQUE (Loi 25) ---
	s.encrypt(&inter)

	if inter.ID > 0 {
		return s.db.UpdateIntervenant(inter)
	}

	_, err := s.db.CreateIntervenant(inter)
	return err
}

// Delete supprime un intervenant
func (s *IntervenantService) Delete(id int) error {
	if id <= 0 {
		return errors.New("ID invalide")
	}
	return s.db.DeleteIntervenant(id)
}

// --- HELPERS PRIVÉS DE CHIFFREMENT COMPLET ---

func (s *IntervenantService) encrypt(i *models.Intervenant) {
	// Emails - TOUS chiffrés
	if i.CourrielPersonnel != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.CourrielPersonnel)
		if err == nil {
			i.CourrielPersonnel = encrypted
		}
	}
	if i.CourrielProfessionnel != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.CourrielProfessionnel)
		if err == nil {
			i.CourrielProfessionnel = encrypted
		}
	}

	// Téléphones - TOUS chiffrés
	if i.Telephone != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.Telephone)
		if err == nil {
			i.Telephone = encrypted
		}
	}
	if i.CellulairePro != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.CellulairePro)
		if err == nil {
			i.CellulairePro = encrypted
		}
	}
	if i.CellulairePerso != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.CellulairePerso)
		if err == nil {
			i.CellulairePerso = encrypted
		}
	}
	if i.Poste != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.Poste)
		if err == nil {
			i.Poste = encrypted
		}
	}

	// Courrier interne
	if i.CourrierInterne != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.CourrierInterne)
		if err == nil {
			i.CourrierInterne = encrypted
		}
	}

	// Note fixe - données sensibles potentielles
	if i.NoteFixe != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.NoteFixe)
		if err == nil {
			i.NoteFixe = encrypted
		}
	}

	// Numéro de permis - donnée professionnelle sensible
	if i.NumeroPermis != "" {
		encrypted, err := s.cryptoSvc.Encrypt(i.NumeroPermis)
		if err == nil {
			i.NumeroPermis = encrypted
		}
	}
	// Remplace le bloc DateNaissance par celui-ci :
	if i.DateNaissance != "" {
		// Plus besoin de .Format() car c'est déjà du texte (YYYY-MM-DD) venant de Vue/FormKit
		encrypted, err := s.cryptoSvc.Encrypt(i.DateNaissance)
		if err == nil {
			i.DateNaissance = encrypted
		}
	}
}

func (s *IntervenantService) decrypt(i *models.Intervenant) {
	// Emails
	if i.CourrielPersonnel != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.CourrielPersonnel)
		if err == nil {
			i.CourrielPersonnel = decrypted
		}
	}
	if i.CourrielProfessionnel != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.CourrielProfessionnel)
		if err == nil {
			i.CourrielProfessionnel = decrypted
		}
	}

	// Téléphones
	if i.Telephone != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.Telephone)
		if err == nil {
			i.Telephone = decrypted
		}
	}
	if i.CellulairePro != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.CellulairePro)
		if err == nil {
			i.CellulairePro = decrypted
		}
	}
	if i.CellulairePerso != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.CellulairePerso)
		if err == nil {
			i.CellulairePerso = decrypted
		}
	}
	if i.Poste != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.Poste)
		if err == nil {
			i.Poste = decrypted
		}
	}

	// Courrier interne
	if i.CourrierInterne != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.CourrierInterne)
		if err == nil {
			i.CourrierInterne = decrypted
		}
	}

	// Date de naissance
	// Remplace le bloc DateNaissance par celui-ci :
	if i.DateNaissance != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.DateNaissance)
		if err == nil {
			i.DateNaissance = decrypted
		}
	}
	// Note fixe
	if i.NoteFixe != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.NoteFixe)
		if err == nil {
			i.NoteFixe = decrypted
		}
	}

	// Numéro de permis
	if i.NumeroPermis != "" {
		decrypted, err := s.cryptoSvc.Decrypt(i.NumeroPermis)
		if err == nil {
			i.NumeroPermis = decrypted
		}
	}
}
