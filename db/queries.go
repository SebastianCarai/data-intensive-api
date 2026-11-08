package db

import (
	"errors"
	"fmt"
)

func GetFeedQuery(id int) (string, error){
	idType := fmt.Sprintf("%T", id)
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
limit 20;`, id)
	return query, nil
}