package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

// TestMain 测试类的入口
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", TestAddUser)
	t.Run("get", TestGetUser)
	t.Run("del", TestDeleteUser)
	t.Run("reGet", testReGetUser)
}

func TestAddUser(t *testing.T) {
	err := AddCredential("avensso", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	credential, err := GetUserCredential("avensso")
	if credential != "123" && err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}
	t.Logf("credential: %v", credential)
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("avensso", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testReGetUser(t *testing.T) {
	credential, err := GetUserCredential("avensso")
	if err != nil {
		t.Errorf("Error of ReGetUser:%v", err)
	}
	if credential != "" {
		t.Errorf("Deleting user test failed")
	}
}
