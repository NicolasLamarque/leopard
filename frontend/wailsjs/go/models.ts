export namespace database {
	
	export class ImportStats {
	    nouveaux: number;
	    mis_a_jour: number;
	    errors: string[];
	
	    static createFrom(source: any = {}) {
	        return new ImportStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nouveaux = source["nouveaux"];
	        this.mis_a_jour = source["mis_a_jour"];
	        this.errors = source["errors"];
	    }
	}

}

export namespace main {
	
	export class ClientFolderInfo {
	    path: string;
	    subfoldersCount: number;
	    filesCount: number;
	
	    static createFrom(source: any = {}) {
	        return new ClientFolderInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.subfoldersCount = source["subfoldersCount"];
	        this.filesCount = source["filesCount"];
	    }
	}
	export class ClientFolderResult {
	    success: boolean;
	    path?: string;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new ClientFolderResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.path = source["path"];
	        this.error = source["error"];
	    }
	}
	export class CreateFirstUserRequest {
	    username: string;
	    password: string;
	    fullName: string;
	    role: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateFirstUserRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	        this.fullName = source["fullName"];
	        this.role = source["role"];
	    }
	}
	export class Result {
	    success: boolean;
	    path: string;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.path = source["path"];
	        this.error = source["error"];
	    }
	}

}

export namespace models {
	
	export class AlerteFinance {
	    type: string;
	    niveau: string;
	    message: string;
	    valeur?: number;
	
