from fastapi import FastAPI
import uvicorn
import models
import numpy as np
import pandas as pd

print("Starting main.py")

app = FastAPI()

reference_abstracts = [
        "Mitochondrial dysfunction is a hallmark of Alzheimer's disease. Impaired mitochondrial biogenesis contributes to neurodegeneration.",
        "Gene editing using CRISPR technology offers new therapeutic possibilities for genetic disorders."
    ]
target_abstracts = [
        "The role of mitochondrial dysfunction in neurodegenerative diseases has been extensively studied. Recent research shows that mitochondrial biogenesis is impaired in Alzheimer's disease.",
        "CRISPR-Cas9 gene editing technology has revolutionized molecular biology research. This study demonstrates its application in correcting genetic mutations.",
        "Machine learning algorithms are increasingly being applied to drug discovery. Deep learning models can predict molecular properties with high accuracy.",
        "The gut microbiome plays a crucial role in human health and disease. Dysbiosis has been linked to various metabolic disorders.",
        "Immunotherapy has shown promising results in cancer treatment. Checkpoint inhibitors have improved survival rates in melanoma patients."
    ]

tokenizer, model = models.load_biobert_model()
print("loaded tokenizer and model")

# Use the same model for multiple operations
ref_embeddings = models.get_biobert_embeddings(reference_abstracts, tokenizer=tokenizer, model=model)
target_embeddings = models.get_biobert_embeddings(target_abstracts, tokenizer=tokenizer, model=model)

print(f"length of ref_embeddings: {len(ref_embeddings)}")
print(f"length of target_embeddings: {len(target_embeddings)}")

# Compute similarities
similarity_matrix = models.compute_similarity_matrix(target_embeddings, ref_embeddings)

df = models.create_dataframe(target_abstracts, similarity_matrix)
print("df:", df)