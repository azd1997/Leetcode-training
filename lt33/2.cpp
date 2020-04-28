#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // 二分查找，根据旋转点在哪边来进行二分
    // 一边有旋转点，另一边有序
    int search(vector<int>& nums, int target) {
        int n = nums.size();
        if (n == 0) return -1;

        int l = 0, r = n-1, mid;
        while (l <= r) {
            mid = (r - l) / 2 + l;

            if (nums[mid] == target) return mid;

            if (nums[mid] >= nums[l]) {  // 左边有序。 必须是 >= 不能是 >
                if (target >= nums[l] && target <= nums[mid]) { // target位于左边有序区间
                    return bs(nums, l, mid-1, target);
                } else {
                    l = mid + 1;
                }
            } else {    // 右边有序
                if (target >= nums[mid] && target <= nums[r]) { // target位于左边有序区间
                    return bs(nums, mid+1, r, target);
                } else {
                    r = mid - 1;
                }
            }
        }
        return -1;
    }

    // 普通二分查找
    int bs(vector<int>& nums, int start, int end, int target) {
        int l = start, r = end, mid;
        while (l <= r) {
            mid = (r - l) / 2 + l;
            if (nums[mid] > target) {
                r = mid - 1; 
            } else if (nums[mid] < target) {
                l = mid + 1;
            } else {    // ==
                return mid;
            }
        }
        return -1;
    }
};

int main() {
	Solution sol;
	vector<int> nums{4,5,6,7,1,2,3};
	int ans = sol.search(nums, 2);
	cout << ans;
} 