	    static createFrom(source: any = {}) {
	        return new AlerteFinance(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.niveau = source["niveau"];
	        this.message = source["message"];
	        this.valeur = source["valeur"];
	    }
	}
	export class Appel {
	    id: number;
	    // Go type: time
	    date_appel: any;
	    heure_appel: string;
	    appelant_nom: string;
	    appelant_prenom: string;
	    appelant_telephone: string;
	    appelant_relation: string;
	    client_id?: number;
	    prospect_nom: string;
	    prospect_prenom: string;
	    prospect_telephone: string;
	    type_demande: string;
	    motif_appel: string;
	    priorite: string;
	    statut: string;
	    notes_internes: string;
	    rdv_date?: string;
	    rdv_heure: string;
	    rdv_lieu: string;
	    recu_par: number;
	    assigne_a?: number;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Appel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date_appel = this.convertValues(source["date_appel"], null);
	        this.heure_appel = source["heure_appel"];
	        this.appelant_nom = source["appelant_nom"];
	        this.appelant_prenom = source["appelant_prenom"];
	        this.appelant_telephone = source["appelant_telephone"];
	        this.appelant_relation = source["appelant_relation"];
	        this.client_id = source["client_id"];
	        this.prospect_nom = source["prospect_nom"];
	        this.prospect_prenom = source["prospect_prenom"];
	        this.prospect_telephone = source["prospect_telephone"];
	        this.type_demande = source["type_demande"];
	        this.motif_appel = source["motif_appel"];
	        this.priorite = source["priorite"];
	        this.statut = source["statut"];
	        this.notes_internes = source["notes_internes"];
	        this.rdv_date = source["rdv_date"];
	        this.rdv_heure = source["rdv_heure"];
	        this.rdv_lieu = source["rdv_lieu"];
	        this.recu_par = source["recu_par"];
	        this.assigne_a = source["assigne_a"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class AppelListItem {
	    id: number;
	    // Go type: time
	    date_appel: any;
	    heure_appel: string;
	    appelant_nom: string;
	    appelant_prenom: string;
	    appelant_telephone: string;
	    prospect_nom: string;
	    prospect_prenom: string;
	    type_demande: string;
	    priorite: string;
	    statut: string;
	    client_id?: number;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new AppelListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date_appel = this.convertValues(source["date_appel"], null);
	        this.heure_appel = source["heure_appel"];
	        this.appelant_nom = source["appelant_nom"];
	        this.appelant_prenom = source["appelant_prenom"];
	        this.appelant_telephone = source["appelant_telephone"];
	        this.prospect_nom = source["prospect_nom"];
	        this.prospect_prenom = source["prospect_prenom"];
	        this.type_demande = source["type_demande"];
	        this.priorite = source["priorite"];
	        this.statut = source["statut"];
	        this.client_id = source["client_id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CHSLD {
	    id: number;
	    Region: string;
	    TitreCHSLD: string;
	    Adresse: string;
	    Municipalite: string;
	    CodePostal: string;
	    Telephone: string;
	    telecopieur: string;
	    Poste_Garde_infirmiere: string;
	    Capacite: number;
	    TypeCHSLD: string;
	    Proprietaire: string;
	    DateCertification: string;
	    Statut: string;
	    SourceURL: string;
	    InfosCHSLD: string;
	    Notes: string;
	    DateAjout: string;
	    DateMaj: string;
	
	    static createFrom(source: any = {}) {
	        return new CHSLD(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.Region = source["Region"];
	        this.TitreCHSLD = source["TitreCHSLD"];
	        this.Adresse = source["Adresse"];
	        this.Municipalite = source["Municipalite"];
	        this.CodePostal = source["CodePostal"];
	        this.Telephone = source["Telephone"];
	        this.telecopieur = source["telecopieur"];
	        this.Poste_Garde_infirmiere = source["Poste_Garde_infirmiere"];
	        this.Capacite = source["Capacite"];
	        this.TypeCHSLD = source["TypeCHSLD"];
	        this.Proprietaire = source["Proprietaire"];
	        this.DateCertification = source["DateCertification"];
	        this.Statut = source["Statut"];
	        this.SourceURL = source["SourceURL"];
	        this.InfosCHSLD = source["InfosCHSLD"];
	        this.Notes = source["Notes"];
	        this.DateAjout = source["DateAjout"];
	        this.DateMaj = source["DateMaj"];
	    }
	}
	export class ChangePasswordRequest {
	    current_password: string;
	    new_password: string;
	
	    static createFrom(source: any = {}) {
	        return new ChangePasswordRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current_password = source["current_password"];
	        this.new_password = source["new_password"];
	    }
	}
	export class Client {
	    id: number;
	    nom: string;
	    prenom: string;
	    date_naissance?: string;
	    sexe?: string;
	    lieu_naissance?: string;
	    statut_marital?: string;
	    occupation?: string;
	    employeur?: string;
	    profession?: string;
	    niveau_scolaire?: string;
	    langue_preferee?: string;
	    origine_ethnique?: string;
	    premiere_nation?: string;
	    identite_genre?: string;
	    orientation_sexuelle?: string;
	    religion?: string;
	    premiere_langue?: string;
	    telephone?: string;
	    cellulaire?: string;
	    email?: string;
	    adresse?: string;
	    appartement?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
	    province?: string;
	    numero_assurance_maladie?: string;
	    numero_securite_sociale?: string;
	    no_hcm?: string;
	    no_chaur?: string;
	    no_dossier_leopard?: string;
	    medecin_famille_No_Licence?: string;
	    notaire_id?: number;
	    pharmacie_id?: number;
	    pivot_id?: number;
	    rpa_id?: number;
	    chsld_id?: number;
	    ri_id?: number;
	    procuration_notariee?: string;
	    procuration_bancaire?: string;
	    note_fixe?: string;
	    actif: number;
	    dcd: number;
	    date_deces?: string;
	    created_by?: number;
	    created_at: string;
	    nom_pere?: string;
	    nom_mere?: string;
	    telephone_pere?: string;
	    telephone_mere?: string;
	
	    static createFrom(source: any = {}) {
	        return new Client(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.date_naissance = source["date_naissance"];
	        this.sexe = source["sexe"];
	        this.lieu_naissance = source["lieu_naissance"];
	        this.statut_marital = source["statut_marital"];
	        this.occupation = source["occupation"];
	        this.employeur = source["employeur"];
	        this.profession = source["profession"];
	        this.niveau_scolaire = source["niveau_scolaire"];
	        this.langue_preferee = source["langue_preferee"];
	        this.origine_ethnique = source["origine_ethnique"];
	        this.premiere_nation = source["premiere_nation"];
	        this.identite_genre = source["identite_genre"];
	        this.orientation_sexuelle = source["orientation_sexuelle"];
	        this.religion = source["religion"];
	        this.premiere_langue = source["premiere_langue"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.appartement = source["appartement"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.province = source["province"];
	        this.numero_assurance_maladie = source["numero_assurance_maladie"];
	        this.numero_securite_sociale = source["numero_securite_sociale"];
	        this.no_hcm = source["no_hcm"];
	        this.no_chaur = source["no_chaur"];
	        this.no_dossier_leopard = source["no_dossier_leopard"];
	        this.medecin_famille_No_Licence = source["medecin_famille_No_Licence"];
	        this.notaire_id = source["notaire_id"];
	        this.pharmacie_id = source["pharmacie_id"];
	        this.pivot_id = source["pivot_id"];
	        this.rpa_id = source["rpa_id"];
	        this.chsld_id = source["chsld_id"];
	        this.ri_id = source["ri_id"];
	        this.procuration_notariee = source["procuration_notariee"];
	        this.procuration_bancaire = source["procuration_bancaire"];
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.dcd = source["dcd"];
	        this.date_deces = source["date_deces"];
	        this.created_by = source["created_by"];
	        this.created_at = source["created_at"];
	        this.nom_pere = source["nom_pere"];
	        this.nom_mere = source["nom_mere"];
	        this.telephone_pere = source["telephone_pere"];
	        this.telephone_mere = source["telephone_mere"];
	    }
	}
	export class ClientPharmacieInfo {
	    id: number;
	    nom: string;
	    prenom: string;
	    no_dossier_leopard: string;
	    dcd: number;
	    actif: number;
	
	    static createFrom(source: any = {}) {
	        return new ClientPharmacieInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.no_dossier_leopard = source["no_dossier_leopard"];
	        this.dcd = source["dcd"];
	        this.actif = source["actif"];
	    }
	}
	export class Contact {
	    id: number;
	    nom: string;
	    prenom: string;
	    telephone?: string;
	    cellulaire?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
	    email?: string;
	    employeur?: string;
	    profession?: string;
	    relation_client?: string;
	    lien_familial?: string;
	    force_lien: number;
	    contact_urgence: number;
	    aidant_naturel: number;
	    type_de_contact?: string;
	    procuration_bancaire: number;
	    procuration_notariee: number;
	    note_fixe?: string;
	    relation_type: string;
	    client_id: number;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Contact(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.email = source["email"];
	        this.employeur = source["employeur"];
	        this.profession = source["profession"];
	        this.relation_client = source["relation_client"];
	        this.lien_familial = source["lien_familial"];
	        this.force_lien = source["force_lien"];
	        this.contact_urgence = source["contact_urgence"];
	        this.aidant_naturel = source["aidant_naturel"];
	        this.type_de_contact = source["type_de_contact"];
	        this.procuration_bancaire = source["procuration_bancaire"];
	        this.procuration_notariee = source["procuration_notariee"];
	        this.note_fixe = source["note_fixe"];
	        this.relation_type = source["relation_type"];
	        this.client_id = source["client_id"];
	        this.created_at = source["created_at"];
	    }
	}
	export class Contrat {
	    ID: number;
	    ClientID: number;
	    ServiceID?: number;
	    TypeContrat: string;
	    Statut: string;
	    DateDebut: string;
	    DateFin: string;
	    RenouvellementAuto: number;
	    ConsentementSigne: number;
	    DateSignature: string;
	    ModeFacturation: string;
	    TauxHoraire: number;
	    TauxForfait: number;
	    DureeSeanceMin?: number;
	    FrequenceSeances: string;
	    ClauseObjet: string;
	    ClauseObjectifs: string;
	    ClauseServicesInclus: string;
	    ClauseServicesExclus: string;
	    ClauseTarification: string;
	    ClausePaiementRetard: string;
	    ClauseAnnulation: string;
	    ClauseAbsenceProlongee: string;
	    ClauseConfidentialite: string;
	    ClauseLimitesConfidentialite: string;
	    ClauseDossierClient: string;
	    ClauseCommunication: string;
	    ClauseUrgences: string;
	    ClauseResiliationClient: string;
	    ClauseResiliationTS: string;
	    ClausesSpecifiques: string;
	    NotesInternes: string;
	    CreatedBy?: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Contrat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.ClientID = source["ClientID"];
	        this.ServiceID = source["ServiceID"];
	        this.TypeContrat = source["TypeContrat"];
	        this.Statut = source["Statut"];
	        this.DateDebut = source["DateDebut"];
	        this.DateFin = source["DateFin"];
	        this.RenouvellementAuto = source["RenouvellementAuto"];
	        this.ConsentementSigne = source["ConsentementSigne"];
	        this.DateSignature = source["DateSignature"];
	        this.ModeFacturation = source["ModeFacturation"];
	        this.TauxHoraire = source["TauxHoraire"];
	        this.TauxForfait = source["TauxForfait"];
	        this.DureeSeanceMin = source["DureeSeanceMin"];
	        this.FrequenceSeances = source["FrequenceSeances"];
	        this.ClauseObjet = source["ClauseObjet"];
	        this.ClauseObjectifs = source["ClauseObjectifs"];
	        this.ClauseServicesInclus = source["ClauseServicesInclus"];
	        this.ClauseServicesExclus = source["ClauseServicesExclus"];
	        this.ClauseTarification = source["ClauseTarification"];
	        this.ClausePaiementRetard = source["ClausePaiementRetard"];
	        this.ClauseAnnulation = source["ClauseAnnulation"];
	        this.ClauseAbsenceProlongee = source["ClauseAbsenceProlongee"];
	        this.ClauseConfidentialite = source["ClauseConfidentialite"];
	        this.ClauseLimitesConfidentialite = source["ClauseLimitesConfidentialite"];
	        this.ClauseDossierClient = source["ClauseDossierClient"];
	        this.ClauseCommunication = source["ClauseCommunication"];
	        this.ClauseUrgences = source["ClauseUrgences"];
	        this.ClauseResiliationClient = source["ClauseResiliationClient"];
	        this.ClauseResiliationTS = source["ClauseResiliationTS"];
	        this.ClausesSpecifiques = source["ClausesSpecifiques"];
	        this.NotesInternes = source["NotesInternes"];
	        this.CreatedBy = source["CreatedBy"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CreateAppelRequest {
	    date_appel: string;
	    heure_appel: string;
	    appelant_nom: string;
	    appelant_prenom: string;
	    appelant_telephone: string;
	    appelant_relation: string;
	    client_id?: number;
	    prospect_nom: string;
	    prospect_prenom: string;
	    prospect_telephone: string;
	    type_demande: string;
	    motif_appel: string;
	    priorite: string;
	    statut: string;
	    notes_internes: string;
	    rdv_date: string;
	    rdv_heure: string;
	    rdv_lieu: string;
	    assigne_a?: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateAppelRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date_appel = source["date_appel"];
	        this.heure_appel = source["heure_appel"];
	        this.appelant_nom = source["appelant_nom"];
	        this.appelant_prenom = source["appelant_prenom"];
	        this.appelant_telephone = source["appelant_telephone"];
	        this.appelant_relation = source["appelant_relation"];
	        this.client_id = source["client_id"];
	        this.prospect_nom = source["prospect_nom"];
	        this.prospect_prenom = source["prospect_prenom"];
	        this.prospect_telephone = source["prospect_telephone"];
	        this.type_demande = source["type_demande"];
	        this.motif_appel = source["motif_appel"];
	        this.priorite = source["priorite"];
	        this.statut = source["statut"];
	        this.notes_internes = source["notes_internes"];
	        this.rdv_date = source["rdv_date"];
	        this.rdv_heure = source["rdv_heure"];
	        this.rdv_lieu = source["rdv_lieu"];
	        this.assigne_a = source["assigne_a"];
	    }
	}
	export class CreateClientRequest {
	    nom: string;
	    prenom: string;
	    date_naissance?: string;
	    sexe?: string;
	    lieu_naissance?: string;
	    statut_marital?: string;
	    occupation?: string;
	    employeur?: string;
	    profession?: string;
	    niveau_scolaire?: string;
	    langue_preferee?: string;
	    origine_ethnique?: string;
	    premiere_nation?: string;
	    identite_genre?: string;
	    orientation_sexuelle?: string;
	    religion?: string;
	    premiere_langue?: string;
	    telephone?: string;
	    cellulaire?: string;
	    email?: string;
	    adresse?: string;
	    appartement?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
	    province?: string;
	    numero_assurance_maladie?: string;
	    numero_securite_sociale?: string;
	    no_hcm?: string;
	    no_chaur?: string;
	    no_dossier_leopard?: string;
	    medecin_famille_No_Licence?: string;
	    notaire_id?: number;
	    pharmacie_id?: number;
	    pivot_id?: number;
	    rpa_id?: number;
	    chsld_id?: number;
	    ri_id?: number;
	    procuration_notariee: string;
	    procuration_bancaire: string;
	    note_fixe?: string;
	    actif: number;
	    dcd: number;
	    date_deces?: string;
	    nom_pere?: string;
	    nom_mere?: string;
	    telephone_pere?: string;
	    telephone_mere?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateClientRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.date_naissance = source["date_naissance"];
	        this.sexe = source["sexe"];
	        this.lieu_naissance = source["lieu_naissance"];
	        this.statut_marital = source["statut_marital"];
	        this.occupation = source["occupation"];
	        this.employeur = source["employeur"];
	        this.profession = source["profession"];
	        this.niveau_scolaire = source["niveau_scolaire"];
	        this.langue_preferee = source["langue_preferee"];
	        this.origine_ethnique = source["origine_ethnique"];
	        this.premiere_nation = source["premiere_nation"];
	        this.identite_genre = source["identite_genre"];
	        this.orientation_sexuelle = source["orientation_sexuelle"];
	        this.religion = source["religion"];
	        this.premiere_langue = source["premiere_langue"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.appartement = source["appartement"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.province = source["province"];
	        this.numero_assurance_maladie = source["numero_assurance_maladie"];
	        this.numero_securite_sociale = source["numero_securite_sociale"];
	        this.no_hcm = source["no_hcm"];
	        this.no_chaur = source["no_chaur"];
	        this.no_dossier_leopard = source["no_dossier_leopard"];
	        this.medecin_famille_No_Licence = source["medecin_famille_No_Licence"];
	        this.notaire_id = source["notaire_id"];
	        this.pharmacie_id = source["pharmacie_id"];
	        this.pivot_id = source["pivot_id"];
	        this.rpa_id = source["rpa_id"];
	        this.chsld_id = source["chsld_id"];
	        this.ri_id = source["ri_id"];
	        this.procuration_notariee = source["procuration_notariee"];
	        this.procuration_bancaire = source["procuration_bancaire"];
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.dcd = source["dcd"];
	        this.date_deces = source["date_deces"];
	        this.nom_pere = source["nom_pere"];
	        this.nom_mere = source["nom_mere"];
	        this.telephone_pere = source["telephone_pere"];
	        this.telephone_mere = source["telephone_mere"];
	    }
	}
	export class CreateContactRequest {
	    nom: string;
	    prenom: string;
	    telephone?: string;
	    cellulaire?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
	    email?: string;
	    employeur?: string;
	    profession?: string;
	    relation_client?: string;
	    lien_familial?: string;
	    force_lien: number;
	    contact_urgence: number;
	    aidant_naturel: number;
	    type_de_contact?: string;
	    procuration_bancaire: number;
	    procuration_notariee: number;
	    note_fixe?: string;
	    relation_type: string;
	    client_id: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateContactRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.email = source["email"];
	        this.employeur = source["employeur"];
	        this.profession = source["profession"];
	        this.relation_client = source["relation_client"];
	        this.lien_familial = source["lien_familial"];
	        this.force_lien = source["force_lien"];
	        this.contact_urgence = source["contact_urgence"];
	        this.aidant_naturel = source["aidant_naturel"];
	        this.type_de_contact = source["type_de_contact"];
	        this.procuration_bancaire = source["procuration_bancaire"];
	        this.procuration_notariee = source["procuration_notariee"];
	        this.note_fixe = source["note_fixe"];
	        this.relation_type = source["relation_type"];
	        this.client_id = source["client_id"];
	    }
	}
	export class CreateContratRequest {
	    client_id: number;
	    service_id?: number;
	    type_contrat: string;
	    date_debut: string;
	    date_fin?: string;
	    renouvellement_auto: number;
	    mode_facturation: string;
	    taux_horaire: number;
	    taux_forfait: number;
	    duree_seance_min?: number;
	    frequence_seances?: string;
	    clause_objet?: string;
	    clause_objectifs?: string;
	    clause_services_inclus?: string;
	    clause_services_exclus?: string;
	    clause_tarification?: string;
	    clause_paiement_retard?: string;
	    clause_annulation?: string;
	    clause_absence_prolongee?: string;
	    clause_confidentialite?: string;
	    clause_limites_confidentialite?: string;
	    clause_dossier_client?: string;
	    clause_communication?: string;
	    clause_urgences?: string;
	    clause_resiliation_client?: string;
	    clause_resiliation_ts?: string;
	    clauses_specifiques?: string;
	    notes_internes?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateContratRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.service_id = source["service_id"];
	        this.type_contrat = source["type_contrat"];
	        this.date_debut = source["date_debut"];
	        this.date_fin = source["date_fin"];
	        this.renouvellement_auto = source["renouvellement_auto"];
	        this.mode_facturation = source["mode_facturation"];
	        this.taux_horaire = source["taux_horaire"];
	        this.taux_forfait = source["taux_forfait"];
	        this.duree_seance_min = source["duree_seance_min"];
	        this.frequence_seances = source["frequence_seances"];
	        this.clause_objet = source["clause_objet"];
	        this.clause_objectifs = source["clause_objectifs"];
	        this.clause_services_inclus = source["clause_services_inclus"];
	        this.clause_services_exclus = source["clause_services_exclus"];
	        this.clause_tarification = source["clause_tarification"];
	        this.clause_paiement_retard = source["clause_paiement_retard"];
	        this.clause_annulation = source["clause_annulation"];
	        this.clause_absence_prolongee = source["clause_absence_prolongee"];
	        this.clause_confidentialite = source["clause_confidentialite"];
	        this.clause_limites_confidentialite = source["clause_limites_confidentialite"];
	        this.clause_dossier_client = source["clause_dossier_client"];
	        this.clause_communication = source["clause_communication"];
	        this.clause_urgences = source["clause_urgences"];
	        this.clause_resiliation_client = source["clause_resiliation_client"];
	        this.clause_resiliation_ts = source["clause_resiliation_ts"];
	        this.clauses_specifiques = source["clauses_specifiques"];
	        this.notes_internes = source["notes_internes"];
	    }
	}
	export class CreateDepenseRequest {
	    date_depense: string;
	    categorie: string;
	    deductible: number;
	    pct_utilisation_affaires: number;
	    est_kilometrage: number;
	    sous_total: number;
	    montant_tps: number;
	    montant_tvq: number;
	    montant_total: number;
	    montant_deductible: number;
	    km_parcourus: number;
	    taux_km_utilise: number;
	    fournisseur?: string;
	    no_recu?: string;
	    description: string;
	    notes?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateDepenseRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date_depense = source["date_depense"];
	        this.categorie = source["categorie"];
	        this.deductible = source["deductible"];
	        this.pct_utilisation_affaires = source["pct_utilisation_affaires"];
	        this.est_kilometrage = source["est_kilometrage"];
	        this.sous_total = source["sous_total"];
	        this.montant_tps = source["montant_tps"];
	        this.montant_tvq = source["montant_tvq"];
	        this.montant_total = source["montant_total"];
	        this.montant_deductible = source["montant_deductible"];
	        this.km_parcourus = source["km_parcourus"];
	        this.taux_km_utilise = source["taux_km_utilise"];
	        this.fournisseur = source["fournisseur"];
	        this.no_recu = source["no_recu"];
	        this.description = source["description"];
	        this.notes = source["notes"];
	    }
	}
	export class CreateEvaluationRequest {
	    client_id: number;
	    no_leopard: string;
	    no_eval_leopard: string;
	    type_evaluation?: string;
	    contexte_evaluation?: string;
	    motif_reference?: string;
	    objet_evaluation?: string;
	    capacites_cognitives?: string;
	    etat_sante_physique?: string;
	    dimensions_psycho_sociales?: string;
	    roles_sociaux?: string;
	    reseau_social_soutien?: string;
	    analyse_clinique?: string;
	    opinion_professionnelle?: string;
	    recommandations?: string;
	    verrouille: number;
	    isDraft: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateEvaluationRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.no_leopard = source["no_leopard"];
	        this.no_eval_leopard = source["no_eval_leopard"];
	        this.type_evaluation = source["type_evaluation"];
	        this.contexte_evaluation = source["contexte_evaluation"];
	        this.motif_reference = source["motif_reference"];
	        this.objet_evaluation = source["objet_evaluation"];
	        this.capacites_cognitives = source["capacites_cognitives"];
	        this.etat_sante_physique = source["etat_sante_physique"];
	        this.dimensions_psycho_sociales = source["dimensions_psycho_sociales"];
	        this.roles_sociaux = source["roles_sociaux"];
	        this.reseau_social_soutien = source["reseau_social_soutien"];
	        this.analyse_clinique = source["analyse_clinique"];
	        this.opinion_professionnelle = source["opinion_professionnelle"];
	        this.recommandations = source["recommandations"];
	        this.verrouille = source["verrouille"];
	        this.isDraft = source["isDraft"];
	    }
	}
	export class CreateFactureLigneRequest {
	    service_id?: number;
	    revenu_id?: number;
	    ordre: number;
	    description: string;
	    date_service?: string;
	    quantite: number;
	    unite?: string;
	    tarif_unitaire: number;
	    montant_ligne: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateFactureLigneRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.service_id = source["service_id"];
	        this.revenu_id = source["revenu_id"];
	        this.ordre = source["ordre"];
	        this.description = source["description"];
	        this.date_service = source["date_service"];
	        this.quantite = source["quantite"];
	        this.unite = source["unite"];
	        this.tarif_unitaire = source["tarif_unitaire"];
	        this.montant_ligne = source["montant_ligne"];
	    }
	}
	export class CreateFactureRequest {
	    client_id: number;
	    date_emission: string;
	    date_echeance: string;
	    exempte_taxes: number;
	    avec_tps: number;
	    avec_tvq: number;
	    jours_paiement: number;
	    montant_sous_total: number;
	    montant_tps: number;
	    montant_tvq: number;
	    montant_total: number;
	    titre_facture?: string;
	    notes_client?: string;
	    notes_internes?: string;
	    tiers_payant_nom?: string;
	    montant_tiers_payant: number;
	    montant_du_client: number;
	    lignes: CreateFactureLigneRequest[];
	
	    static createFrom(source: any = {}) {
	        return new CreateFactureRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.date_emission = source["date_emission"];
	        this.date_echeance = source["date_echeance"];
	        this.exempte_taxes = source["exempte_taxes"];
	        this.avec_tps = source["avec_tps"];
	        this.avec_tvq = source["avec_tvq"];
	        this.jours_paiement = source["jours_paiement"];
	        this.montant_sous_total = source["montant_sous_total"];
	        this.montant_tps = source["montant_tps"];
	        this.montant_tvq = source["montant_tvq"];
	        this.montant_total = source["montant_total"];
	        this.titre_facture = source["titre_facture"];
	        this.notes_client = source["notes_client"];
	        this.notes_internes = source["notes_internes"];
	        this.tiers_payant_nom = source["tiers_payant_nom"];
	        this.montant_tiers_payant = source["montant_tiers_payant"];
	        this.montant_du_client = source["montant_du_client"];
	        this.lignes = this.convertValues(source["lignes"], CreateFactureLigneRequest);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CreateMedecinRequest {
	    licence: string;
	    civilite: string;
	    prenom: string;
	    nom: string;
	    nomComplet: string;
	    statut: string;
	    specialisation: string;
	    service: string;
	    departement: string;
	    installationPrincipale: string;
	    rls: string;
	    telephone: string;
	    extension: string;
	    cellulaire: string;
	    email: string;
	    adresse: string;
	    codePostal: string;
	    ville: string;
	    pays: string;
	    noteFixe: string;
	    actif: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateMedecinRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.licence = source["licence"];
	        this.civilite = source["civilite"];
	        this.prenom = source["prenom"];
	        this.nom = source["nom"];
	        this.nomComplet = source["nomComplet"];
	        this.statut = source["statut"];
	        this.specialisation = source["specialisation"];
	        this.service = source["service"];
	        this.departement = source["departement"];
	        this.installationPrincipale = source["installationPrincipale"];
	        this.rls = source["rls"];
	        this.telephone = source["telephone"];
	        this.extension = source["extension"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.codePostal = source["codePostal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.noteFixe = source["noteFixe"];
	        this.actif = source["actif"];
	    }
	}
	export class CreateNotaireRequest {
	    civilite: string;
	    prenom: string;
	    nom: string;
	    telephone?: string;
	    cellulaire?: string;
	    telecopieur?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    email?: string;
	    secteur_activite?: string;
	    note_fixe?: string;
	    created_by: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateNotaireRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.civilite = source["civilite"];
	        this.prenom = source["prenom"];
	        this.nom = source["nom"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.telecopieur = source["telecopieur"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.email = source["email"];
	        this.secteur_activite = source["secteur_activite"];
	        this.note_fixe = source["note_fixe"];
	        this.created_by = source["created_by"];
	    }
	}
	export class CreateNoteRequest {
	    client_id: number;
	    client_NoRAMQ?: string;
	    client_Nom?: string;
	    client_Prenom?: string;
	    client_date_naissance?: string;
	    client_Telephone?: string;
	    client_Cellulaire?: string;
	    client_NoLeopard?: string;
	    client_Adresse?: string;
	    client_appartement?: string;
	    client_code_postal?: string;
	    client_ville?: string;
	    client_pays?: string;
	    client_province?: string;
	    user_id: number;
	    date_intervention?: string;
	    heure_intervention?: string;
	    duree_intervention?: string;
	    mode_intervention?: string;
	    type_intervention?: string;
	    type_note?: string;
	    titre?: string;
	    contenu?: string;
	    note_tardive: number;
	    note_de_tier: number;
	    note_liee_id?: number;
	    type_lien?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateNoteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.client_NoRAMQ = source["client_NoRAMQ"];
	        this.client_Nom = source["client_Nom"];
	        this.client_Prenom = source["client_Prenom"];
	        this.client_date_naissance = source["client_date_naissance"];
	        this.client_Telephone = source["client_Telephone"];
	        this.client_Cellulaire = source["client_Cellulaire"];
	        this.client_NoLeopard = source["client_NoLeopard"];
	        this.client_Adresse = source["client_Adresse"];
	        this.client_appartement = source["client_appartement"];
	        this.client_code_postal = source["client_code_postal"];
	        this.client_ville = source["client_ville"];
	        this.client_pays = source["client_pays"];
	        this.client_province = source["client_province"];
	        this.user_id = source["user_id"];
	        this.date_intervention = source["date_intervention"];
	        this.heure_intervention = source["heure_intervention"];
	        this.duree_intervention = source["duree_intervention"];
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.titre = source["titre"];
	        this.contenu = source["contenu"];
	        this.note_tardive = source["note_tardive"];
	        this.note_de_tier = source["note_de_tier"];
	        this.note_liee_id = source["note_liee_id"];
	        this.type_lien = source["type_lien"];
	    }
	}
	export class CreateOrganisationRequest {
	    type_org: string;
	    no_fournisseur?: string;
	    conditions_paiement: number;
	    taux_tps_exempt: number;
	    mode_paiement_pref: string;
	    nom?: string;
	    acronyme?: string;
	    no_organisme?: string;
	    adresse?: string;
	    ville?: string;
	    code_postal?: string;
	    province?: string;
	    pays?: string;
	    telephone?: string;
	    telecopieur?: string;
	    email_general?: string;
	    site_web?: string;
	    contact_nom?: string;
	    contact_titre?: string;
	    contact_telephone?: string;
	    contact_email?: string;
	    notes?: string;
	    conditions_speciales?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateOrganisationRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type_org = source["type_org"];
	        this.no_fournisseur = source["no_fournisseur"];
	        this.conditions_paiement = source["conditions_paiement"];
	        this.taux_tps_exempt = source["taux_tps_exempt"];
	        this.mode_paiement_pref = source["mode_paiement_pref"];
	        this.nom = source["nom"];
	        this.acronyme = source["acronyme"];
	        this.no_organisme = source["no_organisme"];
	        this.adresse = source["adresse"];
	        this.ville = source["ville"];
	        this.code_postal = source["code_postal"];
	        this.province = source["province"];
	        this.pays = source["pays"];
	        this.telephone = source["telephone"];
	        this.telecopieur = source["telecopieur"];
	        this.email_general = source["email_general"];
	        this.site_web = source["site_web"];
	        this.contact_nom = source["contact_nom"];
	        this.contact_titre = source["contact_titre"];
	        this.contact_telephone = source["contact_telephone"];
	        this.contact_email = source["contact_email"];
	        this.notes = source["notes"];
	        this.conditions_speciales = source["conditions_speciales"];
	    }
	}
	export class CreatePaiementRequest {
	    facture_id: number;
	    date_paiement: string;
	    source: string;
	    mode_paiement: string;
	    montant: number;
	    notes?: string;
	    reference?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreatePaiementRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.facture_id = source["facture_id"];
	        this.date_paiement = source["date_paiement"];
	        this.source = source["source"];
	        this.mode_paiement = source["mode_paiement"];
	        this.montant = source["montant"];
	        this.notes = source["notes"];
	        this.reference = source["reference"];
	    }
	}
	export class CreatePayeurRequest {
	    type_payeur: string;
	    client_id?: number;
	    contact_id?: number;
	    organisation_id?: number;
	    est_defaut: number;
	    nom_facturation?: string;
	    adresse_facturation?: string;
	    email_facturation?: string;
	    notes?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreatePayeurRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type_payeur = source["type_payeur"];
	        this.client_id = source["client_id"];
	        this.contact_id = source["contact_id"];
	        this.organisation_id = source["organisation_id"];
	        this.est_defaut = source["est_defaut"];
	        this.nom_facturation = source["nom_facturation"];
	        this.adresse_facturation = source["adresse_facturation"];
	        this.email_facturation = source["email_facturation"];
	        this.notes = source["notes"];
	    }
	}
	export class CreatePlanRequest {
	    client_id: number;
	    titre: string;
	    type_plan: string;
	    date_debut?: string;
	    date_fin_prevue?: string;
	    date_revision_prevue?: string;
	    contexte?: string;
	    problematique?: string;
	    forces?: string;
	    objectifs?: string;
	    interventions?: string;
	    resultats?: string;
	    ententes?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreatePlanRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.titre = source["titre"];
	        this.type_plan = source["type_plan"];
	        this.date_debut = source["date_debut"];
	        this.date_fin_prevue = source["date_fin_prevue"];
	        this.date_revision_prevue = source["date_revision_prevue"];
	        this.contexte = source["contexte"];
	        this.problematique = source["problematique"];
	        this.forces = source["forces"];
	        this.objectifs = source["objectifs"];
	        this.interventions = source["interventions"];
	        this.resultats = source["resultats"];
	        this.ententes = source["ententes"];
	    }
	}
	export class CreateRevenuRequest {
	    client_id: number;
	    service_id?: number;
	    facture_id?: number;
	    date_service: string;
	    categorie: string;
	    mode_facturation: string;
	    statut_paiement: string;
	    mode_paiement: string;
	    date_paiement?: string;
	    duree_heures: number;
	    taux_horaire_applique: number;
	    montant_brut: number;
	    montant_tps: number;
	    montant_tvq: number;
	    montant_total: number;
	    description?: string;
	    notes?: string;
	    reference_paiement?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateRevenuRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.service_id = source["service_id"];
	        this.facture_id = source["facture_id"];
	        this.date_service = source["date_service"];
	        this.categorie = source["categorie"];
	        this.mode_facturation = source["mode_facturation"];
	        this.statut_paiement = source["statut_paiement"];
	        this.mode_paiement = source["mode_paiement"];
	        this.date_paiement = source["date_paiement"];
	        this.duree_heures = source["duree_heures"];
	        this.taux_horaire_applique = source["taux_horaire_applique"];
	        this.montant_brut = source["montant_brut"];
	        this.montant_tps = source["montant_tps"];
	        this.montant_tvq = source["montant_tvq"];
	        this.montant_total = source["montant_total"];
	        this.description = source["description"];
	        this.notes = source["notes"];
	        this.reference_paiement = source["reference_paiement"];
	    }
	}
	export class CreateServiceRequest {
	    code: string;
	    categorie: string;
	    type_livraison: string;
	    actif: number;
	    ordre_affichage: number;
	    nom: string;
	    description_courte?: string;
	    description_longue?: string;
	    public_cible?: string;
	    notes_internes?: string;
	    mode_tarification: string;
	    duree_minutes?: number;
	    taux_horaire?: number;
	    tarif_unitaire?: number;
	    tarif_min?: number;
	    tarif_max?: number;
	    exempte_taxes: number;
	    tps_applicable: number;
	    tvq_applicable: number;
	    nb_places_max?: number;
	    nb_seances_forfait?: number;
	    duree_programme_semaines?: number;
	    format_tract?: string;
	    couleur_tract?: string;
	    duree_video_minutes?: number;
	    url_ressource?: string;
	    tags?: string;
	    icone?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateServiceRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.categorie = source["categorie"];
	        this.type_livraison = source["type_livraison"];
	        this.actif = source["actif"];
	        this.ordre_affichage = source["ordre_affichage"];
	        this.nom = source["nom"];
	        this.description_courte = source["description_courte"];
	        this.description_longue = source["description_longue"];
	        this.public_cible = source["public_cible"];
	        this.notes_internes = source["notes_internes"];
	        this.mode_tarification = source["mode_tarification"];
	        this.duree_minutes = source["duree_minutes"];
	        this.taux_horaire = source["taux_horaire"];
	        this.tarif_unitaire = source["tarif_unitaire"];
	        this.tarif_min = source["tarif_min"];
	        this.tarif_max = source["tarif_max"];
	        this.exempte_taxes = source["exempte_taxes"];
	        this.tps_applicable = source["tps_applicable"];
	        this.tvq_applicable = source["tvq_applicable"];
	        this.nb_places_max = source["nb_places_max"];
	        this.nb_seances_forfait = source["nb_seances_forfait"];
	        this.duree_programme_semaines = source["duree_programme_semaines"];
	        this.format_tract = source["format_tract"];
	        this.couleur_tract = source["couleur_tract"];
	        this.duree_video_minutes = source["duree_video_minutes"];
	        this.url_ressource = source["url_ressource"];
	        this.tags = source["tags"];
	        this.icone = source["icone"];
	    }
	}
	export class StatsAnnee {
	    annee: number;
	    revenu_brut_total: number;
	    total_depenses: number;
	    total_deductible: number;
	    benefice_net: number;
	    pct_vers_seuil_tps: number;
	    montant_a_percevoir: number;
	    factures_en_retard: number;
	
