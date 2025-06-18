from transformers import AutoTokenizer, AutoModel
import torch
import numpy as np
from sklearn.metrics.pairwise import cosine_similarity

# Load tokenizer and model
tokenizer = AutoTokenizer.from_pretrained("microsoft/BiomedNLP-PubMedBERT-base-uncased-abstract")
model = AutoModel.from_pretrained("microsoft/BiomedNLP-PubMedBERT-base-uncased-abstract")

# Sample abstracts
abstracts = [
    "Gene expression profiling can identify potential biomarkers for breast cancer.",
    "CRISPR-Cas9 technology enables genome editing in a wide variety of organisms.",
    "Machine learning models predict drug response from tumor genomic data."
]

# Function to compute embeddings
def get_embeddings(text_list):
    embeddings = []
    for text in text_list:
        inputs = tokenizer(text, return_tensors="pt", truncation=True, padding=True, max_length=512)
        with torch.no_grad():
            outputs = model(**inputs)
            # Take the mean of the last hidden state
            mean_embedding = outputs.last_hidden_state.mean(dim=1).squeeze()
            embeddings.append(mean_embedding.numpy())
    return np.array(embeddings)

# Generate embeddings
embeddings = get_embeddings(abstracts)

# Cosine similarity matrix
similarities = cosine_similarity(embeddings)

print("Cosine Similarity Matrix:")
print(similarities)
