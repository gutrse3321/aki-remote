package aki_remote

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/5/29 17:14
 * @Title:
 * --- --- ---
 * @Desc:
 */
var WireSet = wire.NewSet(
	NewUCenterRemote,
)
