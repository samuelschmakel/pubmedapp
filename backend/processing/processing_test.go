package processing

import (
	"strings"
	"testing"
)

func TestFetchAbstracts(t *testing.T) {
	tests := []struct {
		name string
		inputURL string
		expected string
		errorContains string
	}{
		{
		name: "test fetchAbstracts",
		inputURL: "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=pubmed&id=40601938,40601888&rettype=abstract&retmode=text&email=samuel.schmakel@gmail.com",
		expected: `
		1. J Am Chem Soc. 2025 Jul 2. doi: 10.1021/jacs.5c05148. Online ahead of print.

Biomineralized Engineered Bacterial Outer Membrane Vesicles as cGAS-STING 
Nanoagonists Synergize with Lactate Metabolism Modulation to Potentiate 
Immunotherapy.

Li QR(1), Zhang X(1), Zhang C(1), Zhang Y(1), Niu MT(1), Chen Z(1), Zhang SM(1), 
He J(1), Chen WH(1), Zhang XZ(1).

Author information:
(1)Key Laboratory of Biomedical Polymers of Ministry of Education, Department of 
Chemistry, Department of Cardiology, Zhongnan Hospital, Wuhan University, Wuhan 
430072, P. R. China.

The immunosuppressive tumor microenvironment (TME) significantly limits the 
efficacy of cancer immunotherapy. Activation of the cyclic guanosine 
monophosphate-adenosine monophosphate synthase (cGAS) stimulator of interferon 
genes (STING) pathway and depletion of the tumor metabolic byproduct lactate 
(LA) represent promising strategies to reverse the immunosuppressive TME and 
enhance antitumor therapeutic outcomes. Herein, biomineralized engineered 
bacterial outer membrane vesicles (OMVs@MnCaP-FA) are developed to 
synergistically activate the cGAS-STING pathway and modulate LA metabolism for 
antitumor immunotherapy. Upon internalization by 4T1 tumor cells, OMVs@MnCaP-FA 
undergo acid-responsive degradation, releasing Ca2+, Mn2+, and lactate oxidase 
(LOX)-expressing OMVs (OMVs-EcL). These components collectively promote 
mitochondrial DNA (mtDNA) generation, enhance cGAS-mediated mtDNA recognition 
and cyclic GMP-AMP (cGAMP) production, and potentiate the binding of cGAMP to 
STING, leading to robust activation of the cGAS-STING signaling pathway. More 
importantly, OMVs-EcL-mediated LA depletion reprograms the immunosuppressive TME 
into an immunoresponsive state, revitalizing antitumor immunity. In vivo studies 
demonstrate that the combined activation of the cGAS-STING pathway and 
regulation of LA metabolism effectively inhibit primary tumor growth and 
metastatic progression, highlighting the potential of this synergistic strategy 
for advancing antitumor immunotherapy.

DOI: 10.1021/jacs.5c05148
PMID: 40601938


2. J Clin Oncol. 2025 Jul 2:JCO2500234. doi: 10.1200/JCO-25-00234. Online ahead
of  print.

Tumor-Intrinsic and Microenvironmental Determinants of Impaired Antitumor 
Immunity in Chromophobe Renal Cell Carcinoma.

Labaki C(1)(2)(3), Saad E(1)(3), Madsen KN(3), Hobeika C(4)(5), Bi K(1), 
Alchoueiry M(5), Camp S(1), Hou Y(6), Bakouny Z(7), Matar S(8), El Ahmar N(8), 
Nyman J(9), Zhang L(10), Priolo C(5), Rout R(3), Daou M(11), Khabibullin D(5), 
Salem S(5), Schindler N(3), Saliby RM(3)(11), Meli K(12), Wells JC(13), Pimenta 
E(1), Takemura K(14), Park J(1), Eid M(1), Semaan K(1), Fu J(1), Denize T(15), 
El Hajj Chehade R(1), Machaalani M(1), Nawfal R(1), Khatoun WD(1), Saleh M(1), 
El Masri J(1), Haddad NR(16), Xu W(1), McGregor BA(1), Hirsch MS(8), Xie W(1), 
Heng DYC(14), McDermott DF(17), Signoretti S(8), Van Allen EM(1), Shukla SA(18), 
Choueiri TK(1), Henske EP(5), Braun DA(3).

Author information:
(1)Department of Medical Oncology, Dana-Farber Cancer Institute, Boston, MA.
(2)Department of Medicine, Beth Israel Deaconess Medical Center, Harvard Medical 
School, Boston, MA.
(3)Center of Molecular and Cellular Oncology (CMCO), Yale School of Medicine, 
New Haven, CT.
(4)Department of Medicine, Cleveland Clinic Fairview Hospital, Cleveland, OH.
(5)Pulmonary and Critical Care Medicine, Department of Medicine, Brigham and 
Women's Hospital, Harvard Medical School, Boston, MA.
(6)Akoya Biosciences Inc, Marlborough, MA.
(7)Department of Medical Oncology, Memorial Sloan Kettering Cancer Center, New 
York, NY.
(8)Department of Pathology, Brigham and Women's Hospital, Boston, MA.
(9)PathAI, Inc, Boston, MA.
(10)University of Shanghai for Science and Technology, Shanghai, China.
(11)Department of Medicine, Yale University, New Haven, CT.
(12)RBC Capital Markets, New York, NY.
(13)BC Cancer Agency, Vancouver, Canada.
(14)Tom Baker Cancer Centre, University of Calgary, Calgary, AB, Canada.
(15)Department of Pathology, Massachusetts General Hospital, Boston, MA.
(16)Johns Hopkins University School of Medicine, Baltimore, MD.
(17)Department of Medical Oncology, Beth Israel Deaconess Medical Center, 
Harvard Medical School, Boston, MA.
(18)Department of Hematopoietic Biology and Malignancy, The University of Texas 
MD Anderson Cancer Center, Texas, Houston, TX.

PURPOSE: While immune checkpoint inhibition (ICI) has transformed the management 
of many advanced renal cell carcinomas (RCCs), the determinants of effective 
antitumor immunity for chromophobe RCC (ChRCC) and renal oncocytic tumors remain 
an unmet clinical and scientific need.
METHODS: Single-cell transcriptomic and T-cell receptor profiling was performed 
on tumor and adjacent normal tissue of patients with ChRCC and renal oncocytic 
neoplasms. Using machine learning, the cellular origin of renal oncocytic 
neoplasms was evaluated, with analysis of associated oncogenic pathways. Using 
immunohistochemistry, immune infiltration was analyzed in renal oncocytic 
neoplasms in comparison with clear cell RCC (ccRCC). Immune checkpoint 
expression, clonal expansion, and tumor specificity were compared between ChRCC 
and ccRCC. Using the International Metastatic RCC Database Consortium data set, 
clinical outcomes of patients with metastatic ChRCC (mChRCC) treated with 
first-line systemic regimens were compared with those of patients with ccRCC.
RESULTS: We validated α-intercalated cells as the cellular origin of renal 
oncocytic neoplasms. We identified a downregulation of HLA class I molecules 
with enrichment of potentially targetable pathways including mammalian target of 
rapamycin and ferroptosis in ChRCC. The tumor microenvironment of ChRCC showed 
markedly decreased immune infiltration, with a pronounced depletion in 
tumor-infiltrating CD8+ T cells. ChRCC-infiltrating CD8+ T cells demonstrated 
lower immune checkpoint expression, diminished clonal expansion, and decreased 
tumor specificity. Clinical analysis identified poor survival outcomes 
selectively among patients with mChRCC treated with immune-based therapies.
CONCLUSION: Immunogenomic analysis of ChRCC revealed profound depletion of T 
cells, with an immune phenotype marked by a lack of expression of immune 
checkpoints and poor tumor specificity, suggesting that the few T cells in these 
tumor types are likely nonspecific bystanders. This immune-cold environment 
hinders an effective response to immunotherapy and underscores the need for 
ChRCC-tailored treatments designed to improve tumor-specific T-cell infiltration 
into the microenvironment.

DOI: 10.1200/JCO-25-00234
PMID: 40601888
		`,
	},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := FetchAbstracts(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
					t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
					return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}