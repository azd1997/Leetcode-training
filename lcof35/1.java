// 复杂链表的复制

请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

思考：

由于是具备两个指针的复杂链表，问题的麻烦在于，深度拷贝之后，加入根据next指针，拷贝了其next后继，但是其random后继也需要生成
这会导致重复生成两套链表。
为了解决这个问题，必须将复制的链表结点与原先的链表形成映射，使用一个哈希表解决之，Java里边哈希表使用HashMap 

也就是 迭代（按next迭代） + 哈希表

```java
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
class Solution {
    // 迭代做法
    public Node copyRandomList(Node head) {
        if (head == null) return head;

        // 原链表结点 -> 新链表结点
        HashMap<Node, Node> map = new HashMap<>();
        for (Node cur = head; cur != null; cur = cur.next) {
            map.put(cur, new Node(cur.val));
        }

        // 重新构建指针链接
        for (Node cur = head; cur != null; cur = cur.next) {
            if (cur.next != null) map.get(cur).next = map.get(cur.next);
            if (cur.random != null) map.get(cur).random = map.get(cur.random);
        }
        // 返回新的链表头
        return map.get(head);        
    }


    // DFS做法
    public Node copyRandomList(Node head) {
        // 原链表结点 -> 新链表结点
        HashMap<Node, Node> map = new HashMap<>();
        return dfs(head, map);        
    }

    // dfs遍历，返回拷贝后的头结点
    private Node dfs(Node cur, HashMap<Node, Node> map) {
        if (cur == null) return null;

        // 拷贝当前结点
        if (map.containsKey(cur)) return map.get(cur);     // 哈希表已存在
        // 否则新建节点，加入到哈希表，再修改新结点的后继指针
        Node tmp = new Node(cur.val);
        map.put(cur, tmp);
        tmp.next = dfs(cur.next, map);
        tmp.random = dfs(cur.random, map); 
        // 返回深拷贝后的结点tmp
        return tmp;
    }
}