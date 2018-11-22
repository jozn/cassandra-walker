package main

import (
	"github.com/gocql/gocql"
	"github.com/jozn/cassandra-walker/samples/twitter/xc"
)

func main() {
	// create cassandra session
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "twitter"
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	// Create
	twitt1 := xc.Tweet{
		Body:       "Hello World",
		CreateTime: 1566000000,
		TweetId:    "1",
		UserId:     1,
	}

	err := twitt1.Save(session)

	// Delete one object
	twitt1.Delete(session)

	//////////////// For Selector
	tweets, err := xc.NewTweet_Selector().UserId_Eq(1).Limit(5).GetRows(session) // returns and array of tweets ( []*tweet ,err )

	tweet, err := xc.NewTweet_Selector().UserId_Eq(1).Limit(5).GetRows(session) // returns a single tweet ( *tweet ,err )

	//can use clustering columns too
	tweets, err = xc.NewTweet_Selector().UserId_Eq(1).And_TweetId_In("1", "25", "68").GetRows(session)

	//can select just some columns, it will returns *[]Tweet, with just selected columns sets
	tweets, err = xc.NewTweet_Selector().Select_UserId().Select_Body().UserId_Eq(1).And_TweetId_In("1", "25", "68").Limit(12).GetRows(session)

	//for when need to use filtering
	tweets, err = xc.NewTweet_Selector().UserId_LT_Filtering(100).Limit(10).AllowFiltering().GetRows(session)

	//////////////// For Updater
	err = xc.NewTweet_Updater().
		Body("new tweet text").UserId_Eq(1).And_TweetId_In("1", "2", "3").Update(session)

	//////////////// For Deleter
	err = xc.NewTweet_Deleter().UserId_Eq(1).And_TweetId_In("1", "2", "3").Delete(session)
	err = xc.NewTweet_Deleter().UserId_Eq(1).Delete(session)

	_ = err
	_ = tweets
	_ = tweet
}

/* log output - this is produced CQL queries to cassandra:

2018/09/18 22:35:54 CQL:  [insert into tweeter.tweet (body,create_time,tweet_id,user_id) values (?,?,?,?)  [Hello World 1566000000 1 1]]
2018/09/18 22:35:54 CQL:  [DELETE FROM tweeter.tweet WHERE  user_id = ? And tweet_id = ?  [1 1]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM tweeter.tweet WHERE  user_id = ?  LIMIT 5 [1]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM tweeter.tweet WHERE  user_id = ?  LIMIT 5 [1]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM tweeter.tweet WHERE  user_id = ? And tweet_id IN (?,?,?)  [1 1 25 68]]
2018/09/18 22:35:54 CQL:  [SELECT user_id, body FROM tweeter.tweet WHERE  user_id = ? And tweet_id IN (?,?,?)  LIMIT 12 [1 1 25 68]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM tweeter.tweet WHERE  user_id < ?  LIMIT 10  ALLOW FILTERING [100]]
2018/09/18 22:35:54 CQL:  [UPDATE tweeter.tweet SET body = ?  WHERE  user_id = ? And tweet_id IN (?,?,?)  [new tweet text 1 1 2 3]]
2018/09/18 22:35:54 CQL:  [DELETE FROM tweeter.tweet WHERE  user_id = ? And tweet_id IN (?,?,?)  [1 1 2 3]]
2018/09/18 22:35:54 CQL:  [DELETE FROM tweeter.tweet WHERE  user_id = ?  [1]]

*/
