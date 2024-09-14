package graph

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/nsym-m/go-graphql/internal/graph/model"
	"github.com/nsym-m/go-graphql/internal/services"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: srv}

	return &Loaders{
		// dataloader.Loader[string, *model.User]構造体型をセットするために、
		// dataloader.NewBatchedLoader関数を呼び出す
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, ids []string) []*dataloader.Result[*model.User] {
	res := make([]*dataloader.Result[*model.User], len(ids))
	for i := range res {
		res[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}
	if len(ids) == 0 {
		return res
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[string]int, len(ids))
	for i, id := range ids {
		indexs[id] = i
	}

	users, err := u.Srv.ListUserByIDs(ctx, ids)
	if err != nil {
		return res
	}

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for _, user := range users {
		var r *dataloader.Result[*model.User]
		// Note: DB取得でエラーの場合全部取れないからusersがnilでそもそもこのループ回らないのでは？
		if err != nil {
			r = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			r = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		// Note: 何かのミスでuserがゼロ値だったらmapから見つからず0が取得され、res[0]が不正に更新される
		res[indexs[user.ID]] = r
	}

	return res
}
