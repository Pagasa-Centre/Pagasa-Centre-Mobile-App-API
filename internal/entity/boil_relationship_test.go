// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import "testing"

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CellGroupToUserUsingLeader", testCellGroupToOneUserUsingLeader)
	t.Run("MinistryToUserUsingLeader", testMinistryToOneUserUsingLeader)
	t.Run("MinistryToOutreachUsingOutreach", testMinistryToOneOutreachUsingOutreach)
	t.Run("OutreachServiceToOutreachUsingOutreach", testOutreachServiceToOneOutreachUsingOutreach)
	t.Run("UserRoleToRoleUsingRole", testUserRoleToOneRoleUsingRole)
	t.Run("UserRoleToUserUsingUser", testUserRoleToOneUserUsingUser)
	t.Run("UserToUserUsingCellLeader", testUserToOneUserUsingCellLeader)
	t.Run("UserToOutreachUsingOutreach", testUserToOneOutreachUsingOutreach)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CellGroupToUsers", testCellGroupToManyUsers)
	t.Run("OutreachToMinistries", testOutreachToManyMinistries)
	t.Run("OutreachToOutreachServices", testOutreachToManyOutreachServices)
	t.Run("OutreachToUsers", testOutreachToManyUsers)
	t.Run("RoleToUserRoles", testRoleToManyUserRoles)
	t.Run("UserToCellGroups", testUserToManyCellGroups)
	t.Run("UserToLeaderCellGroups", testUserToManyLeaderCellGroups)
	t.Run("UserToLeaderMinistries", testUserToManyLeaderMinistries)
	t.Run("UserToUserRoles", testUserToManyUserRoles)
	t.Run("UserToCellLeaderUsers", testUserToManyCellLeaderUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CellGroupToUserUsingLeaderCellGroups", testCellGroupToOneSetOpUserUsingLeader)
	t.Run("MinistryToUserUsingLeaderMinistries", testMinistryToOneSetOpUserUsingLeader)
	t.Run("MinistryToOutreachUsingMinistries", testMinistryToOneSetOpOutreachUsingOutreach)
	t.Run("OutreachServiceToOutreachUsingOutreachServices", testOutreachServiceToOneSetOpOutreachUsingOutreach)
	t.Run("UserRoleToRoleUsingUserRoles", testUserRoleToOneSetOpRoleUsingRole)
	t.Run("UserRoleToUserUsingUserRoles", testUserRoleToOneSetOpUserUsingUser)
	t.Run("UserToUserUsingCellLeaderUsers", testUserToOneSetOpUserUsingCellLeader)
	t.Run("UserToOutreachUsingUsers", testUserToOneSetOpOutreachUsingOutreach)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("CellGroupToUserUsingLeaderCellGroups", testCellGroupToOneRemoveOpUserUsingLeader)
	t.Run("MinistryToUserUsingLeaderMinistries", testMinistryToOneRemoveOpUserUsingLeader)
	t.Run("UserToUserUsingCellLeaderUsers", testUserToOneRemoveOpUserUsingCellLeader)
	t.Run("UserToOutreachUsingUsers", testUserToOneRemoveOpOutreachUsingOutreach)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CellGroupToUsers", testCellGroupToManyAddOpUsers)
	t.Run("OutreachToMinistries", testOutreachToManyAddOpMinistries)
	t.Run("OutreachToOutreachServices", testOutreachToManyAddOpOutreachServices)
	t.Run("OutreachToUsers", testOutreachToManyAddOpUsers)
	t.Run("RoleToUserRoles", testRoleToManyAddOpUserRoles)
	t.Run("UserToCellGroups", testUserToManyAddOpCellGroups)
	t.Run("UserToLeaderCellGroups", testUserToManyAddOpLeaderCellGroups)
	t.Run("UserToLeaderMinistries", testUserToManyAddOpLeaderMinistries)
	t.Run("UserToUserRoles", testUserToManyAddOpUserRoles)
	t.Run("UserToCellLeaderUsers", testUserToManyAddOpCellLeaderUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("CellGroupToUsers", testCellGroupToManySetOpUsers)
	t.Run("OutreachToUsers", testOutreachToManySetOpUsers)
	t.Run("UserToCellGroups", testUserToManySetOpCellGroups)
	t.Run("UserToLeaderCellGroups", testUserToManySetOpLeaderCellGroups)
	t.Run("UserToLeaderMinistries", testUserToManySetOpLeaderMinistries)
	t.Run("UserToCellLeaderUsers", testUserToManySetOpCellLeaderUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("CellGroupToUsers", testCellGroupToManyRemoveOpUsers)
	t.Run("OutreachToUsers", testOutreachToManyRemoveOpUsers)
	t.Run("UserToCellGroups", testUserToManyRemoveOpCellGroups)
	t.Run("UserToLeaderCellGroups", testUserToManyRemoveOpLeaderCellGroups)
	t.Run("UserToLeaderMinistries", testUserToManyRemoveOpLeaderMinistries)
	t.Run("UserToCellLeaderUsers", testUserToManyRemoveOpCellLeaderUsers)
}
