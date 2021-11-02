package UsersSpendsManager

import "swap.io-ledger/src/database"

type UsersSpendsManager struct {
	database *database.Database
}
type Config struct {
	Database *database.Database
}

func InitialiseUsersSpendsManager(config Config) *UsersSpendsManager {
	return &UsersSpendsManager{
		database: config.Database,
	}
}

func (*UsersSpendsManager) Start() {}
func (*UsersSpendsManager) Status() error {
	return nil
}
func (*UsersSpendsManager) Stop() error {
	return nil
}