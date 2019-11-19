package lt278

// 第一个错误的版本

//
///* The isBadVersion API is defined in the parent class VersionControl.
//      boolean isBadVersion(int version); */
//
//public class Solution extends VersionControl {
//    public int firstBadVersion(int n) {
//
//        // 迭代二分
//        int l = 1, r = n;
//        int mid;
//        while (l<r) {
//            mid = l + (r-l)/2;
//            if (isBadVersion(mid)) {
//                r = mid;
//            } else {
//                l = mid+1;  // 保证左边界一定是错误版本
//            }
//        }
//
//        return l;
//    }
//
//}
