package repository

import "trec/internal/domain"

type Factory struct {
	infra infraDB
}

func NewFactory(infra infraDB) *Factory {
	return &Factory{infra: infra}
}

func (f *Factory) CreateRecordRepository() (domain.RecordRepository, error) {
	db := newRecordsDatabase(f.infra)
	if err := db.ensureTable(); err != nil {
		return nil, err
	}
	return db, nil
}
