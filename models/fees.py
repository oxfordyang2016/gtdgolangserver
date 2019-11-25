#!/usr/bin/env python
# -*- coding: utf-8 -*-
import jieba.posseg as pseg
from datetime import datetime, timedelta  
# coding=utf-8
from flask import Flask,render_template,request,url_for 
#import simplejson as json
import  json
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import Column, String, Integer, Date,Numeric,and_
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
    if days == "-1":
        print("i am here")
        date = datetime.strftime(datetime.now() - timedelta(1), '%y%m%d')
    if days == "0":
        print("------huodonjianshiq")
        date = datetime.strftime(datetime.now(), '%y%m%d') 
    email = request.headers['email']
    #date = content['date']
    session = Session()
    all = session.query(Accounting).filter(and_(Accounting.email == email, Accounting.date == date)).all()
    #写消费统计部分
    allcost = sum([float(row.fee) for row in all])
    for k in all:
        print(k)
    # return "ok"
    result = {"cost":allcost,"income":0}
    return json.dumps(result)
    






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
    if len(fees) >1 or len(fees) == 0:
        return  "请不要在消费中包含两个数字，我不能帮你识别"
    #接下来准备写入到Q数据库
    feefromclient = float(fees[0])
    print(feefromclient) 
    session = Session()
    oneday = Accounting(direction,record,feefromclient,date,email)
    session.add(oneday)
    session.commit()
    session.close()
    return json.dumps({"date":"time","info":"ok"})
    





    

def rowtodict(row):
    dictforsinglerow ={}
    dictforsinglerow["email"] = row.email
    dictforsinglerow["fee"] = row.fee
    dictforsinglerow["direction"]= row.direction
    print(dictforsinglerow)
    return dictforsinglerow






if __name__ == '__main__':
   app.run('0.0.0.0',threaded=True,debug = True,port = 6000)




        
