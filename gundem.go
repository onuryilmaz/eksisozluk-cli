package main

import (
	"strings"
)

// GundemCommand for gathering topics ("gundem")
type GundemCommand struct {
	cli EksiSozlukCLICommand
}

// Help provides the usage guide
func (c *GundemCommand) Help() string {
	helpText := `
				Kullanım: eksisozluk-cli gundem [--page=SAYFA_SAYISI] [--limit=BASLIK_LIMITI] [--output=json|console]
				  Başlık için bulunan entry'leri listelemek için kullanılır.
				Seçenekler:
				  --page=SAYFA_SAYISI			Başlıklar için başlangıç sayfası (varsayılan 1)
				  --limit=BAŞLIK_LIMITI			Toplam listelenen entry limiti (varsayılan 10)
				  --output=json,console			console: Komut satırı çıktısı, json: JSON dosyası çıktısı (varsayılan console)
				`
	return strings.TrimSpace(helpText)
}

// Run handles the main operations of command
func (c *GundemCommand) Run(args []string) int {
	parameter, err := parameterFlagHandler(args, c, c.cli)
	if err != nil {
		return 1
	}
	if parameter.Limit == -1 {
		parameter.Limit = 10
	}

	topicList := GetPopularTopics(parameter)
	WriteTopicList(topicList, parameter)
	return 0
}

// Synopsis provides the usage
func (c *GundemCommand) Synopsis() string {
	return "Gündem başlıklarını listeleme"
}
