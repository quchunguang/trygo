package main

// [Ref]: http://nerglish.tumblr.com/post/90671946112/reading-arrays-from-postgres-in-golang
import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)

type UserScores struct {
	username string
	score    float64
	scores   FloatSlice
}

func (u UserScores) String() string {
	s := fmt.Sprintf("%s\t%3.1f\t%v", u.username, u.score, u.scores)
	return s
}

type FloatSlice []float64

func (s *FloatSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	asString := string(asBytes)
	*s = strToFloatSlice(asString)
	return nil
}
func strToFloatSlice(s string) []float64 {
	r := strings.Trim(s, "{}")
	a := make([]float64, 0, 10)
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.ParseFloat(t, 64)
		a = append(a, i)
	}
	return a
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=test host=localhost password=123456 sslmode=disable")
	checkErr(err)
	defer db.Close()

	var selectStatement = `select name, score, scores from test1 limit 10`
	rows, err := db.Query(selectStatement)
	checkErr(err)

	fmt.Printf("%s\t%s\t%s\n", "用户名", "得分", "得分详细")
	fmt.Println("--------------------------")
	for rows.Next() {
		user := new(UserScores)
		err = rows.Scan(&user.username, &user.score, &user.scores)
		checkErr(err)

		fmt.Println(user)
	}
}
