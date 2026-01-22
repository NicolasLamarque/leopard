export const archiveur = {
  async expedier(doc, typeDossier, client, leopardNumber) {
    const dateStr = new Date().toISOString().split('T')[0];
    const fileName = `${dateStr}_${typeDossier}_${client.nom}.pdf`;
    
    const folderInfo = await window.go.main.App.GetClientFolderInfo(leopardNumber);
    const targetFolder = `${folderInfo.path}/${typeDossier.toLowerCase()}`;
    
    const pdfBase64 = doc.output('datauristring').split(',')[1];
    return await window.go.main.App.SavePDFToFolder(targetFolder, fileName, pdfBase64);
  }
};