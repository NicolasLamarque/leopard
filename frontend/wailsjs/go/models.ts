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

}

export namespace models {
	
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
	    telephone?: string;
	    cellulaire?: string;
	    email?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
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
	    note_fixe?: string;
	    actif: number;
	    dcd: number;
	    created_by?: number;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Client(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.date_naissance = source["date_naissance"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
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
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.dcd = source["dcd"];
	        this.created_by = source["created_by"];
	        this.created_at = source["created_at"];
	    }
	}
	export class CreateClientRequest {
	    nom: string;
	    prenom: string;
	    date_naissance?: string;
	    telephone?: string;
	    cellulaire?: string;
	    email?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
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
	    note_fixe?: string;
	    actif: number;
	    dcd: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateClientRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.date_naissance = source["date_naissance"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
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
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.dcd = source["dcd"];
	    }
	}
	export class CreateMedecinRequest {
	    licence: string;
	    nomComplet: string;
	    specialisation: string;
	    telephone: string;
	    extension: string;
	    cellulaire: string;
	    email: string;
	    adresse: string;
	    code_postal: string;
	    ville: string;
	    pays: string;
	    note_fixe: string;
	    actif: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateMedecinRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.licence = source["licence"];
	        this.nomComplet = source["nomComplet"];
	        this.specialisation = source["specialisation"];
	        this.telephone = source["telephone"];
	        this.extension = source["extension"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	    }
	}
	export class CreateNoteRequest {
	    client_id: number;
	    date_intervention: string;
	    mode_intervention: string;
	    type_intervention: string;
	    type_note: string;
	    sujet: string;
	    contenu: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateNoteRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.date_intervention = source["date_intervention"];
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.sujet = source["sujet"];
	        this.contenu = source["contenu"];
	    }
	}
	export class Medecin {
	    id: number;
	    licence: string;
	    nomComplet: string;
	    specialisation: string;
	    telephone: string;
	    extension: string;
	    cellulaire: string;
	    email: string;
	    adresse: string;
	    code_postal: string;
	    ville: string;
	    pays: string;
	    note_fixe: string;
	    actif: number;
	    created_by: number;
	    // Go type: time
	    created_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Medecin(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.licence = source["licence"];
	        this.nomComplet = source["nomComplet"];
	        this.specialisation = source["specialisation"];
	        this.telephone = source["telephone"];
	        this.extension = source["extension"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.created_by = source["created_by"];
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
	export class Note {
	    id: number;
	    client_id: number;
	    user_id: number;
	    // Go type: time
	    date_note: any;
	    // Go type: time
	    date_intervention?: any;
	    mode_intervention: string;
	    type_intervention: string;
	    type_note: string;
	    sujet: string;
	    contenu: string;
	    verrouille: boolean;
	    signature_nom: string;
	    // Go type: time
	    signature_date?: any;
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
	        this.mode_intervention = source["mode_intervention"];
	        this.type_intervention = source["type_intervention"];
	        this.type_note = source["type_note"];
	        this.sujet = source["sujet"];
	        this.contenu = source["contenu"];
	        this.verrouille = source["verrouille"];
	        this.signature_nom = source["signature_nom"];
	        this.signature_date = this.convertValues(source["signature_date"], null);
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
	    region: string;
	    registre: string;
	    titre: string;
	    municipalite: string;
	    adresse: string;
	    ville: string;
	    code_postal: string;
	    telephone: string;
	    capacite: number;
	    type_resid: string;
	    proprietaires: string;
	    services: string;
	    date_certification: string;
	    statut: string;
	    source_url: string;
	    notes: string;
	    derniere_verification: string;
	    date_ajout: string;
	    date_maj: string;
	    date_fermeture: string;
	
	    static createFrom(source: any = {}) {
	        return new Residence(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.region = source["region"];
	        this.registre = source["registre"];
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
	    telephone?: string;
	    cellulaire?: string;
	    email?: string;
	    adresse?: string;
	    code_postal?: string;
	    ville?: string;
	    pays?: string;
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
	    note_fixe?: string;
	    actif: number;
	    dcd: number;
	
	    static createFrom(source: any = {}) {
	        return new UpdateClientRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.date_naissance = source["date_naissance"];
	        this.telephone = source["telephone"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
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
	        this.note_fixe = source["note_fixe"];
	        this.actif = source["actif"];
	        this.dcd = source["dcd"];
	    }
	}
	export class UpdateMedecinRequest {
	    id: number;
	    licence: string;
	    nomComplet: string;
	    specialisation: string;
	    telephone: string;
	    extension: string;
	    cellulaire: string;
	    email: string;
	    adresse: string;
	    code_postal: string;
	    ville: string;
	    pays: string;
	    note_fixe: string;
	    actif: number;
	
	    static createFrom(source: any = {}) {
	        return new UpdateMedecinRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.licence = source["licence"];
	        this.nomComplet = source["nomComplet"];
	        this.specialisation = source["specialisation"];
	        this.telephone = source["telephone"];
	        this.extension = source["extension"];
	        this.cellulaire = source["cellulaire"];
	        this.email = source["email"];
	        this.adresse = source["adresse"];
	        this.code_postal = source["code_postal"];
	        this.ville = source["ville"];
	        this.pays = source["pays"];
	        this.note_fixe = source["note_fixe"];
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

