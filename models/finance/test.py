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










# 2 - generate database schema
Base.metadata.create_all(engine)




app = Flask(__name__)

@app.route('/hello')
def hello_world():
   return 'Hello World'




if __name__ == '__main__':
    logger.add("./logging/file_1.log", rotation="50 MB",backtrace=True, diagnose=True,colorize=True) 
    app.run('0.0.0.0',threaded=True,port = 6000,debug=True)




        

