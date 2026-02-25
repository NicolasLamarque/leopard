package importer

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (p *Pipeline) ReadExcel() ([][]string, error) {
	f, err := excelize.OpenFile(p.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("aucune feuille trouvée")
	}

	// On prend la première feuille dynamiquement
	return f.GetRows(sheets[0])
}
