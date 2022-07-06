from typing import Optional


def addTwoNumbers(self, l1: Optional[list], l2: Optional[list]) -> Optional[list]:
    # 判断长度补0
    len_dis = len(l1) - len(l2)
    short_li = [0] * abs(len_dis)
    short_li.extend(l1) if len_dis <= 0 else short_li.extend(l2)
    long_li = l2 if len_dis <= 0 else l1
    new_li = []

    a = 0
    for i in range(len(short_li) - 1, -1, -1):
        # 如果进位
        if a == 1:
            val = short_li[i] + long_li[i] + 1
        else:
            val = short_li[i] + long_li[i]
        new_li.append(str(val)[-1])
        if val >= 10:
            a = 1
        else:
            a = 0

        if i == 0 and a == 1:
            new_li.append(1)

    return new_li