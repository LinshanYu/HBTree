package hboperate

/**
hbtree insert
 */

const (
	LH = 1
	EH = 0
	RH = -1
)

type BiTNode struct{
	data int
	bf int
	lchild *BiTNode
	rchild *BiTNode
}
/*
根节点为T，插入一个e
 */
func InsertAVL(T *BiTNode, e int , taller *int )bool {

	if nil != T {
		T = new(BiTNode)
		T.data = e
		T.lchild = nil
		T.rchild = nil
		T.bf = EH
		*taller = 1
		return true
	}else {
		if e == T.data{
			*taller = 0
			return false
		}
		if e < T.data {
			if ! InsertAVL(T.lchild,e,taller) {
				return false
			}
			if *taller !=  0 { //通过子节点的插入方向 来调整父节点的状态
				switch T.bf {
				case LH: //原本T的左边的，那么插入后其父情节点的平衡因子 = 2，此时会长高为2，平衡后会变成0
					LeftBalance(T)
					*taller = 0
					T.bf = EH
					break
				case EH: //如果原本T是平衡的EH，插入后其平衡因子 = 1，长高一个
					T.bf = LH
					*taller = 1
					break
				case RH://T原本有一个右节点，现在插入为左节点 T达到平衡，左右平衡不会长高
					T.bf = EH
					*taller = 0
					break
				}
			}
		}else {
			if ! InsertAVL(T.rchild,e,taller) {
				return false
			}
			if *taller !=  0 {
				switch T.bf {
				case LH: // 原本T是LH，插入的是右节点，刚好平衡，也不长高
					T.bf = EH
					*taller = 0
					break
				case EH:
					T.bf = RH
					*taller = 1
					break
				case RH:
					RightBalance(T)
					*taller = 0
					break
				}
			}
		}
	}

	return false
}

func LeftBalance(T *BiTNode){

	l := T.lchild
	switch l.bf {
	case LH:
		Rotate_left(T)
	case RH: //先进行T的左子树tl和tl的右子树的交换位置，然后在右旋
		l.lchild = new(BiTNode)
		l.lchild.data = l.data
		l.data = l.rchild.data
		l.rchild = nil
		l.lchild.bf = EH
		l.bf = EH
		Rotate_left(T)
	}
}

func RightBalance(T *BiTNode){
	r := T.rchild
	switch r.bf {
	case LH:
		r.rchild = new(BiTNode)
		r.rchild.data = r.data
		r.data = r.lchild.data
		r.lchild = nil
		r.rchild.bf = EH
		r.bf = EH
		Rotate_right(T)
	case RH:
		Rotate_right(T)
	}

}

func Rotate_left(T *BiTNode) {

	T.rchild = new(BiTNode)
	T.rchild.data = T.data
	T.rchild.bf = EH
	T.data = T.lchild.data
	T.lchild.data = T.lchild.lchild.data
	T.lchild.lchild = nil
}

func Rotate_right(T *BiTNode) {
	T.lchild = new(BiTNode)
	T.lchild.data = T.data
	T.lchild.bf = EH
	T.data = T.rchild.data
	T.rchild.data = T.rchild.rchild.data
	T.rchild.rchild = nil

}