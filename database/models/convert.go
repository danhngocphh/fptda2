package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/danhngocphh/fptda2/lib/common"
)

// Post data model
type Convert struct {
	Voice   string `sql:"type:text;"`
	Text   string `sql:"type:text;"`
	Speed   int `sql:"type:int;"`
	Format   string `sql:"type:text;"`
}

c.ghhh()


PlayerInterface.java
class interface PlayerInterface () {
	static function next();
}
PlayerController.java
class PlayerController implement PlayerInterface () {
	static function next(){
		int a = 1;
		return a;
	};
}

main.java
public static function main([]){
	PlayerInterface model = new PlayerController();
}





// Serialize serializes post data
func (c Convert) Serialize() common.JSON {
	return common.JSON{
		"Voice":       c.Voice,
		"Text":       c.Text,
		"Speed":       c.Speed,
		"Format":       c.Format,
	}
}
