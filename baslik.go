package main

import (
	"strings"
)

type BaslikCommand struct {
	cli EksiSozlukCLICommand
}

func (c *BaslikCommand) Help() string {
	helpText := `
				Kullanım: eksisozluk-cli baslik BASLIK [--sukela] [--page=SAYFA_SAYISI] [--limit=ENTRY_LIMITI] [--output=json|console]
				  Başlık için bulunan entry'leri listelemek için kullanılır.
				Seçenekler:
				  --sukela				Başlıktaki entry'leri şukela modunda listelemek
				  --page=SAYFA_SAYISI			Entry'ler için başlangıç sayfası (varsayılan 1)
				  --limit=ENTRY_LIMITI			Toplam listelenen entry limiti (varsayılan 10)
				  --output=json,console	console: 	Komut satırı çıktısı, json: JSON dosyası çıktısı (varsayılan console)
				`
	return strings.TrimSpace(helpText)
}

func (c *BaslikCommand) Run(args []string) int {

	if len(args) < 1 {
		c.cli.UI.Output(c.Help())
		return 1
	}
	baslik := args[0]

	parameter, err := ParameterFlagHandler(args[1:], c, c.cli)

	if err != nil {
		return 1
	}
	if parameter.Limit == -1 {
		parameter.Limit = 10
	}
	entryList := GetEntries(baslik, parameter)

	WriteEntryList(entryList, parameter, baslik)

	return 0
}

func (c *BaslikCommand) Synopsis() string {
	return "Başlık adı ile arama"
}
