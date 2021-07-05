#!/usr/bin/env python
# -*- coding: utf-8 -*-
import jieba.posseg as pseg
from datetime import datetime, timedelta  
from datetime import date 
import re
from flask import Flask, flash, request, redirect, render_template
#import urllib.request
import os
import ast
# coding=utf-8
from flask import Flask,render_template,request,url_for 
#import simplejson as json
import  json
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import *
from loguru import logger
# Column, String, Integer, Date,Numeric,and_,in_
# from models import techniqueanalysis
#from sqlalchemy.ext.declarative import declarative_base
#from flaskjsontools import JsonSerializableBase
import pymysql
pymysql.install_as_MySQLdb()
engine = create_engine('mysql://root:123456@localhost:3306/finance?host=127.0.0.1&charset=utf8mb4')
engine1 = create_engine('mysql://root:123456@localhost:3306/dreamteam_db?host=127.0.0.1&charset=utf8mb4')
Session = sessionmaker(bind=engine)
Session1 = sessionmaker(bind=engine1)
Base = declarative_base()

import decimal

class DecimalEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, decimal.Decimal):
            return float(o)
        return super(DecimalEncoder, self).default(o)








#获取间隔时间
def gen_dates(b_date, days):
    day = timedelta(days=1)
    for i in range(days):
        yield b_date + day*i


def get_date_list(start=None, end=None):
    """
    这里传入时间都可以进行改写
    获取日期列表
    :param start: 开始日期
    :param end: 结束日期
    :return:
    """
    #这里写代码我自已进行解析时间
    '%y%m%d'
    
    start = datetime.strptime(start, '%y%m%d')
    end = datetime.strptime(end, '%y%m%d')
    
    # if start is None:
    #     start = datetime.strptime("2000-01-01", "%Y-%m-%d")
    # if end is None:
    #     end = datetime.now()
    data = []
    #这里传入的时间参数可以进行改变
    for d in gen_dates(start, (end+timedelta(days=1)-start).days):
        data.append(datetime.strftime(d,'%y%m%d'))
    return data



#获取当前这一周时间
def  weektime_current():
    #date_obj = datetime.strptime(date_str, '%Y-%m-%d') 
    from datetime import date  
    date_obj  = date.today() 
    #接下来的两部分是用来获取date类型的时间格式
    start_of_week = date_obj - timedelta(days=date_obj.weekday())  # Monday 
    end_of_week = start_of_week + timedelta(days=6)  # Sunday 
    start_of_week = datetime.strftime(start_of_week, '%y%m%d') 
    end_of_week =datetime.strftime(end_of_week , '%y%m%d') 
    return get_date_list(start_of_week,end_of_week)


#获取这个月所有时间
def  monthtime_current():
    #date_obj = datetime.strptime(date_str, '%Y-%m-%d') 
    # from datetime import date,datetime 
    import datetime, calendar
    date_obj  = date.today() 
    #接下来的两部分是用来获取date类型的时间格式
    year = date_obj.year
    month = date_obj.month
    num_days = calendar.monthrange(year, month)[1]
    days = [datetime.datetime.strftime(datetime.date(year, month, day), '%y%m%d')  for day in range(1, num_days+1)]
    return days

#获取上个月所有的时间
def  lastmonthtime_current():
    #date_obj = datetime.strptime(date_str, '%Y-%m-%d') 
    # from datetime import date,datetime 
    import datetime, calendar
    date_obj  = date.today() 
    #接下来的两部分是用来获取date类型的时间格式
    year = date_obj.year
    month = date_obj.month -1
    if date_obj.month == 1:
        year = year -1
        month = 12
    num_days = calendar.monthrange(year, month)[1]
    days = [datetime.datetime.strftime(datetime.date(year, month, day), '%y%m%d')  for day in range(1, num_days+1)]
    return days


#获取上个月所有的时间
def  thisyear_current():
    #date_obj = datetime.strptime(date_str, '%Y-%m-%d') 
    # from datetime import date,datetime 
    import datetime, calendar
    date_obj  = date.today() 
    #接下来的两部分是用来获取date类型的时间格式
    year = date_obj.year
    start = str(year)[2:]+"0101"
    end = str(year)[2:]+"1231"
    print(year)
    print(start)
    print(end)
    days = get_date_list(start,end)
    return days


