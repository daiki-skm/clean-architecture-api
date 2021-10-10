package repositories

import (
	"context"
	"encoding/json"
	"log"

	"example/domain"
	"example/infrastructure"
	"example/usecase/services"

	"github.com/labstack/echo"
)

type UsersRepository struct {
	echo *echo.Echo
	usersService services.UsersService
}

func NewUsersRepository(
	echo *echo.Echo,
) *UsersRepository {
	return &UsersRepository{
		echo: echo,
	}
}

func (r *UsersRepository) AddUsers(ec echo.Context, user *domain.User) ([]*domain.User, error) {
	ctx := context.Background()
	client, err := infrastructure.FirebaseInit(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// データ追加
	_, err = client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{}{
		"age":     user.Age,
		"address": user.Address,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	// データ読み込み
	allData := client.Collection("users").Documents(ctx)
	// 全ドキュメント取得
	docs, err := allData.GetAll()
	if err != nil {
		log.Fatalf("Failed adding getAll: %v", err)
	}
	// 配列の初期化
	users := make([]*domain.User, 0)
	for _, doc := range docs {
		// 構造体の初期化
		u := new(domain.User)
		// 構造体にFireStoreのデータをセット
		mapToStruct(doc.Data(), &u)
		// ドキュメント名を取得してnameにセット
		u.Name = doc.Ref.ID
		// 配列に構造体をセット
		users = append(users, u)
	}

	// 切断
	defer client.Close()

	// 成功していればusersに値が、失敗の場合はerrに値が入る
	return users, err
}

func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
