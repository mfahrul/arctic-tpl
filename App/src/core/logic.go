package [[ with .ModuleToParse ]][[.Name]][[ end ]]

import (
	"encoding/json"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"[[.Projectpath]]/utils"
)

type service struct {
	repo   Repository
	logger log.Logger
}

//NewService function
func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

[[ with .ModuleToParse.Model ]]

func (s service) Create[[.Name | ToCamel]](itm Create[[.Name | ToCamel]]Req) ([[.Name | ToLower]]ID string, err error) {
	logger := log.With(s.logger, "method", "Register [[.Name | ToCamel]]")

	tn := time.Now()
	createdBy, _ := primitive.ObjectIDFromHex(itm.CreatedBy)
	new[[.Name | ToCamel]] := [[.Name | ToCamel]]{ [[ with .Structures ]][[ range . ]]
		[[.Name | ToCamel]]:       itm.[[.Name | ToCamel]], [[ end ]][[ end ]]
		CreatedBy:      &createdBy,
		CreatedAt:      &tn,
	}

	[[.Name | ToLower]]ID, err = s.repo.Create[[.Name | ToCamel]](new[[.Name | ToCamel]])
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		level.Info(logger).Log("Create [[.Name | ToLower]] success with id ", [[.Name | ToLower]]ID)
	}

	return
}

func (s service) GetAll[[.Name | ToCamel | ToPlural]]() (items [][[.Name | ToCamel]], err error) {
	logger := log.With(s.logger, "method", "Get All [[.Name | ToCamel]]")

	items, err = s.repo.GetAll[[.Name | ToCamel | ToPlural]]()
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		// i, _ := json.Marshal(items)
		level.Info(logger).Log("Get items ", len(items))
	}

	return
}

func (s service) Get[[.Name | ToCamel]]ByID(id string) ([[.Name | ToLower]] [[.Name | ToCamel]], err error) {
	logger := log.With(s.logger, "method", "Get[[.Name | ToCamel]]")

	[[.Name | ToLower]], err = s.repo.Get[[.Name | ToCamel]]ByID(id)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		i, _ := json.Marshal([[.Name | ToLower]])
		level.Info(logger).Log("Get [[.Name | ToLower]] ", string(i))
	}

	return
}

func (s service) Update[[.Name | ToCamel]](uid string, updt Update[[.Name | ToCamel]]Req) ([[.Name | ToLower]] [[.Name | ToCamel]], err error) {
	logger := log.With(s.logger, "method", "Update [[.Name | ToCamel]]")

	tn := time.Now()

	update, _ := utils.ToDoc(updt)

	update = append(update, bson.E{Key: "UpdatedAt", Value: &tn})

	[[.Name | ToLower]], err = s.repo.Update[[.Name | ToCamel]](uid, update)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		i, _ := json.Marshal([[.Name | ToLower]])
		level.Info(logger).Log("Updated [[.Name | ToLower]] ", string(i))
	}

	return
}

func (s service) Delete[[.Name | ToCamel]](uid string, deletedby string) ([[.Name | ToLower]] [[.Name | ToCamel]], err error) {
	logger := log.With(s.logger, "method", "Delete [[.Name | ToCamel]]")

	type deleteStruct struct {
		DeletedBy string     `bson:"id:DeletedBy"`
		IsDeleted bool       `bson:"IsDeleted"`
		DeletedAt *time.Time `bson:"DeletedAt,omitempty"`
	}

	tn := time.Now()
	itm := deleteStruct{
		DeletedBy: deletedby,
		IsDeleted: true,
		DeletedAt: &tn,
	}

	update, _ := utils.ToDoc(itm)

	[[.Name | ToLower]], err = s.repo.Update[[.Name | ToCamel]](uid, update)
	if err != nil {
		level.Error(logger).Log("err", err)
	} else {
		i, _ := json.Marshal([[.Name | ToLower]])
		level.Info(logger).Log("Deleted [[.Name | ToLower]] ", string(i))
	}

	return
} [[ end ]]