	    static createFrom(source: any = {}) {
	        return new StatsAnnee(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.annee = source["annee"];
	        this.revenu_brut_total = source["revenu_brut_total"];
	        this.total_depenses = source["total_depenses"];
	        this.total_deductible = source["total_deductible"];
	        this.benefice_net = source["benefice_net"];
	        this.pct_vers_seuil_tps = source["pct_vers_seuil_tps"];
	        this.montant_a_percevoir = source["montant_a_percevoir"];
	        this.factures_en_retard = source["factures_en_retard"];
	    }
	}
	export class StatsMensuelles {
	    ID: number;
	    Annee: number;
	    Mois: number;
	    RevTotal: number;
	    RevConsultations: number;
	    RevEvaluations: number;
	    RevRapports: number;
	    RevForfaits: number;
	    RevAteliers: number;
	    RevAutres: number;
	    NbSeances: number;
	    NbClientsUniques: number;
	    DepTotal: number;
	    DepDeductible: number;
	    DepBureau: number;
	    DepFormation: number;
	    DepAssurance: number;
	    DepLogiciel: number;
	    DepMateriel: number;
	    DepDeplacement: number;
	    DepHonoraires: number;
	    DepMarketing: number;
	    DepAutres: number;
	    BeneficeNet: number;
	    CumulAnnuelRevenus: number;
	    PctMarge: number;
	    NbFacturesEmises: number;
	    NbFacturesPayees: number;
	    NbFacturesEnRetard: number;
	    MontantAPercevoir: number;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new StatsMensuelles(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Annee = source["Annee"];
	        this.Mois = source["Mois"];
	        this.RevTotal = source["RevTotal"];
	        this.RevConsultations = source["RevConsultations"];
	        this.RevEvaluations = source["RevEvaluations"];
	        this.RevRapports = source["RevRapports"];
	        this.RevForfaits = source["RevForfaits"];
	        this.RevAteliers = source["RevAteliers"];
	        this.RevAutres = source["RevAutres"];
	        this.NbSeances = source["NbSeances"];
	        this.NbClientsUniques = source["NbClientsUniques"];
	        this.DepTotal = source["DepTotal"];
	        this.DepDeductible = source["DepDeductible"];
	        this.DepBureau = source["DepBureau"];
	        this.DepFormation = source["DepFormation"];
	        this.DepAssurance = source["DepAssurance"];
	        this.DepLogiciel = source["DepLogiciel"];
	        this.DepMateriel = source["DepMateriel"];
	        this.DepDeplacement = source["DepDeplacement"];
	        this.DepHonoraires = source["DepHonoraires"];
	        this.DepMarketing = source["DepMarketing"];
	        this.DepAutres = source["DepAutres"];
	        this.BeneficeNet = source["BeneficeNet"];
	        this.CumulAnnuelRevenus = source["CumulAnnuelRevenus"];
	        this.PctMarge = source["PctMarge"];
	        this.NbFacturesEmises = source["NbFacturesEmises"];
	        this.NbFacturesPayees = source["NbFacturesPayees"];
	        this.NbFacturesEnRetard = source["NbFacturesEnRetard"];
	        this.MontantAPercevoir = source["MontantAPercevoir"];
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DashboardData {
	    stats_mois: StatsMensuelles;
	    stats_annee: StatsAnnee;
	    derniers_mois: StatsMensuelles[];
	    alertes: AlerteFinance[];
	
	    static createFrom(source: any = {}) {
	        return new DashboardData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.stats_mois = this.convertValues(source["stats_mois"], StatsMensuelles);
	        this.stats_annee = this.convertValues(source["stats_annee"], StatsAnnee);
	        this.derniers_mois = this.convertValues(source["derniers_mois"], StatsMensuelles);
	        this.alertes = this.convertValues(source["alertes"], AlerteFinance);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Depense {
	    ID: number;
	    DateDepense: string;
	    Categorie: string;
	    Deductible: number;
	    PctUtilisationAffaires: number;
	    EstKilometrage: number;
	    SousTotal: number;
	    MontantTPS: number;
	    MontantTVQ: number;
	    MontantTotal: number;
	    MontantDeductible: number;
	    KMParcourus: number;
	    TauxKMUtilise: number;
	    Fournisseur: string;
	    NoRecu: string;
	    Description: string;
	    Notes: string;
	    CheminRecu: string;
	    CreatedBy?: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Depense(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.DateDepense = source["DateDepense"];
	        this.Categorie = source["Categorie"];
	        this.Deductible = source["Deductible"];
	        this.PctUtilisationAffaires = source["PctUtilisationAffaires"];
	        this.EstKilometrage = source["EstKilometrage"];
	        this.SousTotal = source["SousTotal"];
	        this.MontantTPS = source["MontantTPS"];
	        this.MontantTVQ = source["MontantTVQ"];
	        this.MontantTotal = source["MontantTotal"];
	        this.MontantDeductible = source["MontantDeductible"];
	        this.KMParcourus = source["KMParcourus"];
	        this.TauxKMUtilise = source["TauxKMUtilise"];
	        this.Fournisseur = source["Fournisseur"];
	        this.NoRecu = source["NoRecu"];
	        this.Description = source["Description"];
	        this.Notes = source["Notes"];
	        this.CheminRecu = source["CheminRecu"];
	        this.CreatedBy = source["CreatedBy"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class EvaluationSocialeDetail {
	    id: number;
	    client_id: number;
	    no_eval_leopard: string;
	    type_evaluation?: string;
	    created_by: number;
	    contexte_evaluation?: string;
	    motif_reference?: string;
	    objet_evaluation?: string;
	    capacites_cognitives?: string;
	    etat_sante_physique?: string;
	    dimensions_psycho_sociales?: string;
	    roles_sociaux?: string;
	    reseau_social_soutien?: string;
	    analyse_clinique?: string;
	    opinion_professionnelle?: string;
	    recommandations?: string;
	    signature_nom?: string;
	    verrouille: number;
	    isDraft: number;
	    // Go type: time
	    date_signature?: any;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    client_nom: string;
	    client_prenom: string;
	    client_dn?: string;
	    client_nam?: string;
	    client_leopard?: string;
	    auteur_nom: string;
	
	    static createFrom(source: any = {}) {
	        return new EvaluationSocialeDetail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.client_id = source["client_id"];
	        this.no_eval_leopard = source["no_eval_leopard"];
	        this.type_evaluation = source["type_evaluation"];
	        this.created_by = source["created_by"];
	        this.contexte_evaluation = source["contexte_evaluation"];
	        this.motif_reference = source["motif_reference"];
	        this.objet_evaluation = source["objet_evaluation"];
	        this.capacites_cognitives = source["capacites_cognitives"];
	        this.etat_sante_physique = source["etat_sante_physique"];
	        this.dimensions_psycho_sociales = source["dimensions_psycho_sociales"];
	        this.roles_sociaux = source["roles_sociaux"];
	        this.reseau_social_soutien = source["reseau_social_soutien"];
	        this.analyse_clinique = source["analyse_clinique"];
	        this.opinion_professionnelle = source["opinion_professionnelle"];
	        this.recommandations = source["recommandations"];
	        this.signature_nom = source["signature_nom"];
	        this.verrouille = source["verrouille"];
	        this.isDraft = source["isDraft"];
	        this.date_signature = this.convertValues(source["date_signature"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.client_nom = source["client_nom"];
	        this.client_prenom = source["client_prenom"];
	        this.client_dn = source["client_dn"];
	        this.client_nam = source["client_nam"];
	        this.client_leopard = source["client_leopard"];
	        this.auteur_nom = source["auteur_nom"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class FactureLigne {
	    ID: number;
	    FactureID: number;
	    ServiceID?: number;
	    RevenuID?: number;
	    Ordre: number;
	    Description: string;
	    DateService: string;
	    Quantite: number;
	    Unite: string;
	    TarifUnitaire: number;
	    MontantLigne: number;
	
	    static createFrom(source: any = {}) {
	        return new FactureLigne(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.FactureID = source["FactureID"];
	        this.ServiceID = source["ServiceID"];
	        this.RevenuID = source["RevenuID"];
	        this.Ordre = source["Ordre"];
	        this.Description = source["Description"];
	        this.DateService = source["DateService"];
	        this.Quantite = source["Quantite"];
	        this.Unite = source["Unite"];
	        this.TarifUnitaire = source["TarifUnitaire"];
	        this.MontantLigne = source["MontantLigne"];
	    }
	}
	export class Facture {
	    ID: number;
	    ClientID: number;
	    Numero: string;
	    DateEmission: string;
	    DateEcheance: string;
	    DatePaiement: string;
	    Statut: string;
	    ExempteTaxes: number;
	    AvecTPS: number;
	    AvecTVQ: number;
	    JoursPaiement: number;
	    MontantSousTotal: number;
	    MontantTPS: number;
	    MontantTVQ: number;
	    MontantTotal: number;
	    MontantPaye: number;
	    MontantSolde: number;
	    MontantTiersPayant: number;
	    MontantDuClient: number;
	    TitreFacture: string;
	    NotesClient: string;
	    NotesInternes: string;
	    TiersPayantNom?: string;
	    CreatedBy?: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    Lignes: FactureLigne[];
	
	    static createFrom(source: any = {}) {
	        return new Facture(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.ClientID = source["ClientID"];
	        this.Numero = source["Numero"];
	        this.DateEmission = source["DateEmission"];
	        this.DateEcheance = source["DateEcheance"];
	        this.DatePaiement = source["DatePaiement"];
	        this.Statut = source["Statut"];
	        this.ExempteTaxes = source["ExempteTaxes"];
	        this.AvecTPS = source["AvecTPS"];
	        this.AvecTVQ = source["AvecTVQ"];
	        this.JoursPaiement = source["JoursPaiement"];
	        this.MontantSousTotal = source["MontantSousTotal"];
	        this.MontantTPS = source["MontantTPS"];
	        this.MontantTVQ = source["MontantTVQ"];
	        this.MontantTotal = source["MontantTotal"];
	        this.MontantPaye = source["MontantPaye"];
	        this.MontantSolde = source["MontantSolde"];
	        this.MontantTiersPayant = source["MontantTiersPayant"];
	        this.MontantDuClient = source["MontantDuClient"];
	        this.TitreFacture = source["TitreFacture"];
	        this.NotesClient = source["NotesClient"];
	        this.NotesInternes = source["NotesInternes"];
	        this.TiersPayantNom = source["TiersPayantNom"];
	        this.CreatedBy = source["CreatedBy"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.Lignes = this.convertValues(source["Lignes"], FactureLigne);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class Intervenant {
	    id: number;
	    nom_complet: string;
	    titre_emploi: string;
	    direction: string;
	    installation: string;
	    telephone: string;
	    poste: string;
	    cellulaire_pro: string;
	    cellulaire_perso: string;
	    courriel_personnel: string;
	    courriel_professionnel: string;
	    courrier_interne: string;
	    actif: boolean;
	    is_intervenant_social: boolean;
	    numero_permis: string;
	    ordre_professionnel: string;
	    date_naissance: string;
	    note_fixe: string;
	    personne_ressource_reseau_dir: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Intervenant(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom_complet = source["nom_complet"];
	        this.titre_emploi = source["titre_emploi"];
	        this.direction = source["direction"];
	        this.installation = source["installation"];
	        this.telephone = source["telephone"];
	        this.poste = source["poste"];
	        this.cellulaire_pro = source["cellulaire_pro"];
	        this.cellulaire_perso = source["cellulaire_perso"];
	        this.courriel_personnel = source["courriel_personnel"];
	        this.courriel_professionnel = source["courriel_professionnel"];
	        this.courrier_interne = source["courrier_interne"];
	        this.actif = source["actif"];
	        this.is_intervenant_social = source["is_intervenant_social"];
	        this.numero_permis = source["numero_permis"];
	        this.ordre_professionnel = source["ordre_professionnel"];
	        this.date_naissance = source["date_naissance"];
	        this.note_fixe = source["note_fixe"];
	        this.personne_ressource_reseau_dir = source["personne_ressource_reseau_dir"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Medecin {
	    id: number;
	    licence: string;
	    civilite: string;
	    prenom: string;
	    nom: string;
	    nomComplet: string;
	    statut: string;
	    specialisation: string;
	    service: string;
	    departement: string;
	    installationPrincipale: string;
	    rls: string;
	    telephone: string;
	    extension: string;
	    cellulaire: string;
	    email: string;
	    adresse: string;
	    codePostal: string;
	    ville: string;
	    pays: string;
	    noteFixe: string;
	    actif: number;
	    createdBy: number;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Medecin(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.licence = source["licence"];
	        this.civilite = source["civilite"];
	        this.prenom = source["prenom"];
	        this.nom = source["nom"];
	        this.nomComplet = source["nomComplet"];
	        this.statut = source["statut"];
	        this.specialisation = source["specialisation"];
	        this.service = source["service"];
	        this.departement = source["departement"];
	        this.installationPrincipale = source["installationPrincipale"];
	        this.rls = source["rls"];
	        this.telephone = source["telephone"];
	        this.extension = source["extension"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.codePostal = source["codePostal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.noteFixe = source["noteFixe"];
	        this.actif = source["actif"];
	        this.createdBy = source["createdBy"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Notaire {
	    id: number;
	    civilite: string;
	    prenom: string;
	    nom: string;
	    telephone?: string;
	    cellulaire?: string;
	    telecopieur?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays: string;
	    email?: string;
	    secteur_activite?: string;
	    note_fixe?: string;
	    actif: number;
	    created_by?: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Notaire(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.civilite = source["civilite"];
	        this.prenom = source["prenom"];
	        this.nom = source["nom"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.telecopieur = source["telecopieur"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.email = source["email"];
	        this.secteur_activite = source["secteur_activite"];
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.created_by = source["created_by"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Note {
	    id: number;
	    client_id: number;
	    client_NoRAMQ?: string;
	    client_Nom?: string;
	    client_Prenom?: string;
	    // Go type: time
	    client_date_naissance?: any;
	    client_Telephone?: string;
	    client_Cellulaire?: string;
	    client_NoLeopard?: string;
	    client_Adresse?: string;
	    client_appartement?: string;
	    client_code_postal?: string;
	    client_ville?: string;
	    client_pays?: string;
	    client_province?: string;
	    user_id: number;
	    // Go type: time
	    date_note: any;
	    // Go type: time
	    date_intervention?: any;
	    heure_intervention?: string;
	    duree_intervention?: string;
	    mode_intervention?: string;
	    type_intervention?: string;
	    type_note?: string;
	    titre?: string;
	    contenu?: string;
	    verrouille: number;
	    note_tardive: number;
	    note_de_tier: number;
	    signature_nom?: string;
	    // Go type: time
	    signature_date?: any;
	    note_liee_id?: number;
	    type_lien?: string;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Note(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.client_id = source["client_id"];
	        this.client_NoRAMQ = source["client_NoRAMQ"];
	        this.client_Nom = source["client_Nom"];
	        this.client_Prenom = source["client_Prenom"];
	        this.client_date_naissance = this.convertValues(source["client_date_naissance"], null);
	        this.client_Telephone = source["client_Telephone"];
	        this.client_Cellulaire = source["client_Cellulaire"];
	        this.client_NoLeopard = source["client_NoLeopard"];
	        this.client_Adresse = source["client_Adresse"];
	        this.client_appartement = source["client_appartement"];
	        this.client_code_postal = source["client_code_postal"];
	        this.client_ville = source["client_ville"];
	        this.client_pays = source["client_pays"];
	        this.client_province = source["client_province"];
	        this.user_id = source["user_id"];
	        this.date_note = this.convertValues(source["date_note"], null);
	        this.date_intervention = this.convertValues(source["date_intervention"], null);
	        this.heure_intervention = source["heure_intervention"];
	        this.duree_intervention = source["duree_intervention"];
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.titre = source["titre"];
	        this.contenu = source["contenu"];
	        this.verrouille = source["verrouille"];
	        this.note_tardive = source["note_tardive"];
	        this.note_de_tier = source["note_de_tier"];
	        this.signature_nom = source["signature_nom"];
	        this.signature_date = this.convertValues(source["signature_date"], null);
	        this.note_liee_id = source["note_liee_id"];
	        this.type_lien = source["type_lien"];
	        this.created_at = this.convertValues(source["created_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NoteListItem {
	    id: number;
	    // Go type: time
	    date_note: any;
	    // Go type: time
	    date_intervention?: any;
	    type_note?: string;
	    titre: string;
	    verrouille: number;
	    note_tardive: number;
	    type_lien?: string;
	
	    static createFrom(source: any = {}) {
	        return new NoteListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date_note = this.convertValues(source["date_note"], null);
	        this.date_intervention = this.convertValues(source["date_intervention"], null);
	        this.type_note = source["type_note"];
	        this.titre = source["titre"];
	        this.verrouille = source["verrouille"];
	        this.note_tardive = source["note_tardive"];
	        this.type_lien = source["type_lien"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Organisation {
	    ID: number;
	    TypeOrg: string;
	    Statut: string;
	    NoFournisseur: string;
	    ConditionsPaiement: number;
	    TauxTPSExempt: number;
	    ModePaiementPref: string;
	    Actif: number;
	    Nom?: string;
	    Acronyme?: string;
	    NoOrganisme?: string;
	    Adresse?: string;
	    Ville?: string;
	    CodePostal?: string;
	    Province?: string;
	    Pays?: string;
	    Telephone?: string;
	    Telecopieur?: string;
	    EmailGeneral?: string;
	    SiteWeb?: string;
	    ContactNom?: string;
	    ContactTitre?: string;
	    ContactTelephone?: string;
	    ContactEmail?: string;
	    Notes?: string;
	    ConditionsSpeciales?: string;
	    CreatedBy?: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Organisation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.TypeOrg = source["TypeOrg"];
	        this.Statut = source["Statut"];
	        this.NoFournisseur = source["NoFournisseur"];
	        this.ConditionsPaiement = source["ConditionsPaiement"];
	        this.TauxTPSExempt = source["TauxTPSExempt"];
	        this.ModePaiementPref = source["ModePaiementPref"];
	        this.Actif = source["Actif"];
	        this.Nom = source["Nom"];
	        this.Acronyme = source["Acronyme"];
	        this.NoOrganisme = source["NoOrganisme"];
	        this.Adresse = source["Adresse"];
	        this.Ville = source["Ville"];
	        this.CodePostal = source["CodePostal"];
	        this.Province = source["Province"];
	        this.Pays = source["Pays"];
	        this.Telephone = source["Telephone"];
	        this.Telecopieur = source["Telecopieur"];
	        this.EmailGeneral = source["EmailGeneral"];
	        this.SiteWeb = source["SiteWeb"];
	        this.ContactNom = source["ContactNom"];
	        this.ContactTitre = source["ContactTitre"];
	        this.ContactTelephone = source["ContactTelephone"];
	        this.ContactEmail = source["ContactEmail"];
	        this.Notes = source["Notes"];
	        this.ConditionsSpeciales = source["ConditionsSpeciales"];
	        this.CreatedBy = source["CreatedBy"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class OrganisationListItem {
	    id: number;
	    type_org: string;
	    nom: string;
	    acronyme: string;
	    contact_nom: string;
	    telephone: string;
	    actif: number;
	
	    static createFrom(source: any = {}) {
	        return new OrganisationListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.type_org = source["type_org"];
	        this.nom = source["nom"];
	        this.acronyme = source["acronyme"];
	        this.contact_nom = source["contact_nom"];
	        this.telephone = source["telephone"];
	        this.actif = source["actif"];
	    }
	}
	export class ParametresFinance {
	    ID: number;
	    NomComplet?: string;
	    TitreProfessionnel?: string;
	    NoMembreOrdre?: string;
	    EmailProfessionnel?: string;
	    TelephoneProfessionnel?: string;
	    AdressePratique?: string;
	    NoTPS?: string;
	    NoTVQ?: string;
	    InteracEmail?: string;
	    TauxHoraireDefaut: number;
	    TauxForfaitEvaluation: number;
	    TauxRapportExpertise: number;
	    TauxKM: number;
	    SousSeuilTPS: number;
	    InscriteTPS: number;
	    InscriteTVQ: number;
	    TauxTPS: number;
	    TauxTVQ: number;
	    PrefixeFacture: string;
	    AnneeFacture: number;
	    ProchainNoFacture: number;
	    JoursPaiementDefaut: number;
	    InfosPaiement: string;
	    PolitiqueAnnulation: string;
	    MentionsLegalesPDF: string;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new ParametresFinance(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.NomComplet = source["NomComplet"];
	        this.TitreProfessionnel = source["TitreProfessionnel"];
	        this.NoMembreOrdre = source["NoMembreOrdre"];
	        this.EmailProfessionnel = source["EmailProfessionnel"];
	        this.TelephoneProfessionnel = source["TelephoneProfessionnel"];
	        this.AdressePratique = source["AdressePratique"];
	        this.NoTPS = source["NoTPS"];
	        this.NoTVQ = source["NoTVQ"];
	        this.InteracEmail = source["InteracEmail"];
	        this.TauxHoraireDefaut = source["TauxHoraireDefaut"];
	        this.TauxForfaitEvaluation = source["TauxForfaitEvaluation"];
	        this.TauxRapportExpertise = source["TauxRapportExpertise"];
	        this.TauxKM = source["TauxKM"];
	        this.SousSeuilTPS = source["SousSeuilTPS"];
	        this.InscriteTPS = source["InscriteTPS"];
	        this.InscriteTVQ = source["InscriteTVQ"];
	        this.TauxTPS = source["TauxTPS"];
	        this.TauxTVQ = source["TauxTVQ"];
	        this.PrefixeFacture = source["PrefixeFacture"];
	        this.AnneeFacture = source["AnneeFacture"];
	        this.ProchainNoFacture = source["ProchainNoFacture"];
	        this.JoursPaiementDefaut = source["JoursPaiementDefaut"];
	        this.InfosPaiement = source["InfosPaiement"];
	        this.PolitiqueAnnulation = source["PolitiqueAnnulation"];
	        this.MentionsLegalesPDF = source["MentionsLegalesPDF"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PayeurResolu {
	    payeur_id: number;
	    type_payeur: string;
	    nom: string;
	    detail: string;
	    email: string;
	    telephone: string;
	    adresse: string;
	    mode_paiement: string;
	    jours_paiement: number;
	    exempte_tps_tvq: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PayeurResolu(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.payeur_id = source["payeur_id"];
	        this.type_payeur = source["type_payeur"];
	        this.nom = source["nom"];
	        this.detail = source["detail"];
	        this.email = source["email"];
	        this.telephone = source["telephone"];
	        this.adresse = source["adresse"];
	        this.mode_paiement = source["mode_paiement"];
	        this.jours_paiement = source["jours_paiement"];
	        this.exempte_tps_tvq = source["exempte_tps_tvq"];
	    }
	}
	export class Pays {
	    id: number;
	    pays: string;
	    alpha2: string;
	    alpha3: string;
	    numerique: number;
	
	    static createFrom(source: any = {}) {
	        return new Pays(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.pays = source["pays"];
	        this.alpha2 = source["alpha2"];
	        this.alpha3 = source["alpha3"];
	        this.numerique = source["numerique"];
	    }
	}
	export class PaysListItemForSelect {
	    id: number;
	    pays: string;
	    alpha2: string;
	    alpha3: string;
	    numerique: number;
	
	    static createFrom(source: any = {}) {
	        return new PaysListItemForSelect(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.pays = source["pays"];
	        this.alpha2 = source["alpha2"];
	        this.alpha3 = source["alpha3"];
	        this.numerique = source["numerique"];
	    }
	}
	export class Pharmacie {
	    id: number;
	    NomOrganisation: string;
	    Banniere?: string;
	    Adresse?: string;
	    Ville?: string;
	    Region?: string;
	    Tel?: string;
	    Fax?: string;
	    DimancheOuvert: number;
	    HeureOuvertureDimanche?: string;
	    HeureFermetureDimanche?: string;
	    LundiOuvert: number;
	    HeureOuvertureLundi?: string;
	    HeureFermetureLundi?: string;
	    MardiOuvert: number;
	    HeureOuvertureMardi?: string;
	    HeureFermetureMardi?: string;
	    MercrediOuvert: number;
	    HeureOuvertureMercredi?: string;
	    HeureFermetureMercredi?: string;
	    JeudiOuvert: number;
	    HeureOuvertureJeudi?: string;
	    HeureFermetureJeudi?: string;
	    VendrediOuvert: number;
	    HeureOuvertureVendredi?: string;
	    HeureFermetureVendredi?: string;
	    SamediOuvert: number;
	    HeureOuvertureSamedi?: string;
	    HeureFermetureSamedi?: string;
	    DateMaj?: string;
	    note?: string;
	
	    static createFrom(source: any = {}) {
	        return new Pharmacie(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.NomOrganisation = source["NomOrganisation"];
	        this.Banniere = source["Banniere"];
	        this.Adresse = source["Adresse"];
	        this.Ville = source["Ville"];
	        this.Region = source["Region"];
	        this.Tel = source["Tel"];
	        this.Fax = source["Fax"];
	        this.DimancheOuvert = source["DimancheOuvert"];
	        this.HeureOuvertureDimanche = source["HeureOuvertureDimanche"];
	        this.HeureFermetureDimanche = source["HeureFermetureDimanche"];
	        this.LundiOuvert = source["LundiOuvert"];
	        this.HeureOuvertureLundi = source["HeureOuvertureLundi"];
	        this.HeureFermetureLundi = source["HeureFermetureLundi"];
	        this.MardiOuvert = source["MardiOuvert"];
	        this.HeureOuvertureMardi = source["HeureOuvertureMardi"];
	        this.HeureFermetureMardi = source["HeureFermetureMardi"];
	        this.MercrediOuvert = source["MercrediOuvert"];
	        this.HeureOuvertureMercredi = source["HeureOuvertureMercredi"];
	        this.HeureFermetureMercredi = source["HeureFermetureMercredi"];
	        this.JeudiOuvert = source["JeudiOuvert"];
	        this.HeureOuvertureJeudi = source["HeureOuvertureJeudi"];
	        this.HeureFermetureJeudi = source["HeureFermetureJeudi"];
	        this.VendrediOuvert = source["VendrediOuvert"];
	        this.HeureOuvertureVendredi = source["HeureOuvertureVendredi"];
	        this.HeureFermetureVendredi = source["HeureFermetureVendredi"];
	        this.SamediOuvert = source["SamediOuvert"];
	        this.HeureOuvertureSamedi = source["HeureOuvertureSamedi"];
	        this.HeureFermetureSamedi = source["HeureFermetureSamedi"];
	        this.DateMaj = source["DateMaj"];
	        this.note = source["note"];
	    }
	}
	export class PlanInterventionDetail {
	    id: number;
	    client_id: number;
	    created_by: number;
	    titre: string;
	    type_plan: string;
	    statut: string;
	    date_debut?: string;
	    date_fin_prevue?: string;
	    date_revision_prevue?: string;
	    contexte?: string;
	    problematique?: string;
	    forces?: string;
	    objectifs?: string;
	    interventions?: string;
	    resultats?: string;
	    ententes?: string;
	    verrouille: number;
	    signature_nom?: string;
	    // Go type: time
	    date_signature?: any;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    client_nom: string;
	    client_prenom: string;
	    client_leopard?: string;
	    auteur_nom: string;
	
	    static createFrom(source: any = {}) {
	        return new PlanInterventionDetail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.client_id = source["client_id"];
	        this.created_by = source["created_by"];
	        this.titre = source["titre"];
	        this.type_plan = source["type_plan"];
	        this.statut = source["statut"];
	        this.date_debut = source["date_debut"];
	        this.date_fin_prevue = source["date_fin_prevue"];
	        this.date_revision_prevue = source["date_revision_prevue"];
	        this.contexte = source["contexte"];
	        this.problematique = source["problematique"];
	        this.forces = source["forces"];
	        this.objectifs = source["objectifs"];
	        this.interventions = source["interventions"];
	        this.resultats = source["resultats"];
	        this.ententes = source["ententes"];
	        this.verrouille = source["verrouille"];
	        this.signature_nom = source["signature_nom"];
	        this.date_signature = this.convertValues(source["date_signature"], null);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.client_nom = source["client_nom"];
	        this.client_prenom = source["client_prenom"];
	        this.client_leopard = source["client_leopard"];
	        this.auteur_nom = source["auteur_nom"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RapportFiscalAnnuel {
	    annee: number;
	    revenu_brut_total: number;
	    rev_consultations: number;
	    rev_evaluations: number;
	    rev_rapports: number;
	    rev_forfaits: number;
	    rev_ateliers: number;
	    rev_autres: number;
	    total_depenses_deduct: number;
	    dep_bureau: number;
	    dep_formation: number;
	    dep_assurance: number;
	    dep_logiciel: number;
	    dep_materiel: number;
	    dep_deplacement: number;
	    dep_honoraires: number;
	    dep_marketing: number;
	    revenu_net_entreprise: number;
	    cotisations_rrq: number;
	    deduction_rrq: number;
	    revenu_imposable: number;
	    seuil_tps_atteint: boolean;
	    par_mois: StatsMensuelles[];
	
	    static createFrom(source: any = {}) {
	        return new RapportFiscalAnnuel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.annee = source["annee"];
	        this.revenu_brut_total = source["revenu_brut_total"];
	        this.rev_consultations = source["rev_consultations"];
	        this.rev_evaluations = source["rev_evaluations"];
	        this.rev_rapports = source["rev_rapports"];
	        this.rev_forfaits = source["rev_forfaits"];
	        this.rev_ateliers = source["rev_ateliers"];
	        this.rev_autres = source["rev_autres"];
	        this.total_depenses_deduct = source["total_depenses_deduct"];
	        this.dep_bureau = source["dep_bureau"];
	        this.dep_formation = source["dep_formation"];
	        this.dep_assurance = source["dep_assurance"];
	        this.dep_logiciel = source["dep_logiciel"];
	        this.dep_materiel = source["dep_materiel"];
	        this.dep_deplacement = source["dep_deplacement"];
	        this.dep_honoraires = source["dep_honoraires"];
	        this.dep_marketing = source["dep_marketing"];
	        this.revenu_net_entreprise = source["revenu_net_entreprise"];
	        this.cotisations_rrq = source["cotisations_rrq"];
	        this.deduction_rrq = source["deduction_rrq"];
	        this.revenu_imposable = source["revenu_imposable"];
	        this.seuil_tps_atteint = source["seuil_tps_atteint"];
	        this.par_mois = this.convertValues(source["par_mois"], StatsMensuelles);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RentabiliteParClient {
	    client_id: number;
	    client_nom: string;
	    client_prenom: string;
	    nb_seances: number;
	    revenu_total: number;
	    revenu_moyen: number;
	    dernier_contact: string;
	    montant_en_attente: number;
	
	    static createFrom(source: any = {}) {
	        return new RentabiliteParClient(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.client_nom = source["client_nom"];
	        this.client_prenom = source["client_prenom"];
	        this.nb_seances = source["nb_seances"];
	        this.revenu_total = source["revenu_total"];
	        this.revenu_moyen = source["revenu_moyen"];
	        this.dernier_contact = source["dernier_contact"];
	        this.montant_en_attente = source["montant_en_attente"];
	    }
	}
	export class Residence {
	    id: number;
	    region?: string;
	    registre: string;
	    numero_interne?: string;
	    titre: string;
	    municipalite?: string;
	    adresse?: string;
	    ville?: string;
	    code_postal?: string;
	    telephone?: string;
	    capacite: number;
	    type_resid?: string;
	    proprietaires?: string;
	    services?: string;
	    date_certification?: string;
	    statut: string;
	    source_url?: string;
	    source_url_detaillee?: string;
	    notes?: string;
	    derniere_verification?: string;
	    date_ajout?: string;
	    date_maj?: string;
	    date_fermeture?: string;
	
	    static createFrom(source: any = {}) {
	        return new Residence(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.region = source["region"];
	        this.registre = source["registre"];
	        this.numero_interne = source["numero_interne"];
	        this.titre = source["titre"];
	        this.municipalite = source["municipalite"];
	        this.adresse = source["adresse"];
	        this.ville = source["ville"];
	        this.code_postal = source["code_postal"];
	        this.telephone = source["telephone"];
	        this.capacite = source["capacite"];
	        this.type_resid = source["type_resid"];
	        this.proprietaires = source["proprietaires"];
	        this.services = source["services"];
	        this.date_certification = source["date_certification"];
	        this.statut = source["statut"];
	        this.source_url = source["source_url"];
	        this.source_url_detaillee = source["source_url_detaillee"];
	        this.notes = source["notes"];
	        this.derniere_verification = source["derniere_verification"];
	        this.date_ajout = source["date_ajout"];
	        this.date_maj = source["date_maj"];
	        this.date_fermeture = source["date_fermeture"];
	    }
	}
	export class Revenu {
	    ID: number;
	    ClientID: number;
	    ServiceID?: number;
	    FactureID?: number;
	    DateService: string;
	    Categorie: string;
	    ModeFacturation: string;
	    StatutPaiement: string;
	    ModePaiement: string;
	    DatePaiement: string;
	    DureeHeures: number;
	    TauxHoraireApplique: number;
	    MontantBrut: number;
	    MontantTPS: number;
	    MontantTVQ: number;
	    MontantTotal: number;
	    Description: string;
	    Notes: string;
	    ReferencePaiement?: string;
	    CreatedBy?: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Revenu(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.ClientID = source["ClientID"];
	        this.ServiceID = source["ServiceID"];
	        this.FactureID = source["FactureID"];
	        this.DateService = source["DateService"];
	        this.Categorie = source["Categorie"];
	        this.ModeFacturation = source["ModeFacturation"];
	        this.StatutPaiement = source["StatutPaiement"];
	        this.ModePaiement = source["ModePaiement"];
	        this.DatePaiement = source["DatePaiement"];
	        this.DureeHeures = source["DureeHeures"];
	        this.TauxHoraireApplique = source["TauxHoraireApplique"];
	        this.MontantBrut = source["MontantBrut"];
	        this.MontantTPS = source["MontantTPS"];
	        this.MontantTVQ = source["MontantTVQ"];
	        this.MontantTotal = source["MontantTotal"];
	        this.Description = source["Description"];
	        this.Notes = source["Notes"];
	        this.ReferencePaiement = source["ReferencePaiement"];
	        this.CreatedBy = source["CreatedBy"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Service {
	    ID: number;
	    Code: string;
	    Categorie: string;
	    TypeLivraison: string;
	    Actif: number;
	    OrdreAffichage: number;
	    Nom: string;
	    DescriptionCourte: string;
	    DescriptionLongue: string;
	    PublicCible: string;
	    NotesInternes: string;
	    ModeTarification: string;
	    DureeMinutes?: number;
	    TauxHoraire?: number;
	    TarifUnitaire?: number;
	    TarifMin?: number;
	    TarifMax?: number;
	    ExempteTaxes: number;
	    TPSApplicable: number;
	    TVQApplicable: number;
	    NbPlacesMax?: number;
	    NbSeancesForfait?: number;
	    DureeProgrammeSemaines?: number;
	    FormatTract: string;
	    CouleurTract: string;
	    DureeVideoMinutes?: number;
	    URLRessource: string;
	    Tags: string;
	    Icone: string;
	    CreatedBy?: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Service(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Code = source["Code"];
	        this.Categorie = source["Categorie"];
	        this.TypeLivraison = source["TypeLivraison"];
	        this.Actif = source["Actif"];
	        this.OrdreAffichage = source["OrdreAffichage"];
	        this.Nom = source["Nom"];
	        this.DescriptionCourte = source["DescriptionCourte"];
	        this.DescriptionLongue = source["DescriptionLongue"];
	        this.PublicCible = source["PublicCible"];
	        this.NotesInternes = source["NotesInternes"];
	        this.ModeTarification = source["ModeTarification"];
	        this.DureeMinutes = source["DureeMinutes"];
	        this.TauxHoraire = source["TauxHoraire"];
	        this.TarifUnitaire = source["TarifUnitaire"];
	        this.TarifMin = source["TarifMin"];
	        this.TarifMax = source["TarifMax"];
	        this.ExempteTaxes = source["ExempteTaxes"];
	        this.TPSApplicable = source["TPSApplicable"];
	        this.TVQApplicable = source["TVQApplicable"];
	        this.NbPlacesMax = source["NbPlacesMax"];
	        this.NbSeancesForfait = source["NbSeancesForfait"];
	        this.DureeProgrammeSemaines = source["DureeProgrammeSemaines"];
	        this.FormatTract = source["FormatTract"];
	        this.CouleurTract = source["CouleurTract"];
	        this.DureeVideoMinutes = source["DureeVideoMinutes"];
	        this.URLRessource = source["URLRessource"];
	        this.Tags = source["Tags"];
	        this.Icone = source["Icone"];
	        this.CreatedBy = source["CreatedBy"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class StatsAppels {
	    total: number;
	    nouveaux: number;
	    enAttente: number;
	    rdvPlanifies: number;
	    urgents: number;
	    aujourdhui: number;
	
	    static createFrom(source: any = {}) {
	        return new StatsAppels(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.nouveaux = source["nouveaux"];
	        this.enAttente = source["enAttente"];
	        this.rdvPlanifies = source["rdvPlanifies"];
	        this.urgents = source["urgents"];
	        this.aujourdhui = source["aujourdhui"];
	    }
	}
	
	export class UpdateClientRequest {
	    id: number;
	    nom: string;
	    prenom: string;
	    date_naissance?: string;
	    sexe?: string;
	    lieu_naissance?: string;
	    statut_marital?: string;
	    occupation?: string;
	    employeur?: string;
	    profession?: string;
	    niveau_scolaire?: string;
	    langue_preferee?: string;
	    origine_ethnique?: string;
	    premiere_nation?: string;
	    identite_genre?: string;
	    orientation_sexuelle?: string;
	    religion?: string;
	    premiere_langue?: string;
	    telephone?: string;
	    cellulaire?: string;
	    email?: string;
	    adresse?: string;
	    appartement?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
	    province?: string;
	    numero_assurance_maladie?: string;
	    numero_securite_sociale?: string;
	    no_hcm?: string;
	    no_chaur?: string;
	    no_dossier_leopard?: string;
	    medecin_famille_No_Licence?: string;
	    notaire_id?: number;
	    pharmacie_id?: number;
	    pivot_id?: number;
	    rpa_id?: number;
	    chsld_id?: number;
	    ri_id?: number;
	    procuration_bancaire: string;
	    procuration_notariee: string;
	    note_fixe?: string;
	    actif: number;
	    dcd: number;
	    date_deces?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateClientRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.date_naissance = source["date_naissance"];
	        this.sexe = source["sexe"];
	        this.lieu_naissance = source["lieu_naissance"];
	        this.statut_marital = source["statut_marital"];
	        this.occupation = source["occupation"];
	        this.employeur = source["employeur"];
	        this.profession = source["profession"];
	        this.niveau_scolaire = source["niveau_scolaire"];
	        this.langue_preferee = source["langue_preferee"];
	        this.origine_ethnique = source["origine_ethnique"];
	        this.premiere_nation = source["premiere_nation"];
	        this.identite_genre = source["identite_genre"];
	        this.orientation_sexuelle = source["orientation_sexuelle"];
	        this.religion = source["religion"];
	        this.premiere_langue = source["premiere_langue"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.appartement = source["appartement"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.province = source["province"];
	        this.numero_assurance_maladie = source["numero_assurance_maladie"];
	        this.numero_securite_sociale = source["numero_securite_sociale"];
	        this.no_hcm = source["no_hcm"];
	        this.no_chaur = source["no_chaur"];
	        this.no_dossier_leopard = source["no_dossier_leopard"];
	        this.medecin_famille_No_Licence = source["medecin_famille_No_Licence"];
	        this.notaire_id = source["notaire_id"];
	        this.pharmacie_id = source["pharmacie_id"];
	        this.pivot_id = source["pivot_id"];
	        this.rpa_id = source["rpa_id"];
	        this.chsld_id = source["chsld_id"];
	        this.ri_id = source["ri_id"];
	        this.procuration_bancaire = source["procuration_bancaire"];
	        this.procuration_notariee = source["procuration_notariee"];
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.dcd = source["dcd"];
	        this.date_deces = source["date_deces"];
	    }
	}
	export class UpdateContactRequest {
	    id: number;
	    nom: string;
	    prenom: string;
	    telephone?: string;
	    cellulaire?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
	    email?: string;
	    employeur?: string;
	    profession?: string;
	    relation_client?: string;
	    lien_familial?: string;
	    force_lien: number;
	    contact_urgence: number;
	    aidant_naturel: number;
	    type_de_contact?: string;
	    procuration_bancaire: number;
	    procuration_notariee: number;
	    note_fixe?: string;
	    relation_type: string;
	    client_id: number;
	
	    static createFrom(source: any = {}) {
	        return new UpdateContactRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.email = source["email"];
	        this.employeur = source["employeur"];
	        this.profession = source["profession"];
	        this.relation_client = source["relation_client"];
	        this.lien_familial = source["lien_familial"];
	        this.force_lien = source["force_lien"];
	        this.contact_urgence = source["contact_urgence"];
	        this.aidant_naturel = source["aidant_naturel"];
	        this.type_de_contact = source["type_de_contact"];
	        this.procuration_bancaire = source["procuration_bancaire"];
	        this.procuration_notariee = source["procuration_notariee"];
	        this.note_fixe = source["note_fixe"];
	        this.relation_type = source["relation_type"];
	        this.client_id = source["client_id"];
	    }
	}
	export class UpdateContratRequest {
	    id: number;
	    client_id: number;
	    service_id?: number;
	    type_contrat: string;
	    date_debut: string;
	    date_fin?: string;
	    renouvellement_auto: number;
	    mode_facturation: string;
	    taux_horaire: number;
	    taux_forfait: number;
	    duree_seance_min?: number;
	    frequence_seances?: string;
	    clause_objet?: string;
	    clause_objectifs?: string;
	    clause_services_inclus?: string;
	    clause_services_exclus?: string;
	    clause_tarification?: string;
	    clause_paiement_retard?: string;
	    clause_annulation?: string;
	    clause_absence_prolongee?: string;
	    clause_confidentialite?: string;
	    clause_limites_confidentialite?: string;
	    clause_dossier_client?: string;
	    clause_communication?: string;
	    clause_urgences?: string;
	    clause_resiliation_client?: string;
	    clause_resiliation_ts?: string;
	    clauses_specifiques?: string;
	    notes_internes?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateContratRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.client_id = source["client_id"];
	        this.service_id = source["service_id"];
	        this.type_contrat = source["type_contrat"];
	        this.date_debut = source["date_debut"];
	        this.date_fin = source["date_fin"];
	        this.renouvellement_auto = source["renouvellement_auto"];
	        this.mode_facturation = source["mode_facturation"];
	        this.taux_horaire = source["taux_horaire"];
	        this.taux_forfait = source["taux_forfait"];
	        this.duree_seance_min = source["duree_seance_min"];
	        this.frequence_seances = source["frequence_seances"];
	        this.clause_objet = source["clause_objet"];
	        this.clause_objectifs = source["clause_objectifs"];
	        this.clause_services_inclus = source["clause_services_inclus"];
	        this.clause_services_exclus = source["clause_services_exclus"];
	        this.clause_tarification = source["clause_tarification"];
	        this.clause_paiement_retard = source["clause_paiement_retard"];
	        this.clause_annulation = source["clause_annulation"];
	        this.clause_absence_prolongee = source["clause_absence_prolongee"];
	        this.clause_confidentialite = source["clause_confidentialite"];
	        this.clause_limites_confidentialite = source["clause_limites_confidentialite"];
	        this.clause_dossier_client = source["clause_dossier_client"];
	        this.clause_communication = source["clause_communication"];
	        this.clause_urgences = source["clause_urgences"];
	        this.clause_resiliation_client = source["clause_resiliation_client"];
	        this.clause_resiliation_ts = source["clause_resiliation_ts"];
	        this.clauses_specifiques = source["clauses_specifiques"];
	        this.notes_internes = source["notes_internes"];
	    }
	}
	export class UpdateDepenseRequest {
	    id: number;
	    date_depense: string;
	    categorie: string;
	    deductible: number;
	    pct_utilisation_affaires: number;
	    est_kilometrage: number;
	    sous_total: number;
	    montant_tps: number;
	    montant_tvq: number;
	    montant_total: number;
	    montant_deductible: number;
	    km_parcourus: number;
	    taux_km_utilise: number;
	    fournisseur?: string;
	    no_recu?: string;
	    description: string;
	    notes?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateDepenseRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date_depense = source["date_depense"];
	        this.categorie = source["categorie"];
	        this.deductible = source["deductible"];
	        this.pct_utilisation_affaires = source["pct_utilisation_affaires"];
	        this.est_kilometrage = source["est_kilometrage"];
	        this.sous_total = source["sous_total"];
	        this.montant_tps = source["montant_tps"];
	        this.montant_tvq = source["montant_tvq"];
	        this.montant_total = source["montant_total"];
	        this.montant_deductible = source["montant_deductible"];
	        this.km_parcourus = source["km_parcourus"];
	        this.taux_km_utilise = source["taux_km_utilise"];
	        this.fournisseur = source["fournisseur"];
	        this.no_recu = source["no_recu"];
	        this.description = source["description"];
	        this.notes = source["notes"];
	    }
	}
	export class UpdateMedecinRequest {
	    id: number;
	    licence: string;
	    civilite: string;
	    prenom: string;
	    nom: string;
	    nomComplet: string;
	    statut: string;
	    specialisation: string;
	    service: string;
	    departement: string;
	    installationPrincipale: string;
	    rls: string;
	    telephone: string;
	    extension: string;
	    cellulaire: string;
	    email: string;
	    adresse: string;
	    codePostal: string;
	    ville: string;
	    pays: string;
	    noteFixe: string;
	    actif: number;
	
	    static createFrom(source: any = {}) {
	        return new UpdateMedecinRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.licence = source["licence"];
	        this.civilite = source["civilite"];
	        this.prenom = source["prenom"];
	        this.nom = source["nom"];
	        this.nomComplet = source["nomComplet"];
	        this.statut = source["statut"];
	        this.specialisation = source["specialisation"];
	        this.service = source["service"];
	        this.departement = source["departement"];
	        this.installationPrincipale = source["installationPrincipale"];
	        this.rls = source["rls"];
	        this.telephone = source["telephone"];
	        this.extension = source["extension"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.codePostal = source["codePostal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.noteFixe = source["noteFixe"];
	        this.actif = source["actif"];
	    }
	}
	export class UpdateNoteRequest {
	    id: number;
	    date_intervention?: string;
	    heure_intervention?: string;
	    duree_intervention?: string;
	    mode_intervention?: string;
	    type_intervention?: string;
	    type_note?: string;
	    titre?: string;
	    contenu?: string;
	    note_tardive: number;
	    note_de_tier: number;
	    note_liee_id?: number;
	    type_lien?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateNoteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date_intervention = source["date_intervention"];
	        this.heure_intervention = source["heure_intervention"];
	        this.duree_intervention = source["duree_intervention"];
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.titre = source["titre"];
	        this.contenu = source["contenu"];
	        this.note_tardive = source["note_tardive"];
	        this.note_de_tier = source["note_de_tier"];
	        this.note_liee_id = source["note_liee_id"];
	        this.type_lien = source["type_lien"];
	    }
	}
	export class UpdateOrganisationRequest {
	    id: number;
	    type_org: string;
	    no_fournisseur?: string;
	    conditions_paiement: number;
	    taux_tps_exempt: number;
	    mode_paiement_pref: string;
	    nom?: string;
	    acronyme?: string;
	    no_organisme?: string;
	    adresse?: string;
	    ville?: string;
	    code_postal?: string;
	    province?: string;
	    pays?: string;
	    telephone?: string;
	    telecopieur?: string;
	    email_general?: string;
	    site_web?: string;
	    contact_nom?: string;
	    contact_titre?: string;
	    contact_telephone?: string;
	    contact_email?: string;
	    notes?: string;
	    conditions_speciales?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateOrganisationRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.type_org = source["type_org"];
	        this.no_fournisseur = source["no_fournisseur"];
	        this.conditions_paiement = source["conditions_paiement"];
	        this.taux_tps_exempt = source["taux_tps_exempt"];
	        this.mode_paiement_pref = source["mode_paiement_pref"];
	        this.nom = source["nom"];
	        this.acronyme = source["acronyme"];
	        this.no_organisme = source["no_organisme"];
	        this.adresse = source["adresse"];
	        this.ville = source["ville"];
	        this.code_postal = source["code_postal"];
	        this.province = source["province"];
	        this.pays = source["pays"];
	        this.telephone = source["telephone"];
	        this.telecopieur = source["telecopieur"];
	        this.email_general = source["email_general"];
	        this.site_web = source["site_web"];
	        this.contact_nom = source["contact_nom"];
	        this.contact_titre = source["contact_titre"];
	        this.contact_telephone = source["contact_telephone"];
	        this.contact_email = source["contact_email"];
	        this.notes = source["notes"];
	        this.conditions_speciales = source["conditions_speciales"];
	    }
	}
	export class UpdateProfileRequest {
	    full_name: string;
	    email: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateProfileRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.full_name = source["full_name"];
	        this.email = source["email"];
	    }
	}
	export class UpdateRevenuRequest {
	    id: number;
	    client_id: number;
	    service_id?: number;
	    facture_id?: number;
	    date_service: string;
	    categorie: string;
	    mode_facturation: string;
	    statut_paiement: string;
	    mode_paiement: string;
	    date_paiement?: string;
	    duree_heures: number;
	    taux_horaire_applique: number;
	    montant_brut: number;
	    montant_tps: number;
	    montant_tvq: number;
	    montant_total: number;
	    description?: string;
	    notes?: string;
	    reference_paiement?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateRevenuRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.client_id = source["client_id"];
	        this.service_id = source["service_id"];
	        this.facture_id = source["facture_id"];
	        this.date_service = source["date_service"];
	        this.categorie = source["categorie"];
	        this.mode_facturation = source["mode_facturation"];
	        this.statut_paiement = source["statut_paiement"];
	        this.mode_paiement = source["mode_paiement"];
	        this.date_paiement = source["date_paiement"];
	        this.duree_heures = source["duree_heures"];
	        this.taux_horaire_applique = source["taux_horaire_applique"];
	        this.montant_brut = source["montant_brut"];
	        this.montant_tps = source["montant_tps"];
	        this.montant_tvq = source["montant_tvq"];
	        this.montant_total = source["montant_total"];
	        this.description = source["description"];
	        this.notes = source["notes"];
	        this.reference_paiement = source["reference_paiement"];
	    }
	}
	export class UpdateServiceRequest {
	    id: number;
	    code: string;
	    categorie: string;
	    type_livraison: string;
	    actif: number;
	    ordre_affichage: number;
	    nom: string;
	    description_courte?: string;
	    description_longue?: string;
	    public_cible?: string;
	    notes_internes?: string;
	    mode_tarification: string;
	    duree_minutes?: number;
	    taux_horaire?: number;
	    tarif_unitaire?: number;
	    tarif_min?: number;
	    tarif_max?: number;
	    exempte_taxes: number;
	    tps_applicable: number;
	    tvq_applicable: number;
	    nb_places_max?: number;
	    nb_seances_forfait?: number;
	    duree_programme_semaines?: number;
	    format_tract?: string;
	    couleur_tract?: string;
	    duree_video_minutes?: number;
	    url_ressource?: string;
	    tags?: string;
	    icone?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateServiceRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.code = source["code"];
	        this.categorie = source["categorie"];
	        this.type_livraison = source["type_livraison"];
	        this.actif = source["actif"];
	        this.ordre_affichage = source["ordre_affichage"];
	        this.nom = source["nom"];
	        this.description_courte = source["description_courte"];
	        this.description_longue = source["description_longue"];
	        this.public_cible = source["public_cible"];
	        this.notes_internes = source["notes_internes"];
	        this.mode_tarification = source["mode_tarification"];
	        this.duree_minutes = source["duree_minutes"];
	        this.taux_horaire = source["taux_horaire"];
	        this.tarif_unitaire = source["tarif_unitaire"];
	        this.tarif_min = source["tarif_min"];
	        this.tarif_max = source["tarif_max"];
	        this.exempte_taxes = source["exempte_taxes"];
	        this.tps_applicable = source["tps_applicable"];
	        this.tvq_applicable = source["tvq_applicable"];
	        this.nb_places_max = source["nb_places_max"];
	        this.nb_seances_forfait = source["nb_seances_forfait"];
	        this.duree_programme_semaines = source["duree_programme_semaines"];
	        this.format_tract = source["format_tract"];
	        this.couleur_tract = source["couleur_tract"];
	        this.duree_video_minutes = source["duree_video_minutes"];
	        this.url_ressource = source["url_ressource"];
	        this.tags = source["tags"];
	        this.icone = source["icone"];
	    }
	}
	export class UpdateSettingsRequest {
	    theme: string;
	    language: string;
	    notifications_enabled: boolean;
	    email_notifications: boolean;
	
	    static createFrom(source: any = {}) {
	        return new UpdateSettingsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.language = source["language"];
	        this.notifications_enabled = source["notifications_enabled"];
	        this.email_notifications = source["email_notifications"];
	    }
	}
	export class User {
	    id: number;
	    username: string;
	    fullName: string;
	    role: string;
	    noMembreOrdre: string;
	    email: string;
	    telephone: string;
	    cellulaire: string;
	    telecopieur: string;
	    adresse: string;
	    codePostal: string;
	    ville: string;
	    pays: string;
	    createdAt: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.fullName = source["fullName"];
	        this.role = source["role"];
	        this.noMembreOrdre = source["noMembreOrdre"];
	        this.email = source["email"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.telecopieur = source["telecopieur"];
	        this.adresse = source["adresse"];
	        this.codePostal = source["codePostal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.createdAt = source["createdAt"];
	    }
	}
	export class UserSettings {
	    id: number;
	    user_id: number;
	    theme: string;
	    language: string;
	    notifications_enabled: boolean;
	    email_notifications: boolean;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new UserSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.user_id = source["user_id"];
	        this.theme = source["theme"];
	        this.language = source["language"];
	        this.notifications_enabled = source["notifications_enabled"];
	        this.email_notifications = source["email_notifications"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}

}

