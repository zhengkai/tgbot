package pb

import "strconv"

func (x *Chat) IDStr() string {
	return strconv.FormatInt(x.GetId(), 10)
}
