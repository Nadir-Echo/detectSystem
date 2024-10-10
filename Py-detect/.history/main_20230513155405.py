import json

from fastapi import FastAPI
import cv2
import ml
import numpy as np
import requests

from fastapi import FastAPI, File, UploadFile

app = FastAPI()


# uvicorn main:app --reload
@app.get("")

@app.post("/steel/defectimg")
async def root(file: UploadFile):
    print('777')
    # return {"filename": file.filename}
    img = await file.read()
    # return {"file_size": len(file)}
    # 用opencv读取file字节流图片


    img = cv2.imdecode(np.frombuffer(img, np.uint8), cv2.IMREAD_COLOR)
    img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
    # cv2.imshow('img', img)
    # cv2.waitKey(0)
    # cv2.destroyAllWindows()

    # detect image in detecr function,and return the json data
    data = ml.detect(img)
    print(data)
    
        # data = {"class": "无缺陷", "conf": 0.0, "position": [0, 0, 0, 0],err_result:0}
    # add a field called err_result to the 'data'
    if len(data)==0:
        data.append(err_)
    else:
        data.err_result=1
    # print("json:")
    #convert data to json format
    # data = json.dumps(data)
    # print(data)
    #return response
    return data
