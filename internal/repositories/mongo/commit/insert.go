
	package commit

	import (
		"context"
		"log"
		"os"

		"github.com/fxivan/commitnamegen_ia/internal/domain"
	)

	func (r Repository) Insert(commit domain.Commit) (id interface{}, err error) {

		collection := r.Client.Database(os.Getenv("MONGO_NAME_DB")).Collection(os.Getenv("MONGO_NAME_COLLECTION"))
		insertResult, err := collection.InsertOne(context.Background(), commit)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		return insertResult.InsertedID, nil
	}
