#!/usr/bin/env python
# -*- coding: utf-8 -*-
import jieba.posseg as pseg
from datetime import datetime, timedelta  
from datetime import date 
# coding=utf-8
from flask import Flask,render_template,request,url_for 
#import simplejson as json
import  json
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import *
# Column, String, Integer, Date,Numeric,and_,in_
# from models import techniqueanalysis
#from sqlalchemy.ext.declarative import declarative_base
#from flaskjsontools import JsonSerializableBase
engine = create_engine('mysql://root:123456@localhost:3306/finance?host=127.0.0.1&charset=utf8mb4')
Session = sessionmaker(bind=engine)
Base = declarative_base()


import decimal

class DecimalEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, decimal.Decimal):
            return float(o)
        return super(DecimalEncoder, self).default(o)





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



def  weektime_current():
    #date_obj = datetime.strptime(date_str, '%Y-%m-%d') 
    from datetime import date  
    date_obj  = date.today() 
    #接下来的两部分是用来获取date类型的时间格式
    start_of_week = date_obj - timedelta(days=date_obj.weekday())  # Monday 
    end_of_week = start_of_week + timedelta(days=6)  # Sunday 
    start_of_week =  date = datetime.strftime(start_of_week, '%y%m%d') 
    end_of_week =  date = datetime.strftime(end_of_week , '%y%m%d') 
    return get_date_list(start_of_week,end_of_week)






def getmoney(sentence):
    #这里其实也要检测语言
    words =pseg.cut(sentence)
    money = []
    for w in words: 
        if w.flag=="m":
            money.append(float(w.word))
            print(w.word,w.flag,type(w.flag)) 
    return money









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






@app.route('/finance/statistics',methods=["POST","GET","PUT"])
def staticticsformoney():
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
    email = request.headers['email']
    #date = content['date']
    session = Session()
    print(date)
    #这里使用级别链接的方式重新设计
    #all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
    #使用级别链接的方式
    all = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.date.in_(date)).all()
    #写消费统计部分
    allcost = sum([float(row.fee) for row in all if row.direction == "buy"])
    allincome = sum([float(row.fee) for row in all if row.direction == "sell"])
    for k in all:
        print(k)
    # return "ok"
    result = {"cost":allcost,"income":allincome}
    return json.dumps(result)
    


#这里是为了获取详详细的消费情况便于展示，这里要使用细节的cookie
@app.route('/finance/getfeesdetail',methods=["POST","GET","PUT"])
def getfeesdetail():
    # email = request.headers['email']
    #email = request.cookies["email"]
    print(request.cookies)
    email = "yang756260386@gmail.com"
    #date = content['date']
    session = Session()
    date = weektime_current()
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
    resultfromserver = {"allfees":result}
    return json.dumps(resultfromserver)











@app.route('/finance/uploadfees',methods=["POST","GET","PUT"])
def createfees():
    email = request.headers['email']
    content = request.json
    record = content['inbox']
    direction = content['direction']
    date = content['date']
    print(record)
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
    return json.dumps({"info":"记账成功","status":"ok"})
    


@app.route('/finance/updatefees',methods=["POST","GET","PUT"])
def updatefees():
    email = "yang756260386@gmail.com"
    
    content = request.json
    print(request)
    print(content)
    record = content['record']
    direction = content['direction']
    date = content['date']
    fee = content['fee']
    id = content['id']
    print(record)
    #record = "我们吃了10块钱的晚饭"
    session = Session()
    recordfromdb = session.query(Accounting).filter(Accounting.email == email).filter(Accounting.id==id).first()
    recordfromdb.record = record
    recordfromdb.direction = direction
    recordfromdb.date = date
    recordfromdb.fee = fee
    session.commit()
    session.close()
    return json.dumps({"info":"记账成功","status":"ok"})









    

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
   app.run('0.0.0.0',threaded=True,debug = True,port = 6000)




        
