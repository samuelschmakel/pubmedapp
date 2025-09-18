from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
import uvicorn
import models
import numpy as np

class ArticleInfo(BaseModel):
    title: str
    abstract: str
    url: str

class PythonAPIInput(BaseModel):
    articleInfo: List[ArticleInfo]
    context: List[str]

print("Starting main.py")

app = FastAPI()

@app.post("/process-list")
async def process_list(input_data: PythonAPIInput):
    print("in process_list")
    reference_abstracts = input_data.context
    target_abstracts = [item.abstract for item in input_data.articleInfo]
    print(f'reference_abstracts: {reference_abstracts}')
    print(f'target_abstracts: {target_abstracts}')
    results = models.get_similarity_df(reference_abstracts, target_abstracts)
    print(f'length of results: {len(results)}')

    # Transform DataFrame to a list, a format Go expects
    dataframe_rows = []

    for index, row in results.iterrows():
        dataframe_rows.append({
            "abstract": row['abstract'],
            "similarity_score": float(row['avg_similarity'])  # Convert to native Python float
        })

    print(f"Transformed data: {dataframe_rows}")  # Debug output
    return dataframe_rows  # Return array directly, not wrapped in an object

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8001)