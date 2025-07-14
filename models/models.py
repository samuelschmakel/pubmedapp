import numpy as np
from transformers import AutoTokenizer, AutoModel
import torch
from sklearn.metrics.pairwise import cosine_similarity
from typing import List, Tuple, Optional
import pandas as pd

# Global variables to cache model and tokenizer
_cached_model = None
_cached_tokenizer = None
_cached_model_name = None

def load_biobert_model(model_name: str = "dmis-lab/biobert-v1.1"):
    """
    Load BioBERT model and tokenizer.
    Args:
        model_name: BioBERT model name from Hugging Face
        
    Returns:
        tuple: (tokenizer, model)
    """
    global _cached_model, _cached_tokenizer, _cached_model_name
    
    # Use cached model if same model_name
    if _cached_model is not None and _cached_model_name == model_name:
        return _cached_tokenizer, _cached_model

    print(f"Loading BioBERT model: {model_name}")
    tokenizer = AutoTokenizer.from_pretrained(model_name)
    model = AutoModel.from_pretrained(model_name)
    model.eval()
    
    # Cache for future use
    _cached_tokenizer = tokenizer
    _cached_model = model
    _cached_model_name = model_name
    
    return tokenizer, model

def get_biobert_embeddings(
    texts: List[str], 
    model_name: str = "dmis-lab/biobert-v1.1",
    batch_size: int = 8,
    tokenizer: Optional[AutoTokenizer] = None,
    model: Optional[AutoModel] = None
) -> np.ndarray:
    """
    Get BioBERT embeddings for a list of texts.
    
    Args:
        texts: List of text strings
        model_name: BioBERT model name (ignored if tokenizer/model provided)
        batch_size: Batch size for processing
        tokenizer: Pre-loaded tokenizer (optional)
        model: Pre-loaded model (optional)
        
    Returns:
        numpy array of embeddings
    """
    # Load model if not provided
    if tokenizer is None or model is None:
        tokenizer, model = load_biobert_model(model_name)
    
    embeddings = []
    
    with torch.no_grad():
        for i in range(0, len(texts), batch_size):
            batch_texts = texts[i:i + batch_size]
            
            # Tokenize batch
            inputs = tokenizer(
                batch_texts,
                padding=True,
                truncation=True,
                max_length=512,
                return_tensors="pt"
            )
            
            # Get model outputs
            outputs = model(**inputs)
            
            # Use [CLS] token embedding as sentence representation
            batch_embeddings = outputs.last_hidden_state[:, 0, :].cpu().numpy()
            embeddings.extend(batch_embeddings)
    
    return np.array(embeddings)

def compute_similarity_matrix(
    target_embeddings: np.ndarray, 
    reference_embeddings: np.ndarray
) -> np.ndarray:
    """
    Compute cosine similarity matrix between target and reference embeddings.
    
    Args:
        target_embeddings: Embeddings for target abstracts
        reference_embeddings: Embeddings for reference abstracts
        
    Returns:
        Similarity matrix of shape (n_targets, n_references)
    """
    return cosine_similarity(target_embeddings, reference_embeddings)