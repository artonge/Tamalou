# Install
go get github.com/artonge/Tamalou

# Dependencies
- [mysql - *mySQL drivers*](https://godoc.org/database/sql/driver#Conn)
- [sqlite3 - *SQLite drivers*](https://godoc.org/github.com/mattn/go-sqlite3)
- [couchdb - *couchDB drivers*](https://godoc.org/github.com/rhinoman/couchdb-go)
- [bleve - *indexing*](github.com/blevesearch/bleve)


# Data Sources

## Sickness
- OrphaData, *couchDB* : `http://couchdb.telecomnancy.univ-lorraine.fr/orphadatabase`
- OMIM, *csv* : `omim_onto.csv`

## Synonymes
- HPO, *SQLite* : `hpo_annotations.sqlite`

## After effect
- Sider, *mySQL* : `gmd-read:esial@neptune.telecomnancy.univ-lorraine.fr/gmd`

# Use
```
symptoms --> [Tamalou]

[Tamalou] --> {

  sicknesses {
    name
    drugs
  }

  drugs {
    name
    drugs
  }
}
```

# On request
- Get synonymes with HPO
- Query OrphaData and OMIM for sickness
- Query Sider for after effects
- Map the fetched datas
- Return results
