package model

import "gopkg.in/mgo.v2/bson"

const (
	ANSWER_RECORD_COLLECTION_NAME = "answer_record"
)

type Result int

const (
	RIGHT         Result = 1
	WRONG         Result = 2
	NOTBEST       Result = 3
	WAITTINGCHECK Result = 4
)

type AnswerRecord struct {
	ID         bson.ObjectId `json:"_id" bson:"_id"`
	UserID     bson.ObjectId `json:"userid" bson:"userid"`
	UserName   string        `json:"username" bson:"username"`
	QuestionID bson.ObjectId `json:"question" bson:"question"`
	Answer     []string      `json:"answer" bson:"answer"`
	AImage     bool          `json:"aimage" bson:"aimage"`
	IsTrue     Result        `json:"istrue" bson:"istrue"`
	// Knowledge  string        `json:"knowledge" bson:"knowledge"`
	Subject    string `json:"subject" bson:"subject"`
	HomeWorkId string `json:"homeworkid" bson:"homeworkid"`
	PaperId    string `json:"paperid" bson:"paperid"`
	Score      int64  `json:"score" bson:"socre"`
	CreateTime int64  `json:"createtime" bson:"createtime"`
}
