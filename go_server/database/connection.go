package database


import(
"go_server/database"
 "gorm.io/gorm"
  "gorm.io/driver/mysql"

)

func Connect(){
connection, err := gorm.Open(mysql.Open("root:@/rocknroll"), &gorm.Config{})

if err != nil {
panic("Could not connect to the database")
}
connection.AutoMigrate(database.User{})

}
