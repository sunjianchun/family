#!/Users/sunjianchun/work/environment/python/Assets/bin/python
# -*- coding: utf-8 -*-
import xlrd
import sys, json, re
from datetime import date,datetime
reload(sys)

sys.setdefaultencoding('utf8')
rd_filename = sys.argv[1]
 
def Opera_Excel(rd_filename):
    data = xlrd.open_workbook(rd_filename,formatting_info = True)
    for i in range(len(data.sheets())):
        table = data.sheets()[i]
        merge_list = table.merged_cells
        #排序,读取下标
        if len(merge_list) > 0:
            merge_list = sorted(merge_list,key=lambda ml: ml[0])
            merge_list = sorted(merge_list,key=lambda ml: ml[2])
    
            name = table.cell(1,1).value
            fellowRank = table.cell(1,4).value
            compatriotRank = table.cell(1,6).value
            phone = ""

            age = table.cell(2,2).value
            if table.cell(2,2).ctype == 2:
                age = str(int(age))
            else:
                p = re.compile(r'([0-9]{2,4})')
                m = p.match(age)
                if m:
                    age = m.group(0)
                table.cell(2,2).value
            if table.cell(2,4).ctype == 3:
                date_value = xlrd.xldate_as_tuple(table.cell(2,4).value,data.datemode)
                birthday = date(*date_value[:3]).strftime('%Y-%m-%d')
            else:
                birthday = str(table.cell(2,4).value)

            p = re.compile(r'([0-9]{2,4})[^0-9]+([0-9]{1,2})[^0-9]+([0-9]{1,2})')
            match = p.match(birthday)
            if match and len(match.groups()) == 3:
                birthday = match.group(1) + "-" + match.group(2) + "-" + match.group(3)
            else:
                p = re.compile(r'([0-9]{2,4})[^0-9]+([0-9]{1,2})')
                match = p.match(birthday)
                if match and len(match.groups()) == 2:
                    tmp = match.group(2)
                    if str(tmp) == "0":
                        tmp = "01"
                    birthday = match.group(1) + "-" + tmp + "-01"
                else:
                    p = re.compile(r'([0-9]{2,4})[^0-9]+')
                    match = p.match(birthday)
                    if match and len(match.groups()) == 1:
                        birthday = match.group(1) + "-01-01"
                    else:
                        p = re.compile(r'([0-9]{2,4})')
                        match = p.match(birthday)
                        if match and len(match.groups()) == 1:
                            birthday = match.group(1) + "-01-01"

            selfIntroduce = table.cell(3,1).value

            spouseIntroduce = table.cell(4,1).value

            list = table.cell(5,2).value.split("：")
            dad = ""
            mom = ""
            if len(list) > 2:
                dad = str(list[1]).strip('母').strip()
                mom = "".join(list[2:]).strip()


            brothers = table.cell(6,2).value
            sisters = table.cell(7,2).value
            children1 = table.cell(8,1).value
            children2 = table.cell(8,2).value
            children = ""
            if children1 != "" and children1 != None:
                children = children1

            if children2 != "" and children2 != None:
                children = children2
            
            remark = table.cell(9,2).value

            dict = {}
            dict["name"] = name 
            dict["fellowRank"] = fellowRank
            if fellowRank != "" and fellowRank != None:
                dict["fellowRank"] = str(int(fellowRank))
            dict["compatriotRank"] = compatriotRank
            if compatriotRank != "" and compatriotRank != None:
                dict["compatriotRank"] = str(int(compatriotRank))
            dict["phone"] = phone
            if phone != "" and phone != None:
                dict["phone"] = str(int(phone))
            dict["age"] = age
            dict["dad"] = dad
            dict["mom"] = mom
            if "不详" not in birthday:
                dict["birthday"] = birthday
            dict["selfIntroduce"] = selfIntroduce
            dict["spouseIntroduce"] = spouseIntroduce
            dict["brothers"] = brothers
            dict["sisters"] = sisters
            dict["children"] = children
            dict["remark"] = remark

            for k, v in dict.items():
                dict[k] = str(v)


            

if __name__ == '__main__':
    Opera_Excel(rd_filename)
    
