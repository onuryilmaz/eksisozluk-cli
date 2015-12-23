package main

import (
	"strings"
)

type DebeCommand struct {
	cli EksiSozlukCLICommand
}

func (c *DebeCommand) Help() string {
	helpText := `
				Kullanım: eksisozluk-cli debe [--limit=DEBE_LIMITI] [--output=json|console]
				  Başlık için bulunan entry'leri listelemek için kullanılır.
				Seçenekler:
				  --limit=ENTRY_LIMITI			Toplam listelenen entry limiti (varsayılan 100)
				  --output=json,console	console: 	Komut satırı çıktısı, json: JSON dosyası çıktısı (varsayılan console)
				`
	return strings.TrimSpace(helpText)
}

func (c *DebeCommand) Run(args []string) int {
	parameter, err := ParameterFlagHandler(args, c, c.cli)
	if err != nil {
		return 1
	}

	if parameter.Limit == -1 {
		parameter.Limit = 100
	}
	debeList := GetDEBE(parameter)
	WriteDebeList(debeList, parameter)
	return 0
}

func (c *DebeCommand) Synopsis() string {
	return "Dünün En Beğenilen Entry'lerini listeleme"
}
