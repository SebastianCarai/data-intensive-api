package db

import (
	"errors"
	"fmt"
)

func GetFeedQuery(userId int) (string, error){
	idType := fmt.Sprintf("%T", userId)
	if idType != "int"{
		return "", errors.New("Invalid id")
	}

	query := fmt.Sprintf(`select p.id, p.author, p.content, p.likes_count, p.comments_count
from follows f
cross join lateral(
    select *
    from posts p
    where p.author = f.followee_id
    order by p.created_at desc, p.id desc
    limit 20
) p
where f.follower_id = %v
order by p.created_at desc, p.id desc
limit 20;`, userId)
	return query, nil
}


func GetPostDetailsQuery(postId int) (string, error){
	idType := fmt.Sprintf("%T", postId)
	if idType != "int"{
		return "", errors.New("Invalid post id")
	}

	query := fmt.Sprintf(`select p.id, p.author, p.content, p.likes_count, p.comments_count 
	from posts p 
	where p.id = %v`, postId)

	return query, nil
}