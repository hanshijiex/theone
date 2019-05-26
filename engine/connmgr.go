package engine

type Connmgr struct {
	 conns map[uint32]*Conn
	 count uint32
}

func (connmgr *Connmgr) AddAConn(conn *Conn) bool {
	connmgr.conns[conn.Id] = conn
	connmgr.count++
	return true
}
