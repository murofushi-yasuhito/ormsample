package test

import (
	"github.com/ant0ine/go-json-rest/rest"
	//"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" // エイリアスでprefixを省略できる
	"net/http"
	"ormSample/domain"
)

// パスパラメータ:idの国の該当データを出力
func (g *Gorm) GetUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := domain.User{}

	if g.Db.Debug().Table("USER_MST").Where("USER_ID = ?", id).Find(&user).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&user)
}

// countriesテーブル内のデータ全出力
func (g *Gorm) GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	country := []domain.Country{}
	g.Db.Table("country").Find(&country)
	w.WriteJson(country)
}

// json形式のデータをPOST {name:国名}
func (g *Gorm) PostUser(w rest.ResponseWriter, r *rest.Request) {
	user := domain.User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = g.Db.Debug().Table("USER_MST").Create(&user).Error
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&user)
}

// パスパラメータ:idの国の該当データのNameを変更し出力
func (g *Gorm) PutUser(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	user := domain.User{}
	if g.Db.Debug().Table("USER_MST").First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	updated := domain.User{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.UserNm = updated.UserNm

	if err := g.Db.Table("USER_MST").Debug().Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&user)
}

// パスパラメータ:idの国の該当データを削除
func (g *Gorm) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := domain.User{}
	if g.Db.Table("USER_MST").First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := g.Db.Table("USER_MST").Debug().Delete(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	w.WriteHeader(http.StatusOK)
}
