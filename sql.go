package main

import "log"

func (me *post) storeMessage() {

	ldb.stmtInsertMessage.Exec(me.username, me.time, me.message)
}

func getNote() (result []post) {
	sqlReadall := `
	SELECT * FROM chats
	`

	rows, err := ldb.db.Query(sqlReadall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		post := post{}
		err = rows.Scan(&post.username, &post.time, &post.message)
		if err != nil {
			panic(err)
		}
		result = append(result, post)
	}
	log.Printf("\nThese are the results: %s\n ", result)

	return result
}
