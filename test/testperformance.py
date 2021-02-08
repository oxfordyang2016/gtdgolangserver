import requests
def testperformance(k):
    for time in range(k):
        print("i am requests the "+str(time)+"times")
        a=requests.post("http://localhost:8080/createtask",data={"status":"finish","inbox":"i am requests "+str(time)+"times","project":"gtd","plantime":"sdf"}).text
        print(a)
        import time
        time.sleep(2)
