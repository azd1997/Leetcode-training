/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    //按层遍历二叉树，用一个队列保存每一层的节点
    public String serialize(TreeNode root) {
        if (root == null)
            return "";
        StringBuilder builder = new StringBuilder();
        Deque<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            int count = queue.size();
            while (count-- > 0) {
                TreeNode node = queue.removeFirst();
                //这个优化点很重要，当遍历到null时，不要往队列写入null子节点，减少生成的字符串大小
                if (node != null) {
                    queue.add(node.left);
                    queue.add(node.right);
                    builder.append(",").append(node.val);
                } else {
                    builder.append(",null");
                }
            }
        }
        //删除第一个逗号
        builder.deleteCharAt(0);
        //下面的代码通过正则匹配字符串末尾的null,这些null可以省略
        //Pattern pattern = Pattern.compile("[,null]+$");
        //Matcher matcher = pattern.matcher(builder.toString());
        //return matcher.replaceFirst("");
        return builder.toString();
    }


    //反序列化的时候，用两个指针，第一个指向父节点，第二个指向子节点
    //由于字符串是按层遍历，同一层的节点相邻，子节点排在父节点后面
    //用一个队列记录按层遍历的树节点，不包含null,队列头节点为整个树的根节点
    public TreeNode deserialize(String data) {
        if (data == null || data.length() == 0)
            return null;
        String[] values = data.split(",");
        List<TreeNode> list = new LinkedList<>();
        TreeNode head = createNode(values[0]);
        list.add(head);
        int rootIndex = 0;
        int valueIndex = 1;
        while (rootIndex < list.size()) {
            TreeNode root = list.get(rootIndex++);
            if (valueIndex < values.length){
                root.left = createNode(values[valueIndex++]);
                root.right = createNode(values[valueIndex++]);
            }
            if (root.left != null)
                list.add(root.left);
            if (root.right != null)
                list.add(root.right);
        }
        return head;
    }

    private TreeNode createNode(String str) {
        if (str == null) {
            return null;
        }
        if (str.equalsIgnoreCase("null")) {
            return null;
        } else {
            int integer = Integer.parseInt(str);
            return new TreeNode(integer);
        }
    }

}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));