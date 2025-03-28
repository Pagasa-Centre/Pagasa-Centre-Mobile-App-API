// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("CellGroups", testCellGroups)
	t.Run("GooseDBVersions", testGooseDBVersions)
	t.Run("Media", testMedia)
	t.Run("Ministries", testMinistries)
	t.Run("Outreaches", testOutreaches)
	t.Run("Roles", testRoles)
	t.Run("UserRoles", testUserRoles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("CellGroups", testCellGroupsDelete)
	t.Run("GooseDBVersions", testGooseDBVersionsDelete)
	t.Run("Media", testMediaDelete)
	t.Run("Ministries", testMinistriesDelete)
	t.Run("Outreaches", testOutreachesDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("UserRoles", testUserRolesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("CellGroups", testCellGroupsQueryDeleteAll)
	t.Run("GooseDBVersions", testGooseDBVersionsQueryDeleteAll)
	t.Run("Media", testMediaQueryDeleteAll)
	t.Run("Ministries", testMinistriesQueryDeleteAll)
	t.Run("Outreaches", testOutreachesQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("UserRoles", testUserRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("CellGroups", testCellGroupsSliceDeleteAll)
	t.Run("GooseDBVersions", testGooseDBVersionsSliceDeleteAll)
	t.Run("Media", testMediaSliceDeleteAll)
	t.Run("Ministries", testMinistriesSliceDeleteAll)
	t.Run("Outreaches", testOutreachesSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("UserRoles", testUserRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("CellGroups", testCellGroupsExists)
	t.Run("GooseDBVersions", testGooseDBVersionsExists)
	t.Run("Media", testMediaExists)
	t.Run("Ministries", testMinistriesExists)
	t.Run("Outreaches", testOutreachesExists)
	t.Run("Roles", testRolesExists)
	t.Run("UserRoles", testUserRolesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("CellGroups", testCellGroupsFind)
	t.Run("GooseDBVersions", testGooseDBVersionsFind)
	t.Run("Media", testMediaFind)
	t.Run("Ministries", testMinistriesFind)
	t.Run("Outreaches", testOutreachesFind)
	t.Run("Roles", testRolesFind)
	t.Run("UserRoles", testUserRolesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("CellGroups", testCellGroupsBind)
	t.Run("GooseDBVersions", testGooseDBVersionsBind)
	t.Run("Media", testMediaBind)
	t.Run("Ministries", testMinistriesBind)
	t.Run("Outreaches", testOutreachesBind)
	t.Run("Roles", testRolesBind)
	t.Run("UserRoles", testUserRolesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("CellGroups", testCellGroupsOne)
	t.Run("GooseDBVersions", testGooseDBVersionsOne)
	t.Run("Media", testMediaOne)
	t.Run("Ministries", testMinistriesOne)
	t.Run("Outreaches", testOutreachesOne)
	t.Run("Roles", testRolesOne)
	t.Run("UserRoles", testUserRolesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("CellGroups", testCellGroupsAll)
	t.Run("GooseDBVersions", testGooseDBVersionsAll)
	t.Run("Media", testMediaAll)
	t.Run("Ministries", testMinistriesAll)
	t.Run("Outreaches", testOutreachesAll)
	t.Run("Roles", testRolesAll)
	t.Run("UserRoles", testUserRolesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("CellGroups", testCellGroupsCount)
	t.Run("GooseDBVersions", testGooseDBVersionsCount)
	t.Run("Media", testMediaCount)
	t.Run("Ministries", testMinistriesCount)
	t.Run("Outreaches", testOutreachesCount)
	t.Run("Roles", testRolesCount)
	t.Run("UserRoles", testUserRolesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("CellGroups", testCellGroupsHooks)
	t.Run("GooseDBVersions", testGooseDBVersionsHooks)
	t.Run("Media", testMediaHooks)
	t.Run("Ministries", testMinistriesHooks)
	t.Run("Outreaches", testOutreachesHooks)
	t.Run("Roles", testRolesHooks)
	t.Run("UserRoles", testUserRolesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("CellGroups", testCellGroupsInsert)
	t.Run("CellGroups", testCellGroupsInsertWhitelist)
	t.Run("GooseDBVersions", testGooseDBVersionsInsert)
	t.Run("GooseDBVersions", testGooseDBVersionsInsertWhitelist)
	t.Run("Media", testMediaInsert)
	t.Run("Media", testMediaInsertWhitelist)
	t.Run("Ministries", testMinistriesInsert)
	t.Run("Ministries", testMinistriesInsertWhitelist)
	t.Run("Outreaches", testOutreachesInsert)
	t.Run("Outreaches", testOutreachesInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("UserRoles", testUserRolesInsert)
	t.Run("UserRoles", testUserRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("CellGroups", testCellGroupsReload)
	t.Run("GooseDBVersions", testGooseDBVersionsReload)
	t.Run("Media", testMediaReload)
	t.Run("Ministries", testMinistriesReload)
	t.Run("Outreaches", testOutreachesReload)
	t.Run("Roles", testRolesReload)
	t.Run("UserRoles", testUserRolesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("CellGroups", testCellGroupsReloadAll)
	t.Run("GooseDBVersions", testGooseDBVersionsReloadAll)
	t.Run("Media", testMediaReloadAll)
	t.Run("Ministries", testMinistriesReloadAll)
	t.Run("Outreaches", testOutreachesReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("UserRoles", testUserRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("CellGroups", testCellGroupsSelect)
	t.Run("GooseDBVersions", testGooseDBVersionsSelect)
	t.Run("Media", testMediaSelect)
	t.Run("Ministries", testMinistriesSelect)
	t.Run("Outreaches", testOutreachesSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("UserRoles", testUserRolesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("CellGroups", testCellGroupsUpdate)
	t.Run("GooseDBVersions", testGooseDBVersionsUpdate)
	t.Run("Media", testMediaUpdate)
	t.Run("Ministries", testMinistriesUpdate)
	t.Run("Outreaches", testOutreachesUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("UserRoles", testUserRolesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("CellGroups", testCellGroupsSliceUpdateAll)
	t.Run("GooseDBVersions", testGooseDBVersionsSliceUpdateAll)
	t.Run("Media", testMediaSliceUpdateAll)
	t.Run("Ministries", testMinistriesSliceUpdateAll)
	t.Run("Outreaches", testOutreachesSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("UserRoles", testUserRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
