package interfaces

// TODO:
// func TestFindAll(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	const query = `
// 		SELECT
// 			id,
// 			title
// 		FROM
// 			posts
// 		LIMIT 10
// 	`

// 	columns := []string{"id", "title"}
// 	rows := sqlmock.NewRows(columns).
// 		AddRow(1, "foo-1").
// 		AddRow(2, "foo-2")
// 	mock.ExpectQuery(query).WillReturnRows(rows)

// 	pr := &PostRepository{
// 		Conn: db,
// 	}

// 	actual, err := pr.FindAll(1, 1)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// 	actual, ok := interface{}(actual).(domain.Posts)
// 	if !ok {
// 		t.Errorf("%v doesn't implement domain.Posts.", actual)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	expected := domain.Posts{
// 		domain.Post{
// 			ID:    1,
// 			Title: "foo-1",
// 		},
// 		domain.Post{
// 			ID:    2,
// 			Title: "foo-2",
// 		},
// 	}

// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Errorf("actual:%v expected:%v", actual, expected)
// 	}
// }
