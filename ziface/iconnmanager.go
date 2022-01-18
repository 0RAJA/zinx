package ziface

type IConnManager interface {
	Add(connection IConnection)             //新增连接
	Remove(connection IConnection)          //删除连接
	Get(connID uint32) (IConnection, error) //获取连接
	Len() int                               //连接数
	Clear()                                 //清空连接
}
