// Package dao contains the types for schema 'mj'.
package dao

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// Role represents a row from 'mj.role'.
type Role struct {
	IndexID     int32  `json:"index_id"`    // index_id
	Role        string `json:"role"`        // role
	Comment     string `json:"comment"`     // comment
	Permissions string `json:"permissions"` // permissions

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Role exists in the database.
func (r *Role) Exists() bool {
	return r._exists
}

// Deleted provides information if the Role has been deleted from the database.
func (r *Role) Deleted() bool {
	return r._deleted
}

// Insert inserts the Role to the database.
func (r *Role) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO mj.role (` +
		`role, comment, permissions` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, r.Role, r.Comment, r.Permissions)
	res, err := db.Exec(sqlstr, r.Role, r.Comment, r.Permissions)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	r.IndexID = int32(id)
	r._exists = true

	return nil
}

// Update updates the Role in the database.
func (r *Role) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE mj.role SET ` +
		`role = ?, comment = ?, permissions = ?` +
		` WHERE index_id = ?`

	// run query
	XOLog(sqlstr, r.Role, r.Comment, r.Permissions, r.IndexID)
	_, err = db.Exec(sqlstr, r.Role, r.Comment, r.Permissions, r.IndexID)
	return err
}

// Save saves the Role to the database.
func (r *Role) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Delete deletes the Role from the database.
func (r *Role) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM mj.role WHERE index_id = ?`

	// run query
	XOLog(sqlstr, r.IndexID)
	_, err = db.Exec(sqlstr, r.IndexID)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// RoleByIndexID retrieves a row from 'mj.role' as a Role.
//
// Generated from index 'role_index_id_pkey'.
func RoleByIndexID(db XODB, indexID int32) (*Role, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, role, comment, permissions ` +
		`FROM mj.role ` +
		`WHERE index_id = ?`

	// run query
	XOLog(sqlstr, indexID)
	r := Role{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, indexID).Scan(&r.IndexID, &r.Role, &r.Comment, &r.Permissions)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// RoleByRole retrieves a row from 'mj.role' as a Role.
//
// Generated from index 'uidx_role'.
func RoleByRole(db XODB, role string) (*Role, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, role, comment, permissions ` +
		`FROM mj.role ` +
		`WHERE role = ?`

	// run query
	XOLog(sqlstr, role)
	r := Role{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, role).Scan(&r.IndexID, &r.Role, &r.Comment, &r.Permissions)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
