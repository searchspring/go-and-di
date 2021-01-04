package campaigns

import (
	"database/sql"
	"github.com/searchspring/go-basics/clients/sscore"
)

type Deps struct {
	SSCore sscore.SSCore
}

type Campaigns interface {
	Exists(id string) (bool, error)
}

type impl struct {
	deps *Deps
}

func New(deps *Deps) Campaigns {
	return &impl{deps: deps}
}

func getCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err:= rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	return count
}

func (impl *impl) Exists(id string) (bool, error) {
	rows, err := impl.deps.SSCore.Query(`SELECT COUNT(*) FROM MerchandisingCampaigns WHERE id = ?`, id)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	return getCount(rows) > 0, nil
}
