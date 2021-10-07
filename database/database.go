package database

import(
	"log"
	"app/config"
	"database/sql"
	_"github.com/lib/pq"
)

type Airport struct {
	Id 		int
	Status	bool
	Description string
}

const SQL_GET_QUERY = `
	select
		plan_id,
		status,
		description
	from plan
` 


func GetBase() []Airport {

	db, err := sql.Open("postgres", config.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("%v", err)
	}

	rows, err := db.Query(SQL_GET_QUERY)

	if err != nil {
		log.Fatalf("%v", err)
	}

	var plans []Airport

	for rows.Next(){

		var plan Airport

		rows.Scan(
			&plan.Id,
			&plan.Status,
			&plan.Description,
		)

		plans = append(plans, plan)
	}

	return plans
}