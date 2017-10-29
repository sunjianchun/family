#!/Users/sunjianchun/work/environment/python/Assets/bin/python
# -*- coding: utf-8 -*-
import xlrd
import sys, json

reload(sys)

sys.setdefaultencoding('utf8')
#rd_filename = "/Users/sunjianchun/ddd.xls"
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
            phone = table.cell(1,9).value

            age = table.cell(2,2).value
            birthday = table.cell(2,4).value
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
            children = table.cell(8,1).value
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
            dict["age"] = str(int(age))
            if age != "" and age != None:    
                dict["age"] = str(int(age))
            dict["dad"] = dad
            dict["mom"] = mom
            dict["birthday"] = birthday
            dict["selfIntroduce"] = selfIntroduce
            dict["spouseIntroduce"] = spouseIntroduce
            dict["brothers"] = brothers
            dict["sisters"] = sisters
            dict["children"] = children
            dict["remark"] = remark

            for k, v in dict.items():
                dict[k] = str(v)
            print json.dumps(dict, sort_keys=True, indent=4, separators=(',', ': '))

            

if __name__ == '__main__':
    Opera_Excel(rd_filename)
    