#获取今天的日期
from datetime import date   
from datetime import datetime  
def gettoday():
    date_obj  = date.today()
    today = datetime.strftime(date_obj, '%y%m%d') 
    return today

   # 这里是将评价孙法与消费结果进行捆绑，建立良好的惩罚与训练机制
def rewardorpunishment(starttime="200101",times=10,email="yang756260386@gmail.com"):
    #获取算指定日期以后的日期
    computesdates = get_date_list("200101",gettoday())
    print("------yangming is here--------")
    print(computesdates)
    #获取指定日期以后的消费总额
    #获取指定日期以后的评价分数
    # email = request.headers['email']
    #date = content['date']
    session = Session()
    email = "yang756260386@gmail.com"
    #date = content['date']
    # session = Session()
    lastmonth = lastmonthtime_current()
    thismonth=  monthtime_current()
    date = lastmonth + thismonth
    print(date)
    #这里使用级别链接的方式重新设计
    #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
    #使用级别链接的方式
    all = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.date.in_(date)).all()
    date = computesdates
    print(date)
    #这里使用级别链接的方式重新设计
    #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
    #使用级别链接的方式
    print(email)
    all = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.date.in_(date)).all()
    #写消费统计部分
    allcost = sum([float(row.fee) for row in all if row.direction == "buy"])
    allincome = sum([float(row.fee) for row in all if row.direction == "sell"])
    print("=======yangming=======")
    print(allcost)
    #获取评价算法的部分
    # 这里新开了一个链接不知道是否有影响
    session1 = Session1()
    #这里使用级别链接的方式重新设计
    #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
    #使用级别链接的方式
    all = session1.query(Reviewofdays).filter(Reviewofdays.email == email).filter(Reviewofdays.date.in_(date)).all()
    print("长度是%s"%(len(all)))
    session.close()
    session1.close()
    all = [k.details for k in all]
    print(all[0])
    print(ast.literal_eval(all[0])["totalscore"])
    bc = [ast.literal_eval(k)["totalscore"] for k in all if len(k)>0]
    print(len(bc))
    print(bc)
    sum_algo = sum(bc)
    print("=================aaa===========")
    print(sum_algo)
    #乘一个系数
    sumalgomultiple = sum_algo*10
    #返回值【负债，可使用资金，奖励剩余，还有】
    """
    计算总的投入时间
    计算每天睡眠时间
    计算休息时间
    总24*60 -8*60
    计算应该的休息时间=总的分数*30
    """




    left = sumalgomultiple - allcost
    return {"left":left,"thisyear":allcost}
  












def getmoney(sentence):
    #这里其实也要检测语言
    words =pseg.cut(sentence)
    money = []
    for w in words: 
        if w.flag=="m":
            money.append(float(w.word))
            print(w.word,w.flag,type(w.flag)) 
    return money



class Reviewofdays(Base):
    __tablename__ = 'reviewofdays'
    id = Column(Integer, primary_key=True)
    # scores=Column(String(32))
    email = Column(String(32))
    date = Column(String(32))
    details = Column(String(32))
    scores=Column('scores', Numeric)
    def as_dict(self):
        return {c.name: getattr(self, c.name) for c in self.__table__.columns}
    def __init__(self,email,detail,date,scores):
        # self.direction= direction
        self.scores = scores
        self.date = date
        self.details = details
        self.email = email





class BalanceSheet(Base):
    __tablename__ = 'banlancesheet'
    id = Column(Integer, primary_key=True)
    # scores=Column(String(32))
    email = Column(String(32))
    date = Column(String(32))
    asset=Column('asset', Numeric)
    debt=Column('debet', Numeric) 
    def as_dict(self):
        return {c.name: getattr(self, c.name) for c in self.__table__.columns}
    def __init__(self,email,date,asset,debt):
        # self.direction= direction
        self.date = date
        self.asset = asset
        self.debt = debt
        self.email = email






class Accounting(Base):
    __tablename__ = 'accounting'
    id = Column(Integer, primary_key=True)
    direction=Column(String(32))
    record = Column(String(32))
    email = Column(String(32))
    date = Column(String(32))
    fee=Column('fee', Numeric)
    def as_dict(self):
        return {c.name: getattr(self, c.name) for c in self.__table__.columns}
    def __init__(self,direction,record,fee,date,email):
        self.direction= direction
        self.record = record
        self.date = date
        self.fee = fee
        self.email = email
        








