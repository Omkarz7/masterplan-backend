package database

import (
	"database/sql"
	"errors"
	"masterplan-backend/models"
)

// VerifyCredentials authenticates user credentials
func VerifyCredentials(cred models.Credentials) (err error) {
	query := "select username from UserCredentials WHERE username = ? and password = ?"
	row := DBconn.QueryRow(query, cred.Username, cred.PasswordHash)

	err = row.Scan(&cred.Username)
	switch err {
	case sql.ErrNoRows:
		return errors.New("Invalid Username or Password")
	case nil:
		return nil
	default:
		return err
	}

}

//GetMasterplanFromDB retrieve data from DB to ber written in excel file
//Excel file will be created only for the one project for simplicity purpose
func GetMasterplanFromDB(username string) (masterplanList []models.Masterplan, err error) {
	projectID, err := getFirstProjectIDofUser(username)
	query := "SELECT sr_no, activity, start_date, end_date from Masterplan WHERE project_id = ?"

	rows, err := DBconn.Query(query, projectID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var masterplan models.Masterplan

	for rows.Next() {
		err := rows.Scan(&masterplan.SlNo, &masterplan.Activity, &masterplan.StartDate, &masterplan.EndDate)
		if err != nil {

			return nil, err
		}
		masterplanList = append(masterplanList, masterplan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return masterplanList, nil
}

func getFirstProjectIDofUser(username string) (projectID int, err error) {
	query := "SELECT pd.id from UserCredentials uc inner join ProjectDetails pd on uc.id = pd.project_manager_id WHERE uc.username = ?"
	row := DBconn.QueryRow(query, username)

	err = row.Scan(&projectID)
	switch err {
	case sql.ErrNoRows:
		return -1, errors.New("Invalid Username or Password")
	case nil:
		return projectID, nil
	default:
		return -1, err
	}
}
