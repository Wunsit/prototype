package user

import "github.com/dgraph-io/dgo"

type DGraphStore struct {
	db *dgo.Dgraph
}

func NewDGraphStore(db *dgo.Dgraph) {

}

func (ds *DGraphStore) Get(id int) User {
	return User{}
}

func (ds *DGraphStore) Add(u User) error {
	return nil
}
