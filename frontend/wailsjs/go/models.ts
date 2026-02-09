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

}

export namespace models {
	
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
	export class CreateEvaluationRequest {
	    client_id: number;
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
	
	    static createFrom(source: any = {}) {
	        return new CreateEvaluationRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
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
	export class EvaluationSocialeDetail {
	    id: number;
	    client_id: number;
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

