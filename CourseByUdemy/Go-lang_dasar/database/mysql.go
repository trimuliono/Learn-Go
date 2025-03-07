package database

var connection string // variabel ini tidak bisa diakses dari luar package

// fungsi ini akan dijalankan pertama kali
func init() { 
	connection = "MySQL"
}

func GetDatabase() string { // fungsi ini bisa diakses dari luar package
	return connection
}