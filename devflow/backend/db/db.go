package db

const MongoDBName = "MONGO_DB_NAME"

type Store struct {
	Question QuestionStore
}