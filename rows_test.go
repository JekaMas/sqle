// Copyright 2017 Lazada South East Asia Pte. Ltd.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqle

import (
	"testing"
	"time"

	"github.com/lazada/sqle/testdata"
)

func TestRows_ScanMap(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	m := make(map[string]interface{})
	if err = rows.Scan(m); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "%#v\n", m)
}

func BenchmarkRows_ScanMap(b *testing.B) {
	rows, err := db.Query(testdata.SelectUserStmt, 1)
	if err != nil {
		b.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m := make(map[string]interface{})
		if err = rows.Scan(m); err != nil {
			b.Errorf("(%T).Scan() failed: %s", rows, err)
		}
		_ = m
	}
}

func TestRows_ScanPtrMap(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	m := make(map[string]interface{})
	if err = rows.Scan(&m); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "%#v\n", m)
}

func TestRows_ScanVarMap(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	id, name, m := 0, ``, make(map[string]interface{})
	if err = rows.Scan(&id, &name, m); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "id: %#v, name: %#v, %#v\n", id, name, m)
}

func TestRows_ScanAnonStruct(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	u := struct {
		Id      int
		Name    string
		Email   *string
		Created time.Time
		Updated *time.Time
	}{}
	if err = rows.Scan(&u); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "%#v\n", u)
}

func BenchmarkRows_ScanAnonStruct(b *testing.B) {
	rows, err := db.Query(testdata.SelectUserStmt, 1)
	if err != nil {
		b.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		u := struct {
			Id      int
			Name    string
			Email   *string
			Created time.Time
			Updated *time.Time
		}{}
		if err = rows.Scan(&u); err != nil {
			b.Errorf("(%T).Scan() failed: %s", rows, err)
		}
		_ = u
	}
}

func TestRows_ScanVarAnonStruct(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	id, name := 0, ``
	u := struct {
		Email   *string
		Created time.Time
		Updated *time.Time
	}{}
	if err = rows.Scan(&id, &name, &u); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "id: %#v, name: %#v, %#v\n", id, name, u)
}

func TestRows_ScanVarAnonStructVar(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	id, updated := 0, new(time.Time)
	u := struct {
		Name    string
		Email   *string
		Created time.Time
	}{}
	if err = rows.Scan(&id, &u, &updated); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "id: %#v, updated: %#v, %#v\n", id, updated, u)
}

func TestRows_ScanAnonPart(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	p := struct{ Name string }{}
	if err = rows.Scan(&p); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "%#v\n", p)
}

func TestRows_ScanVarAnonPart(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	id, p := 0, struct{ Name string }{}
	if err = rows.Scan(&id, &p); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "id: %#v, %#v\n", id, p)
}

func TestRows_ScanVarAnonPart2(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	id, p1, p2 := 0, struct {
		Name  string
		Email *string
	}{}, struct {
		Created time.Time
		Updated *time.Time
	}{}
	if err = rows.Scan(&id, &p1, &p2); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "id: %#v, %#v, %#v\n", id, p1, p2)
}

func TestRows_ScanStruct(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	u := testdata.User{}
	if err = rows.Scan(&u); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "%#v\n", u)
}

func BenchmarkRows_ScanStruct(b *testing.B) {
	rows, err := db.Query(testdata.SelectUserStmt, 1)
	if err != nil {
		b.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		u := new(testdata.User)
		if err = rows.Scan(u); err != nil {
			b.Errorf("(%T).Scan() failed: %s", rows, err)
		}
		_ = u
	}
}

func TestRows_ScanVarStruct(t *testing.T) {
	rows, err := db.Query(testdata.SelectUserStmt, nextUserId())
	if err != nil {
		t.Fatalf("(%T).Query() failed: %s", db, err)
	}
	defer rows.Close()
	rows.Next()

	id, u := 0, testdata.User{}
	if err = rows.Scan(&id, &u); err != nil {
		t.Errorf("(%T).Scan() failed: %s", rows, err)
	}
	debugf(t, "id: %#v, %#v\n", id, u)
}
