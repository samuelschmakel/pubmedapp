import models

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

results = models.get_similarity_df(reference_abstracts, target_abstracts)
print(f"results: {results}")