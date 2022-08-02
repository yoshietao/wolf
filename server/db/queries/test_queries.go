package queries

import (
  // Import this so we don't have to use qm.Limit etc.
  . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

db, err := sql.Open("mysql", "dbname=wolf user=wolf")
defer fb.close()
if err != nil {
  return err
}