# 2 - generate database schema
Base.metadata.create_all(engine)




app = Flask(__name__)

@app.route('/hello')
def hello_world():
   return 'Hello World'


# 获取奖励剩余的人民币和时间，时间接库还没实现
@app.route('/finance/getrewardleft',methods=["POST","GET","PUT"])
def rewardfun():
    try:
        finance = rewardorpunishment()
        finance={"left":finance["left"],"available":40,"budget":90,"thisyear":finance["thisyear"],"code":200}
        return json.dumps(finance)
    except:
        logger.exception("获取回馈数据异常")
        return json.dumps({'code':403,'info':"server eror"})






@app.route('/finance/banlance',methods=["POST","GET","PUT"])
def staticticsforbanlancetable():
    try:
        # days = request.args.get('days')
        # email = request.headers['email']
        email= "yang756260386@gmail.com"
        #date = content['date']
        session = Session()
        #这里使用级别链接的方式重新设计
        #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
        #使用级别链接的方式
        all = session.query(BalanceSheet).filter(BalanceSheet.email == email).all()
        #写消费统计部分
        session.close()
        allcost = sum([float(row.debt) for row in all ])
        allincome = sum([float(row.asset) for row in all ])
        result = {"debt":allcost,"asset":allincome,"code":200}
        return json.dumps(result)
    except:
        logger.exception("获取统计数据异常")
        return json.dumps({"code":403,"info":"server error"})







@app.route('/finance/statistics',methods=["POST","GET","PUT"])
def staticticsformoney():
    try:
        days = request.args.get('days')
        print(days)
        date = datetime.strftime(datetime.now() - timedelta(1), '%y%m%d')
        if days == "-1":
            print("i am here")
            date = [datetime.strftime(datetime.now() - timedelta(1), '%y%m%d')]
        if days == "0":
            print("------")
            date = [datetime.strftime(datetime.now(), '%y%m%d')] 
        if days =="7":
            date = weektime_current()
        if days =="31":
            date = monthtime_current()
        if days =="-31":
            date = lastmonthtime_current()
        if days =="365":
            date = thisyear_current()
        email = request.headers['email']
        #date = content['date']
        session = Session()
        print(date)
        #这里使用级别链接的方式重新设计
        #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
        #使用级别链接的方式
        all = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.date.in_(date)).all()
        #写消费统计部分
        session.close()
        allcost = sum([float(row.fee) for row in all if row.direction == "buy"])
        allincome = sum([float(row.fee) for row in all if row.direction == "sell"])
        for k in all:
            print(k)
        # return "ok"
        result = {"cost":allcost,"income":allincome,"code":200}
        return json.dumps(result)
    except:
        logger.exception("获取统计数据异常")
        return json.dumps({"code":403,"info":"server error"})


#这里是为了获取详详细的消费情况便于展示，这里要使用细节的cookie
@app.route('/finance/getfeesdetail',methods=["POST","GET","PUT"])
def getfeesdetail():
    try:
        # email = request.headers['email']
        #email = request.cookies["email"]
        print(request.cookies)
        # 这里需要读取request部分的client！！！
        # 获取到实际的email
        email = "yang756260386@gmail.com"
        #date = content['date']
        session = Session()
        lastmonth = lastmonthtime_current()
        thismonth=  monthtime_current()
        date = lastmonth + thismonth
        print(date)
        #这里使用级别链接的方式重新设计
        #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
        #使用级别链接的方式
        all = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.date.in_(date)).all()
        #写消费统计部分
        #建立一个大的数组，返回的数据格式应该是
        #allrecords = {"data":[{"name":"191102","records":[row]},{},{},{},{}]}
        from collections import defaultdict
        alldays_records = defaultdict(list)
        for k in all:
            print(k.id)
            alldays_records[k.date].append(rowtodict(k))
        
        result = []
        for k in date:
            if len(alldays_records[k])>0:
                result.append({"date":k,"allrecordsinaday":alldays_records[k]})
        
        # allcost = sum([float(row.fee) for row in all if row.direction == "buy"])
        
        # allcost = sum([float(row.fee) for row in all if row.direction == "buy"])
        # allincome = sum([float(row.fee) for row in all if row.direction == "sell"])
        resultfromserver = {"allfees":result,"code":200}
        return json.dumps(resultfromserver)
    except:
        logger.exception("获取详细费用异常")
        return json.dumps({"info":"server error","code":403})






