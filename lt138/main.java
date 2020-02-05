/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/

class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}

// 回溯解法
class Solution {

    // 哈希表记录旧节点到新节点的映射
    HashMap<Node, Node> visited = new HashMap<Node, Node>();

    public Node copyRandomList(Node head) {
        // 递归停止条件
        if (head == null) return null;

        // 如果已经拷贝过(访问过)该节点则直接将该节点的那一份拷贝返回
        if (this.visited.containsKey(head)) {
            return this.visited.get(head);
        }

        // 创造一个新节点用于拷贝旧节点
        Node node = new Node(head.val);

        // 将 旧->新 节点映射存入 visited
        this.visited.put(head, node);

        // 递归调用
        node.next = copyRandomList(head.next);
        node.random = copyRandomList(head.random);

        return node;
    }
}