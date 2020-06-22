package dao

import "github.com/wudaoluo/sonic/model"

//




type ImUserService struct {}

var DBImUser *ImUserService
func init() {
    DBImUser = &ImUserService{}
}



func (t *ImUserService) SelectByUserName(username string) (*model.ImUser,error) {
    sqlText := "SELECT  avatar,email,password,uid,username FROM IM_USER WHERE username = ? limit 1"
    row := db.QueryRow(sqlText,username)

    msg := new(model.ImUser)
    err := row.Scan(
        &msg.Avatar,
        &msg.Email,
        &msg.Password,
        &msg.Uid,
        &msg.Username,
    )
    if err != nil {
        return nil,err
    }

    return msg,nil
}


func (t *ImUserService) Select() ([]*model.ImUser,error) {
    return nil,nil
}

func (t *ImUserService) Insert(msg *model.ImUser) (int64,error) {
    sqlText := "INSERT INTO IM_USER (  email,  password,  uid,  username) " +
        "VALUE (  ?,  ?,  ?,  ?)"

    ret, err := db.Exec(sqlText,
            
            
            &msg.Email,
            &msg.Password,
            &msg.Uid,
            &msg.Username,)

    if err != nil {
        return 0,err
    }
    return ret.LastInsertId()
}

func (t *ImUserService) DeleteById (id int64) error {
    sqlText := "DELETE FROM IM_USER where id = ?"
    _, err := db.Exec(sqlText,id)
    return err
}