@app.route('/finance/uploadfees',methods=["POST","GET","PUT"])
def createfees():
    try:
        print("----i开始---------")
        print(dict(request.headers))
        print("----后来---------")
        email = request.headers['email']
        content = request.json
        record = content['inbox']
        direction = content['direction']
        date = content['date']
        feefromclient = 0.0
        print(request.headers)
        if request.headers['client'] == "iosnotsiri":
            feefromclient = float(content['fee'])
            print(record.encode("utf8"))
        else:
            #record = "我们吃了10块钱的晚饭"
            fees = getmoney(record)
            print(fees)
            #这里因该给出明确的告警信息！否则下面这句话会出问题
            if len(fees) >1 or len(fees) == 0:
                return json.dumps({"info":"请不要在消费中包含两个数字，我不能帮你识别","status":"fail"})
            #接下来准备写入到Q数据库
            feefromclient = float(fees[0])
            print(feefromclient) 
        session = Session()
        oneday = Accounting(direction,record,feefromclient,date,email)
        session.add(oneday)
        session.commit()
        session.close()
        return json.dumps({"info":"记账成功","status":"ok","code":200})
    except:
        logger.exception("上传费用异常")
        return json.dumps({"info":"server error","code":403})


@app.route('/finance/updatefees',methods=["POST","GET","PUT"])
def updatefees():
    try:
        email = "yang756260386@gmail.com"
        # print(request.method)
        print("----i开始---------")
        print(dict(request.headers))
        print("----后来---------")
        content = request.json
        print(request)
        print(content)
        print("----i am here---------")
        record = content['record']
        direction = content['direction']
        date = content['date']
        fee = content['fee']
        id = content['id']
        
        print(record)
        print("-----i am here---------")
        #record = "我们吃了10块钱的晚饭"
        session = Session()
        recordfromdb = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.id==id).first()
        recordfromdb.record = record
        recordfromdb.direction = direction
        recordfromdb.date = date
        recordfromdb.fee = fee
        session.commit()
        session.close()
        return json.dumps({"info":"记账成功","status":"ok","code":200})
    except:
        logger.exception("更新费用异常")
        return json.dumps({"code":405,"info":"server error"})

import time

@app.route('/finance/updatebalancetable',methods=["POST","GET","PUT"])
def updatebanlancetable():
    try:
        email = "yang756260386@gmail.com"
        # print(request.method)
        print("----i开始---------")
        print(dict(request.headers))
        print("----后来---------")
        content = request.json
        print(request)
        print(content)
        print("----i am here---------")
        debt = content['debt']
        asset = content['asset']
        email = content['email']

        
        # print(record)
        print("-----i am here---------")
        #record = "我们吃了10块钱的晚饭"
        session = Session()
        recordfromdb = session.query(BalanceSheet).filter(BalanceSheet.email == email).first()
        print(recordfromdb)
        # time.sleep(10)
        #TODO 这里要注意更新日期的保存工作
        if recordfromdb ==  None:
            one= BalanceSheet(email,'2021/06/08',asset,debt)
            session.add(one)   
            pass
        else:
            recordfromdb.debt = debt
            recordfromdb.asset = asset
        
        session.commit()
        session.close()
        return json.dumps({"info":"记账成功","status":"ok","code":200})
    except:
        logger.exception("更新费用异常")
        return json.dumps({"code":405,"info":"server error"})






    

def rowtodict(row):
    dictforsinglerow ={}
    dictforsinglerow["email"] = row.email
    dictforsinglerow["fee"] = float(row.fee)
    dictforsinglerow["direction"]= row.direction
    dictforsinglerow["date"] = row.date
    dictforsinglerow["record"] = row.record
    dictforsinglerow["id"] = row.id
    return dictforsinglerow






if __name__ == '__main__':
    logger.add("./logging/file_1.log", rotation="50 MB",backtrace=True, diagnose=True,colorize=True) 
    app.run('0.0.0.0',threaded=True,port = 6000,debug=True)




        

