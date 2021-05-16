package main

import (
	"about-data-stores/internal/infra/dynamodb"
	"about-data-stores/internal/infra/elasticache"
	"about-data-stores/internal/infra/neptune"
	"about-data-stores/internal/infra/rds"
	"fmt"
)

func main() {
	// RDS
	rdsHandler, _ := rds.NewRDBHandler()
	db := rdsHandler.DbService
	db.Set("gorm:table_options", "ENGINE = InnoDB")

	rdsRes, _ := db.Raw("SELECT name FROM users WHERE user_id = 'dummy-user-1'").Rows()
	defer db.Close()
	fmt.Print(rdsRes)

	// DynamoDB
	dynamodbHandler := dynamodb.NewDynamodbHandler()

	dynamoRes, _ := dynamodbHandler.Get("about-data-store", "user_id", "dummy-user-1")
	fmt.Print(dynamoRes)

	// Neptune
	neptuneHandler, _ := neptune.NewNeptuneHandler()
	neptuneRes, _ := neptuneHandler.Execute("g.V().has('name','dummy-user-1').out('favorite')", nil, nil)
	fmt.Print(neptuneRes)

	// ElastiCache
	elasticacheHandler, _ := elasticache.NewElastiCacheHandler()
	elasticacheRes, _ := elasticacheHandler.Get("dummy-user-1")
	fmt.Print(elasticacheRes)
}
