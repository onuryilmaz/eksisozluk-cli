[![Build Status](https://travis-ci.org/onuryilmaz/eksisozluk-cli.svg?branch=master)](https://travis-ci.org/onuryilmaz/eksisozluk-cli)


### Build
```go
go build
```

### Run

Başlık ile arama:
```go
eksisozluk-cli baslik "golang" [--page=PAGE_NUMBER] [--limit=ENTRY_LIMIT]
```

Gündem Başlıkları:
```go
eksisozluk-cli gundem
```

DEBE:
```go
eksisozluk-cli debe
```

### TODO
- [ ] Gündem: static & live with limitation like top 10 & coloring
- [ ] Başlık: "şükela mod"
- [ ] Debe: Limit
- [ ] JSON, CSV export
- [ ] Logger