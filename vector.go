package toolbox

import "sync"

type Pool struct{
	pool chan interface{}
	OnEmpty func(p *Pool)
	mu sync.Mutex
}

func NewPool(size int) Pool{
	return Pool{
		pool:make(chan interface{},size),
		mu:sync.Mutex{},
		OnEmpty:nil,
	}
}

func (this *Pool) Get() interface{}{
	if this.OnEmpty==nil{
		return <-this.pool
	}else{
		for{
			select{
			case elem:= <-this.pool:
				return elem
			default:
				this.mu.Lock()
				if len(this.pool) > 0{
					this.mu.Unlock()
					break
				}
				this.OnEmpty(this)
				this.mu.Unlock()
			}
		}
	}
}

func (this *Pool) Put(elem interface{}){
	this.pool<-elem
}

