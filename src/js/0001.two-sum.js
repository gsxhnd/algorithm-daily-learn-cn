/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
const twoSum = function (nums, target) {
    const a = [];
    for (let i = 0; i < nums.length; i++) {
        const tmp = target - nums[i];
        if (a[tmp] !== undefined) return [a[tmp], i];
        a[nums[i]] = i;
    }
};


const main = function () {
    let nums = [2, 7, 11, 15], target = 9
    const resultAnswer = [0, 1]
    let result = twoSum(nums,target)
    if (result !== resultAnswer) {
        console.log(true)
    } else {
        console.log(false)
    }
};

main();
