package utils

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ToDoc : Convert struct to bson document
func ToDoc(v interface{}) (doc bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	var update bson.D
	err = bson.Unmarshal(data, &update)
	if err != nil {
		return
	}

	for k, v := range update.Map() {
		if v != nil {
			fmt.Println(k, v)
			s := strings.Split(k, ":")
			if len(s) > 0 && s[0] == "id" {
				k = s[1]
				v, _ = primitive.ObjectIDFromHex(v.(string))
			}
			doc = append(doc, bson.E{Key: k, Value: v})
		}
	}

	return
}
