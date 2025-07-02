from Bio import Entrez

Entrez.email = "samuel.schmakel@gmail.com"

def fetch_pubmed_abstracts(query, max_results=10):
    # Search for PubMed IDs (PMIDs)
    search_handle = Entrez.esearch(db="pubmed", term=query, retmax=max_results)
    search_results = Entrez.read(search_handle)
    search_handle.close()
    id_list = search_results["IdList"]

    if not id_list:
        print("No results found.")
        return []
    
    # Fetch abstracts using the PMIDs
    fetch_handle = Entrez.efetch(db="pubmed", id=",".join(id_list), rettype="abstract", retmode="text")
    abstracts = fetch_handle.read()
    fetch_handle.close()

    return abstracts

abstracts = fetch_pubmed_abstracts("CRISPR-Cas9", 4)
print(f"Result: {abstracts}")