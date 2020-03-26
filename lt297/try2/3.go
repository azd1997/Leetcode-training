package lt297

// 二叉树的序列化与反序列化

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {

}

// Serializes a tree to a single string.
// 序列化的难点在于如何控制空指针是该打印null呢还是不打
// 一个粗暴易懂的做法是，将所有空指针（最末一层以及之前的）都打成null
// 再向前切割字符串，把多余的null去掉
// 而且遍历的顺序其实按照层序遍历会更好做。普通的前序遍历还要想办法直到当前层是不是最后一层
func (this *Codec) serialize(root *TreeNode) string {

}

func serhelp(root *TreeNode, str *string) {

}

// Deserializes your encoded data to tree.
// 序列化是层序遍历，反序列化根据结点数量，先将字符串补齐到满二叉树状态
// 这样就保证了树的均衡，同时没有改变原本字符串所表示的树的形态
// 其实也不用这样做
// 虽然最后一层未必是满的，但是前面肯定是满的，依赖可以按照每层结点数来按层还原
func (this *Codec) deserialize(data string) *TreeNode {

}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
