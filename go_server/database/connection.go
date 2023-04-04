package database


import(
 "gorm.io/gorm"
  "gorm.io/driver/mysql"

)

func Connect(){
_, err := gorm.Open(mysql.Open("root:@/rocknroll"), &gorm.Config{})

if err != nil {
panic("Could not connect to the database")
}

}
