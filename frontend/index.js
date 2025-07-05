const form = document.getElementById('myForm');

function init() {
    form.addEventListener('submit', handleClick);
}

async function handleClick(e) {
    console.log(`This event: ${e} happened.`);
    e.preventDefault(); // Prevent page reload

    const query = document.getElementById("query");

    // TODO: Make this less hideous than the alert box
    if (!query.value.trim()) {
        alert('The query field must not be empty.')
        return;
    }
    
    const formData = new FormData(e.target);
    for (const pair of formData.entries()) {
        console.log(`${pair[0]}: ${pair[1]}`);
      }
    const params = new URLSearchParams(formData);
    console.log(`params: ${params}`)

    let url = `http://localhost:8080/api/data?${params.toString()}`

    console.log(`url sent to backend: ${url}`)

    const response = await fetch(url, {
        method: 'GET',
    })

    const result = await response.json();
    result.forEach((paper, index) => {
        console.log(`Paper ${index+1}:`);
        console.log(`Title: ${paper.title}`);
        console.log(`Abstract: ${paper.abstract}`);
        console.log(`URL: ${paper.url}`);
    });
    
    const container = document.getElementById('papers-container');

    result.forEach((paper) => {
    const paperElement = document.createElement('div');
    paperElement.className = 'paper-card';
    paperElement.innerHTML = `
        <h3>${paper.title} 
            <a href="url">${paper.url}</a>
        </h3>
        <p><strong>Abstract:</strong> ${paper.abstract}</p>
    `;
    container.appendChild(paperElement);
    });
}

init();