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
	export class CreateNoteRequest {
	    client_id: number;
	    date_intervention: string;
	    heure_intervention: string;
	    duree_intervention: string;
	    mode_intervention: string;
	    type_intervention: string;
	    type_note: string;
	    sujet: string;
	    contenu: string;
	    note_liee_id?: number;
	    type_lien?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateNoteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.date_intervention = source["date_intervention"];
	        this.heure_intervention = source["heure_intervention"];
	        this.duree_intervention = source["duree_intervention"];
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.sujet = source["sujet"];
	        this.contenu = source["contenu"];
	        this.note_liee_id = source["note_liee_id"];
	        this.type_lien = source["type_lien"];
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
	export class Note {
	    id: number;
	    client_id: number;
	    user_id: number;
	    // Go type: time
	    date_note: any;
	    // Go type: time
	    date_intervention?: any;
	    heure_intervention?: string;
	    duree_intervention?: string;
	    mode_intervention: string;
	    type_intervention: string;
	    type_note: string;
	    sujet: string;
	    contenu: string;
	    verrouille: boolean;
	    signature_nom: string;
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
	        this.user_id = source["user_id"];
	        this.date_note = this.convertValues(source["date_note"], null);
	        this.date_intervention = this.convertValues(source["date_intervention"], null);
	        this.heure_intervention = source["heure_intervention"];
	        this.duree_intervention = source["duree_intervention"];
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.sujet = source["sujet"];
	        this.contenu = source["contenu"];
	        this.verrouille = source["verrouille"];
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
	    type_note: string;
	    sujet: string;
	    // Go type: time
	    date_note: any;
	    signature_nom: string;
	    verrouille: boolean;
	    note_liee_id?: number;
	    type_lien?: string;
	    note_liee_titre?: string;
	    note_liee_date?: string;
	
	    static createFrom(source: any = {}) {
	        return new NoteListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.type_note = source["type_note"];
	        this.sujet = source["sujet"];
	        this.date_note = this.convertValues(source["date_note"], null);
	        this.signature_nom = source["signature_nom"];
	        this.verrouille = source["verrouille"];
	        this.note_liee_id = source["note_liee_id"];
	        this.type_lien = source["type_lien"];
	        this.note_liee_titre = source["note_liee_titre"];
	        this.note_liee_date = source["note_liee_date"];
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
	export class NotesFilter {
	    client_id: number;
	    search_query: string;
	
	    static createFrom(source: any = {}) {
	        return new NotesFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.search_query = source["search_query"];
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

