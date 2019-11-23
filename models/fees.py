#!/usr/bin/env python
# -*- coding: utf-8 -*-
import jieba.posseg as pseg

# coding=utf-8
from flask import Flask,render_template,request,url_for 
#import simplejson as json
import  json
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import Column, String, Integer, Date,Numeric
# from models import techniqueanalysis
#from sqlalchemy.ext.declarative import declarative_base
#from flaskjsontools import JsonSerializableBase
engine = create_engine('mysql://root:123456@localhost:3306/finance')
Session = sessionmaker(bind=engine)
Base = declarative_base()


# class DecimalEncoder(json.JSONEncoder):
#     def _iterencode(self, o, markers=None):
#         if isinstance(o, decimal.Decimal):
#             # wanted a simple yield str(o) in the next line,
#             # but that would mean a yield on the line with super(...),
#             # which wouldn't work (see my comment below), so...
#             return (str(o) for o in [o])
#         return super(DecimalEncoder, self)._iterencode(o, markers)


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


@app.route('/createfee')
def getfees():
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
    dictforsinglerow["id"] = row.id
    dictforsinglerow["note"] = row.note
    dictforsinglerow["classname"]= row.classname
    dictforsinglerow["time"] = str(row.time)
    dictforsinglerow["highest"] =row.highest
    dictforsinglerow["lowest"] = row.lowest
    dictforsinglerow["opening"] = row.opening
    dictforsinglerow["closing"]  = row.closing
    dictforsinglerow["volume"] = row.volume
    print(dictforsinglerow)
    return dictforsinglerow






if __name__ == '__main__':
   app.run('0.0.0.0',threaded=True,debug = True,port = 6000)




        
