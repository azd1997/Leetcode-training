/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/

// 中序遍历
class Solution {

    // 全局变量
    private Node pre;
    private Node head, tail;

    public Node treeToDoublyList(Node root) {
        // 特殊情况
        if (root == null) return null;

        // 中序遍历，将中间处理好，还剩下head.left和tail.right没处理
        inorder(root);

        head.left = tail;
        tail.right = head;

        return head;
    }

    // 中序遍历，对root子树作中序遍历
    private void inorder(Node root) {
        if (root == null) return;   // 递归终止：到了叶节点以下

        // 左子树遍历
        inorder(root.left);

        // 处理当前

        root.left = pre;     // 先指向pre，接下来要修改pre.right,需要分情况
        if (pre == null) {
            // 那说明是head
            head = root;
        } else {
            // 说明不是head
            pre.right = root;
        }

        pre = root;     // pre后移
        tail = root;    // tail也后移。 （其实pre,tail只要一个就可以，这里只是方便阅读）

        // 右子树遍历
        inorder(root.right);
    } 
}