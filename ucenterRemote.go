package aki_remote

import (
	"context"
	"github.com/gutrse3321/aki/persit/remote"
	"github.com/gutrse3321/aki/pkg/transports/rpc"
	"github.com/pkg/errors"
	"github.com/smallnest/rpcx/client"
)

/**
 * @Author: Tomonori
 * @Date: 2020/5/29 17:09
 * @Title:
 * --- --- ---
 * @Desc:
 */
type UCenterRemote struct {
	client  *rpc.Client
	xClient client.XClient
}

func NewUCenterRemote(client *rpc.Client) (*UCenterRemote, error) {
	xClient, err := client.Connect("UCenterProvider")
	if err != nil {
		return nil, errors.Wrap(err, "connected ucenter-provider error")
	}
	return &UCenterRemote{client, xClient}, nil
}

func (u *UCenterRemote) GetUserBaseInfo(args *string) (*remote.Remote, error) {
	codeRemote := &remote.Remote{}
	err := u.xClient.Call(context.Background(), "GetUserBaseInfo", args, codeRemote)
	if err != nil {
		return nil, err
	}

	return codeRemote, nil
}
