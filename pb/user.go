package pb

import "strconv"

func (x *User) IDStr() string {
	return strconv.FormatUint(x.GetId(), 10)
}
