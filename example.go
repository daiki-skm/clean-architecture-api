package main

import (
	//"context"
	//"fmt"
	//"log"
	"net/http"

	"github.com/labstack/echo"
	//"cloud.google.com/go/firestore"
	//firebase "firebase.google.com/go"
	//"google.golang.org/api/iterator"
	//"google.golang.org/api/option"
)

//type User struct {
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Response struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	e := echo.New()
	//e.GET("/users/:name", getUserName)
	//e.GET("/show", show)
	//e.POST("/save", save)
	//e.POST("/users", saveUser)
	e.POST("/send", sendMessage)
	e.Logger.Fatal(e.Start(":1323"))

	// 初期化
	//ctx := context.Background()
	//sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	//app, err := firebase.NewApp(ctx, nil, sa)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//client, err := app.Firestore(ctx)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// データ追加
	//_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
	//	"first": "Ada",
	//	"middle": "Mathison",
	//	"last":  "Lovelace",
	//	"born":  1815,
	//})
	//if err != nil {
	//	log.Fatalf("Failed adding alovelace: %v", err)
	//}

	// データ読み取り
	//iter := client.Collection("users").Documents(ctx)
	//for {
	//	doc, err := iter.Next()
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Failed to iterate: %v", err)
	//	}
	//	fmt.Println(doc.Data())
	//}

	// データ追加
	//_, err = client.Collection("users").Doc("user2").Set(ctx, map[string]interface{}{
	//	"first":  "Ada",
	//	"middle": "Mathison",
	//	"last":   "Lovelace",
	//	"born":   1815,
	//})
	//if err != nil {
	//	log.Fatalf("Failed adding alovelace: %v", err)
	//}

	// データ更新
	//_, updateError := client.Collection("users").Doc("user2").Set(ctx, map[string]interface{}{
	//	"first": "Yeah",
	//}, firestore.MergeAll)
	//if updateError != nil {
	//	// Handle any errors in an appropriate way, such as returning them.
	//	log.Printf("An error has occurred: %s", err)
	//}

	// フィールド削除
	//_, errorDelete := client.Collection("users").Doc("user2").Update(ctx, []firestore.Update{
	//	{
	//		Path:  "middle",
	//		Value: firestore.Delete,
	//	},
	//})
	//if errorDelete != nil {
	//	// Handle any errors in an appropriate way, such as returning them.
	//	log.Printf("An error has occurred: %s", err)
	//}

	// ドキュメント削除
	//_, errorDelete := client.Collection("users").Doc("uesr2").Delete(ctx)
	//if errorDelete != nil {
	//	// Handle any errors in an appropriate way, such as returning them.
	//	log.Printf("An error has occurred: %s", err)
	//}

	//ref := client.Collection("users")
	//deleteCollection(ctx, client, ref, 10)

	// 切断
	//defer client.Close()
}

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if error := c.Bind(m); error != nil {
		return error
	}
	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(http.StatusOK, r)
}

//func deleteCollection(ctx context.Context, client *firestore.Client,
//	ref *firestore.CollectionRef, batchSize int) error {
//
//	for {
//		// Get a batch of documents
//		iter := ref.Limit(batchSize).Documents(ctx)
//		numDeleted := 0
//
//		// Iterate through the documents, adding
//		// a delete operation for each one to a
//		// WriteBatch.
//		batch := client.Batch()
//		for {
//			doc, err := iter.Next()
//			if err == iterator.Done {
//				break
//			}
//			if err != nil {
//				return err
//			}
//
//			batch.Delete(doc.Ref)
//			numDeleted++
//		}
//
//		// If there are no documents to delete,
//		// the process is over.
//		if numDeleted == 0 {
//			return nil
//		}
//
//		_, err := batch.Commit(ctx)
//		if err != nil {
//			return err
//		}
//	}
//}
