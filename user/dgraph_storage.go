package user

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo/protos/api"

	"github.com/dgraph-io/dgo"
)

type DGraphStore struct {
	db *dgo.Dgraph
}

func NewDGraphStore(db *dgo.Dgraph) *DGraphStore {
	return &DGraphStore{db}
}

const queryByEmail = `query User($email: string){
	root(func: eq(user.email, $email)) {
		email: user.email
		firstname: user.firstname
		lastname: user.lastname
	}
}`

func (ds *DGraphStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	resp, err := ds.db.NewTxn().QueryWithVars(
		ctx,
		queryByEmail,
		map[string]string{"$email": email},
	)

	if err != nil {
		return nil, err
	}

	r := struct {
		Root []User `json:"root"`
	}{}

	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		return nil, err
	}
	return &r.Root[0], nil
}

func (ds *DGraphStore) Add(ctx context.Context, u User) (err error) {
	mu := &api.Mutation{CommitNow: true}
	if mu.SetJson, err = json.Marshal(u); err != nil {
		return err
	}

	if _, err = ds.db.NewTxn().Mutate(ctx, mu); err != nil {
		return err
	}
	return nil
}
