from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()



def confirmation():
    print("Connected to backend from models!")