// Author: xufei
// Date: 2019-12-23

package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

var (
	// 时间偏移量，从2010-11-04 01:42:54
	Epoch int64 = 1288834974657

	// NodeBits 节点ID 使用 10位 = 2^10 = 1024 个节点
	NodeBits uint8 = 10

	// StepBits 顺序ID 使用 12位 2^12=4096, 支持每毫秒生产 4096 个。
	StepBits uint8 = 12
)

type Node struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64 // 时间信息
	node  int64 // 节点信息
	step  int64 // 顺序信息

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8 // 时间位移位置
	nodeShift uint8 // 节点位移位置
}

type ID int64

func NewNode(node int64) (*Node, error) {
	n := Node{}
	n.node = node
	n.nodeMax = -1 ^ (-1 << NodeBits)
	n.nodeMask = n.nodeMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.timeShift = NodeBits + StepBits
	n.nodeShift = StepBits

	if n.node < 0 || n.node > n.nodeMax {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(n.nodeMax, 10))
	}

	var curTime = time.Now()
	n.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &n, nil
}

func (n *Node) NewID() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Since(n.epoch).Nanoseconds() / 1000000
	// 如果是同一时间生成的，则进行毫秒内序列
	if now == n.time {
		n.step = (n.step + 1) & n.stepMask
		if n.step == 0 {
			// 顺序ID在同一毫秒被用完，阻塞到下一个时间节点
			for now <= n.time {
				now = time.Since(n.epoch).Nanoseconds() / 1000000
			}
		}
	} else {
		// 时间戳改变，毫秒内序列重置
		n.step = 0
	}

	// 更新时间戳
	n.time = now

	// now 的 n.timeShift 次方 OR n.node 的 n.nodeShift 次方 OR n.step
	r := ID((now)<<n.timeShift | (n.node << n.nodeShift) | n.step)

	return r
}

func (f ID) Int64() int64 {
	return int64(f)
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}
