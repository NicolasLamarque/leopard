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

